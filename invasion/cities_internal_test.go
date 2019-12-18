package invasion

import (
	"testing"
)

var invadedCitiesTests = []struct {
	d string   // description of test case
	a []string // alien names
	c []string // city names

}{
	{
		d: "expect no aliens to be added",
		a: []string{},
		c: []string{"Foo", "Bar", "Baz", "Qu-ux"},
	},
	{
		d: "expect one alien to be added",
		a: []string{"alien 1"},
		c: []string{"Foo", "Bar", "Baz", "Qu-ux"},
	},
	{
		d: "expect all aliens to be added",
		a: []string{"alien 1", "alien 2", "alien 3"},
		c: []string{"Foo", "Bar", "Baz", "Qu-ux"},
	},
	{
		d: "expect all cites to be invaded by 2 aliens",
		a: []string{"alien 1", "alien 2", "alien 3", "alien 4"},
		c: []string{"Foo", "Bar"},
	},
}

func TestInitInvadedCities(t *testing.T) {
	for _, tt := range invadedCitiesTests {
		a := make([]string, 0)
		i := initInvadedCities(tt.a, tt.c)
		for cityName, aliens := range i {
			// consider MaxInvaders
			if want, got := MaxInvaders, len(aliens); want < got {
				t.Errorf("%s: max invader count %d but got %d", tt.d, want, got)
			}
			// all added cities are valid
			if !contains(tt.c, cityName) {
				t.Errorf("%s: unknown city %s", tt.d, cityName)
			}
			// all added aliens are valid
			for alien := range aliens {
				a = append(a, alien)
				if !contains(tt.a, alien) {
					t.Errorf("%s: unknown alien %s", tt.d, alien)
				}
			}
		}
		// all aliens have been added
		for _, alien := range tt.a {
			if !contains(a, alien) {
				t.Errorf("%s: missing %s", tt.d, alien)
			}
		}
	}
}

// helper to check if e is contained by s.
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
