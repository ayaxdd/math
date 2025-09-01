package items

import (
	"fmt"
	"io"

	"github.com/ayaxdd/math/types/root"
)

type TItemNum struct {
	mData int
}

func CreateNum(n int) *TItemNum {
	return &TItemNum{
		mData: n,
	}
}

func (n *TItemNum) Compare(arg root.TItem) root.TCompare {
	a, ok := arg.(*TItemNum)
	if n == a {
		return root.CmpEq
	}

	if !ok {
		return root.CmpIncomp
	}

	d1, d2 := n.mData, a.mData
	if d1 < d2 {
		return root.CmpLess
	} else if d1 > d2 {
		return root.CmpGreat
	}

	return root.CmpEq
}

func (n *TItemNum) Copy() root.TItem {
	return CreateNum(n.mData)
}

func (n *TItemNum) Print(w io.Writer) {
	fmt.Fprint(w, n.mData)
}

func (n *TItemNum) GetData() int {
	return n.mData
}
