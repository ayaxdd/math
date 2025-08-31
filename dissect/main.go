package main

import (
	"fmt"
	"os"

	"github.com/ayaxdd/math/types/items"
	"github.com/ayaxdd/math/types/root"
)

func main() {
	test := []string{"a", "b", "c"} // , "e", "f", "g", "h"
	set := root.CreateSet()
	for _, s := range test {
		item := items.CreateStr(s)
		set.Insert(item)
	}

	all := GenAllDissections(set)
	all.Print(os.Stdin)
	fmt.Println()
	fmt.Println(all.GetCount())
}

func GenAllDissections(aSet *root.TSet) *root.TSet {
	que := items.CreateBuffer()
	ssPut := root.CreateSet()
	que.Put(ssPut)

	for i := range aSet.GetCount() {
		item := aSet.GetItem(i)
		cpQue := que.Copy().(*items.TBuffer)
		for range cpQue.GetCount() {
			ssGet := cpQue.Get().(*root.TSet)
			localHandle(item, ssGet, que)
		}
	}

	res := root.CreateSet()
	for que.GetCount() > 0 {
		res.Insert(que.Get())
	}

	return res
}

func localHandle(item root.TItem, ssGet *root.TSet, que *items.TBuffer) {
	// add item to all existing substets
	for i := range ssGet.GetCount() {
		ssPut := ssGet.Copy().(*root.TSet)
		s := ssPut.GetItem(i).(*root.TSet)
		ssPut.Delete(s)
		s.Insert(item)
		ssPut.Insert(s)
		que.Put(ssPut)
	}

	// create and add new {item} subset
	ssPut := ssGet.Copy().(*root.TSet)
	s := root.CreateSet()
	s.Insert(item)
	ssPut.Insert(s)
	que.Put(ssPut)
}
