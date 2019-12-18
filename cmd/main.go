package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fgrimme/alien-invasion/encoding"
	"github.com/fgrimme/alien-invasion/invasion"
)

var (
	version = "unknown" // will be compiled into the binary, see Makefile

	inFile       string
	numAliens    int
	printVersion bool
)

func main() {
	flag.StringVar(&inFile, "in", "", "input file holding the map definition")
	flag.IntVar(&numAliens, "n", 0, "count of aliens invading the world")
	flag.BoolVar(&printVersion, "v", false, "print version")
	flag.Parse()

	if printVersion {
		fmt.Println(version)
		os.Exit(0)
	} else if len(inFile) == 0 {
		fmt.Println("no input file specified")
		os.Exit(1)
	} else if numAliens == 0 {
		fmt.Println("count of aliens must be greater than zero")
		os.Exit(1)
	}

	file, err := os.Open(inFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	worldMap, err := encoding.Decode(file)
	if err != nil {
		fmt.Printf("failed to build world: %+v", err)
		file.Close()
		os.Exit(1) // does not execute deferred funcs so we need to close the file
	}
	file.Close()

	if count, max := numAliens/invasion.MaxInvaders, len(worldMap); count > max {
		fmt.Printf("too many aliens, max count for %d cities is %d\n", max, max*invasion.MaxInvaders)
		os.Exit(1)
	}

	// execute the main loop of the invasion
	if err := invasion.New(worldMap, numAliens).Iterate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := "invasion completed"
	w := encoding.Encode(worldMap)
	if len(w) > 0 {
		s = fmt.Sprintf("%s, state of the world:\n%s", s, w)
	}
	fmt.Println(s)
}
