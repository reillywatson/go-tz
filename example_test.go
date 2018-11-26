package gotz_test

import (
	"fmt"

	"github.com/ugjka/go-tz"
)

func ExampleGetZone() {
	// Loading Zone for Line Islands, Kiritimati
	p := gotz.Point{
		Lon: -157.21328, Lat: 1.74294,
	}
	zone, err := gotz.GetZone(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(zone[0])
	// Output: Pacific/Kiritimati
}
