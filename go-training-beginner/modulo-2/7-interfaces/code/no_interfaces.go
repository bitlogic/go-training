package main

import "fmt"

type mobile struct {
	brand string
}

func (m mobile) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d\n", m, m.brand, power)
}

type laptop struct {
	cpu string
}

func (l laptop) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d\n", l, l.cpu, power)
}

type toaster struct {
	amount int
}

func (t toaster) Draw(power int) {
	fmt.Printf("%T -> brand: %v, power: %d\n", t, t.amount, power)
}

type kettle struct {
	quantity string
}

func (k kettle) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d\n", k, k.quantity, power)
}

type socket struct{}

func (socket) Plug(device mobile, power int) {
	device.Draw(power)
}

func (socket) PlugLaptop(device laptop, power int) {
	device.Draw(power)
}

func (socket) PlugToaster(device toaster, power int) {
	device.Draw(power)
}

func (socket) PlugKettle(device kettle, power int) {
	device.Draw(power)
}

func main() {
	m := mobile{"Apple"}
	l := laptop{"Intel i7"}
	t := toaster{4}
	k := kettle{"50%"}

	s := socket{}

	s.Plug(m, 10)
	s.PlugLaptop(l, 50)
	s.PlugToaster(t, 30)
	s.PlugKettle(k, 25)
}
