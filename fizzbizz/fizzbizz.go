package fizzbizz

import (
	"strconv"
	"sync"
)

// SyncBlock structure holds the synchronization constructs and
// data required to solve the fizzbizz problem.
// current contains the current value to be processed.
// max is the value upto which the goroutines should process.
// result holds the output of the assignment.
// wg is the WaitGroup used to wait for the completion of the goroutines.
type SyncBlock struct {
	sync.Mutex
	cond    *sync.Cond
	wg      sync.WaitGroup
	current int
	max     int
	result  string
}

// newSyncBlock initializes the SyncBlock.
func newSyncBlock(max int) *SyncBlock {
	block := &SyncBlock{}
	block.max = max
	block.current = 1
	block.cond = sync.NewCond(block)
	return block
}

// appendToResult appends partialResult to the result
// generated by the goroutines and increments s.current.
// Must only be called when holding the lock on s.
func (s *SyncBlock) appendToResult(partialResult string) {
	s.result = s.result + partialResult
	s.current++
}

// fizz appends "Fizz" to the result if
// s.current is divisible by 3 and not divisible by 5,
// and increments the value of s.current.
//
// Otherwise, if the above condition is not satisfied,
// fizz waits for other goroutines to execute and
// increment the value of s.current.
func (s *SyncBlock) fizz() {
	for i := 1; i <= s.max; i++ {
		s.cond.L.Lock()
		for !(s.current%3 == 0 && s.current%5 != 0) && s.current <= s.max {
			s.cond.Wait()
		}
		if s.current > s.max {
			s.cond.Broadcast()
			s.Unlock()
			break
		} else {
			s.appendToResult("Fizz")
			s.cond.Broadcast()
			s.cond.L.Unlock()

		}
	}
	s.wg.Done()
}

// bizz appends "Bizz" to the result if
// s.current is divisible by 5 and not divisible by 3,
// and increments the value of s.current.
//
// Otherwise, if the above condition is not satisfied,
// bizz waits for other goroutines to execute and
// increment the value of s.current.
func (s *SyncBlock) bizz() {
	for i := 1; i <= s.max; i++ {
		s.cond.L.Lock()
		for !(s.current%5 == 0 && s.current%3 != 0) && s.current <= s.max {
			s.cond.Wait()
		}
		if s.current > s.max {
			s.cond.Broadcast()
			s.Unlock()
			break
		} else {

			s.appendToResult("Bizz")
			s.cond.Broadcast()
			s.cond.L.Unlock()
		}
	}
	s.wg.Done()
}

// number appends s.current (as a string) to the result if
// s.current is not divisible by 3 and 5,
// and increments the value of s.current.
//
// Otherwise, if the above condition is not satisfied,
// number waits for other goroutines to execute and
// increment the value of s.current.
func (s *SyncBlock) number() {
	for i := 1; i <= s.max; i++ {
		s.cond.L.Lock()
		for !(s.current%5 != 0 && s.current%3 != 0) && s.current <= s.max {
			s.cond.Wait()
		}
		if s.current > s.max {
			s.cond.Broadcast()
			s.Unlock()
			break
		} else {
			c := strconv.Itoa(s.current)
			s.appendToResult(c)
			s.cond.Broadcast()
			s.cond.L.Unlock()

		}

	}
	s.wg.Done()
}

// fizzBizz appends "FizzBizz" to the result if
// s.current is divisible by 3 and 5,
// and increments the value of s.current.
//
// Otherwise, if the above condition is not satisfied,
// fizzBizz waits for other goroutines to execute and
// increment the value of s.current.
func (s *SyncBlock) fizzBizz() {
	for i := 1; i <= s.max; i++ {
		s.cond.L.Lock()
		for !(s.current%5 == 0 && s.current%3 == 0) && s.current <= s.max {
			s.cond.Wait()
		}
		if s.current > s.max {
			s.cond.Broadcast()
			s.Unlock()
			break
		} else {
			s.appendToResult("FizzBizz")
			s.cond.Broadcast()
			s.cond.L.Unlock()
		}
	}
	s.wg.Done()
}

// FizzBizz returns the result of the fizzbizz algorithm for the given max.
// The output is produced when all goroutines have completed.
// DO NOT edit the FizzBizz function or modify the signatures of the other methods used by FizzBizz.
func FizzBizz(max int) string {
	const numGoroutines = 4
	s := newSyncBlock(max)
	s.wg.Add(numGoroutines)
	go s.fizz()
	go s.bizz()
	go s.number()
	go s.fizzBizz()
	s.wg.Wait()
	return s.result
}
