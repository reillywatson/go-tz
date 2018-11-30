package gotz

import (
	"reflect"
	"testing"
)

type result struct {
	zones []string
	err   error
}

var tt = []struct {
	name  string
	point Point
	result
}{
	{
		"Riga",
		Point{24.105078, 56.946285},
		result{
			zones: []string{"Europe/Riga"},
			err:   nil,
		},
	},
	{
		"Tokyo",
		Point{139.7594549, 35.6828387},
		result{
			zones: []string{"Asia/Tokyo"},
			err:   nil,
		},
	},
	{
		"Urumqi/Shanghai",
		Point{87.319461, 43.419754},
		result{
			zones: []string{"Asia/Shanghai", "Asia/Urumqi"},
			err:   nil,
		},
	},
	{
		"Tuvalu",
		Point{178.1167698, -7.768959},
		result{
			zones: []string{"Pacific/Funafuti"},
			err:   nil,
		},
	},
	{
		"Baker Island",
		Point{-176.474331436, 0.190165906},
		result{
			zones: nil,
			err:   ErrNoZoneFound,
		},
	},
}

func TestGetZone(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tzid, err := GetZone(tc.point)
			if err != tc.err {
				t.Errorf("expected err %v; got %v", tc.err, err)
			}
			if !reflect.DeepEqual(tzid, tc.zones) {
				t.Errorf("expected zones %v; got %v", tc.zones, tzid)
			}
		})
	}
}

func BenchmarkZones(b *testing.B) {
	b.Run("polygon centers", func(b *testing.B) {
	Loop:
		for n := 0; n < b.N; {
			for _, v := range centerCache {
				for i := range v {
					if n > b.N {
						break Loop
					}
					_, err :=GetZone(v[i])
					if err != nil {
						b.Errorf("point %v did not return a zone", v[i])
					}
					n++
				}
			}
		}
	})
	b.Run("test cases", func(b *testing.B) {
	Loop:
		for n := 0; n < b.N; {
			for _, tc := range tt {
				if n > b.N {
					break Loop
				}
				GetZone(tc.point)
				n++
			}

		}
	})
}
