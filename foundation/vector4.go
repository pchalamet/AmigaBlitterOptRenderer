package foundation

type Vector4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

func NewVector4(x, y, z, w float64) Vector4 {
	return Vector4{ X: x, 
					Y: y, 
					Z: z, 
					W: w }
}
 
func (this Vector4) ToVector3() Vector3 {
	return NewVector3(this.X, this.Y, this.Z)
}
