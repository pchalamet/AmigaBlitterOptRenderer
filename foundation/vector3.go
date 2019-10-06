package foundation

import "math"

type Vector3 struct {
	X float64
	Y float64
	Z float64
}


func NewVector3(x, y, z float64) Vector3 {
    return Vector3{ X: x, 
                    Y: y, 
                    Z: z }
}


func (v1 Vector3) Add (v2 Vector3) Vector3 {
    return NewVector3(v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z)    
}

func (v1 Vector3) Sub (v2 Vector3) Vector3 {
    return NewVector3(v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z)    
}

func (v Vector3) Neg () Vector3 {
    return NewVector3(-v.X, -v.Y, -v.Z)    
}

func (v1 Vector3) Scale (value float64) Vector3 {
    return NewVector3(v1.X*value, v1.Y*value, v1.Z*value)    
}

func (v1 Vector3) Div (value float64) Vector3 {
    return NewVector3(v1.X/value, v1.Y/value, v1.Z/value)    
}

func (v1 Vector3) Cross (v2 Vector3) Vector3 {
    x := v1.Y * v2.Z - v2.Y * v1.Z
    y := -(v1.X * v2.Z - v2.X * v1.Z)
    z := v1.X * v2.Y - v2.X * v1.Y
    return NewVector3(x, y, z)
}

func (v1 Vector3) Dot (v2 Vector3) float64 {
    return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
}

func (this Vector3) Length() float64 {
    return math.Sqrt(this.X*this.X + this.Y*this.Y + this.Z*this.Z)
}

func (this Vector3) Normalize() Vector3 {
    len := this.Length()
    return this.Div(len)
}

func (v Vector3) Mult(m Matrix) Vector3 {
    x := (v.X * m.M11) + (v.Y * m.M21) + (v.Z * m.M31) + m.M41
    y := (v.X * m.M12) + (v.Y * m.M22) + (v.Z * m.M32) + m.M42
    z := (v.X * m.M13) + (v.Y * m.M23) + (v.Z * m.M33) + m.M43
    return NewVector3(x, y, z)    
}
