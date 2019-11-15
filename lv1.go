package main
import(
	"fmt"
	"strconv"
	"time"
)
func main(){
	inputs :=make([]int,0,10)
	var t string
	fmt.Println("请输入时间戳：")

	for{
		_, _ = fmt.Scan(&t)
		if t =="result" {
			fmt.Println("the result are :")
			for _, i := range inputs {
			fmt.Println(time.Unix(int64(i), 1))
			}
			break
		}
		n,_:=strconv.Atoi(t)
		inputs= append(inputs,n )
		fmt.Println("input ok!")
	}
}