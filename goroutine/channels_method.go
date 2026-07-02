package main

func Unbuffered_Channels_method() {
	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	// result := <-ch

	// fmt.Println("Final returned value is ", result)
}
