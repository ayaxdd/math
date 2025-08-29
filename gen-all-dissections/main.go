package main

import (
	"fmt"

	"github.com/ayaxdd/math/gen-all-dissections/types"
)

const (
	msg = "%d: %v\n"
)

func main() {
	test := []string{"a", "b", "c", "e", "f", "g", "h"}
	all := GenAllDissections(test)
	for i, subsets := range all {
		fmt.Printf(msg, i+1, subsets)
	}
}

func GenAllDissections(aSet []string) []*types.NSet {
	que := types.CreateBuffer()
	ssPut := types.CreateNSet()
	que.Put(ssPut)

	for _, item := range aSet {
		for range que.Size() {
			ssGet := que.Get().(*types.NSet)
			localHandle(item, ssGet, que)
		}
	}

	res := make([]*types.NSet, 0, que.Size())
	for que.Size() > 0 {
		res = append(res, que.Get().(*types.NSet))
	}

	return res
}

func localHandle(item string, ssGet *types.NSet, que *types.Buffer) {
	// add item to all existing substets
	for i := range ssGet.Size() {
		ssPut := ssGet.Copy()
		s := ssPut.GetItem(i)
		ssPut.Delete(i)
		s.Insert(item)
		ssPut.Insert(s)
		que.Put(ssPut)
	}

	// create and add new {item} subset
	ssPut := ssGet.Copy()
	s := types.CreateSet()
	s.Insert(item)
	ssPut.Insert(s)
	que.Put(ssPut)
}
