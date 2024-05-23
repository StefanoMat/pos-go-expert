package main

const a = "Hello, World!"

type ID int

var (
	b bool
	c int
	d string = "Y"
	f ID     = 1
)

func main() {
	a := "X"
	println(a)
	println(b)
	println(c)
	println(d)
	println(f)
}
