package main

import (
	"fmt"
)

func main() {

	// ch := make(chan int) // unbuffered

	// go func() {
	// 	fmt.Println("Goroutine: trying to send 10")
	// 	ch <- 10 // <- unbuffered send
	// 	fmt.Println("Goroutine: send completed")
	// }()
	// // v := <-ch
	// // fmt.Println(v)

	// fmt.Println("main")

	// errors.New("Stop")

}

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 10
// 		ch <- 20
// 		ch <- 30

// 	}()
// 	for c := range ch {
// 		fmt.Println(c)
// 	}
// }

// func print() {
// 	fmt.Println("C")
// }

// func main() {
// 	var wait sync.WaitGroup
// 	wait.Add(3)
// 	go func() {
// 		fmt.Println("A")
// 		wait.Done()
// 	}()
// 	fxn := func() {
// 		fmt.Println("B")
// 		wait.Done()
// 	}
// 	go fxn()
// 	go func(){
// 		 print()
// 		 wait.Done()
// 	}()
// 	wait.Wait()
// 	fmt.Println("all done")
// }

// func main() {
// 	ch := make(chan string)

// 	go func() {
// 		ch <- "done with work"
// 	}()

// 	cha := <-ch
// 	fmt.Println(cha)
// }

// func main() {
// 	var wait sync.WaitGroup
// 	wait.Add(1)

// 	go func() {
// 		fmt.Println("I am working")
// 		wait.Done()
// 	}()
// 	wait.Wait()
// 	fmt.Println("main done")
// }

// func say(s string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(s)
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func main() {
// 	go say("goroutine") // runs concurrently
// 	say("main")         // runs in main goroutine
// }
