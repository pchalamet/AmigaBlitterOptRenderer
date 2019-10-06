package foundation

import "testing"


func TestBoundingBoxAdd(t *testing.T) {
	bb := NewBoundingBox()
	bb.Add(Point{1, 1})
	bb.Add(Point{3, 3})
	bb.Add(Point{2, 2})
	bb.Add(Point{4, 5})

	if ! equalWithUncertainty(bb.min.X, 1.0) {
		t.Errorf("Unexpected value MinX %f", bb.min.X)
	}

	if ! equalWithUncertainty(bb.min.Y, 1.0) {
		t.Errorf("Unexpected value MinY %f", bb.min.Y)
	}

	if ! equalWithUncertainty(bb.max.X, 4.0) {
		t.Errorf("Unexpected value MaxX %f", bb.max.X)
	}

	if ! equalWithUncertainty(bb.max.Y, 5.0) {
		t.Errorf("Unexpected value MaxY %f", bb.max.Y)
	}
}

func TestBoundingBoxInside(t *testing.T) {
	bb := BoundingBox{ min:Point{1, 1}, max:Point{10, 10} }

	if !bb.Inside(Point{1, 1}) {
		t.Error("Unexpected result for Inside")
	}

	if !bb.Inside(Point{2, 2}) {
		t.Error("Unexpected result for Inside")
	}

	if bb.Inside(Point{0, 0}) {
		t.Error("Unexpected result for Inside")
	}
}

func TestBoundingBoxInsersect(t *testing.T) {
	bb1 := BoundingBox{ min:Point{1, 1}, max:Point{10, 10} }
	bb2 := BoundingBox{ min:Point{2, 2}, max:Point{10, 10} }
	bb3 := BoundingBox{ min:Point{20, 20}, max:Point{30, 30} }

	if !bb1.Intersect(bb2) {
		t.Error("BoundingBox's should intersect")
	}

	if bb1.Intersect(bb3) {
		t.Error("BoundingBox's should not intersect")
	}
}
