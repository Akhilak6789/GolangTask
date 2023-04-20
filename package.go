package caluclator

import ( "math"
)

func Add(a int, b int) int{
	return a+b
}
func Sub(a int, b int) int {
	return a-b
}
func Div(a int, b int) int {
	return a/b
}
func Mul(a int, b int) int {
	return a*b
}
func Rem(a int, b int) int {
	return a%b
}
func Sqrt(a float64) float64 {
	return math.Sqrt(a)
}
func Sine(a float64) float64 {
	return math.Sin(a)
}
func Cos(a float64) float64 {
	return math.Cos(a)
}