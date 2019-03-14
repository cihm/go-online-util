package designpattern

import (
	"fmt"
)

type Account interface {
	Query(id string) int
	Update(id string, value int)
}

type AccountImpl struct {
	Id    string
	Name  string
	Value int
}

func (a *AccountImpl) Query(_ string) int {
	fmt.Println("AccountImpl.Query")
	return 100
}

func (a *AccountImpl) Update(_ string, _ int) {
	fmt.Println("AccountImpl.Update")
}

// var New = func(id, name string, value int) Account {
// 	return &AccountImpl{id, name, value}
// }

type Proxy interface {
	PQuery(context, id string) int
	PUpdate(context, id string, value int)
}
type ProxyImpl struct {
	Account AccountImpl
}

// var PNew = func(Account AccountImpl) Proxy {
// 	return &ProxyImpl{Account}
// }

func (p *ProxyImpl) PQuery(context, id string) int {
	fmt.Println("Proxy.Query begin:" + context)
	value := p.Account.Query(id)
	fmt.Println("Proxy.Query end")
	return value
}

func (p *ProxyImpl) PUpdate(context, id string, value int) {
	fmt.Println("Proxy.Update begin" + context)
	fmt.Println("Proxy.Update begin")
	p.Account.Update(id, value)
	fmt.Println("Proxy.Update end")
}

var PPNew func(d, name string, value int) Proxy

func Proxyd() {
	id := "100111"
	a := PPNew(id, "ZhangSan", 100)
	a.PQuery("context1", id)
	a.PUpdate("context1-1", id, 500)
}

func init() {
	PPNew = func(id, name string, value int) Proxy {
		a := &AccountImpl{id, name, value}
		p := &ProxyImpl{*a}
		return p
	}
}
