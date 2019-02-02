package listener

import (
	"bufio"
	"fmt"
	"go-online-util/listener/event"
	"os"
)

//https://github.com/pocke/goevent
var (
	String string
	Number int
	Input  string
)

func EventFlow() {
	table := event.NewTable()
	table.On("foo", func(i int) {
		fmt.Printf("foo: %d\n", i)
	})
	table.On("bar", func(s string) {
		fmt.Printf("bar: %s\n", s)
	})

	var eRR error
	eRR = table.Trigger("foo", 1)
	eRR = table.Trigger("bar", "hoge")
	eRR = table.Trigger("bar", 38) // retrun error
	fmt.Println(eRR)

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

		eRR = table.Trigger("bar", Input)
		fmt.Println(eRR)
		fmt.Printf(Input, "!!!", String, "!!!", Number)
	}
}
