package timing

import (
	"fmt"
	"time"
)

//https://github.com/lwl1989/timing
func Flow1() {
	scheduler := NewScheduler()

	scheduler.Start()
	fmt.Println(time.Now().Unix())
	scheduler.AddFunc(time.Now().Unix()+10, func() {
		fmt.Println("one second after")
	})

	scheduler.AddTask(&Task{
		Job: FuncJob(func() {
			fmt.Println("hello task2")
		}),
		RunTime: 1549083300,
	})

	fmt.Scanln()
}
