package gotz

import (
	"fmt"
	"testing"
	"time"
)

func TestGetZone(t *testing.T) {
	//Test Riga
	p := Point{24.105078, 56.946285}
	start := time.Now()
	zone, err := GetZone(p)
	if err != nil {
		t.Error("Could not find Europe/Riga")
	}
	if len(zone) != 0 && zone[0] != "Europe/Riga" {
		t.Error("Zone not Europe/Riga but", zone[0])
	}
	fmt.Println(zone, time.Since(start))
	//Test Tokyo
	p = Point{139.7594549, 35.6828387}
	start = time.Now()
	zone, err = GetZone(p)
	if err != nil {
		t.Error("Could not find Asia/Tokyo")
	}
	if len(zone) != 0 && zone[0] != "Asia/Tokyo" {
		t.Error("Zone not Asia/Tokyo but", zone[0])
	}
	fmt.Println(zone, time.Since(start))
	//Test Urumqi
	p = Point{87.319461, 43.419754}
	start = time.Now()
	zone, err = GetZone(p)
	if len(zone) < 2 {
		t.Error("didn't detect overlapping zone in Urumqi")
	}
	if err != nil {
		t.Error("Could not find Asia/Shanghai Asia/Urumqi")
	}
	if len(zone) > 0 && zone[0] != "Asia/Shanghai" {
		t.Error("Zone not Asia/Shanghai but", zone[0])
	}
	if len(zone) > 1 && zone[1] != "Asia/Urumqi" {
		t.Error("Zone not Asia/Urumqi but", zone[0])
	}
	fmt.Println(zone, time.Since(start))
	//Tuvalu testing center cache
	p = Point{178.1167698, -7.768959}
	start = time.Now()
	zone, err = GetZone(p)
	if err != nil {
		t.Error("Could not find Pacific/Funafuti")
	}
	if len(zone) != 0 && zone[0] != "Pacific/Funafuti" {
		t.Error("Zone not Pacific/Funafuti but", zone[0])
	}
	fmt.Println(zone, time.Since(start))
	//Baker Island AoE. Should error out
	p = Point{-176.474331436, 0.190165906}
	start = time.Now()
	_, err = GetZone(p)
	if err == nil {
		t.Error("Baker island didn't error")
	}
	fmt.Println("Not found", time.Since(start))
}

func BenchmarkZones(b *testing.B) {
Loop:
	for n := 0; n < b.N; {
		for _, v := range centerCache {
			for i := range v {
				if n > b.N {
					break Loop
				}
				GetZone(v[i])
				n++
			}
		}
	}
}
