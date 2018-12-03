package main

import (
	"fmt"
)

import (
	"time"
	"github.com/afeng0007/studygolang/multi_service"
)

type MainContext struct{
	i string
}

func (ctx MainContext)Handler() {
	fmt.Println(ctx.i)
}

func main(){
	multiService := multi_service.MultiService{}
	ctx := MainContext{
		i: "aaaa",
	}

	fmt.Println("bbbbbb")
	multiService.AddService(ctx)
	time.Sleep(time.Second*10)
}