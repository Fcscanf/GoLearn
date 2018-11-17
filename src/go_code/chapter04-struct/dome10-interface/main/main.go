package main

import "fmt"

/*
接口不能有任何变量
接口方法都没有任何方法体
引用类型
*/

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

func (p Phone) Call() {
	fmt.Println("打电话...")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	/*类型断言*/
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)

	var usbArr [3]Usb
	usbArr[0] = Phone{"VIVO"}
	usbArr[1] = Phone{"OPPO"}
	usbArr[2] = Phone{"索尼"}

	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
}
