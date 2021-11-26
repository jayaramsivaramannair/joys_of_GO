package main

import (
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

//Structs are analogous to classes in other object oriented languages
type car struct {
	gas_pedal uint16
	brake_pedal uint16 
	steering_wheel int16
	top_speed_kmh float64
}

//Value receivers method makes a copy of the original struct which is very time consuming when the struct is large

//This method is available for call on the car struct - A Value receivers method
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax /kmh_multiple)
}

//This is a pointer receivers method which is used to modify the data inside a struct
func(c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func newer_top_speed(c car, speed float64) car {
	c.top_speed_kmh = speed
	return c
}

func main() {
	a_car := car {gas_pedal: 65000, brake_pedal: 0, steering_wheel: 12561, top_speed_kmh: 225.0}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())


}