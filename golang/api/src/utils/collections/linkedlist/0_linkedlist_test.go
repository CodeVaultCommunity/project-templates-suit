// Package linkedlist provides linkedlist implementation
package linkedlist

import (
	"reflect"
	"testing"

	errorsapp "mod_name/error"
)

/* --------------------------
   Helpers
---------------------------*/

func mustPanic(t *testing.T, expected error, f func()) {
	t.Helper()

	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("expected panic, got none")
		}

		err, ok := r.(error)
		if !ok {
			t.Fatalf("expected panic(error), got %T", r)
		}

		if reflect.TypeOf(err) != reflect.TypeOf(expected) {
			t.Fatalf(
				"expected panic type %T, got %T",
				expected,
				err,
			)
		}
	}()

	f()
}

func alwaysTrue(_ int) bool {
	return true
}

func alwaysFalse(_ int) bool {
	return false
}

/* --------------------------
   Constructor
---------------------------*/

func TestNew(t *testing.T) {
	list := New[int]()
	if list == nil {
		t.Fatal("expected non-nil list")
	}
	if list.GetSize() != 0 {
		t.Fatalf("expected size 0, got %d", list.GetSize())
	}
	if !list.IsEmpty() {
		t.Fatal("expected empty list")
	}
}

/* --------------------------
   AddFront / AddBack
---------------------------*/

func TestAddFrontAndBack(t *testing.T) {
	list := New[int]()

	list.AddFront(2)
	list.AddFront(1)
	list.AddBack(3)

	if list.GetSize() != 3 {
		t.Fatalf("expected size 3, got %d", list.GetSize())
	}

	if list.Get(0) != 1 || list.Get(1) != 2 || list.Get(2) != 3 {
		t.Fatal("unexpected list order")
	}
}

func TestAddNilListPanics(t *testing.T) {
	var list *LinkedList[int]

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.AddFront(1)
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.AddBack(1)
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.AddAt(0, 0)
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.AddAtRev(0, 0)
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.PopLast()
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.GetRev(0)
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.GetSize()
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.Filter(nil)
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.AddAt(0, 0)
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.AddAtRev(0, 0)
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		l := New[int]()
		l.Filter(nil)
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.FilterAsNewList(nil)
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		l := New[int]()
		l.Filter(nil)
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() { list.IsEmpty() })
	mustPanic(t, errorsapp.ErrNilPointer(), func() { list.Iter() })
}

/* --------------------------
   AddAt / AddAtRev
---------------------------*/

func TestAddAt(t *testing.T) {
	list := New[int]()

	list.AddAt(0, 1)
	list.AddAt(1, 3)
	list.AddAt(1, 2)

	if list.Get(0) != 1 || list.Get(1) != 2 || list.Get(2) != 3 || list.size != 3 {
		t.Fatal("AddAt failed")
	}
}

func TestAddAtRev(t *testing.T) {
	list := New[int]()

	list.AddAtRev(0, 1) // [1]
	list.AddAtRev(1, 3) // [3,1]
	list.AddAtRev(1, 2) // [3,2,1]

	if list.Get(0) != 3 || list.Get(1) != 2 || list.Get(2) != 1 || list.size != 3 {
		t.Fatal("AddAtRev failed")
	}
}

/* --------------------------
   PopFirst / PopLast
---------------------------*/

func TestPopFirst(t *testing.T) {
	list := New[int]()
	list.AddBack(1)
	list.AddBack(2)

	v := list.PopFirst()
	if v != 1 {
		t.Fatalf("expected 1, got %d", v)
	}

	if list.GetSize() != 1 {
		t.Fatal("size not updated")
	}
}

func TestPopLast(t *testing.T) {
	list := New[int]()
	list.AddBack(1)

	// test single item on list block
	func() {
		v := list.PopLast()
		if v != 1 {
			t.Fatalf("expected 1, got %d", v)
		}
		list.AddBack(v)
	}()

	list.AddBack(2)
	list.AddBack(3)

	v := list.PopLast()
	if v != 3 {
		t.Fatalf("expected 3, got %d", v)
	}

	if list.GetSize() != 2 {
		t.Fatal("size not updated")
	}
}

func TestPopPanics(t *testing.T) {
	var nilList *LinkedList[int]
	empty := New[int]()

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		nilList.PopFirst()
	})

	mustPanic(t, errorsapp.ErrListIsEmpty(), func() {
		empty.PopFirst()
	})

	mustPanic(t, errorsapp.ErrListIsEmpty(), func() {
		empty.PopLast()
	})
}

/* --------------------------
   Get / GetRev / Size
---------------------------*/

func TestGetAndGetRev(t *testing.T) {
	list := New[int]()
	list.AddBack(1)
	list.AddBack(2)
	list.AddBack(3)

	if list.Get(1) != 2 {
		t.Fatal("Get failed")
	}

	if list.GetRev(0) != 3 {
		t.Fatal("GetRev failed")
	}

	mustPanic(t, errorsapp.ErrIndexOutOfBound(), func() { list.get(20) })
	mustPanic(t, errorsapp.ErrIndexOutOfBound(), func() { list.getRev(20) })
	head := list.getRev(list.size)
	if head != list.head {
		t.Fatal("getRev faild: expect to return hean on getRev(list.size)")
	}

	list.AddBack(4)
	list.AddBack(5)
	list.AddBack(6)
	if list.GetRev(3) != 3 {
		t.Fatal("GetRev failed, expect 3")
	}
}

func TestGetPanics(t *testing.T) {
	var nilList *LinkedList[int]
	list := New[int]()

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		nilList.Get(0)
	})

	mustPanic(t, errorsapp.ErrIndexOutOfBound(), func() {
		list.Get(0)
	})
}

/* --------------------------
   Iter
---------------------------*/

func TestIter(t *testing.T) {
	list := New[int]()
	list.AddBack(1)
	list.AddBack(2)
	list.AddBack(3)

	sum := 0
	for v := range list.Iter() {
		sum += v
	}

	if sum != 6 {
		t.Fatalf("expected sum 6, got %d", sum)
	}
}

/* --------------------------
   MapLinkedList
---------------------------*/

func TestMapLinkedList(t *testing.T) {
	list := New[int]()
	list.AddBack(1)
	list.AddBack(2)

	mapped := Map(list, func(x int) int {
		return x * 2
	})

	if mapped.Get(0) != 2 || mapped.Get(1) != 4 {
		t.Fatal("MapLinkedList failed")
	}
}

func TestMapLinkedListPanics(t *testing.T) {
	list := New[int]()

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		Map[int, int](nil, func(x int) int { return x })
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		Map[int, any](list, nil)
	})
}

/* --------------------------
   FilterLinkedList
---------------------------*/

func TestFilterLinkedList(t *testing.T) {
	list := New[int]()
	list.AddBack(1)
	list.AddBack(2)
	list.AddBack(3)

	list.Filter(func(x int) bool {
		return x%2 == 1
	})

	if list.GetSize() != 2 {
		t.Fatal("FilterLinkedList size wrong")
	}

	if list.Get(0) != 1 || list.Get(1) != 3 {
		t.Fatal("FilterLinkedList values wrong")
	}
}

func TestFilterLinkedListAllRemoved(t *testing.T) {
	list := New[int]()
	list.AddBack(2)
	list.AddBack(4)

	list.Filter(func(x int) bool {
		return x%2 == 1
	})

	if !list.IsEmpty() {
		t.Fatal("expected empty list")
	}
}

/* --------------------------
   MapFilterLinkedList
---------------------------*/

func TestMapFilterLinkedList(t *testing.T) {
	list := New[int]()
	list.AddBack(1)
	list.AddBack(2)

	newList := list.FilterAsNewList(func(x int) bool {
		return x == 1
	})

	if newList.GetSize() != 2 {
		t.Fatal("MapFilterLinkedList size wrong")
	}
}

/* --------------------------
   Some / All
---------------------------*/

func TestSomeAndAll(t *testing.T) {
	list := New[int]()
	list.AddBack(1)
	list.AddBack(2)
	list.AddBack(3)

	if !list.Some(func(x int) bool { return x == 2 }) {
		t.Fatal("SomeInLinkedList failed")
	}

	if list.Some(alwaysFalse) {
		t.Fatal("SomeInLinkedList failed")
	}

	if list.All(func(x int) bool { return x%2 == 0 }) {
		t.Fatal("AllInLinkedList failed")
	}

	if !list.All(alwaysTrue) {
		t.Fatal("AllInLinkedList failed")
	}
}

func TestSomeAllPanics(t *testing.T) {
	var list *LinkedList[int]

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.Some(func(x int) bool { return true })
	})

	mustPanic(t, errorsapp.ErrNilPointer(), func() {
		list.All(func(x int) bool { return true })
	})
}

func TestLinkedList_FullScenario_WithPointers(t *testing.T) {
	list := New[int]()

	/* ----------------
	   Construção
	------------------*/

	list.AddBack(2)     // [2]
	list.AddFront(1)    // [1,2]
	list.AddBack(4)     // [1,2,4]
	list.AddAt(2, 3)    // [1,2,3,4]
	list.AddAtRev(0, 5) // [1,2,3,4,5]

	if list.GetSize() != 5 {
		t.Fatalf("expected size 5, got %d", list.GetSize())
	}

	/* ----------------
	   Ponteiros básicos
	------------------*/

	if list.head == nil || list.tail == nil {
		t.Fatal("head or tail is nil")
	}

	if list.tail.GetNext() != nil {
		t.Fatal("tail.Next must be nil")
	}

	/* ----------------
	   Encadeamento
	------------------*/

	expected := []int{1, 2, 3, 4, 5}
	node := list.head

	for i, v := range expected {
		if node == nil {
			t.Fatalf("node %d is nil", i)
		}
		if node.GetValue() != v {
			t.Fatalf("expected %d at index %d, got %d", v, i, node.GetValue())
		}
		node = node.GetNext()
	}

	if node != nil {
		t.Fatal("list has extra nodes")
	}

	/* ----------------
	   Iterador
	------------------*/

	sum := 0
	for v := range list.Iter() {
		sum += v
	}

	if sum != 15 {
		t.Fatalf("expected sum 15, got %d", sum)
	}

	/* ----------------
	   Predicados
	------------------*/

	if !list.Some(func(x int) bool { return x == 3 }) {
		t.Fatal("SomeInLinkedList failed")
	}

	if !list.All(func(x int) bool { return x > 0 }) {
		t.Fatal("AllInLinkedList failed")
	}

	/* ----------------
	   Filter (remove pares)
	------------------*/

	list.Filter(func(x int) bool {
		return x%2 == 1
	}) // [1,3,5]

	if list.GetSize() != 3 {
		t.Fatalf("expected size 3 after filter, got %d", list.GetSize())
	}

	if list.head.GetValue() != 1 {
		t.Fatal("head value incorrect after filter")
	}
	if list.tail.GetValue() != 5 {
		t.Fatal("tail value incorrect after filter")
	}
	if list.tail.GetNext() != nil {
		t.Fatal("tail.Next must be nil after filter")
	}

	/* ----------------
	   Pop
	------------------*/

	if list.PopFirst() != 1 {
		t.Fatal("PopFirst failed")
	}

	if list.PopLast() != 5 {
		t.Fatal("PopLast failed")
	}

	if list.GetSize() != 1 {
		t.Fatalf("expected size 1, got %d", list.GetSize())
	}

	if list.head != list.tail {
		t.Fatal("head and tail must point to same node")
	}

	/* ----------------
	   Map
	------------------*/

	mapped := Map(list, func(x int) int {
		return x * 10
	})

	if mapped.GetSize() != 1 {
		t.Fatal("mapped list size incorrect")
	}

	if mapped.Get(0) != 30 {
		t.Fatalf("expected mapped value 30, got %d", mapped.Get(0))
	}

	/* ----------------
	   Esvaziamento final
	------------------*/

	list.PopFirst()

	if !list.IsEmpty() {
		t.Fatal("list should be empty")
	}

	if list.head != nil || list.tail != nil {
		t.Fatal("head and tail must be nil on empty list")
	}
}
