package world

// we use a directed graph to represent the world map.
// since each city/vertex has a max vertex degree of 4, we assume to have a
// sparse graph and thus use an adjacency list to reference neighboring vertices.

// adjacency list with cardinal directions indexed by neighboring city name
type link map[string]string

// World implements a generalized finite directed graph.
// where cities act as nodes and roads as edges.
// roads represent an adjacency list of a cities neighbors indexed by city ids.
// cities indexed by name
type Map map[string]link

// Connect adds a road from city a to city b in direction d.
// for convenience we store the inversed connection between all cities too.
// thus we ensure that all connected cities get created.
func (m Map) Link(origin, destination, direction string) {
	// outbound link a => b in direction d
	if _, ok := m[origin]; !ok {
		m[origin] = make(link)
	}
	m[origin][destination] = direction

	// we need to initialize the linked destination since it may have inbound links only.
	// if it has outbound links, it either exists already or will get overwritten
	// in a subsequent call.
	// inbound only cities still can get invaded and destroyed, so we need to pu them on the map.
	if _, ok := m[destination]; !ok {
		m[destination] = make(link)
	}
}

// destroyCity removes a city with given name from the world, as well as any
// roads attached to it. no-op if a city with the given name does not exist.
func (m Map) DestroyCity(cityName string) {
	// destroy inbound links
	for linked := range m {
		delete(m[linked], cityName)
	}
	// destroy city and outbound links
	delete(m, cityName)
}

func (m Map) Cities() []string {
	cityNames := make([]string, len(m))
	var i int
	for cityName := range m {
		cityNames[i] = cityName
		i++
	}
	return cityNames
}
