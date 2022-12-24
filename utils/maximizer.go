package utils

// Maximizer defines interface necessary to use maximization algorithm
type Maximizer interface {
	Goal(interface{}) (bool, int)
	PossibleNextStates(interface{}, int) []interface{}
}

// MaxFinder provides the ability to find the best solution
type MaxFinder struct {
	Maximizer Maximizer
}

// Max finds the maximum solution through the search domain
func (m *MaxFinder) Max(s interface{}, sMax interface{}, maxVal int) (interface{}, int) {
	goal, val := m.Maximizer.Goal(s)
	if goal {
		if val > maxVal {
			return s, val
		}
		return sMax, maxVal
	}
	for _, next := range m.Maximizer.PossibleNextStates(s, maxVal) {
		sMax, maxVal = m.Max(next, sMax, maxVal)
	}
	return sMax, maxVal
}
