package items

import (
	"fmt"
	"io"

	"github.com/ayaxdd/math/types/root"
)

type TItemChar struct {
	mData rune
}

func CreateChar(c rune) *TItemChar {
	return &TItemChar{
		mData: c,
	}
}

func (c *TItemChar) Compare(arg root.TItem) root.TCompare {
	a, ok := arg.(*TItemChar)
	if c == a {
		return root.CmpEq
	}

	if !ok {
		return root.CmpIncomp
	}

	d1, d2 := c.mData, a.mData
	if d1 < d2 {
		return root.CmpLess
	} else if d1 > d2 {
		return root.CmpGreat
	}

	return root.CmpEq
}

func (c *TItemChar) Copy() root.TItem {
	return CreateChar(c.mData)
}

func (c *TItemChar) Print(w io.Writer) {
	fmt.Fprint(w, c.mData)
}

func (c *TItemChar) GetData() rune {
	return c.mData
}
