package parser

func Integrate(e Expression, vars map[string]float64, wrt string, from, to float64) float64 {
	return IntegrateV(e, vars, wrt, from, to, 100)
}
func IntegrateV(e Expression, vars map[string]float64, wrt string, from, to float64, NumBins int) float64 {
	originalWRT := vars[wrt]
	dx := (to - from) / float64(NumBins)
	sum := 0.0
	for i := 0; i < NumBins; i++ {
		xi := float64(i) * dx
		vars[wrt] = xi
		yi := e.Evaluate(vars)
		xf := float64(i+1) * dx
		vars[wrt] = xf
		yf := e.Evaluate(vars)
		//trapezoidal sum
		area := dx * ((yi + yf) / 2)
		sum += area
	}
	vars[wrt] = originalWRT
	return sum
}
