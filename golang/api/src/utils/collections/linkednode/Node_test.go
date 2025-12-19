package linkednode

import (
	"errors"
	"testing"
)

func TestSLN(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		value int
		want  *SimpleLinkedNode[int]
	}{
		{
			value: 5,
			want:  &SimpleLinkedNode[int]{value: 5, next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.value)
			if got.value != tt.want.value {
				t.Errorf("NewSLN() = %v, want %v", got, tt.want)
			}

			if got.next != tt.want.next {
				t.Errorf("NewSLN() = %v, want %v", got.next, tt.want.next)
			}

			if got.GetValue() != tt.want.value {
				t.Errorf("GetValue() = %v, want %v", got.GetValue(), tt.want.value)
			}

			lastValue := got.SetValue(tt.want.value + tt.want.value)

			if lastValue != tt.value {
				t.Errorf("SetValue() returns = %v, want %v", lastValue, tt.want.value)
			}

			if got.GetValue() != (tt.want.value + tt.want.value) {
				t.Errorf("SetValue() = %v, want %v", got.GetValue(), (tt.want.value + tt.want.value))
			}

			if got.GetNext() != nil {
				t.Errorf("GetNext() = %v, want %v", got.GetNext(), nil)
			}

			last := got.SetNext(got)

			if last != nil {
				t.Errorf("SetNext() return = %v, want %v", last, nil)
			}

			if got.GetNext() != got {
				t.Errorf("SetNext/GetNext() = %v, want %v", got.GetNext(), got)
			}

			got.SetValue(tt.value)
			other := New(tt.value + tt.value)
			got.SetNext(other)

			for element := range got.Iterator() {
				if element != tt.value {
					if element != (tt.value + tt.value) {
						t.Errorf("Iterator() = %v, want %v", element, tt.value+tt.value)
					}
				}
			}

			new := Map(got, func(element int) int { return element * 2 })

			for element := range new.Iterator() {
				if element != tt.value*2 {
					if element != (tt.value+tt.value)*2 {
						t.Errorf("Iterator() = %v, want %v", element, (tt.value+tt.value)*2)
					}
				}
			}

			if got.value == new.value {
				t.Errorf("MapSNL() modify the original arr")
			}

			got.Swap(new)
			if got.value != tt.value*2 {
				t.Errorf("Swap() not modify correctly node. got: %d; expect: %d", got.value, tt.value*2)
			}
		})
	}
}

func TestPanicOnSLN(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		method func(node *SimpleLinkedNode[any])
	}{
		{
			name:   "Testing Panic: GetValue",
			method: func(node *SimpleLinkedNode[any]) { node.GetValue() },
		},
		{
			name:   "Testing Panic: SetValue",
			method: func(node *SimpleLinkedNode[any]) { node.SetValue(nil) },
		},
		{
			name:   "Testing Panic: GetNext",
			method: func(node *SimpleLinkedNode[any]) { node.GetNext() },
		},
		{
			name:   "Testing Panic: SetNext",
			method: func(node *SimpleLinkedNode[any]) { node.SetNext(nil) },
		},
		{
			name:   "Testing Panic: Swap",
			method: func(node *SimpleLinkedNode[any]) { node.Swap(nil) },
		},
		{
			name:   "Testing Panic: Map - NodeNil",
			method: func(node *SimpleLinkedNode[any]) { Map[any, any](nil, nil) },
		},
		{
			name:   "Testing Panic: Map - FuncNil",
			method: func(node *SimpleLinkedNode[any]) { Map[any, any](&SimpleLinkedNode[any]{}, nil) },
		},
		{
			name: "Testing Panic: Map - FuncNil",
			method: func(node *SimpleLinkedNode[any]) {
				start := &SimpleLinkedNode[any]{value: 4, next: nil}
				chain := Map(start, func(o any) any { return o })
				if start == chain || start.next != chain.next || start.value != chain.value {
					panic(errors.New("Erro, start == chain || start.next != chain.next || start.value != chain.value."))
				}
			},
		},
		{
			name:   "Testing Panic: Iterator",
			method: func(node *SimpleLinkedNode[any]) { node.Iterator() },
		},
	}
	for _, tt := range tests {
		func() {
			defer func() {
				if r := recover(); r != nil {
					switch e := r.(type) {
					case error:
						if e.Error() == "system.nil_pointer_error" {
							return
						}
						panic(e)
					}
				}
			}()
			tt.method(nil)
		}()
	}
}

func TestSimpleLinkedNode_Shift(t *testing.T) {
	tests := []struct {
		name string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := New(1)
			end := New(3)
			middle_1 := New(2)
			middle_2 := New(4)

			start.next = middle_1
			middle_1.next = middle_2
			middle_2.next = end

		})
	}
}
