package stack

type stackOperation int

const (
	length stackOperation = iota
	push
	pop
)

type stackCommand struct {
	op               stackOperation
	value            interface{}
	response_element chan interface{}
}

// CspStack is a struct with methods needed to implement the Stack interface.
type CspStack struct {
	size    int
	top     *Element
	channel chan stackCommand
}

// NewCspStack returns an empty CspStack.
func NewCspStack() *CspStack {
	stack := &CspStack{
		channel: make(chan stackCommand),
	}
	go stack.run()
	return stack
}

// Size returns the size of the stack.
func (cs *CspStack) Size() int {
	responseChan := make(chan interface{})
	cs.channel <- stackCommand{op: length, response_element: responseChan}
	size := <-responseChan
	return size.(int)
}

// Push pushes value onto the stack.
func (cs *CspStack) Push(value interface{}) {
	cs.channel <- stackCommand{op: push, value: value}
}

// Pop pops the value at the top of the stack and returns it.
func (cs *CspStack) Pop() (value interface{}) {
	responseChan2 := make(chan interface{})
	cs.channel <- stackCommand{op: pop, response_element: responseChan2}
	return <-responseChan2
}

func (cs *CspStack) run() {
	for sc := range cs.channel {
		switch sc.op {
		case length:
			sc.response_element <- cs.size

		case push:
			cs.size++
			cs.top = &Element{sc.value, cs.top}
		case pop:
			if cs.size > 0 {
				value := cs.top.value
				cs.top = cs.top.next
				cs.size = cs.size - 1
				sc.response_element <- value
			} else {
				sc.response_element <- nil
			}
		}
	}
}
