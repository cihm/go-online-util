package timing

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

//https://github.com/lwl1989/timing
//https://www.freeformatter.com/epoch-timestamp-to-date-converter.html
func Flow1() {
	scheduler := NewScheduler()

	scheduler.Start()
	// fmt.Println(time.Now().Unix())
	// scheduler.AddFunc(time.Now().Unix()+10, func() {
	// 	fmt.Println("one second after")
	// })

	scheduler.AddTask(&Task{
		Job: FuncJob(func() {
			fmt.Println("hello task2")
		}),
		RunTime: 1549114440,
	})

	scheduler.AddTask(&Task{
		Job: FuncJob(func() {
			fmt.Println("hello task3")
		}),
		RunTime: 1549114500,
	})

	var Input string
	var String string
	var Number int
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	for {
		fmt.Print("请输入一些字符串>")
		Input, _ = f.ReadString('\n') //定义一行输入的内容分隔符。
		if len(Input) == 1 {
			continue //如果用户输入的是一个空行就让用户继续输入。
		}
		fmt.Printf("您输入的是:%s", Input)
		fmt.Sscan(Input, &String, &Number) //将Input
		if String == "stop" {
			break
		}
		itmeee, _ := strconv.Atoi(String)
		fmt.Println("itmeee:", itmeee)
		go newAction(itmeee)

		fmt.Printf(Input, "!!!", String, "!!!", Number)
	}

	//fmt.Scanln()
}

func newAction(se int) {
	fmt.Println("newAction:", se)
	done := make(chan bool)
	time.AfterFunc(time.Second*time.Duration(se), func() {
		fmt.Println("hello", se)
		done <- true
	})
	<-done
}
