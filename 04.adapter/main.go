package main

import "fmt"

type IConnect interface {
	Conn()
}

type ChinaConnect struct {
	name string
}

func (c ChinaConnect) Conn() {
	fmt.Println(c.name + "充电中")
}

type HKConnect struct {
	name string
}

func (c HKConnect) ConnInHK() {
	fmt.Println(c.name + "用港版插头充电中")
}

type HKConnectAdapter struct {
	HKConnect
}

func (c HKConnectAdapter) Conn() {
	c.HKConnect.ConnInHK()
}

func Charge(c IConnect) {
	c.Conn()
}

func main() {
	chinaConn := ChinaConnect{"国行"}
	Charge(chinaConn)

	hkConn := HKConnect{"港行"}
	hkConnAdp := HKConnectAdapter{hkConn}
	Charge(hkConnAdp)
}
