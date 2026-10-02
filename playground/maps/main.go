package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := map[string]Vertex{
		"Tokyo":    {35.6850, 139.7514},
		"Shanghai": {31.2325, 121.4692},
	}
	m["Mumbai"] = Vertex{19.0758, 72.8775}

	fmt.Println("Map:  ", m)

	fmt.Println("Delete Mumbai")
	delete(m, "Mumbai")
	fmt.Println("Map", m)

	v, ok := m["Mumbai"]
	if ok != true {
		fmt.Println("Mumbai - not found")
	}
	fmt.Println(m)

	v, ok = m["Tokyo"]
	if ok != true {
		fmt.Println("Tokyo - not found")
	}
	fmt.Println("Get Tokyo", v)
}
