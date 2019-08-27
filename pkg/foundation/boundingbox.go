package foundation

import (
	"math"
)

type BoundingBox struct {
	Min Point
	Max Point
}

func NewBoundingBox() BoundingBox {
	bb := BoundingBox{
		Point{math.MaxFloat64, math.MaxFloat64},
		Point{math.MaxFloat64, math.MaxFloat64}}

	return bb
}

func (bb1 BoundingBox) Intersect(bb2 BoundingBox) bool {
	return (bb2.Max.X >= bb1.Min.X) && (bb2.Min.X <= bb1.Max.X) && (bb2.Max.Y >= bb1.Min.Y) && (bb2.Min.Y <= bb1.Max.Y)
}

func (bb BoundingBox) Inside(pt Point) bool {
	return (bb.Min.X <= pt.X) && (bb.Min.Y <= pt.Y) && (pt.X <= bb.Max.X) && (pt.Y <= bb.Max.Y)
}

func (bb BoundingBox) Add(pt Point) BoundingBox {
	minX := math.Min(pt.X, bb.Min.X)
	minY := math.Min(pt.Y, bb.Min.Y)
	maxX := math.Max(pt.X, bb.Max.X)
	maxY := math.Max(pt.Y, bb.Max.Y)
	return BoundingBox{
		Point{minX, minY},
		Point{maxX, maxY}}
}
