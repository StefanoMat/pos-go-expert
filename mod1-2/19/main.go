package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "3"}
	for k, v := range numeros {
		println(k, v)
	}

	for {
		println("Infinite Hello World!")
	}
}
