package root

import (
	"io"
)

type TCompare int

const (
	CmpEq TCompare = iota
	CmpLess
	CmpGreat
	CmpIncomp
)

type TItem interface {
	Compare(TItem) TCompare
	Copy() TItem
	Print(io.Writer)
}
