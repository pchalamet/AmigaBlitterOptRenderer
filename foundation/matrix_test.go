package foundation

import "fmt"
import "testing"
import "math"



func (matrix Matrix) printMatrix (msg string) {
	fmt.Println(msg)
	fmt.Println("|", matrix.Item2(0, 0), matrix.Item2(0, 1), matrix.Item2(0, 2), matrix.Item2(0, 3), "|")
	fmt.Println("|", matrix.Item2(1, 0), matrix.Item2(1, 1), matrix.Item2(1, 2), matrix.Item2(1, 3)  , "|")
	fmt.Println("|", matrix.Item2(2, 0), matrix.Item2(2, 1), matrix.Item2(2, 2), matrix.Item2(2, 3), "|")
	fmt.Println("|", matrix.Item2(3, 0), matrix.Item2(3, 1), matrix.Item2(3, 2), matrix.Item2(3, 3), "|")
}


func TestMatrixAccess(t *testing.T) {
    // expected matrix
    // | 1  0     0      0 | t     | 1  0      0     0 |
    // | 0  1/2   -V3/2  0 |       | 0  1/2    V3/2  0 |
    // | 0  V3/2  1/2    0 |  ==>  | 0  -V3/2  1/2   0 |
    // | 0  0     0      1 |       | 0  0      0     1 |
	matrix := CreateRotationX(MathConsts.Pi/3.0)
	matrix.printMatrix("Access")

	if ! equalWithUncertainty(matrix.M11, 1.0) {
		t.Error("Value for term M11 is unexpected")
	}

	if ! equalWithUncertainty(matrix.M22, 0.5) {
		t.Error("Value for term M22 is unexpected")
	}

	if ! equalWithUncertainty(matrix.M31, 0.0) {
		t.Error("Value for term M31 is unexpected")
	}

	if ! equalWithUncertainty(matrix.M32, math.Sqrt(3.0)/2.0) {
		t.Error("Value for term M32 is unexpected")
	}

	if ! equalWithUncertainty(matrix.Item2(2, 1), -math.Sqrt(3.0)/2.0) {
		t.Error("Value for term M21 is unexpected")
	}
}


func TestMatrixComposition(t *testing.T) {
    // expected matrix:
    // | 1/2    0     V3/2   0 | t     | 1/2   3/4    -V3/4  0 |
    // | 3/4    1/2   -V3/4  0 |       | 0     1/2    V3/2   0 |
    // | -V3/4  V3/2  1/4    0 |  ==>  | V3/2  -V3/4  1/4    0 |
    // | 0      0     0      1 |       | 0      0     0      1 |
	matrix1 := CreateRotationX(MathConsts.Pi / 3.0)
	matrix2 := CreateRotationY(MathConsts.Pi / 3.0)
	matrix := matrix2.Compose(matrix1)
	matrix.printMatrix("Composition")

	if ! equalWithUncertainty(matrix.M11 , 0.5) {
		t.Error("Value for term M11 is unexpected")
	}

	if ! equalWithUncertainty(matrix.M12 , 0.75) {
		t.Error("Value for term M12 is unexpected")
	}

	if ! equalWithUncertainty(matrix.M13 , -math.Sqrt(3.0) / 4.0) {
		t.Error("Value for term M13 is unexpected")
	}

	if ! equalWithUncertainty(matrix.M32 , -math.Sqrt(3.0) / 4.0) {
		t.Error("Value for term M32 is unexpected")
	}
}


func TestMatrixNeg(t *testing.T) {
	matrix1 := CreateRotationX(MathConsts.Pi / 3.0)
	matrix2 := matrix1.Neg()
	matrix3 := matrix2.Neg()

	for i:=0; i<4; i++ {
		for j:=0; j<4; j++ {
			if matrix3.Item2(i,j) != matrix3.Item2(i,j) {
				t.Errorf("Value for term M%d%d is unexpected", i,j)
			}
		}
	}
}


func TestMatrixInvert(t *testing.T) {
    matrix := CreateRotationX(MathConsts.Pi / 3.0)
    matrixInverted := matrix.Invert()
	matrixRes := matrix.Compose(matrixInverted)
	matrixRes.printMatrix("Invert")

	for i:=0; i<4; i++ {
		for j:=0; j<4; j++ {
			if i == j {
				if !equalWithUncertainty(matrixRes.Item2(i,j), 1.0) {
					t.Errorf("Value for term M%d%d is unexpected", i,j)
				}
			} else if !equalWithUncertainty(matrixRes.Item2(i,j), 0.0)  {
				t.Errorf("Value for term M%d%d is unexpected", i,j)
			}
		}
	}
}


func TestVectorTransform(t *testing.T) {
    // expected result:
    // | 0    |
    // | 1/2  |
    // | V3/2 |
    // | 1    |
    matrix := CreateRotationX(MathConsts.Pi / 3.0)
    vector := NewVector4(0.0, 1.0, 0.0, 1.0)
	result := vector.Mult(matrix)
	
	if !equalWithUncertainty(result.X, 0.0) {
		t.Errorf("Value for Vector.X is unexpected")
	}

	if !equalWithUncertainty(result.Y, 0.5) {
		t.Errorf("Value for Vector.Y is unexpected")
	}

	if !equalWithUncertainty(result.Z, math.Sqrt(3.0) / 2.0) {
		t.Errorf("Value for Vector.Z is unexpected")
	}

	if !equalWithUncertainty(result.W, 1.0) {
		t.Errorf("Value for Vector.Z is unexpected")
	}
}
