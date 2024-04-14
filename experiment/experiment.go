// Package containing the methods 
// used to simulate the experiments.

package experiment

import (
	"math/rand"
	"sync"
)

type Experiment struct {
	Number   int
	Prisoner int
	Result   bool
	Try      int
}

type ExperimentResult struct {
	Experiment int
	Success    bool
}

type ExperimentResults []ExperimentResult

func fillArray(array *[]int, size int) {
	for i := 0; i < size; i++ {
		*array = append(*array, i+1)
	}
}

func randInRange(min, max int) int {
	return rand.Intn(max+1-min) + min
}

func shuffleArray(array *[]int, size int) {
	fillArray(array, size)
	var k int
	temp := new(int)
	for i := len(*array) - 1; i >= 1; i-- {
		k = randInRange(0, i)
		*temp = (*array)[i]
		(*array)[i] = (*array)[k]
		(*array)[k] = *temp
	}
}

func FindNumber(boxes *[]int, limit int, n int) (bool, int) {
	i := 1
	var found bool = false
	var open int = n - 1
	for i <= limit {
		if (*boxes)[open] == n {
			found = true
			break
		} else {
			open = (*boxes)[open] - 1
			i++
		}
	}
	return found, i
}

func isSuccess(array []bool) bool {
	var a0 bool = array[0]
	if !a0 {
		return false
	}
	for i := 1; i < len(array); i++ {
		if array[i] != a0 || !array[i] {
			return false
		}
	}
	return true
}

func RunExperiment(number int, boxes int) ExperimentResult {
	var wg sync.WaitGroup
	li := make([]int, 0)
	shuffleArray(&li, boxes)
	resl := make([]bool, 0)
	for i := 1; i <= boxes; i++ {
		wg.Add(1)
		go func(i int) {
			var expr Experiment
			defer wg.Done()
			expr.Number = number
			expr.Prisoner = i
			expr.Result, expr.Try = FindNumber(&li, boxes/2, i)
			resl = append(resl, expr.Result)
		}(i)
	}
	wg.Wait()
	var res ExperimentResult
	res.Experiment = number
	res.Success = isSuccess(resl)
	li = nil
	resl =nil
	return res
}

func RunExperiments(number int, boxes int) ExperimentResults {
	var resls ExperimentResults
	for i := 1; i <= number; i++ {
		resls = append(resls, RunExperiment(i, boxes))
	}
	return resls
}

func Stats(res ExperimentResults) float64 {
	var s int
	var f int
	for i := 0; i < len(res); i++ {
		if res[i].Success {
			s++
		} else {
			f++
		}
	}
	return float64(s) / float64(len(res))
}
