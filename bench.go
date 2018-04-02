package main

import (
	"encoding/json"
)

var areas = []struct {
	name string
	area []float64
}{
	{
		"Толе Би - Муканова - Абая - Сейфуллина",
		[]float64{43.252865, 76.914907, 43.241822, 76.934134},
	},
	{
		"Розыбакиева - Жамбыла - Жарокова - Абая",
		[]float64{43.244174, 76.888961, 43.239499, 76.899476},
	},
	{
		"Жарокова - Шевченко - Манаса - Абая",
		[]float64{43.243043, 76.898871, 43.240047, 76.908569},
	},
}

type Point struct {
	Lat, Lng float64
}

func RawPointInBoundingBox(raw []byte) (string, bool) {
	var p Point
	json.Unmarshal(raw, &p)
	
	for _, a := range areas {
		if p.Lat < a.area[0] && p.Lat > a.area[2] && p.Lng < a.area[1] && p.Lng > a.area[3] {
			return a.name, true
		}
	}
	
	return "район не найден", false
}
