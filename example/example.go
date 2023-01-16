package main

import (
	"fmt"
	"github.com/slclub/goevents"
)

func main() {
	ev1 := func(args ...interface{}) {
		fmt.Println("ev1----", args)
		//ev1 do something.
	}
	ev2 := func(args ...interface{}) {
		//ev2 do something.
		fmt.Println("ev2---", args)
	}
	events := goevents.Classic()
	//连贯写法 绑定事件及参数
	events.Bind("123", "dfd").On("a", ev1)

	//绑定事件和参数分开写
	events.Bind("event2 serial running")
	events.On("a", ev2)

	//Trigger
	//events.Bind(...)//可以在这里绑定全局参数，当有事件绑定参数，这里可以重新绑定，只是不能针对所有串行事件
	//按模块名 执行事件队列。无参数 执行所有串行 事件
	events.Trigger("a")
	//触发所有事件，串行，并行的全部触发
	//events.Emit()
}
