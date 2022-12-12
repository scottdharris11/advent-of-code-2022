package utils

// Searcher defines interface necessary to use search algorithm
type Searcher interface {
	Goal(interface{}) bool
	PossibleNextMoves(interface{}) []SearchMove
	DistanceFromGoal(interface{}) int
}

// SearchMove represents the state of a particular move in the search
type SearchMove struct {
	Cost  int
	State interface{}
}

// SearchSolution represents the solution to a particular search path
type SearchSolution struct {
	Cost int
	Path []interface{}
}

// Search provides the ability to find the best path solution
type Search struct {
	Searcher Searcher
}

// Best finds the lowest cost solution through the search domain
func (s *Search) Best(init SearchMove) *SearchSolution {
	// Utilize a-star search approach to find the path to the
	// goal with the lowest cost.
	searchQueue := PriorityQueue{}
	searchQueue.Queue(init.State, init.Cost)
	cost := make(map[interface{}]int)
	from := make(map[interface{}]interface{})
	cost[init] = init.Cost
	var goal interface{}
	for !searchQueue.Empty() {
		var current = searchQueue.Next()
		if s.Searcher.Goal(current) {
			goal = current
			break
		}

		for _, next := range s.Searcher.PossibleNextMoves(current) {
			nCost := cost[current] + next.Cost
			cCost, ok := cost[next.State]
			if !ok || nCost < cCost {
				cost[next.State] = nCost
				priority := nCost + s.Searcher.DistanceFromGoal(next.State)
				searchQueue.Queue(next.State, priority)
				from[next.State] = current
			}
		}
	}

	// if no goal, return nil since we didn't reach it
	if goal == nil {
		return nil
	}

	// Construct path
	path := []interface{}{goal}
	current := goal
	for current != init.State {
		current = from[current]
		path = append(path, current)
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	// Print path
	return &SearchSolution{
		Cost: cost[goal],
		Path: path,
	}
}
