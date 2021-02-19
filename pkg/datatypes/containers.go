package datatypes

//MinHeap represents a heap data structure
type MinHeap struct {
	container []float64
}

//NewMinHeap returns a new empty MinHeap
func NewMinHeap() *MinHeap {
	res := make([]float64, 0, 0)
	return &MinHeap{res}
}

//Push pushes a new element to the heap and maintain the heap heuristic
func (m *MinHeap) Push(el float64) {
	m.container = append(m.container, el)
	heapify(m, len(m.container)-1)
}

//PushAll pushes all the elements provided to the heap and maintains the heap heuristic
func (m *MinHeap) PushAll(el ...float64) {
	for i := 0; i < len(el); i++ {
		m.container = append(m.container, el[i])
		heapify(m, len(m.container)-1)
	}
}

//Get() returns the underlying array of the heap
func (m *MinHeap) Get() []float64 {
	return m.container
}

//GetMin() returns the min of the MinHeap
func (m *MinHeap) GetMin() float64 {
	return m.container[0]
}

func heapify(m *MinHeap, i int) {
	parent := (i - 2) / 2
	if parent >= 0 {
		if m.container[i] < m.container[parent] {
			temp := m.container[parent]
			m.container[parent] = m.container[i]
			m.container[i] = temp

			heapify(m, parent)
		}
	}
}
