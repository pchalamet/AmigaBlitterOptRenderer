package foundation

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

func IdentityMatrix() Matrix {
	return Matrix{1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0}
}

func ZeroMatrix() Matrix {
	return Matrix{}
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


// static member CreateLookAt (cameraPosition : Vector3) (cameraTarget : Vector3) (cameraUpVector : Vector3) =

// static member CreatePerspective (width : float) (height : float) (nearPlaneDistance : float) (farPlaneDistance : float) =
//     if nearPlaneDistance <= 0.0 then failwith "nearPlaneDistance <= 0"
//     if farPlaneDistance <= 0.0 then failwith "farPlaneDistance <= 0"
//     if (nearPlaneDistance >= farPlaneDistance) then failwith "nearPlaneDistance >= farPlaneDistance"

//     Matrix((2.0 * nearPlaneDistance) / width, 0.0, 0.0, 0.0,
//            0.0, (2.0 * nearPlaneDistance) / height, 0.0, 0.0,
//            0.0, 0.0, farPlaneDistance / (nearPlaneDistance - farPlaneDistance), -1.0,
//            0.0, 0.0, (nearPlaneDistance * farPlaneDistance) / (nearPlaneDistance - farPlaneDistance), 0.0)

// static member CreatePerspectiveOffCenter (left : float) (right : float) (bottom : float) (top : float) (nearPlaneDistance : float) (farPlaneDistance : float) =
//     if nearPlaneDistance <= 0.0 then failwith "nearPlaneDistance <= 0"
//     if farPlaneDistance <= 0.0 then failwith "farPlaneDistance <= 0"
//     if nearPlaneDistance >= farPlaneDistance then failwith "nearPlaneDistance >= farPlaneDistance"

//     let numW = (2.0 * nearPlaneDistance) / (right - left)
//     let numH = (2.0 * nearPlaneDistance) / (top - bottom)
//     Matrix(numW, 0.0, 0.0, 0.0,
//            0.0, numH, 0.0, 0.0,
//            (left + right) / (right - left), (top + bottom) / (top - bottom), (nearPlaneDistance + farPlaneDistance) / (nearPlaneDistance - farPlaneDistance), -1.0,
//            0.0, 0.0, (2.0 * nearPlaneDistance * farPlaneDistance) / (nearPlaneDistance - farPlaneDistance), 0.0)

// static member CreatePerspectiveFieldOfView (fieldOfView : float) (aspectRatio : float) (nearPlaneDistance : float) (farPlaneDistance : float) =
//     if fieldOfView <= 0.0 || fieldOfView >= 3.141593 then failwith "fieldOfView <= 0 or >= PI"
//     if nearPlaneDistance <= 0.0 then failwith "nearPlaneDistance <= 0"
//     if farPlaneDistance <= 0.0 then failwith "farPlaneDistance <= 0"
//     if nearPlaneDistance >= farPlaneDistance then failwith "nearPlaneDistance >= farPlaneDistance"
//     let top = nearPlaneDistance * tan (fieldOfView / 2.0)
//     let bottom = -top
//     Matrix.CreatePerspectiveOffCenter (bottom*aspectRatio) (top*aspectRatio) bottom top nearPlaneDistance farPlaneDistance

// static member CreateRotationX (radians : float) =
//     let val1 = cos radians
//     let val2 = sin radians
//     Matrix(1.0, 0.0, 0.0, 0.0,
//            0.0, val1, val2, 0.0,
//            0.0, -val2, val1, 0.0,
//            0.0, 0.0, 0.0, 1.0)

// static member CreateRotationY (radians : float) =
//     let val1 = cos radians
//     let val2 = sin radians
//     Matrix(val1, 0.0, -val2, 0.0,
//            0.0, 1.0, 0.0, 0.0,
//            val2, 0.0, val1, 0.0,
//            0.0, 0.0, 0.0, 1.0)

// static member CreateRotationZ (radians : float) =
//     let val1 = cos radians
//     let val2 = sin radians
//     Matrix(val1, val2, 0.0, 0.0,
//            -val2, val1, 0.0, 0.0,
//            0.0, 0.0, 1.0, 0.0,
//            0.0, 0.0, 0.0, 1.0)

// static member CreateScale (scaleX : float) (scaleY : float ) (scaleZ : float) =
//     Matrix(scaleX, 0.0, 0.0, 0.0,
//            0.0, scaleY, 0.0, 0.0,
//            0.0, 0.0, scaleZ, 0.0,
//            0.0, 0.0, 0.0, 1.0)

// static member CreateTranslation positionX positionY positionZ =
//     Matrix(1.0, 0.0, 0.0, 0.0,
//            0.0, 1.0, 0.0, 0.0,
//            0.0, 0.0, 1.0, 0.0,
//            positionX, positionY, positionZ, 1.0)

// static member inline (+) (m1 : Matrix, m2 : Matrix) =
//     Matrix(m1.M11 + m2.M11, m1.M12 + m2.M12, m1.M13 + m2.M13, m1.M14 + m2.M14,
//            m1.M21 + m2.M21, m1.M22 + m2.M22, m1.M23 + m2.M23, m1.M24 + m2.M24,
//            m1.M31 + m2.M31, m1.M32 + m2.M32, m1.M33 + m2.M33, m1.M34 + m2.M34,
//            m1.M41 + m2.M41, m1.M42 + m2.M42, m1.M43 + m2.M43, m1.M44 + m2.M44)

// static member inline (-) (m1 : Matrix, m2 : Matrix) =
//     Matrix(m1.M11 - m2.M11, m1.M12 - m2.M12, m1.M13 - m2.M13, m1.M14 - m2.M14,
//            m1.M21 - m2.M21, m1.M22 - m2.M22, m1.M23 - m2.M23, m1.M24 - m2.M24,
//            m1.M31 - m2.M31, m1.M32 - m2.M32, m1.M33 - m2.M33, m1.M34 - m2.M34,
//            m1.M41 - m2.M41, m1.M42 - m2.M42, m1.M43 - m2.M43, m1.M44 - m2.M44)

// static member inline (~-) (m : Matrix) =
//     Matrix(-m.M11, -m.M12, -m.M13, -m.M14,
//            -m.M21, -m.M22, -m.M23, -m.M24,
//            -m.M31, -m.M32, -m.M33, -m.M34,
//            -m.M41, -m.M42, -m.M43, -m.M44)

// static member inline (*) (m1 : Matrix, m2 : Matrix) =
//     let m11 = (m1.M11 * m2.M11) + (m1.M12 * m2.M21) + (m1.M13 * m2.M31) + (m1.M14 * m2.M41)
//     let m12 = (m1.M11 * m2.M12) + (m1.M12 * m2.M22) + (m1.M13 * m2.M32) + (m1.M14 * m2.M42)
//     let m13 = (m1.M11 * m2.M13) + (m1.M12 * m2.M23) + (m1.M13 * m2.M33) + (m1.M14 * m2.M43)
//     let m14 = (m1.M11 * m2.M14) + (m1.M12 * m2.M24) + (m1.M13 * m2.M34) + (m1.M14 * m2.M44)
//     let m21 = (m1.M21 * m2.M11) + (m1.M22 * m2.M21) + (m1.M23 * m2.M31) + (m1.M24 * m2.M41)
//     let m22 = (m1.M21 * m2.M12) + (m1.M22 * m2.M22) + (m1.M23 * m2.M32) + (m1.M24 * m2.M42)
//     let m23 = (m1.M21 * m2.M13) + (m1.M22 * m2.M23) + (m1.M23 * m2.M33) + (m1.M24 * m2.M43)
//     let m24 = (m1.M21 * m2.M14) + (m1.M22 * m2.M24) + (m1.M23 * m2.M34) + (m1.M24 * m2.M44)
//     let m31 = (m1.M31 * m2.M11) + (m1.M32 * m2.M21) + (m1.M33 * m2.M31) + (m1.M34 * m2.M41)
//     let m32 = (m1.M31 * m2.M12) + (m1.M32 * m2.M22) + (m1.M33 * m2.M32) + (m1.M34 * m2.M42)
//     let m33 = (m1.M31 * m2.M13) + (m1.M32 * m2.M23) + (m1.M33 * m2.M33) + (m1.M34 * m2.M43)
//     let m34 = (m1.M31 * m2.M14) + (m1.M32 * m2.M24) + (m1.M33 * m2.M34) + (m1.M34 * m2.M44)
//     let m41 = (m1.M41 * m2.M11) + (m1.M42 * m2.M21) + (m1.M43 * m2.M31) + (m1.M44 * m2.M41)
//     let m42 = (m1.M41 * m2.M12) + (m1.M42 * m2.M22) + (m1.M43 * m2.M32) + (m1.M44 * m2.M42)
//     let m43 = (m1.M41 * m2.M13) + (m1.M42 * m2.M23) + (m1.M43 * m2.M33) + (m1.M44 * m2.M43)
//     let m44 = (m1.M41 * m2.M14) + (m1.M42 * m2.M24) + (m1.M43 * m2.M34) + (m1.M44 * m2.M44)
//     Matrix(m11, m12, m13, m14,
//            m21, m22, m23, m24,
//            m31, m32, m33, m34,
//            m41, m42, m43, m44)

// member this.Transpose () =
//     Matrix(this.M11, this.M21, this.M31, this.M41,
//            this.M12, this.M22, this.M32, this.M42,
//            this.M13, this.M23, this.M33, this.M43,
//            this.M14, this.M24, this.M34, this.M44)

// static member inline (*) (v : Vector3, m : Matrix) =
//     let x = (v.X * m.M11) + (v.Y * m.M21) + (v.Z * m.M31) + m.M41
//     let y = (v.X * m.M12) + (v.Y * m.M22) + (v.Z * m.M32) + m.M42
//     let z = (v.X * m.M13) + (v.Y * m.M23) + (v.Z * m.M33) + m.M43
//     Vector3.create(x, y, z)

// static member inline (*) (v : Vector4, m : Matrix) =
//     let x = (v.X * m.M11) + (v.Y * m.M21) + (v.Z * m.M31) + (v.W * m.M41)
//     let y = (v.X * m.M12) + (v.Y * m.M22) + (v.Z * m.M32) + (v.W * m.M42)
//     let z = (v.X * m.M13) + (v.Y * m.M23) + (v.Z * m.M33) + (v.W * m.M43)
//     let w = (v.X * m.M14) + (v.Y * m.M24) + (v.Z * m.M34) + (v.W * m.M44)
//     Vector4.create(x, y, z, w)

// static member inline (*) (m : Matrix, v : Vector4) =
//     let x = (m.M11 * v.X) + (m.M12 * v.Y) + (m.M13 * v.Z) + (m.M14 * v.W)
//     let y = (m.M21 * v.X) + (m.M22 * v.Y) + (m.M23 * v.Z) + (m.M24 * v.W)
//     let z = (m.M31 * v.X) + (m.M32 * v.Y) + (m.M33 * v.Z) + (m.M34 * v.W)
//     let w = (m.M41 * v.X) + (m.M42 * v.Y) + (m.M43 * v.Z) + (m.M44 * v.W)
//     Vector4.create(x, y, z, w)

// member this.Invert() =
//     let num1 = this.M11
//     let num2 = this.M12
//     let num3 = this.M13
//     let num4 = this.M14
//     let num5 = this.M21
//     let num6 = this.M22
//     let num7 = this.M23
//     let num8 = this.M24
//     let num9 = this.M31
//     let num10 = this.M32
//     let num11 = this.M33
//     let num12 = this.M34
//     let num13 = this.M41
//     let num14 = this.M42
//     let num15 = this.M43
//     let num16 = this.M44
//     let num17 = num11*num16 - num12*num15
//     let num18 = num10*num16 - num12*num14
//     let num19 = num10*num15 - num11*num14
//     let num20 = num9*num16 - num12*num13
//     let num21 = num9*num15 - num11*num13
//     let num22 = num9*num14 - num10*num13
//     let num23 = num6*num17 - num7*num18 + num8*num19
//     let num24 = -(num5*num17 - num7*num20 + num8*num21)
//     let num25 = num5*num18 - num6*num20 + num8*num22
//     let num26 = -(num5*num19 - num6*num21 + num7*num22)
//     let num27 = 1.0 / (num1*num23 + num2*num24 + num3*num25 + num4*num26)
//     let num28 = num7*num16 - num8*num15
//     let num29 = num6*num16 - num8*num14
//     let num30 = num6*num15 - num7*num14
//     let num31 = num5*num16 - num8*num13
//     let num32 = num5*num15 - num7*num13
//     let num33 = num5*num14 - num6*num13
//     let num34 = num7*num12 - num8*num11
//     let num35 = num6*num12 - num8*num10
//     let num36 = num6*num11 - num7*num10
//     let num37 = num5*num12 - num8*num9
//     let num38 = num5*num11 - num7*num9
//     let num39 = num5*num10 - num6*num9

//     let m11 = num23 * num27
//     let m21 = num24 * num27
//     let m31 = num25 * num27
//     let m41 = num26 * num27

//     let m12 = -(num2*num17 - num3*num18 + num4*num19)*num27
//     let m22 = (num1*num17 - num3*num20 + num4*num21)*num27
//     let m32 = -(num1*num18 - num2*num20 + num4*num22)*num27
//     let m42 = (num1*num19 - num2*num21 + num3*num22)*num27

//     let m13 = (num2*num28 - num3*num29 + num4*num30)*num27
//     let m23 = -(num1*num28 - num3*num31 + num4*num32)*num27
//     let m33 = (num1*num29 - num2*num31 + num4*num33)*num27
//     let m43 = -(num1*num30 - num2*num32 + num3*num33)*num27

//     let m14 = -(num2*num34 - num3*num35 + num4*num36)*num27
//     let m24 = (num1*num34 - num3*num37 + num4*num38)*num27
//     let m34 = -(num1*num35 - num2*num37 + num4*num39)*num27
//     let m44 = (num1*num36 - num2*num38 + num3*num39)*num27

//     Matrix(m11, m12, m13, m14,
//            m21, m22, m23, m24,
//            m31, m32, m33, m34,
//            m41, m42, m43, m44)
