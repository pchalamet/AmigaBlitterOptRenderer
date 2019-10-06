package foundation

func equalWithUncertainty(value, expected float64) bool {
	return value - expected < 0.001
}
