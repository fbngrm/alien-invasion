package invasion

import (
	"reflect"
	"testing"

	"github.com/fgrimme/alien-invasion/world"
)

var moveTests = []struct {
	d string    // description of test case
	i *Invasion // state before fights
	r *Invasion // expected state after fights
}{
	{
		d: "expect alien to be trapped and removed from counter",
		i: &Invasion{
			worldMap:   world.Map{"Foo": {"Bar": "east"}},
			alienMoves: alienMoves{"alien 1": 0},
			invadedCities: invadedCities{
				"Foo": {"alien 1": struct{}{}},
				"Bar": {"alien 2": struct{}{}, "alien 3": struct{}{}},
			},
		},
		r: &Invasion{
			worldMap:   world.Map{"Foo": {"Bar": "east"}},
			alienMoves: alienMoves{},
			invadedCities: invadedCities{
				"Foo": {"alien 1": struct{}{}},
				"Bar": {"alien 2": struct{}{}, "alien 3": struct{}{}},
			},
		},
	},
	{
		d: "expect alien to be trapped and removed from counter",
		i: &Invasion{
			worldMap:   world.Map{"Foo": {}},
			alienMoves: alienMoves{"alien 1": maxMoves - 1},
			invadedCities: invadedCities{
				"Foo": {"alien 1": struct{}{}},
			},
		},
		r: &Invasion{
			worldMap:   world.Map{"Foo": {}},
			alienMoves: alienMoves{},
			invadedCities: invadedCities{
				"Foo": {"alien 1": struct{}{}},
			},
		},
	},
	{
		d: "expect alien to move to a new city",
		i: &Invasion{
			worldMap:   world.Map{"Foo": {"Bar": "east"}},
			alienMoves: alienMoves{"alien 1": 0},
			invadedCities: invadedCities{
				"Foo": {"alien 1": struct{}{}},
			},
		},
		r: &Invasion{
			worldMap:   world.Map{"Foo": {"Bar": "east"}},
			alienMoves: alienMoves{"alien 1": 1},
			invadedCities: invadedCities{
				"Bar": {"alien 1": struct{}{}},
			},
		},
	},
}

func TestMove(t *testing.T) {
	for _, tt := range moveTests {
		tt.i.move()
		if !reflect.DeepEqual(tt.r, tt.i) {
			t.Errorf("%s\nwant\n%+v\ngot\n%+v", tt.d, tt.r, tt.i)
		}
	}
}

var fightTests = []struct {
	d string    // description of test case
	i *Invasion // state before fights
	r *Invasion // expected state after fights
}{
	{
		d: "expect invaded cities and aliens to be destroyed",
		i: &Invasion{
			worldMap:   world.Map{"Foo": {"Bar": "east"}, "Bar": {"Baz": "north"}, "Baz": {}},
			alienMoves: alienMoves{"alien 1": 0, "alien 2": 0, "alien 3": 0, "alien 4": 0},
			invadedCities: invadedCities{
				"Bar": {"alien 1": struct{}{}, "alien 2": struct{}{}},
				"Foo": {"alien 3": struct{}{}, "alien 4": struct{}{}},
			},
		},
		r: &Invasion{
			worldMap:      world.Map{"Baz": {}},
			alienMoves:    alienMoves{},
			invadedCities: invadedCities{},
		},
	},
	{
		d: "expect one invaded city and it's invader to survive",
		i: &Invasion{
			worldMap:   world.Map{"Foo": {"Bar": "east"}, "Bar": {"Baz": "north"}, "Baz": {}},
			alienMoves: alienMoves{"alien 1": 0, "alien 2": 0, "alien 3": 0},
			invadedCities: invadedCities{
				"Bar": {"alien 1": struct{}{}, "alien 2": struct{}{}},
				"Foo": {"alien 3": struct{}{}},
			},
		},
		r: &Invasion{
			worldMap:   world.Map{"Foo": {}, "Baz": {}},
			alienMoves: alienMoves{"alien 3": 0},
			invadedCities: invadedCities{
				"Foo": {"alien 3": struct{}{}},
			},
		},
	},
	{
		d: "expect nothing to be destroyed",
		i: &Invasion{
			worldMap:   world.Map{"Foo": {"Bar": "east"}, "Bar": {"Baz": "north"}, "Baz": {}},
			alienMoves: alienMoves{"alien 1": 0, "alien 2": 0, "alien 3": 0},
			invadedCities: invadedCities{
				"Bar": {"alien 1": struct{}{}},
				"Baz": {"alien 2": struct{}{}},
				"Foo": {"alien 3": struct{}{}},
			},
		},
		r: &Invasion{
			worldMap:   world.Map{"Foo": {"Bar": "east"}, "Bar": {"Baz": "north"}, "Baz": {}},
			alienMoves: alienMoves{"alien 1": 0, "alien 2": 0, "alien 3": 0},
			invadedCities: invadedCities{
				"Bar": {"alien 1": struct{}{}},
				"Baz": {"alien 2": struct{}{}},
				"Foo": {"alien 3": struct{}{}},
			},
		},
	},
}

func TestFight(t *testing.T) {
	for _, tt := range fightTests {
		tt.i.fight()
		if !reflect.DeepEqual(tt.r, tt.i) {
			t.Errorf("%s\nwant\n%+v\ngot\n%+v", tt.d, tt.r, tt.i)
		}
	}
}
