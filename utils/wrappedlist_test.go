package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrappedList(t *testing.T) {
	list := WrappedList{}

	item1 := list.Add(1)
	assert.Equal(t, 1, item1.Value)
	assert.Equal(t, item1, item1.next)
	assert.Equal(t, item1, item1.prev)
	assert.Equal(t, []interface{}{1}, list.Items())

	item2 := list.Add(2)
	assert.Equal(t, 2, item2.Value)
	assert.Equal(t, item1, item2.next)
	assert.Equal(t, item1, item2.prev)
	assert.Equal(t, item2, item1.next)
	assert.Equal(t, item2, item1.prev)
	assert.Equal(t, []interface{}{1, 2}, list.Items())

	item3 := list.Add(3)
	assert.Equal(t, 3, item3.Value)
	assert.Equal(t, item1, item3.next)
	assert.Equal(t, item2, item3.prev)
	assert.Equal(t, item3, item2.next)
	assert.Equal(t, item1, item2.prev)
	assert.Equal(t, item2, item1.next)
	assert.Equal(t, item3, item1.prev)
	assert.Equal(t, []interface{}{1, 2, 3}, list.Items())

	item4 := list.Add(4)
	assert.Equal(t, 4, item4.Value)
	assert.Equal(t, item1, item4.next)
	assert.Equal(t, item3, item4.prev)
	assert.Equal(t, item4, item3.next)
	assert.Equal(t, item2, item3.prev)
	assert.Equal(t, item3, item2.next)
	assert.Equal(t, item1, item2.prev)
	assert.Equal(t, item2, item1.next)
	assert.Equal(t, item4, item1.prev)
	assert.Equal(t, []interface{}{1, 2, 3, 4}, list.Items())
}

func TestWrappedList_Move(t *testing.T) {
	list := WrappedList{}
	item1 := list.Add(1)
	item2 := list.Add(2)
	item3 := list.Add(-3)
	item4 := list.Add(3)
	item5 := list.Add(-2)
	item6 := list.Add(0)
	item7 := list.Add(4)

	assert.Equal(t, []interface{}{1, 2, -3, 3, -2, 0, 4}, list.Items())

	list.Move(item1, 1)
	assert.Equal(t, []interface{}{2, 1, -3, 3, -2, 0, 4}, list.Items())

	list.Move(item2, 2)
	assert.Equal(t, []interface{}{1, -3, 2, 3, -2, 0, 4}, list.Items())

	list.Move(item3, -3)
	assert.Equal(t, []interface{}{1, 2, 3, -2, -3, 0, 4}, list.Items())

	list.Move(item4, 3)
	assert.Equal(t, []interface{}{1, 2, -2, -3, 0, 3, 4}, list.Items())

	list.Move(item5, -2)
	assert.Equal(t, []interface{}{1, 2, -3, 0, 3, 4, -2}, list.Items())

	list.Move(item6, 0)
	assert.Equal(t, []interface{}{1, 2, -3, 0, 3, 4, -2}, list.Items())

	list.Move(item7, 4)
	assert.Equal(t, []interface{}{1, 2, -3, 4, 0, 3, -2}, list.Items())

	list.Move(item5, -1)
	assert.Equal(t, []interface{}{1, 2, -3, 4, 0, -2, 3}, list.Items())

	list.Add(9)
	assert.Equal(t, []interface{}{1, 2, -3, 4, 0, -2, 3, 9}, list.Items())

	list.Move(item4, 1)
	assert.Equal(t, []interface{}{3, 1, 2, -3, 4, 0, -2, 9}, list.Items())
}

func TestWrappedList_Move_OffsetsAboveLength(t *testing.T) {
	list := WrappedList{}
	item1 := list.Add(1)
	item2 := list.Add(2)
	item3 := list.Add(-3)
	item4 := list.Add(3)
	item5 := list.Add(-2)
	item6 := list.Add(0)
	item7 := list.Add(4)

	assert.Equal(t, []interface{}{1, 2, -3, 3, -2, 0, 4}, list.Items())

	list.Move(item1, 1+list.Length())
	assert.Equal(t, []interface{}{2, 1, -3, 3, -2, 0, 4}, list.Items())

	list.Move(item2, 2+list.Length()*2)
	assert.Equal(t, []interface{}{1, -3, 2, 3, -2, 0, 4}, list.Items())

	list.Move(item3, -3-list.Length()*3)
	assert.Equal(t, []interface{}{1, 2, 3, -2, -3, 0, 4}, list.Items())

	list.Move(item4, 3+list.Length()*4)
	assert.Equal(t, []interface{}{1, 2, -2, -3, 0, 3, 4}, list.Items())

	list.Move(item5, -2-list.Length()*5)
	assert.Equal(t, []interface{}{1, 2, -3, 0, 3, 4, -2}, list.Items())

	list.Move(item6, 0+list.Length()*6)
	assert.Equal(t, []interface{}{1, 2, -3, 0, 3, 4, -2}, list.Items())

	list.Move(item6, 0-list.Length()*6)
	assert.Equal(t, []interface{}{1, 2, -3, 0, 3, 4, -2}, list.Items())

	list.Move(item7, 4+list.Length()*7)
	assert.Equal(t, []interface{}{1, 2, -3, 4, 0, 3, -2}, list.Items())

	list.Move(item5, -1-list.Length()*8)
	assert.Equal(t, []interface{}{1, 2, -3, 4, 0, -2, 3}, list.Items())

	list.Add(9)
	assert.Equal(t, []interface{}{1, 2, -3, 4, 0, -2, 3, 9}, list.Items())
}

func TestWrappedList_ItemAtIndex(t *testing.T) {
	list := WrappedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	tests := []struct {
		index    int
		expected int
	}{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 1},
		{5, 2},
		{6, 3},
		{7, 4},
		{-1, 4},
		{-2, 3},
		{-3, 2},
		{-4, 1},
		{-5, 4},
		{-6, 3},
		{-7, 2},
		{-8, 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Index %d", tt.index), func(t *testing.T) {
			assert.Equal(t, tt.expected, list.ItemAtIndex(tt.index))
		})
	}
}

func TestWrappedList_AdjustedIndex(t *testing.T) {
	list := WrappedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	tests := []struct {
		index    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 0},
		{5, 1},
		{6, 2},
		{7, 3},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Index %d", tt.index), func(t *testing.T) {
			assert.Equal(t, tt.expected, list.adjustedIndex(tt.index))
		})
	}
}
