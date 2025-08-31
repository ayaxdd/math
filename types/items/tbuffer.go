package items

import (
	"io"

	"github.com/ayaxdd/math/types/root"
)

type TBuffer struct {
	que []root.TItem
}

func CreateBuffer() *TBuffer {
	return &TBuffer{
		que: make([]root.TItem, 0),
	}
}

func (b *TBuffer) Compare(arg root.TItem) root.TCompare {
	a, ok := arg.(*TBuffer)

	if !ok {
		return root.CmpIncomp
	}

	if b == a {
		return root.CmpEq
	}

	bLen := len(b.que)
	aLen := len(a.que)
	if bLen > aLen {
		return root.CmpGreat
	} else if bLen < aLen {
		return root.CmpLess
	}

	for i := range bLen {
		res := b.que[i].Compare(a.que[i])
		if res != root.CmpEq {
			return res
		}
	}

	return root.CmpEq
}

func (b *TBuffer) Copy() root.TItem {
	cp := CreateBuffer()
	for range b.que {
		cp.Put(b.Get().Copy())
	}

	return cp
}

func (b *TBuffer) Print(w io.Writer) {
	panic("print not implemented")
}

func (b *TBuffer) Put(item root.TItem) {
	b.que = append([]root.TItem{item}, b.que...)
}

func (b *TBuffer) Get() root.TItem {
	if len(b.que) == 0 {
		return nil
	}

	item := b.que[0]
	b.que = b.que[1:]

	return item
}

func (b *TBuffer) Push(item root.TItem) {
	b.que = append(b.que, item)
}

func (b *TBuffer) Pop() root.TItem {
	if len(b.que) == 0 {
		return nil
	}

	n := len(b.que) - 1
	item := b.que[n]
	b.que = b.que[:n]

	return item
}

func (b *TBuffer) Top() root.TItem {
	if len(b.que) == 0 {
		return nil
	}

	return b.que[len(b.que)-1]
}

func (b *TBuffer) Reversion() {
	panic("reversion not implemented")
}

func (b *TBuffer) GetByIndex(i int) root.TItem {
	if len(b.que) == 0 || i < 0 || i > len(b.que)-1 {
		return nil
	}

	return b.que[i]
}

func (b *TBuffer) GetCount() int {
	return len(b.que)
}

func (b *TBuffer) IsPresent(item root.TItem) bool {
	for _, v := range b.que {
		if v.Compare(item) == root.CmpEq {
			return true
		}
	}

	return false
}
