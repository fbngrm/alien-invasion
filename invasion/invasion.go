package invasion

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fgrimme/alien-invasion/world"
)

const (
	maxMoves    = 10000
	MaxInvaders = 2
)

type Invasion struct {
	worldMap      world.Map     // links by city name
	alienMoves    alienMoves    // moves by alien names
	invadedCities invadedCities // alien names by city names
}

func New(worldMap world.Map, numAliens int) *Invasion {
	alienNames := initAlienNames(numAliens)
	invasion := &Invasion{
		worldMap:      worldMap,
		alienMoves:    initAlienMoves(alienNames),
		invadedCities: initInvadedCities(alienNames, worldMap.Cities()),
	}
	return invasion
}

func (i *Invasion) Iterate() error {
	i.fight()
	if i.peace() {
		return nil
	}
	if !i.move() {
		return errors.New("all aliens trapped")
	}
	if i.peace() {
		return nil
	}
	return i.Iterate()
}

func (i Invasion) peace() bool {
	return len(i.invadedCities) == 0 || len(i.alienMoves) == 0
}

// we move one alien at a time
func (i Invasion) move() bool {
	for current, aliens := range i.invadedCities {
		linkedCities := i.worldMap[current]
		for alienName := range aliens {
			for next := range linkedCities {
				if len(i.invadedCities[next]) >= MaxInvaders { // city can't be invaded
					continue
				}
				// remove the formerly invaded city from the invaded cities if
				// not invaded by any other alien.
				// note, this is always the case for maxInvaders=2
				if len(i.invadedCities[current]) == 1 {
					delete(i.invadedCities, current)
				} else {
					// remove the alien from the currently invaded city
					delete(i.invadedCities[current], alienName)
				}

				// add the alien to the next city to be invaded
				if _, ok := i.invadedCities[next]; !ok {
					i.invadedCities[next] = make(map[string]struct{})
				}
				i.invadedCities[next][alienName] = struct{}{}

				// increase the move counter
				if _, ok := i.alienMoves[alienName]; ok {
					i.alienMoves[alienName]++
				}
				// if the alien reached the max move counter threshold, stop
				// counting by just removing it from the counter index
				if i.alienMoves[alienName] == maxMoves {
					delete(i.alienMoves, alienName)
				}
				return true
			}
			// remove trapped aliens from the move counter
			delete(i.alienMoves, alienName)
		}
	}
	return false
}

func (i Invasion) fight() {
	for cityName, alienNames := range i.invadedCities {
		if len(alienNames) == MaxInvaders {
			i.worldMap.DestroyCity(cityName)
			delete(i.invadedCities, cityName)

			aliens := make([]string, len(alienNames))
			var y int
			for alienName := range alienNames {
				delete(i.alienMoves, alienName)
				aliens[y] = alienName
				y++
			}
			fmt.Printf("%s has been destroyed by %s!\n", cityName, strings.Join(aliens, " and "))
		}
	}
}
