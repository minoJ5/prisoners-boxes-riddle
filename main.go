package main

import (
	"fmt"
	exp "prisoners-dilemma/experiment" 
	"time"
)


func main() {
 t := time.Now()
 res := exp.RunExperiments(10_000,100)
 fmt.Printf("%.7f\n",exp.Stats(res))
 fmt.Println(time.Since(t))
}
