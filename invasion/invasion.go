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

// Invasion provides the publicly exposed API of the lib.
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

// Iterate is the main loop of the invasion. It checks if we can proceed to
// invade cities, executes fights and moves. It returns an error if all reamining
// aliens are trapped.
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

// peace indicates if the invasion has finished.
func (i Invasion) peace() bool {
	return len(i.invadedCities) == 0 || len(i.alienMoves) == 0
}

// move moves one alien at a time. If no alien can be moved, false is returned.
// The criteria for moving is that a neighboring city is reachable via an inbound
// link from the currently invaded city and that the city is not invaded by more
// than `maxInvaders` aliens.
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
				// note, this is always the case for maxInvaders=2 but we support
				// higher maxInvader counts as well.
				if len(i.invadedCities[current]) == 1 {
					delete(i.invadedCities, current)
				} else { // maxInvaders > 2
					// remove the alien from the currently invaded city.
					// here we have at least one invader left so we remove the
					// alien from the invaded city only.
					delete(i.invadedCities[current], alienName)
				}

				// add the alien to the next city to be invaded.
				if _, ok := i.invadedCities[next]; !ok {
					i.invadedCities[next] = make(map[string]struct{})
				}
				i.invadedCities[next][alienName] = struct{}{}

				// increase the move counter of the alien.
				if _, ok := i.alienMoves[alienName]; ok {
					i.alienMoves[alienName]++
				}
				// if the alien reached the max move counter threshold, stop
				// counting by just removing it from the moves counter index.
				// note, this does not mean that it is trapped.
				if i.alienMoves[alienName] == maxMoves {
					delete(i.alienMoves, alienName)
				}
				return true
			}
			// no city to go to so we remove the
			// trapped alien from the move counter.
			delete(i.alienMoves, alienName)
		}
	}
	return false
}

// fight executes fights and desrtoys all cities invaded by maxInvaders aliens.
func (i Invasion) fight() {
	for cityName, alienNames := range i.invadedCities {
		if len(alienNames) != MaxInvaders { // no fight
			return
		}

		// remove the destroyed city from the world
		// map which removes the aliens as well.
		i.worldMap.DestroyCity(cityName)
		delete(i.invadedCities, cityName)

		// clean the move counter.
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
