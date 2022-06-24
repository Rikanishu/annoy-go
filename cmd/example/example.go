package main

import (
	"fmt"
	"math/rand"

	annoyindex "github.com/Rikanishu/annoy-go"
)

func main() {
	f := 40
	t := annoyindex.NewAnnoyIndexAngular(f)
	for i := 0; i < 1000; i++ {
		item := make([]float32, 0, f)
		for x := 0; x < f; x++ {
			item = append(item, rand.Float32())
		}
		t.AddItem(i, item)
	}
	t.Build(10)
	t.Save("test.ann")

	annoyindex.DeleteAnnoyIndexAngular(t)

	t = annoyindex.NewAnnoyIndexAngular(f)
	t.Load("test.ann")

	result := annoyindex.NewAnnoyVectorInt()
	defer result.Free()
	t.GetNnsByItem(0, 1000, -1, result)
	fmt.Printf("%v\n", result.ToSlice())
}
