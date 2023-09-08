package stack

import "sync"

// DefaultCap is the default stack capacity.
const DefaultCap = 10

// SliceStack is a struct with methods needed to implement the Stack interface.
type SliceStack struct {
	slice []interface{}
	top   int
	mu    sync.Mutex
}

// NewSliceStack returns an empty SliceStack.
func NewSliceStack() *SliceStack {
	return &SliceStack{
		slice: make([]interface{}, DefaultCap),
		top:   -1,
	}
}

// Size returns the size of the stack.
func (ss *SliceStack) Size() int {
	ss.mu.Lock()
	top := ss.top
	ss.mu.Unlock()

	return top + 1
}

// Push pushes value onto the stack.
func (ss *SliceStack) Push(value interface{}) {
	ss.mu.Lock()
	ss.top++
	top := ss.top
	slice := ss.slice
	ss.mu.Unlock()
	if top == len(slice) {
		// Reallocate
		newSlice := make([]interface{}, len(slice)*2)
		copy(newSlice, slice)

		ss.mu.Lock()
		ss.slice = newSlice
		ss.mu.Unlock()
	}
	ss.mu.Lock()
	ss.slice[ss.top] = value
	ss.mu.Unlock()
}

// Pop pops the value at the top of the stack and returns it.
func (ss *SliceStack) Pop() (value interface{}) {
	ss.mu.Lock()
	top := ss.top
	ss.mu.Unlock()
	if top > -1 {
		ss.mu.Lock()
		defer func() { ss.top-- }()
		sliceval := ss.slice[ss.top]
		ss.mu.Unlock()
		return sliceval
	}
	return nil
}
