package types

import "strings"

type Set struct {
	items map[string]bool
}

func CreateSet() *Set {
	return &Set{
		items: make(map[string]bool, 0),
	}
}

func (s *Set) String() string {
	items := make([]string, 0, s.Size()+2)
	items = append(items, "{")
	items = append(items, s.GetItems()...)
	items = append(items, "}")
	return strings.Join(items, "")
}

func (s *Set) Insert(item string) {
	s.items[item] = true
}

func (s *Set) Delete(item string) {
	delete(s.items, item)
}

func (s *Set) GetItems() []string {
	items := make([]string, len(s.items))
	for k := range s.items {
		items = append(items, k)
	}
	return items
}

func (s *Set) Size() int {
	return len(s.items)
}

func (s *Set) Contains(item string) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set) Copy() *Set {
	cp := CreateSet()
	for k := range s.items {
		cp.Insert(k)
	}
	return cp
}

type NSet struct {
	subsets []*Set
}

func CreateNSet() *NSet {
	return &NSet{
		subsets: make([]*Set, 0, 1),
	}
}

func (ns *NSet) String() string {
	items := make([]string, 0, len(ns.subsets))
	items = append(items, "{")
	for i, set := range ns.subsets {
		items = append(items, set.String())
		if i != ns.Size()-1 {
			items = append(items, ", ")
		}
	}
	items = append(items, "}")
	return strings.Join(items, "")
}

func (ns *NSet) Insert(item *Set) {
	ns.subsets = append(ns.subsets, item)
}

func (ns *NSet) Delete(i int) {
	ns.subsets = append(ns.subsets[:i], ns.subsets[i+1:]...)
}

func (ns *NSet) GetItem(i int) *Set {
	return ns.subsets[i]
}

func (ns *NSet) Size() int {
	return len(ns.subsets)
}

func (ns *NSet) Copy() *NSet {
	cp := CreateNSet()
	for _, s := range ns.subsets {
		cp.Insert(s.Copy())
	}
	return cp
}

type Buffer struct {
	que []any
}

func CreateBuffer() *Buffer {
	return &Buffer{
		que: make([]any, 0, 1),
	}
}

func (b *Buffer) Put(item any) {
	b.que = append(b.que, item)
}

func (b *Buffer) Get() any {
	if len(b.que) == 0 {
		return nil
	}

	item := b.que[0]
	b.que = b.que[1:]
	return item
}

func (b *Buffer) Size() int {
	return len(b.que)
}
