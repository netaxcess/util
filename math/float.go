package math

import (
	"math"
)

/*
对浮点数进行四舍五入
例子:Round(1.99)
返回：2
*/
func Round(val float64) float64 {
	return math.Floor(val + 0.5)
}

/*
舍去法取整，返回不大于 value 的最接近的整数，将 value 的小数部分舍去取整。floor() 返回的类型仍然是 float，因为 float 值的范围通常比 integer 要大。
例子:Floor(9.999)返回9， Floor(-3.14)返回-4
*/
func Floor(val float64) float64 {
	return math.Floor(val)
}

/*
进一法取整，返回不小于 value 的下一个整数，value 如果有小数部分则进一位。
例子:Ceil(4.3)返回5
例子:Ceil(9.9)返回10
例子:Ceil(-3.14)返回-3

*/
func Ceil(val float64) float64 {
	return math.Ceil(val)
}