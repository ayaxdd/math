package root

import (
	"fmt"
	"io"
	"strings"
)

const (
	openBracket  = "{"
	closeBracket = "}"
	tab          = "  "
)

type TSet struct {
	items []TItem
}

func CreateSet() *TSet {
	return &TSet{
		items: make([]TItem, 0),
	}
}

func (s *TSet) Compare(item TItem) TCompare {
	a, ok := item.(*TSet)

	if !ok {
		return CmpIncomp
	}

	if s == a {
		return CmpEq
	}

	sLen := len(s.items)
	aLen := len(a.items)
	if sLen > aLen {
		return CmpGreat
	} else if sLen < aLen {
		return CmpLess
	}

	for _, v := range s.items {
		if a.Exist(v) == -1 {
			return CmpIncomp
		}
	}

	return CmpEq
}

func (s *TSet) Copy() TItem {
	cp := CreateSet()
	for _, item := range s.items {
		cp.Insert(item.Copy())
	}

	return cp
}

var printLevel = 0

func (s *TSet) Print(w io.Writer) {
	if s == nil || len(s.items) == 0 {
		fmt.Fprint(w, "{}")
		return
	}

	currTab := strings.Repeat(tab, printLevel)
	fmt.Fprint(w, currTab+openBracket)

	printLevel++
	flag := false

	for _, item := range s.items {
		_, isSet := item.(*TSet)
		flag = isSet
		if flag {
			fmt.Fprintln(w)
		}
		item.Print(w)

	}

	if flag {
		fmt.Fprintf(w, currTab+closeBracket+": %d\n", len(s.items))
	} else {
		fmt.Fprintf(w, closeBracket+": %d\n", len(s.items))
	}
	printLevel--
}

// func (s *TSet) Add(arg *TSet) {
// }
//
// func (s *TSet) Sub(arg *TSet) {
// }
//
// func (s *TSet) Mul(arg *TSet) {
// }
//
// func (s *TSet) ExoR(arg *TSet) {
// }
//
//	func (s *TSet) TestIntersect(arg *TSet) bool {
//		return false
//	}
//
//	func (s *TSet) CompareSet(arg *TSet) TCompare {
//		return CmpIncomp
//	}
func (s *TSet) Exist(item TItem) int {
	for i, v := range s.items {
		if v.Compare(item) == CmpEq {
			return i
		}
	}

	return -1
}

func (s *TSet) Insert(item TItem) {
	s.items = append(s.items, item)
}

func (s *TSet) Delete(item TItem) {
	i := s.Exist(item)
	if i != -1 {
		s.items = append(s.items[:i], s.items[i+1:]...)
	}
}

func (s *TSet) CopyItems(item *TSet) {
	cp := make([]TItem, len(item.items))
	copy(cp, item.items)
	s.items = cp
}

func (s *TSet) GetItem(i int) TItem {
	if len(s.items) == 0 || i < 0 || i > len(s.items)-1 {
		return nil
	}

	return s.items[i]
}

func (s *TSet) GetObject(item TItem) TItem {
	i := s.Exist(item)
	if i != -1 {
		return s.items[i]
	}

	return nil
}

func (s *TSet) CoverToDissect() {
}

func (s *TSet) GetCount() int {
	return len(s.items)
}
