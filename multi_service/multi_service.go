package multi_service

import (
	"fmt"
	"sync"
)

type MultiService struct{
	waitGroup sync.WaitGroup
}

type IHandler interface{
	Handler()
}

func (pMultiSvr* MultiService) runHandler(handler IHandler) {
	defer pMultiSvr.waitGroup.Done()
	fmt.Println("run handler")
	handler.Handler()
}

func hello(){
	fmt.Println("ccccc")
}

func (pMultiSvr* MultiService)AddService(handler IHandler) error{
	pMultiSvr.waitGroup.Add(1)
	fmt.Println("before run handler")
	go pMultiSvr.runHandler(handler)
	fmt.Println("after run handler")
	return nil
}