package encoding

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/fgrimme/alien-invasion/world"
)

func Decode(r io.Reader) (world.Map, error) {
	world := make(world.Map)
	scanner := bufio.NewScanner(r)
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 0 {
			return nil, fmt.Errorf("failed to parse line: %d", i)
		}
		cityName := parts[0]
		if len(parts) > 1 {
			for _, link := range parts[1:] {
				linkParts := strings.Split(link, "=")
				if len(linkParts) != 2 {
					return nil, fmt.Errorf("failed to parse line: %d", i)
				}
				world.Link(cityName, linkParts[1], linkParts[0])
			}
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return world, nil
}

func Gen(w, h int) world.Map {
	worldMap := make(world.Map)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			name := fmt.Sprintf("%d:%d", x, y)
			if y-1 >= 0 {
				north := fmt.Sprintf("%d:%d", x, y-1)
				worldMap.Link(name, north, "north")
			}
			if x+1 < w {
				east := fmt.Sprintf("%d:%d", x+1, y)
				worldMap.Link(name, east, "east")
			}
			if y+1 < h {
				south := fmt.Sprintf("%d:%d", x, y+1)
				worldMap.Link(name, south, "south")
			}
			if x-1 >= 0 {
				west := fmt.Sprintf("%d:%d", x-1, y)
				worldMap.Link(name, west, "west")
			}
		}
	}
	return worldMap
}
