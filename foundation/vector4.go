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

func (v Vector4) Mult(m Matrix) Vector4 {
    x := (v.X * m.M11) + (v.Y * m.M21) + (v.Z * m.M31) + (v.W * m.M41)
    y := (v.X * m.M12) + (v.Y * m.M22) + (v.Z * m.M32) + (v.W * m.M42)
    z := (v.X * m.M13) + (v.Y * m.M23) + (v.Z * m.M33) + (v.W * m.M43)
    w := (v.X * m.M14) + (v.Y * m.M24) + (v.Z * m.M34) + (v.W * m.M44)
	return NewVector4(x, y, z, w)
}
