package world

// We use a directed graph to represent the world map. Since each city/vertex
// has a max vertex degree of 4, we assume to have a sparse graph and thus,
// use an adjacency list to reference neighboring cities/vertices.

// Adjacency list with cardinal directions
// indexed by neighboring city name.
type link map[string]string

// Map implements a generalized finite directed graph,
// where cities act as nodes and roads as edges.
type Map map[string]link

// Link adds a link from `origin` to `destination` in direction `direction`.
func (m Map) Link(origin, destination, direction string) {
	// outbound link
	if _, ok := m[origin]; !ok {
		m[origin] = make(link)
	}
	m[origin][destination] = direction

	// we need to initialize the linked destination since it may have inbound
	// links only. if it has outbound links, it either exists already or will
	// get overwritten in a subsequent call.
	// inbound only cities still can get invaded and destroyed, so we need to
	// put them on the map.
	if _, ok := m[destination]; !ok {
		m[destination] = make(link)
	}
}

// DestroyCity removes a city with given name from the map, as well as any
// roads attached to it. no-op if a city with the given name does not exist.
func (m Map) DestroyCity(cityName string) {
	// destroy inbound links
	for linked := range m {
		delete(m[linked], cityName)
	}
	// destroy city and outbound links
	delete(m, cityName)
}

// Cities returns a list of all cities on the map.
func (m Map) Cities() []string {
	cityNames := make([]string, len(m))
	var i int
	for cityName := range m {
		cityNames[i] = cityName
		i++
	}
	return cityNames
}
