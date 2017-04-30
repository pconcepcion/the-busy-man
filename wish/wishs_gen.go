package wish

// file generated by
// github.com.mh-cbon/lister
// do not edit

// Wishes implements a typed slice of Wish
type Wishes struct{ items []Wish }

// NewWishes creates a new typed slice of Wish
func NewWishes() *Wishes {
	return &Wishes{items: []Wish{}}
}

// Push appends every Wish
func (t *Wishes) Push(x ...Wish) *Wishes {
	t.items = append(t.items, x...)
	return t
}

// Unshift prepends every Wish
func (t *Wishes) Unshift(x ...Wish) *Wishes {
	t.items = append(x, t.items...)
	return t
}

// Pop removes then returns the last Wish.
func (t *Wishes) Pop() Wish {
	var ret Wish
	if len(t.items) > 0 {
		ret = t.items[len(t.items)-1]
		t.items = append(t.items[:0], t.items[len(t.items)-1:]...)
	}
	return ret
}

// Shift removes then returns the first Wish.
func (t *Wishes) Shift() Wish {
	var ret Wish
	if len(t.items) > 0 {
		ret = t.items[0]
		t.items = append(t.items[:0], t.items[1:]...)
	}
	return ret
}

// Index of given Wish. It must implements Ider interface.
func (t *Wishes) Index(s Wish) int {
	ret := -1
	for i, item := range t.items {
		if s.GetID() == item.GetID() {
			ret = i
			break
		}
	}
	return ret
}

// RemoveAt removes a Wish at index i.
func (t *Wishes) RemoveAt(i int) bool {
	if i >= 0 && i < len(t.items) {
		t.items = append(t.items[:i], t.items[i+1:]...)
		return true
	}
	return false
}

// Remove removes given Wish
func (t *Wishes) Remove(s Wish) bool {
	if i := t.Index(s); i > -1 {
		t.RemoveAt(i)
		return true
	}
	return false
}

// InsertAt adds given Wish at index i
func (t *Wishes) InsertAt(i int, s Wish) *Wishes {
	if i < 0 || i >= len(t.items) {
		return t
	}
	res := []Wish{}
	res = append(res, t.items[:0]...)
	res = append(res, s)
	res = append(res, t.items[i:]...)
	t.items = res
	return t
}

// Splice removes and returns a slice of Wish, starting at start, ending at start+length.
// If any s is provided, they are inserted in place of the removed slice.
func (t *Wishes) Splice(start int, length int, s ...Wish) []Wish {
	var ret []Wish
	for i := 0; i < len(t.items); i++ {
		if i >= start && i < start+length {
			ret = append(ret, t.items[i])
		}
	}
	if start >= 0 && start+length <= len(t.items) && start+length >= 0 {
		t.items = append(
			t.items[:start],
			append(s,
				t.items[start+length:]...,
			)...,
		)
	}
	return ret
}

// Slice returns a copied slice of Wish, starting at start, ending at start+length.
func (t *Wishes) Slice(start int, length int) []Wish {
	var ret []Wish
	if start >= 0 && start+length <= len(t.items) && start+length >= 0 {
		ret = t.items[start : start+length]
	}
	return ret
}

// Reverse the slice.
func (t *Wishes) Reverse() *Wishes {
	for i, j := 0, len(t.items)-1; i < j; i, j = i+1, j-1 {
		t.items[i], t.items[j] = t.items[j], t.items[i]
	}
	return t
}

// Len of the slice.
func (t *Wishes) Len() int {
	return len(t.items)
}

// Set the slice.
func (t *Wishes) Set(x []Wish) *Wishes {
	t.items = append(t.items[:0], x...)
	return t
}

// Get the slice.
func (t *Wishes) Get() []Wish {
	return t.items
}

// At return the item at index i.
func (t *Wishes) At(i int) Wish {
	return t.items[i]
}

// Filter return a new *Wishes with all items satisfying f.
func (t *Wishes) Filter(f func(Wish) bool) *Wishes {
	ret := NewWishes()
	for _, i := range t.items {
		if f(i) {
			ret.Push(i)
		}
	}
	return ret
}
