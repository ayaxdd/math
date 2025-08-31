package items

import (
	"fmt"
	"io"

	"github.com/ayaxdd/math/types/root"
)

type TItemStr struct {
	mData string
}

func CreateStr(s string) *TItemStr {
	return &TItemStr{
		mData: s,
	}
}

func (s *TItemStr) Compare(arg root.TItem) root.TCompare {
	a, ok := arg.(*TItemStr)
	if s == a {
		return root.CmpEq
	}

	if !ok {
		return root.CmpIncomp
	}

	d1, d2 := s.mData, a.mData
	if d1 < d2 {
		return root.CmpLess
	} else if d1 > d2 {
		return root.CmpGreat
	}

	return root.CmpEq
}

func (s *TItemStr) Copy() root.TItem {
	return CreateStr(s.mData)
}

func (s *TItemStr) Print(w io.Writer) {
	fmt.Fprint(w, s.mData)
}

func (n *TItemStr) GetData() string {
	return n.mData
}
