package math
import (
    "math/rand"
    "time"
)

/*
产生随机数播下随机种子
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

/*产生公共随机数
min最小值
max最大值
返回随机数int类型
例子:Rands(100,200)
*/
func Rands(min, max int) int {
    if min > max {
        panic("the min is greater than max!")
    }
    return rand.Intn(max -min + 1) + min
}