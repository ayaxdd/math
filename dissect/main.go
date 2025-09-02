package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/ayaxdd/math/types/items"
	"github.com/ayaxdd/math/types/root"
)

func main() {
	test := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		// "k",
		// "l",
		// "m",
		// "n",
		// "o",
		// "p",
		// "q",
		// "r",
		// "s",
		// "t",
	}
	set := root.CreateSet()
	for _, s := range test {
		item := items.CreateStr(s)
		set.Insert(item)
	}
	buf := GenSubsRand(set, 10, 0, 80)

	set.Print(os.Stdin)
	fmt.Fprintln(os.Stdin)
	buf.Print(os.Stdin)
}

// func GenInput() {
// 	in := "a b c d e f g h i j k l m n o p q r s t"
// 	for _, r := range in {
// 		if r != ' ' {
// 			fmt.Printf(`"%s", `, string(r))
// 		}
// 	}
// }

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

func GenRandSet(aMin, aMax int, aSet *root.TSet) *root.TSet {
	res := root.CreateSet()

	for aSet.GetCount() > 0 {
		n := aMin + rand.Intn(aMax-aMin)
		s := root.CreateSet()

		for range n {
			k := rand.Intn(aSet.GetCount())
			t := aSet.GetItem(k)
			s.Insert(t)
			aSet.Delete(t)

			if aSet.GetCount() == 0 {
				break
			}
		}

		if s.GetCount() > 0 {
			res.Insert(s)
		}
	}

	return res
}

func GenSubsRand(aSet *root.TSet, aCnt, aMin, aMax int) *items.TBuffer {
	nMin := (aSet.GetCount() * aMin) / 100
	if nMin < 1 {
		nMin = 1
	}
	if nMin >= aSet.GetCount() {
		nMin = aSet.GetCount() - 1
	}

	nMax := (aSet.GetCount() * aMax) / 100
	if nMax < nMin {
		nMax = nMin
	}
	if nMax > aSet.GetCount() {
		nMax = aSet.GetCount()
	}

	res := items.CreateBuffer()

	for aCnt > 0 {
		cp := aSet.Copy().(*root.TSet)
		ss := GenRandSet(nMin, nMax, cp)

		for i := range ss.GetCount() {
			subset := ss.GetItem(i).(*root.TSet)

			if rand.Intn(3) != 0 {
				res.Put(subset.Copy().(*root.TSet))
				aCnt--
			}

			if aCnt == 0 {
				break
			}
		}
	}

	return res
}
