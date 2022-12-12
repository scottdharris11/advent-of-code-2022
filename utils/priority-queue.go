package utils

type PriorityQueue struct {
	items []priorityQueueItem
}

// Empty determines if there is any item in the queue
func (q *PriorityQueue) Empty() bool {
	return len(q.items) == 0
}

// Next retrieves the item to process based on priority
func (q *PriorityQueue) Next() interface{} {
	n := len(q.items) - 1
	next := q.items[n]
	q.items = q.items[:n]
	return next.item
}

// Queue the supplied item based on lowest priority value being next
func (q *PriorityQueue) Queue(i interface{}, priority int) {
	qi := priorityQueueItem{item: i, priority: priority}

	insertIdx := -1
	for idx, item := range q.items {
		if item.priority < priority {
			insertIdx = idx
			break
		}
	}

	// If item not before anything, just append to end, otherwise
	// splice the existing array to add spot for new item
	if insertIdx == -1 {
		q.items = append(q.items, qi)
	} else {
		q.items = append(q.items[:insertIdx+1], q.items[insertIdx:]...)
		q.items[insertIdx] = qi
	}
}

type priorityQueueItem struct {
	item     interface{}
	priority int
}
