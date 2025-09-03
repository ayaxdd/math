package root

import (
	"fmt"
	"io"
)

type TCostSet struct {
	mCost int
	mSet  *TSet
	mFlag bool
}

func CreateTCSet(aCost int, aSet TSet) *TCostSet {
	return &TCostSet{
		mCost: aCost,
		mSet:  &aSet,
	}
}

func CreateEmptyTCSet() *TCostSet {
	return &TCostSet{
		mSet: CreateSet(),
	}
}

func (tc *TCostSet) Compare(item TItem) TCompare {
	a, ok := item.(*TCostSet)
	if !ok {
		return CmpIncomp
	}

	if tc == a {
		return CmpEq
	}

	// Сравнение при равной общей стоимости
	if tc.mCost == a.mCost {
		if tc.mSet.GetCount() > a.mSet.GetCount() {
			return CmpLess
		}
		if tc.mSet.GetCount() < a.mSet.GetCount() {
			return CmpGreat
		}
		return tc.mSet.Compare(a.mSet)
	}

	// Сравнение удельной стоимости
	tcUnit := tc.mCost * a.mSet.GetCount()
	aUnit := a.mCost * tc.mSet.GetCount()

	if tcUnit < aUnit {
		return CmpLess
	}
	if tcUnit > aUnit {
		return CmpGreat
	}

	// При равной удельной стоимости сравниваем по количеству
	if tc.mSet.GetCount() > a.mSet.GetCount() {
		return CmpLess
	}
	if tc.mSet.GetCount() < a.mSet.GetCount() {
		return CmpGreat
	}

	return CmpEq
}

func (tc *TCostSet) Copy() TItem {
	return CreateTCSet(tc.mCost, *tc.mSet)
}

func (tc *TCostSet) Print(w io.Writer) {
	t := tc.mSet.GetItem(0)

	_, ok := t.(*TCostSet)
	if ok {
		fmt.Fprintln(w, "Cost= ", tc.mCost)
	} else {
		fmt.Fprintf(w, " Cost=%4d ", tc.mCost)
	}

	tc.mSet.Print(w)
}

func (tc *TCostSet) Append(item *TCostSet) {
	if item == nil {
		return
	}

	for i := range item.mSet.GetCount() {
		cs := item.mSet.GetItem(i).(*TCostSet)
		tc.Insert(cs)
	}
}

func (tc *TCostSet) Insert(item *TCostSet) {
	if item == nil {
		return
	}

	tc.mCost += item.mCost
	tc.mSet.Insert(item)
}
