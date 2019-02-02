package nsqexample

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	nsq "github.com/nsqio/go-nsq"
)

func Producerflow() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	for i := 0; i < 10; i++ {
		w.Publish("Topic_string", []byte(fmt.Sprintf("string%d", i)))
	}

	jsonData := []string{}
	jsonData = append(jsonData, `
								{
								    "func_name":"BarFuncAdd",
								    "params":[
								        0.5,
								        0.51
								    ]
								}
								`)

	jsonData = append(jsonData, `
								{
								    "func_name":"FooFuncSwap",
								    "params":[
								        "a",
								        "b"
								    ]
								}
								`)

	jsonData = append(jsonData, `
								{
								    "func_name":"FooFuncSwap",
								    "params":[
								        1,
								        2
								    ]
								}
								`)

	jsonData = append(jsonData, `
								{
								    "func_name":"FakeMethod",
								    "params":[
								        "a",
								        "b"
								    ]
								}
								`)

	for _, j := range jsonData {
		w.Publish("Topic_json", []byte(j))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	//SIGINT	2	Term	用户发送INTR字符(Ctrl+C)触发
	//SIGTERM	15	Term	结束程序(可以被捕获、阻塞或忽略)
	<-c

	w.Stop()
}
