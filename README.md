
**annoy-go** is a version of [annoy](https://github.com/spotify/annoy/) that was generated for Golang according [this instruction](https://github.com/spotify/annoy/blob/master/README_GO.rst).

This is a forked version with fixed memory leaks.

* Please note, it changes the interface, new wrappers for results vectors were added: `AnnoyVectorInt` / `AnnoyVectorFloat`.
* Always call `Free()` method in vector wrappers when you don't need them anymore. If you don't do that, you'll have memory leaks. I warned you.
* You can make a copy of vector content to slice through methods `Copy(inputSlice)` or `ToSlice()`. Use the first to copy the values to already existed slice, or use the second method to create a copy to a new allocated one.
* Do not reuse them in different threads since it's not thread safe.
* Input slices which are not used to passing the result were left the same.

---

* Also note that indexes will be kept as int32, so keep in mind there is a count limit for items.

__Go code example__

```go
package main

import (
    "github.com/Rikanishu/annoy-go"
    "fmt"
    "math/rand"
)

func main() {
     f := 40
     t := annoyindex.NewAnnoyIndexAngular(f)
     for i := 0; i < 1000; i++ {
       item := make([]float32, 0, f)
       for x:= 0; x < f; x++ {
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
```

