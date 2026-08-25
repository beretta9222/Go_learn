package main

import "fmt"

type People struct {
	speed     int
	disctance int
}

type Car struct {
	speed      int
	disctance  int
	speedlimit int
}
type Moveble interface {
	speedUp(v int)
	speedDown(v int)
}

func (p *People) speedUp(v int) {
	p.speed++
}
func (p *People) speedDown(v int) {
	if p.speed < v {
		p.speed = 0
	} else {
		p.speed -= v
	}
}

func (c *Car) speedUp(v int) {
	if c.speed+v > c.speedlimit {
		c.speed = c.speedlimit
	} else {
		c.speed += v
	}
}
func (c *Car) speedDown(v int) {
	if c.speed < v {
		c.speed = 0
	} else {
		c.speed -= v
	}
}

func main() {
	var people Moveble = &People{}
	var car Moveble = &Car{0, 0, 250}

	people.speedUp(2)
	fmt.Println(people)

	car.speedUp(220)
	fmt.Println(car)
	car.speedUp(50)
	fmt.Println(car)
}
