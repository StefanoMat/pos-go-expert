package main

// T1
func main() {
	forever := make(chan bool) //Vazio

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	<-forever
}
