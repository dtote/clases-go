package main

import "fmt"

type Day uint
type Second uint
type Minute uint
type Hour uint

// Function receivers

func (d Day) ToHours() Hour {
	return Hour(d * 24)
}

func (h Hour) ToMinutes() Minute {
	return Minute(h * 60)
}

func (m Minute) ToSeconds() Second {
	return Second(m * 60)
}

func main() {
	var d Day = 5

	fmt.Printf("\n5 dias son %d horas", d.ToHours())
	fmt.Printf("\n5 dias son %d minutos", d.ToHours().ToMinutes())
	fmt.Printf("\n5 dias son %d segundos", d.ToHours().ToMinutes().ToSeconds())
}
