//there is no code for the classes instead go uses structs
//there can be two types of methods in  go
// 1 value receivers
// 2 pointer receivers

package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.6094

type car struct {
	gas_pedal      uint16 //min 0 max 65535
	brake_pedal    uint16
	steering_wheel int16
	top_Speed_kmh  float64
}

func (c car) kmh() float64 {
	return (float64(c.gas_pedal) * (c.top_Speed_kmh / usixteenbitmax))
}
func (c car) mph() float64 {
	return (float64(c.gas_pedal) * (c.top_Speed_kmh / usixteenbitmax / kmh_multiple))
}

//chage the top speed of the cae
func (c *car) new_topspeed(newspeed float64) {
	c.top_Speed_kmh = newspeed
}
func main() {
	a_car := car{gas_pedal: 2212,
		brake_pedal:    0,
		steering_wheel: 1253,
		top_Speed_kmh:  225.0}
	/*
	 this could have a short_hand also
	 b_car := car{1252,5552,552, 552}

	 but probably this is a worse digest of the good old formulae
	*/
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	a_car.new_topspeed(234)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

}
