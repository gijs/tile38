package rtree

import (
	"math/rand"
	"runtime"
	"testing"
)

func randf(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func randMinMax() (min, max []float64) {
	minX, maxX := randf(-180, 180), randf(-180, 180)
	minY, maxY := randf(-90, 90), randf(-90, 90)
	minZ, maxZ := randf(0, 1000), randf(0, 1000)
	min4, max4 := randf(0, 1000), randf(0, 1000)
	if maxX < minX {
		minX, maxX = maxX, minX
	}
	if maxY < minY {
		minY, maxY = maxY, minY
	}
	if maxZ < minZ {
		minZ, maxZ = maxZ, minZ
	}
	if max4 < min4 {
		min4, max4 = max4, min4
	}
	return []float64{minX, minY, minZ, min4}, []float64{maxX, maxY, maxZ, max4}
}

func wp(min, max []float64) *Rect {
	return &Rect{
		MinX: min[0],
		MinY: min[1],
		MaxX: max[0],
		MaxY: max[1],
	}
}

func TestA(t *testing.T) {
	tr := New()
	item1 := wp([]float64{10, 10, 10, 10}, []float64{20, 20, 20, 20})
	item2 := wp([]float64{5, 5, 5, 5}, []float64{25, 25, 25, 25})
	tr.Insert(item1)
	tr.Insert(item2)
	var itemA Item
	tr.Search(21, 20, 25, 25, func(item Item) bool {
		itemA = item
		return true
	})
	if tr.Count() != 2 {
		t.Fatalf("tr.Count() == %d, expect 2", tr.Count())
	}
	if itemA != item2 {
		t.Fatalf("itemA == %v, expect %v", itemA, item2)
	}
}

func TestMemory(t *testing.T) {
	rand.Seed(0)
	tr := New()
	for i := 0; i < 100000; i++ {
		min, max := randMinMax()
		tr.Insert(wp(min, max))
	}
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	println(int(m.HeapAlloc)/tr.Count(), "bytes/rect")
}

func BenchmarkInsert(b *testing.B) {
	rand.Seed(0)
	tr := New()
	for i := 0; i < b.N; i++ {
		min, max := randMinMax()
		tr.Insert(wp(min, max))
	}
	// count := 0
	// tr.Search([]float64{-116, 32, 20}, []float64{-114, 34, 800}, func(id int) bool {
	// 	count++
	// 	return true
	// })
	// println(count)
	// //println(tr.Count())
}
