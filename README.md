# Alien-Invasion

This program implements the requirements defined in the [task description](https://github.com/fbngrm/alien-invasion/blob/master/Alien%20Invasion.pdf).

### Assumptions

We assume that links between cities are uni-directional, even though most
real-life roads are bi-directional. This assumption is based on the fact that
in the given example, `Foo` declares a link to `Bar` as well as the other way
around.

Meaning, if aliens can travel from city `Foo` to city `Bar` directly, they
do not necessarily be able to travel from `Bar` to `Foo` directly.
Hence, the program distinguishes between inbound and outbound links.

Based on the aforementioned assumptions, the world map is implemented as a directed graph.
We use an adjacency list to represent adjacent cities/vertices since we assume
the graph is sparse given the max degree/edge count of 4.

The definition of the input/output format of the world map implies, that
cities with no links cannot be represented. Thus, when initializing the invasion,
isolated cities with no links are ignored by the invaders.
Cities with inbound connections only, may get invaded or destroyed. In
any case they can, by format definition, only be present in the output as
links from other cities. Aliens can get trapped only in a city with inbound only links.

We assume that the program is executed with no more than `maxInvaders (default=2)`
aliens of the number of cities in the map.
The program supports any `maxInvader` value > 0.

### Requirements

* [Go](https://golang.org/dl/)
* [Git](https://git-scm.com/downloads)

### Test

```
$ make test
```

### Run

```
$ make build
$ ./bin/alien-invasion --in=<INPUT_FILE> --n=<NUMBER_OF_ALIENS>
```

### Termination

The program terminates with `exit status 0` if:

* All cities have been destroyed
* All aliens moved 10000 times

The program terminates with `exit status 1` if:

* A wrong input is provided or the input cannot be read (map file, alien count)
* All aliens are trapped

### Alternative approach and optimizations

If links would be considered bi-directional, the world map could be
implemented using a undirected port graph. Where the port of a link/edge
would be the cardinal direction.

For initial invasion and moving to neighboring cities, a "livable city" index
could be used, distinguished by current occupancy and outbound links.
