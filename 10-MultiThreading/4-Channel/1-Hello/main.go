package main

// T1
func main() {
	canal := make(chan string) //Vazio

	// T2
	go func() {
		canal <- "Olá Canal!" //Cheia
	}()

	msg := <-canal //Esvazia aqui
	println(msg)

}
