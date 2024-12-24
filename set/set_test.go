package set

import (
	"testing"
)

func TestUnion(t *testing.T) {
	a := NewSet[int]()
	b := NewSet[int]()
	c := NewSet[int]()
	a.Add(1)
	a.Add(2)

	b.Add(2)
	b.Add(3)

	c.Add(4)
	c.Add(5)
	d := Union(a, b)
	if d.Size() != 3 {
		t.Errorf("Expected size to be 3 got %d", d.Size())
	}
	d = Union(b, c)
	if d.Size() != 4 {
		t.Errorf("Expected size to be 4 got %d", d.Size())
	}
}

func TestIntersection(t *testing.T) {
	a := NewSet[int]()
	b := NewSet[int]()
	c := NewSet[int]()
	a.Add(1)
	a.Add(2)

	b.Add(2)
	b.Add(3)

	c.Add(4)
	c.Add(5)
	d := Intersection(a, b)
	if d.Size() != 1 {
		t.Errorf("Expected size to be 1 got %d", d.Size())
	}
	if !d.Contains(2) {
		t.Errorf("Expected to contain 2")
	}
	d = Intersection(b, c)
	if d.Size() != 0 {
		t.Errorf("Expected size to be 0 got %d", d.Size())
	}
}

func TestAll(t *testing.T) {
	a := NewSet[int]()
	b := NewSet[int]()
	c := NewSet[int]()
	a.Add(1)
	a.Add(2)

	b.Add(2)
	b.Add(3)

	c.Add(4)
	c.Add(5)
	d := Union(a, b)
	d = Union(d, c)
	if d.Size() != 5 {
		t.Errorf("Expected size to be 5 got %d", d.Size())
	}
	for element := range d.All() {
		if !d.Contains(element) {
			t.Errorf("Expected to contain %d", element)
		}
	}
}
