package main

func main() {
	var (
		a bool = false
		b bool = false
		c bool = false
	)
	if a && b || !c {
		println("true")
	} else {
		println("false")
	}
}
