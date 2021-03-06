package main

import (
	//"encoding/json"
	"fmt"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/event"
	"log"
)

func init() {
	pingpp.Key = "sk_test_ibbTe5jLGCi5rzfH4OqPW9KC"
	fmt.Println("Go SDK Version:", pingpp.Version())
	pingpp.AcceptLanguage = "zh-CN"
}

func ExampleEvent_get() {
	eve, err := event.Get("evt_zRFRk6ekazsH7t7yCqEeovhk")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(eve)
}

func ExampleEvent_list() {
	params := &pingpp.EventListParams{}
	params.Filters.AddFilter("type", "", "charge.succeeded")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	i := event.List(params)
	for i.Next() {
		c := i.Event()
		fmt.Println(c)
	}
}

func main() {
	ExampleEvent_get()
}
