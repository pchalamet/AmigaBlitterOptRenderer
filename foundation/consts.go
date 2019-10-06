package foundation

type vector3Consts struct {
    Zero Vector3
    One Vector3
    UnitX Vector3
    UnitY Vector3
    UnitZ Vector3
    Up Vector3
    Down Vector3
    Right Vector3
    Left Vector3
    Forward Vector3
    Backward Vector3
}

var Vector3Consts = vector3Consts {
    Zero: NewVector3(0.0, 0.0, 0.0),
    One: NewVector3(1.0, 1.0, 1.0),
    UnitX: NewVector3(1.0, 0.0, 0.0),
    UnitY: NewVector3(0.0, 1.0, 0.0),
    UnitZ: NewVector3(0.0, 0.0, 1.0),
    Up: NewVector3(0.0, 1.0, 0.0),
    Down: NewVector3(0.0, -1.0, 0.0),
    Right: NewVector3(1.0, 0.0, 0.0),
    Left: NewVector3(-1.0, 0.0, 0.0),
    Forward: NewVector3(0.0, 0.0, -1.0),
    Backward: NewVector3(0.0, 0.0, 1.0),
}



type matrixConsts struct {
	Zero Matrix
	Identity Matrix
}

var MatrixConsts = matrixConsts {
	Zero: Matrix{},
	Identity: Matrix{ 1.0, 0.0, 0.0, 0.0,
					  0.0, 1.0, 0.0, 0.0,
					  0.0, 0.0, 1.0, 0.0,
					  0.0, 0.0, 0.0, 1.0 },
}

