package main

import (
	"fmt"
	"time"

	"go-core-task/8/wg"
)

func main() {
	myWG := wg.New()

	for i := 1; i <= 3; i++ {
		myWG.Add(1)

		go func(id int) {
			defer myWG.Done()

			fmt.Println("start", id)
			time.Sleep(time.Second)
			fmt.Println("end", id)
		}(i)
	}

	myWG.Wait()
	fmt.Println("all done")
}
