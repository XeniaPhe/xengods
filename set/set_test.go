package set

import "testing"

func TestSetConstructors(t *testing.T) {
	s := New[int]()

	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", s.Size())
	}

	set := Of(1, 2, 3, 4, 5)

	if set.Size() != 5 {
		t.Errorf("Expected size 5, got %d instead", set.Size())
	}

	containsAll := set.Contains(1) && set.Contains(2) && set.Contains(3) && set.Contains(4) && set.Contains(5)

	if !containsAll {
		t.Error("Set does not contain all the elements it should have contained")
	}
}

func TestSetAddRemoveContainsSize(t *testing.T) {
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

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", s.Size())
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