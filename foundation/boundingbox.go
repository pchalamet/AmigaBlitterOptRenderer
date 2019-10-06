package foundation

import (
	"math"
)

type BoundingBox struct {
	min Point
	max Point
}

func NewBoundingBox() BoundingBox {
	bb := BoundingBox{
		Point{math.MaxFloat64, math.MaxFloat64},
		Point{-math.MaxFloat64, -math.MaxFloat64}}

	return bb
}

func (bb1 BoundingBox) Intersect(bb2 BoundingBox) bool {
	return (bb2.max.X >= bb1.min.X) && (bb2.min.X <= bb1.max.X) && (bb2.max.Y >= bb1.min.Y) && (bb2.min.Y <= bb1.max.Y)
}

func (bb BoundingBox) Inside(pt Point) bool {
	return (bb.min.X <= pt.X) && (bb.min.Y <= pt.Y) && (pt.X <= bb.max.X) && (pt.Y <= bb.max.Y)
}

func (bb* BoundingBox) Add(pt Point) {
	minX := math.Min(pt.X, bb.min.X)
	minY := math.Min(pt.Y, bb.min.Y)
	maxX := math.Max(pt.X, bb.max.X)
	maxY := math.Max(pt.Y, bb.max.Y)

	bb.min = Point{minX, minY}
	bb.max = Point{maxX, maxY}
}
