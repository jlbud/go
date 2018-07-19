package main

import (
	"fmt"
	"unsafe"
	"text/template"
)

type NodeInfo struct {
	Domain     string
	Nodemode   string
	IsValidity bool
}

type MyNodeInfo struct {
	Domain     string
	Nodemode   string
	IsValidity bool
}

func main() {
	nodeInfo := &NodeInfo{
		"123456",
		"123456",
		true,
	}
	//NodeInfo和MyNodeInfo之间进行强制类型转换
	myNodeInfo := &MyNodeInfo{
		"1",
		"1",
		false,
	}
	myNodeInfo = (*MyNodeInfo)(unsafe.Pointer(nodeInfo))
	fmt.Println(myNodeInfo)    //说明已经从NodeInfo转为MyNodeInfo类型了
	myNodeInfo = &MyNodeInfo{} //说明已经从NodeInfo转为MyNodeInfo类型了
	fmt.Println(myNodeInfo)
}

// MyTemplate 定义和 template.Template 只是形似
type MyTemplate struct {
	name       string
	parseTree  *unsafe.Pointer
	common     *unsafe.Pointer
	leftDelim  string
	rightDelim string
}

func main() {
	t := template.New("Foo")
	p := (*MyTemplate)(unsafe.Pointer(t))
	p.name = "Bar" // 关键在这里，突破私有成员
	fmt.Println(p, t)
}