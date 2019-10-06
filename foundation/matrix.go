package foundation

import "math"

type Matrix struct {
	M11 float64
	M12 float64
	M13 float64
	M14 float64
	M21 float64
	M22 float64
	M23 float64
	M24 float64
	M31 float64
	M32 float64
	M33 float64
	M34 float64
	M41 float64
	M42 float64
	M43 float64
	M44 float64
}


func (this Matrix) Item(index int) float64 {
	switch index {
	case 0:
		return this.M11
	case 1:
		return this.M12
	case 2:
		return this.M13
	case 3:
		return this.M14
	case 4:
		return this.M21
	case 5:
		return this.M22
	case 6:
		return this.M23
	case 7:
		return this.M24
	case 8:
		return this.M31
	case 9:
		return this.M32
	case 10:
		return this.M33
	case 11:
		return this.M34
	case 12:
		return this.M41
	case 13:
		return this.M42
	case 14:
		return this.M43
	case 15:
		return this.M44
	default:
		panic("Invalid index")
	}
}

func (this Matrix) Item2(row, column int) float64 {
	return this.Item((row * 4) + column)
}

func CreateLookAt(cameraPosition, cameraTarget, cameraUpVector Vector3) Matrix {
    vector := cameraPosition.Sub(cameraTarget).Normalize()
    vector2 := cameraUpVector.Cross(vector).Normalize()
    vector3 := vector.Cross(vector2)
    return Matrix{ vector2.X, vector3.X, vector.X, 0.0,
                   vector2.Y, vector3.Y, vector.Y, 0.0,
                   vector2.Z, vector3.Z, vector.Z, 0.0,
                   -vector2.Dot(cameraPosition), -vector3.Dot(cameraPosition), -vector.Dot(cameraPosition), 1.0 }
}

func CreatePerspective (width, height, nearPlaneDistance, farPlaneDistance float64) Matrix {
	if nearPlaneDistance <= 0.0 {
		panic("nearPlaneDistance <= 0")
	}
    if farPlaneDistance <= 0.0 {
		panic("farPlaneDistance <= 0")
	}
    if nearPlaneDistance >= farPlaneDistance {
		panic("nearPlaneDistance >= farPlaneDistance")
	}

    return Matrix{ (2.0 * nearPlaneDistance) / width, 0.0, 0.0, 0.0,
           		   0.0, (2.0 * nearPlaneDistance) / height, 0.0, 0.0,
           		   0.0, 0.0, farPlaneDistance / (nearPlaneDistance - farPlaneDistance), -1.0,
           		   0.0, 0.0, (nearPlaneDistance * farPlaneDistance) / (nearPlaneDistance - farPlaneDistance), 0.0 }
}


func CreatePerspectiveOffCenter (left, right, bottom, top, nearPlaneDistance, farPlaneDistance float64) Matrix {
    if nearPlaneDistance <= 0.0 {
		panic("nearPlaneDistance <= 0")
	}
    if farPlaneDistance <= 0.0 {
		panic("farPlaneDistance <= 0")
	}
    if nearPlaneDistance >= farPlaneDistance {
		panic("nearPlaneDistance >= farPlaneDistance")
	}

    numW := (2.0 * nearPlaneDistance) / (right - left)
    numH := (2.0 * nearPlaneDistance) / (top - bottom)
    return Matrix{ numW, 0.0, 0.0, 0.0,
           		   0.0, numH, 0.0, 0.0,
                   (left + right) / (right - left), (top + bottom) / (top - bottom), (nearPlaneDistance + farPlaneDistance) / (nearPlaneDistance - farPlaneDistance), -1.0,
           		   0.0, 0.0, (2.0 * nearPlaneDistance * farPlaneDistance) / (nearPlaneDistance - farPlaneDistance), 0.0 }
}

func CreatePerspectiveFieldOfView (fieldOfView, aspectRatio, nearPlaneDistance, farPlaneDistance float64) Matrix {
    if fieldOfView <= 0.0 || fieldOfView >= 3.141593 {
		panic("fieldOfView <= 0 or >= PI")
	}
    if nearPlaneDistance <= 0.0 {
		panic("nearPlaneDistance <= 0")
	}
    if farPlaneDistance <= 0.0 {
		panic("farPlaneDistance <= 0")
	}
    if nearPlaneDistance >= farPlaneDistance {
		panic("nearPlaneDistance >= farPlaneDistance")
	}

    top := nearPlaneDistance * math.Tan(fieldOfView / 2.0)
    bottom := -top
	return CreatePerspectiveOffCenter((bottom*aspectRatio), (top*aspectRatio), bottom, top, nearPlaneDistance, farPlaneDistance)
}

func CreateRotationX (radians float64) Matrix {
    val1 := math.Cos(radians)
    val2 := math.Sin(radians)
    return Matrix{ 1.0, 0.0, 0.0, 0.0,
            	   0.0, val1, val2, 0.0,
                   0.0, -val2, val1, 0.0,
		           0.0, 0.0, 0.0, 1.0 }
}

func CreateRotationY (radians float64) Matrix {
    val1 := math.Cos(radians)
    val2 := math.Sin(radians)
    return Matrix{ val1, 0.0, -val2, 0.0,
           		   0.0, 1.0, 0.0, 0.0,
           		   val2, 0.0, val1, 0.0,
           		   0.0, 0.0, 0.0, 1.0 }
}

func CreateRotationZ (radians float64) Matrix {
    val1 := math.Cos(radians)
    val2 := math.Sin(radians)
    return Matrix{ val1, val2, 0.0, 0.0,
           		   -val2, val1, 0.0, 0.0,
           		   0.0, 0.0, 1.0, 0.0,
	 			   0.0, 0.0, 0.0, 1.0 }
}

func CreateScale (scaleX, scaleY, scaleZ float64) Matrix {
    return Matrix{ scaleX, 0.0, 0.0, 0.0,
           		   0.0, scaleY, 0.0, 0.0,
           		   0.0, 0.0, scaleZ, 0.0,
					  0.0, 0.0, 0.0, 1.0 }
}

func CreateTranslation(positionX, positionY, positionZ float64) Matrix {
    return Matrix{ 1.0, 0.0, 0.0, 0.0,
           		   0.0, 1.0, 0.0, 0.0,
           		   0.0, 0.0, 1.0, 0.0,
		   		   positionX, positionY, positionZ, 1.0 }
}

func (m1 Matrix) Add(m2 Matrix) Matrix {
    return Matrix{ m1.M11 + m2.M11, m1.M12 + m2.M12, m1.M13 + m2.M13, m1.M14 + m2.M14,
           		   m1.M21 + m2.M21, m1.M22 + m2.M22, m1.M23 + m2.M23, m1.M24 + m2.M24,
           		   m1.M31 + m2.M31, m1.M32 + m2.M32, m1.M33 + m2.M33, m1.M34 + m2.M34,
		   		   m1.M41 + m2.M41, m1.M42 + m2.M42, m1.M43 + m2.M43, m1.M44 + m2.M44 }
}

func (m1 Matrix) Sub(m2 Matrix) Matrix {
    return Matrix{ m1.M11 - m2.M11, m1.M12 - m2.M12, m1.M13 - m2.M13, m1.M14 - m2.M14,
           		   m1.M21 - m2.M21, m1.M22 - m2.M22, m1.M23 - m2.M23, m1.M24 - m2.M24,
           		   m1.M31 - m2.M31, m1.M32 - m2.M32, m1.M33 - m2.M33, m1.M34 - m2.M34,
		   		   m1.M41 - m2.M41, m1.M42 - m2.M42, m1.M43 - m2.M43, m1.M44 - m2.M44 }
}

func (m Matrix) Neg() Matrix {
    return Matrix{ -m.M11, -m.M12, -m.M13, -m.M14,
           		   -m.M21, -m.M22, -m.M23, -m.M24,
           		   -m.M31, -m.M32, -m.M33, -m.M34,
		   		   -m.M41, -m.M42, -m.M43, -m.M44 }
}

func (m1 Matrix) Compose (m2 Matrix) Matrix {
    m11 := (m1.M11 * m2.M11) + (m1.M12 * m2.M21) + (m1.M13 * m2.M31) + (m1.M14 * m2.M41)
    m12 := (m1.M11 * m2.M12) + (m1.M12 * m2.M22) + (m1.M13 * m2.M32) + (m1.M14 * m2.M42)
    m13 := (m1.M11 * m2.M13) + (m1.M12 * m2.M23) + (m1.M13 * m2.M33) + (m1.M14 * m2.M43)
    m14 := (m1.M11 * m2.M14) + (m1.M12 * m2.M24) + (m1.M13 * m2.M34) + (m1.M14 * m2.M44)
    m21 := (m1.M21 * m2.M11) + (m1.M22 * m2.M21) + (m1.M23 * m2.M31) + (m1.M24 * m2.M41)
    m22 := (m1.M21 * m2.M12) + (m1.M22 * m2.M22) + (m1.M23 * m2.M32) + (m1.M24 * m2.M42)
    m23 := (m1.M21 * m2.M13) + (m1.M22 * m2.M23) + (m1.M23 * m2.M33) + (m1.M24 * m2.M43)
    m24 := (m1.M21 * m2.M14) + (m1.M22 * m2.M24) + (m1.M23 * m2.M34) + (m1.M24 * m2.M44)
    m31 := (m1.M31 * m2.M11) + (m1.M32 * m2.M21) + (m1.M33 * m2.M31) + (m1.M34 * m2.M41)
    m32 := (m1.M31 * m2.M12) + (m1.M32 * m2.M22) + (m1.M33 * m2.M32) + (m1.M34 * m2.M42)
    m33 := (m1.M31 * m2.M13) + (m1.M32 * m2.M23) + (m1.M33 * m2.M33) + (m1.M34 * m2.M43)
    m34 := (m1.M31 * m2.M14) + (m1.M32 * m2.M24) + (m1.M33 * m2.M34) + (m1.M34 * m2.M44)
    m41 := (m1.M41 * m2.M11) + (m1.M42 * m2.M21) + (m1.M43 * m2.M31) + (m1.M44 * m2.M41)
    m42 := (m1.M41 * m2.M12) + (m1.M42 * m2.M22) + (m1.M43 * m2.M32) + (m1.M44 * m2.M42)
    m43 := (m1.M41 * m2.M13) + (m1.M42 * m2.M23) + (m1.M43 * m2.M33) + (m1.M44 * m2.M43)
    m44 := (m1.M41 * m2.M14) + (m1.M42 * m2.M24) + (m1.M43 * m2.M34) + (m1.M44 * m2.M44)
    return Matrix{ m11, m12, m13, m14,
				   m21, m22, m23, m24,
				   m31, m32, m33, m34,
				   m41, m42, m43, m44 }
}

func (m Matrix) Transpose() Matrix {
    return Matrix{ m.M11, m.M21, m.M31, m.M41,
           		   m.M12, m.M22, m.M32, m.M42,
           		   m.M13, m.M23, m.M33, m.M43,
		   		   m.M14, m.M24, m.M34, m.M44 }
}

func (m Matrix) Mult(v Vector4) Vector4 {
    x := (m.M11 * v.X) + (m.M12 * v.Y) + (m.M13 * v.Z) + (m.M14 * v.W)
    y := (m.M21 * v.X) + (m.M22 * v.Y) + (m.M23 * v.Z) + (m.M24 * v.W)
    z := (m.M31 * v.X) + (m.M32 * v.Y) + (m.M33 * v.Z) + (m.M34 * v.W)
    w := (m.M41 * v.X) + (m.M42 * v.Y) + (m.M43 * v.Z) + (m.M44 * v.W)
	return NewVector4(x, y, z, w)
}

func (this Matrix) Invert() Matrix {
    num1 := this.M11
    num2 := this.M12
    num3 := this.M13
    num4 := this.M14
    num5 := this.M21
    num6 := this.M22
    num7 := this.M23
    num8 := this.M24
    num9 := this.M31
    num10 := this.M32
    num11 := this.M33
    num12 := this.M34
    num13 := this.M41
    num14 := this.M42
    num15 := this.M43
    num16 := this.M44
    num17 := num11*num16 - num12*num15
    num18 := num10*num16 - num12*num14
    num19 := num10*num15 - num11*num14
    num20 := num9*num16 - num12*num13
    num21 := num9*num15 - num11*num13
    num22 := num9*num14 - num10*num13
    num23 := num6*num17 - num7*num18 + num8*num19
    num24 := -(num5*num17 - num7*num20 + num8*num21)
    num25 := num5*num18 - num6*num20 + num8*num22
    num26 := -(num5*num19 - num6*num21 + num7*num22)
    num27 := 1.0 / (num1*num23 + num2*num24 + num3*num25 + num4*num26)
    num28 := num7*num16 - num8*num15
    num29 := num6*num16 - num8*num14
    num30 := num6*num15 - num7*num14
    num31 := num5*num16 - num8*num13
    num32 := num5*num15 - num7*num13
    num33 := num5*num14 - num6*num13
    num34 := num7*num12 - num8*num11
    num35 := num6*num12 - num8*num10
    num36 := num6*num11 - num7*num10
    num37 := num5*num12 - num8*num9
    num38 := num5*num11 - num7*num9
    num39 := num5*num10 - num6*num9

    m11 := num23 * num27
    m21 := num24 * num27
    m31 := num25 * num27
    m41 := num26 * num27

    m12 := -(num2*num17 - num3*num18 + num4*num19)*num27
    m22 := (num1*num17 - num3*num20 + num4*num21)*num27
    m32 := -(num1*num18 - num2*num20 + num4*num22)*num27
    m42 := (num1*num19 - num2*num21 + num3*num22)*num27

    m13 := (num2*num28 - num3*num29 + num4*num30)*num27
    m23 := -(num1*num28 - num3*num31 + num4*num32)*num27
    m33 := (num1*num29 - num2*num31 + num4*num33)*num27
    m43 := -(num1*num30 - num2*num32 + num3*num33)*num27

    m14 := -(num2*num34 - num3*num35 + num4*num36)*num27
    m24 := (num1*num34 - num3*num37 + num4*num38)*num27
    m34 := -(num1*num35 - num2*num37 + num4*num39)*num27
    m44 := (num1*num36 - num2*num38 + num3*num39)*num27

    return Matrix{ m11, m12, m13, m14,
           		   m21, m22, m23, m24,
           		   m31, m32, m33, m34,
           		   m41, m42, m43, m44 }
}
