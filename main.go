package main

import (
	"fmt"
	exp "prisoners-dilemma/experiment"
	tc "prisoners-dilemma/tcolor"
	"time"
)


func main() {
	tc.Cprint("Experiment Started:", tc.FgGreen)
	t := time.Now()
	res := exp.RunExperiments(10_000,100)
	var end  time.Duration = time.Since(t)
	tc.Cprintf("Empirical P =", tc.FgGreen)
	fmt.Printf(" %.7f\n",exp.Stats(res))
	tc.Cprintf("Took: ", tc.FgGreen)
	fmt.Printf("%s\n", end)
}
