package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// mapはKeyとvalueを関連付けする
var m map[string]Vertex

var k = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
var n = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

	fmt.Println(k)

	fmt.Println(n)
}
