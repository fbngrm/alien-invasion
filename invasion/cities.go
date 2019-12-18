package invasion

import (
	"math/rand"
	"time"
)

// alien names by city name
type invadedCities map[string]map[string]struct{}

// distribute aliens pseudo randomly.
func initInvadedCities(names alienNames, cityNames []string) invadedCities {
	rand.Seed(time.Now().UTC().UnixNano())
	invaded := make(invadedCities)
	for _, alienName := range names {

		// shuffle, potentially expensive for large slices.
		// todo: use random index or something more efficient
		// and an interface to make test order deterministic.
		for i := range cityNames {
			j := rand.Intn(i + 1)
			cityNames[i], cityNames[j] = cityNames[j], cityNames[i]
		}

		for _, cityName := range cityNames {
			if len(invaded[cityName]) < MaxInvaders {
				if _, ok := invaded[cityName]; !ok {
					invaded[cityName] = make(map[string]struct{})
				}
				invaded[cityName][alienName] = struct{}{}
				break
			}
		}
	}
	return invaded
}
