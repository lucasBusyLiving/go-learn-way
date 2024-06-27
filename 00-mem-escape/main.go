package main

type Car struct {
	Name string
}

func main() {
	incr := incrementer()
	x := incr()
	x++
}

func Speak() Car {
	bigCar := Car{"Big Car"}
	return bigCar
}

func incrementer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
