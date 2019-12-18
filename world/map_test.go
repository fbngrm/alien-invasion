package world_test

import (
	"reflect"
	"testing"

	"github.com/fgrimme/alien-invasion/world"
)

var linkTests = []struct {
	d string    // description of test case
	m world.Map // map before linking
	r world.Map // expected state after linking
}{
	{
		d: "expect link from Foo to Bar to be added to empty map",
		m: world.Map{},
		r: world.Map{"Foo": {"Bar": "east"}, "Bar": {}},
	},
	{
		d: "expect link from Foo to Bar to be added to non-empty map",
		m: world.Map{"Baz": {"Foo": "south"}},
		r: world.Map{"Baz": {"Foo": "south"}, "Foo": {"Bar": "east"}, "Bar": {}},
	},
	{
		d: "expect link from Foo to Baz to be added to existing Foo links",
		m: world.Map{"Baz": {"Foo": "south"}, "Foo": {"Bar": "east"}, "Bar": {}},
		r: world.Map{"Baz": {"Foo": "south"}, "Foo": {"Bar": "east", "Baz": "north"}, "Bar": {}},
	},
	{
		d: "expect link to be added to existing inbound Bar",
		m: world.Map{"Baz": {"Foo": "south"}, "Foo": {"Bar": "east"}, "Bar": {}},
		r: world.Map{"Baz": {"Foo": "south"}, "Foo": {"Bar": "east", "Baz": "north"}, "Bar": {"Foo": "west"}},
	},
}

func TestLink(t *testing.T) {
	for _, tt := range linkTests {
		for origin, links := range tt.r {
			for destination, direction := range links {
				tt.m.Link(origin, destination, direction)
			}
		}
		if !reflect.DeepEqual(tt.m, tt.r) {
			t.Errorf("test case failed: %s\nwant:\n%+v\ngot:\n%+v", tt.d, tt.r, tt.m)
		}
	}
}

var destructionTests = []struct {
	d string    // description of test case
	m world.Map // map before destruction
	r world.Map // expected state after destruction
	c []string  // city names to destroy
}{
	{
		d: "expect no cities to be destroyed",
		m: world.Map{"Foo": {"Bar": "east"}},
		r: world.Map{"Foo": {"Bar": "east"}},
		c: []string{},
	},
	{
		d: "expect Foo to be destroyed",
		m: world.Map{"Foo": {"Bar": "east"}},
		r: world.Map{},
		c: []string{"Foo"},
	},
	{
		d: "expect Foo to be destroyed but Bar to survive",
		m: world.Map{"Foo": {"Bar": "east"}, "Bar": {"Baz": "north"}},
		r: world.Map{"Bar": {"Baz": "north"}},
		c: []string{"Foo"},
	},
	{
		d: "expect Foo to be destroyed and removed from Bar's links",
		m: world.Map{"Foo": {"Bar": "east"}, "Bar": {"Foo": "west"}},
		r: world.Map{"Bar": {}},
		c: []string{"Foo"},
	},
}

func TestDestroyCity(t *testing.T) {
	for _, tt := range destructionTests {
		for _, cityName := range tt.c {
			tt.m.DestroyCity(cityName)
			if !reflect.DeepEqual(tt.m, tt.r) {
				t.Errorf("test case failed: %s", tt.d)
			}
		}
	}
}
