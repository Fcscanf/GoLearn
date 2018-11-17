package main

import (
	"fmt"
	"strconv"
	"time"
)

/*日期相关操作*/
func main() {

	now := time.Now()
	fmt.Printf("now=%v now type = %T", now, now)

	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	/*格式化日期*/
	fmt.Printf("当前年月日 %d-%d-%d-%d-%d-%d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	/*格式化日期*/
	fmt.Printf(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	/*定时任务-时间常量和休眠*/
	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠
		time.Sleep(time.Second)
		if i == 100 {
			break
		}
	}

	/*Unix和UnixNano的使用*/
	fmt.Printf("unix时间戳=%v ，unixnano时间戳=%v\n", now.Unix(), now.UnixNano())

	/*统计函数执行的时间*/
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("执行test方法耗费时间为%v秒\n", end-start)

}

func test() {
	str := ""
	for i := 0; i < 10000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
