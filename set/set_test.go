package set

import "testing"

func TestSetConstructors(t *testing.T) {
	var uninitialized Set[int]

	if uninitialized.IsInitialized() {
		t.Error("Expected false, got true")
	}

	uninitialized.InitializeIfNot()

	if !uninitialized.IsInitialized() {
		t.Error("Expected true, got false")
	}

	s := New[int]()

	if !s.IsInitialized() {
		t.Error("Expected true, got false")
	}

	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", s.Size())
	}

	sh := New[int](10)

	if !sh.IsInitialized() {
		t.Error("Expected true, got false")
	}

	if sh.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", sh.Size())
	}

	set := Of(1, 2, 3, 4, 5)

	if !set.IsInitialized() {
		t.Error("Expected true, got false")
	}

	if set.Size() != 5 {
		t.Errorf("Expected size 5, got %d instead", set.Size())
	}

	if !set.ContainsAll(1, 2, 3, 4, 5) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	hashmap := make(map[int32]int32)
	hashmap[0] = 0
	hashmap[1] = 1
	hashmap[2] = 2
	hashmap[3] = 2

	keySet := FromKeys(hashmap)
	valueSet := FromValues(hashmap)

	if !keySet.IsInitialized() {
		t.Error("Expected true, got false")
	}

	if !valueSet.IsInitialized() {
		t.Error("Expected true, got false")
	}

	if keySet.Size() != 4 {
		t.Errorf("Expected size 4, got %d instead", keySet.Size())
	}

	if valueSet.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", valueSet.Size())
	}

	if !keySet.ContainsAll(0, 1, 2, 3) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if !valueSet.ContainsAll(0, 1, 2) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	clone := set.Clone()

	if !clone.IsInitialized() {
		t.Error("Expected true, got false")
	}

	if clone.Size() != 5 {
		t.Errorf("Expected size 5, got %d instead", clone.Size())
	}

	if !clone.ContainsAll(1, 2, 3, 4, 5) {
		t.Error("Set does not contain all the elements it should have contained")
	}
}

func TestSetClear(t *testing.T) {
	set := New[int]()
	set.Add(5)
	set.Add(6)
	set.Add(7)

	set.Clear()

	if set.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", set.Size())
	}
}

func TestSetAddRemovePopOneContainsSizeIsEmpty(t *testing.T) {
	s := New[int]()
	s.Add(5)

	if !s.Contains(5) {
		t.Error("Set does not contain expected item: 5")
	}

	s.Add(6)
	s.Remove(6)

	if s.Contains(6) {
		t.Error("Set contains an unexpected item: 6")
	}

	s = New[int]()

	if !s.IsEmpty() {
		t.Error("The newly initialized set is not empty")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", s.Size())
	}

	if !s.ContainsAll(1, 2, 3) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if !s.ContainsSome(7, 8, 9, 10, 11, 3) {
		t.Error("Set does not contain some of the elemenets it should have contained")
	}

	clone := s.Clone()
	val := s.PopOne()

	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d instead", s.Size())
	}

	if !clone.Contains(val) {
		t.Errorf("Set does not contain expected item: %d", val)
	}

	clone.Remove(val)
	val = s.PopOne()

	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d instead", s.Size())
	}

	if !clone.Contains(val) {
		t.Errorf("Set does not contain expected item: %d", val)
	}

	clone.Remove(val)
	val = s.PopOne()

	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", s.Size())
	}

	if !clone.Contains(val) {
		t.Errorf("Set does not contain expected item: %d", val)
	}
}

func TestSetUnion(t *testing.T) {
	set1 := Of(1, 2, 3)
	set2 := Of(4, 5, 6)

	union := set1.Union(set2)

	if union.Size() != 6 {
		t.Errorf("Expected size 6, got %d instead", union.Size())
	}

	if !union.Contains(1) || !union.Contains(2) || !union.Contains(3) || !union.Contains(4) || !union.Contains(5) || !union.Contains(6) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if set1.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", set1.Size())
	}
	
	if set2.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", set2.Size())
	}

	if !set1.Contains(1) || !set1.Contains(2) || !set1.Contains(3) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if !set2.Contains(4) || !set2.Contains(5) || !set2.Contains(6) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	set1.UnionWith(set2)

	if set1.Size() != 6 {
		t.Errorf("Expected size 6, got %d instead", set1.Size())
	}

	if !set1.Contains(1) || !set1.Contains(2) || !set1.Contains(3) || !set1.Contains(4) || !set1.Contains(5) || !set1.Contains(6) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if set1.Size() != 6 {
		t.Errorf("Expected size 6, got %d instead", set1.Size())
	}
	
	if set2.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", set2.Size())
	}

	if !set2.Contains(4) || !set2.Contains(5) || !set2.Contains(6) {
		t.Error("Set does not contain all the elements it should have contained")
	}
}

func TestSetIntersection(t *testing.T) {
	set1 := Of(1, 2, 3)
	set2 := Of(2, 3, 4)

	intersection := set1.Intersection(set2)

	if intersection.Size() != 2 {
		t.Errorf("Expected size 2, got %d instead", intersection.Size())
	}

	if !intersection.Contains(2) || !intersection.Contains(3) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if set1.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", set1.Size())
	}
	
	if set2.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", set2.Size())
	}

	if !set1.Contains(1) || !set1.Contains(2) || !set1.Contains(3) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if !set2.Contains(2) || !set2.Contains(3) || !set2.Contains(4) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	set1.IntersectWith(set2)

	if set1.Size() != 2 {
		t.Errorf("Expected size 2, got %d instead", set1.Size())
	}

	if !set1.Contains(2) || !set1.Contains(3) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if set2.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", set2.Size())
	}

	if !set2.Contains(2) || !set2.Contains(3) || !set2.Contains(4) {
		t.Error("Set does not contain all the elements it should have contained")
	}
}

func TestSetExcept(t *testing.T) {
	set1 := Of(1, 2, 3)
	set2 := Of(3, 4, 5)

	except := set1.Except(set2)

	if except.Size() != 2 {
		t.Errorf("Expected size 2, got %s instead", except.String())
	}

	if !except.Contains(1) || !except.Contains(2) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if set1.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", set1.Size())
	}
	
	if set2.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", set2.Size())
	}

	if !set1.Contains(1) || !set1.Contains(2) || !set1.Contains(3) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if !set2.Contains(3) || !set2.Contains(4) || !set2.Contains(5) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	set1.ExceptWith(set2)

	if set1.Size() != 2 {
		t.Errorf("Expected size 2, got %d instead", set1.Size())
	}

	if !set1.Contains(1) || !set1.Contains(2) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if set2.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", set2.Size())
	}

	if !set2.Contains(3) || !set2.Contains(4) || !set2.Contains(5) {
		t.Error("Set does not contain all the elements it should have contained")
	}
}

func TestSetSymmetricExcept(t *testing.T) {
	set1 := Of(1, 2, 3)
	set2 := Of(2, 3, 4)

	set3 := set1.SymmetricExcept(set2)

	if set3.Size() != 2 {
		t.Errorf("Expected size 2, got %s instead", set3.String())
	}

	if !set3.Contains(1) || !set3.Contains(4) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if set1.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", set1.Size())
	}
	
	if set2.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", set2.Size())
	}

	if !set1.Contains(1) || !set1.Contains(2) || !set1.Contains(3) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if !set2.Contains(2) || !set2.Contains(3) || !set2.Contains(4) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	set1.SymmetricExceptWith(set2)

	if set1.Size() != 2 {
		t.Errorf("Expected size 2, got %d instead", set1.Size())
	}

	if !set1.Contains(1) || !set1.Contains(4) {
		t.Error("Set does not contain all the elements it should have contained")
	}

	if set2.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", set2.Size())
	}

	if !set2.Contains(2) || !set2.Contains(3) || !set2.Contains(4) {
		t.Error("Set does not contain all the elements it should have contained")
	}
}

func TestSetOverlaps(t *testing.T) {
	set1 := Of(1, 2, 3)
	set2 := Of(4, 5, 6)
	set3 := Of(2, 4, 6)

	if set1.Overlaps(set2) {
		t.Error("Expected false, got true")
	}

	if set2.Overlaps(set1) {
		t.Error("Expected false, got true")
	}

	if !set1.Overlaps(set3) {
		t.Error("Expected true, got false")
	}

	if !set3.Overlaps(set1) {
		t.Error("Expected true, got false")
	}

	if !set2.Overlaps(set3) {
		t.Error("Expected true, got false")
	}

	if !set3.Overlaps(set2) {
		t.Error("Expected true, got false")
	}
}

func TestSetEquals(t *testing.T) {
	set1 := Of(1, 2, 3, 4, 5)
	set2 := Of(1, 2, 3, 4, 5)

	if !set1.SetEquals(set2) {
		t.Error("Expected true, got false")
	}

	if !set2.SetEquals(set1) {
		t.Error("Expected true, got false")
	}

	set2.Remove(5)

	if set1.SetEquals(set2) {
		t.Error("Expected false, got true")
	}

	if set2.SetEquals(set1) {
		t.Error("Expected false, got true")
	}

	set1.Remove(1)

	if set1.SetEquals(set2) {
		t.Error("Expected false, got true")
	}

	if set2.SetEquals(set1) {
		t.Error("Expected false, got true")
	}
}

func TestSetSubsetSuperset(t *testing.T) {
	set1 := Of(1, 2, 3)
	set2 := Of(1, 2, 3)
	set3 := Of(1, 2, 3, 4)

	if !set1.IsSubsetOf(set2) {
		t.Error("Expected true, got false")
	}

	if !set2.IsSubsetOf(set1) {
		t.Error("Expected true, got false")
	}

	if !set1.IsSupersetOf(set2) {
		t.Error("Expected true, got false")
	}

	if !set2.IsSupersetOf(set1) {
		t.Error("Expected true, got false")
	}

	if set1.IsProperSubsetOf(set2) {
		t.Error("Expected false, got true")
	}

	if set2.IsProperSubsetOf(set1) {
		t.Error("Expected false, got true")
	}

	if set1.IsProperSupersetOf(set2) {
		t.Error("Expected false, got true")
	}

	if set2.IsProperSupersetOf(set1) {
		t.Error("Expected false, got true")
	}

	if !set1.IsSubsetOf(set3) {
		t.Error("Expected true, got false")
	}

	if set3.IsSubsetOf(set1) {
		t.Error("Expected false, got true")
	}

	if set1.IsSupersetOf(set3) {
		t.Error("Expected false, got true")
	}

	if !set3.IsSupersetOf(set1) {
		t.Error("Expected true, got false")
	}

	if !set1.IsProperSubsetOf(set3) {
		t.Error("Expected true, got false")
	}

	if set3.IsProperSubsetOf(set1) {
		t.Error("Expected false, got true")
	}

	if set1.IsProperSupersetOf(set3) {
		t.Error("Expected false, got true")
	}

	if !set3.IsProperSupersetOf(set1) {
		t.Error("Expected true, got false")
	}
}

func TestSetToSlice(t *testing.T) {
	set := Of(1, 2, 3)
	slice := set.ToSlice()

	for _, val := range slice {
		set.Remove(val)
	}

	if set.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", set.Size())
	}
}

func TestSetString(t *testing.T) {
	set := Of(1, 2)
	str := set.String()

	if str != "Set{1, 2}" && str != "Set{2, 1}" {
		t.Errorf("Expected 'Set{1, 2}' or 'Set{2, 1}', got '%s' instead", str)
	}
}