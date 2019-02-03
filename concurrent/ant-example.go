package concurrent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func AntFlow() {
	defer ants.Release()

	runTimes := 10

	// // Use the common pool
	var wg sync.WaitGroup
	// syncCalculateSum := func() {
	// 	demoFunc()
	// 	wg.Done()
	// }
	// for i := 0; i < runTimes; i++ {
	// 	wg.Add(1)
	// 	ants.Submit(syncCalculateSum)
	// }
	// wg.Wait()
	// fmt.Printf("running goroutines: %d\n", ants.Running())
	// fmt.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the size of goroutine pool and 1 second for expired duration

	setNewTimerFunc := func(i interface{}) {
		myFunc(i)
		wg.Done()
	}
	p, _ := ants.NewPoolWithFunc(10, setNewTimerFunc)
	defer p.Release()
	// Submit tasks
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		p.Serve(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)

	var Input string
	var String string
	var Number int
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	for {
		fmt.Println("请输入一些字符串>")
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
		wg.Add(1)
		p.Serve(int32(itmeee))

		fmt.Printf("running goroutines: %d\n", p.Running())
		fmt.Printf("finish all tasks, result is %d\n", sum)

	}

}
