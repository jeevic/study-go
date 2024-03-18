package main

import (
	"fmt"
	"time"

	"github.com/juju/ratelimit"
)

// main r
func main() {
	bucket := ratelimit.NewBucketWithQuantum(1*time.Second, 10, 10)

	fmt.Printf("bucket available: %d \n", bucket.Available())
	fmt.Printf("bucket rate: %f \n", bucket.Rate())

	/*for i := 0; i < 200; i++ {
	   t := bucket.Take(1)
	   fmt.Printf("bucket take: %s \n", t.String())
	}*/

	for i := 0; i < 50; i++ {
		t := bucket.TakeAvailable(1)
		fmt.Printf("bucket take Avaliable : %d \n", t)
	}

	for i := 0; i < 50; i++ {
		d, b := bucket.TakeMaxDuration(1, 5*time.Second)
		fmt.Printf("bucket take Duration : %s, %#v \n", d.String(), b)
	}

	for i := 0; i < 50; i++ {
		b := bucket.WaitMaxDuration(1, 5*time.Second)
		fmt.Printf("bucket wait duration : %#v \n", b)
	}

	for i := 0; i < 50; i++ {
		bucket.Wait(1)
		fmt.Printf("bucket wait \n")
	}

}
