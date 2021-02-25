package geometry

//MinHeap represents a heap data structure
type MinHeap struct {
	container []*Intersection
}

//NewMinHeap returns a new empty MinHeap
func NewMinHeap() *MinHeap {
	res := make([]*Intersection, 0, 0)
	return &MinHeap{res}
}

//Push pushes a new element to the heap and maintain the heap heuristic
func (m *MinHeap) Push(el *Intersection) {
	m.container = append(m.container, el)
	heapify(m, len(m.container)-1)
}

//PushAll pushes all the elements provided to the heap and maintains the heap heuristic
func (m *MinHeap) PushAll(el ...*Intersection) {
	for i := 0; i < len(el); i++ {
		m.container = append(m.container, el[i])
		heapify(m, len(m.container)-1)
	}
}

//Get returns the underlying array of the heap
func (m *MinHeap) Get() []*Intersection {
	return m.container
}

//Copy returns a copy of the minheap
func (m *MinHeap) Copy() *MinHeap{
	return createMinHeap(m.container)
}

//sets the minheap container to the parameter without checking heap property :only use on arrays already satisfying heap property
func createMinHeap(input []*Intersection) *MinHeap{
	h := NewMinHeap()
	res := make([]*Intersection, 0, 0)
	for i := 0; i < len(input); i++{
		in := NewIntersection(input[i].Object, input[i].T)
		res = append(res, in)
	}
	h.container = res
	return h
}

//GetMin returns the min of the MinHeap
func (m *MinHeap) GetMin() *Intersection {
	return m.container[0]
}

//ExtractMin returns the min and extracts it from the container
func (m *MinHeap) ExtractMin() *Intersection{
	if len(m.container) == 0 {return nil}
	min := m.container[0]
	m.container[0] = m.container[len(m.container) -1]
	m.container = m.container[:len(m.container) -1]
	minHeapify(m, 0)
	return min

}

func heapify(m *MinHeap, i int) {
	parent := (i - 2) / 2
	if parent >= 0 {
		if m.container[i].T < m.container[parent].T {
			temp := m.container[parent]
			m.container[parent] = m.container[i]
			m.container[i] = temp

			heapify(m, parent)
		}
	}
}

func minHeapify(m *MinHeap, i int){
	l := left(i)
	r := right(i)
	smallest := i
	if l < len(m.container) && m.container[l].T < m.container[smallest].T{
		smallest = l
	}
	if r < len(m.container) && m.container[r].T < m.container[smallest].T{
		smallest = r
	}
	if smallest != i {
		temp := m.container[i]
		m.container[i] = m.container[smallest]
		m.container[smallest] = temp
		minHeapify(m, smallest)
	}
}

func left(i int) int{
	return 2 * i + 1
}

func right(i int) int{
 	return 2 * i + 2
}