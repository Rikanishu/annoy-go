package main

import (
	annoy "github.com/Rikanishu/annoy-go"
	"github.com/Rikanishu/annoy-go/utils"
	"math/rand"
	"os"
)

func main() {
	rand.Seed(42)
	f := 100
	t := annoy.NewAnnoyIndexAngular(f)
	for i := 0; i < 1000; i++ {
		item := make([]float32, 0, f)
		for x := 0; x < f; x++ {
			item = append(item, rand.Float32())
		}
		item[i%f] = 1
		t.AddItem(i, item)
	}
	t.Build(10)
	// we can save ANN index as a file and reload it
	t.Save("test.ann")
	annoy.DeleteAnnoyIndexAngular(t)

	t = annoy.NewAnnoyIndexAngular(f)
	t.Load("test.ann")
	defer func() {
		t.Unload()
		_ = os.Remove("test.ann")
	}()

	result := annoy.NewAnnoyVectorInt()
	defer result.Free()
	for i := 0; i < 100000; i++ {
		result = annoy.NewAnnoyVectorInt()
		searchVector := make([]float32, 100)
		for j := 0; j < len(searchVector); j++ {
			searchVector[j] = rand.Float32()
		}
		t.GetNnsByVector(searchVector, 1000, -1, result)
		result.Free()

		if i%10000 == 0 {
			utils.DebugPrintStats()
		}
	}
}
