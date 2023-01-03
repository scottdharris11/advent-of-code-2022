package utils

import "math"

type WrappedList struct {
	count int
	first *WrappedListItem
	last  *WrappedListItem
}

func (w *WrappedList) Length() int {
	return w.count
}

func (w *WrappedList) Add(value interface{}) *WrappedListItem {
	item := &WrappedListItem{Value: value}
	if w.first == nil {
		w.first = item
		w.last = item
	}
	w.last.next = item
	w.first.prev = item
	item.prev = w.last
	item.next = w.first
	w.last = item
	w.count++
	return item
}

func (w *WrappedList) Move(toMove *WrappedListItem, places int) {
	aPlaces := w.adjustedIndex(int(math.Abs(float64(places))))
	if aPlaces == 0 {
		return
	}

	prev := places < 0
	if prev {
		aPlaces++
	}

	item := toMove
	for i := 0; i < aPlaces; i++ {
		if prev {
			item = item.prev
			continue
		}
		item = item.next
	}

	if w.first == toMove {
		w.first = toMove.next
	}
	if w.last == toMove {
		w.last = toMove.prev
	}

	toMove.prev.next = toMove.next
	toMove.next.prev = toMove.prev
	toMove.next = item.next
	toMove.prev = item
	item.next.prev = toMove
	item.next = toMove

	if !prev && w.first == toMove.next {
		w.first = toMove
		w.last = toMove.prev
	}
}

func (w *WrappedList) Items() []interface{} {
	if w.count == 0 {
		return nil
	}
	items := make([]interface{}, w.count)
	item := w.first
	for i := 0; i < w.count; i++ {
		items[i] = item.Value
		item = item.next
	}
	return items
}

func (w *WrappedList) ItemAtIndex(index int) interface{} {
	if w.count == 0 {
		return nil
	}

	prev := index < 0
	aIndex := w.adjustedIndex(int(math.Abs(float64(index))))

	item := w.first
	for i := 0; i < aIndex; i++ {
		if prev {
			item = item.prev
			continue
		}
		item = item.next
	}
	return item.Value
}

func (w *WrappedList) adjustedIndex(index int) int {
	if index < w.count {
		return index
	}
	return index % w.count
}

type WrappedListItem struct {
	Value interface{}
	prev  *WrappedListItem
	next  *WrappedListItem
}
