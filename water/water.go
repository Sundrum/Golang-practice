package water

import (
	"sync"
)

// Water structure holds the synchronization primitives and
// data required to solve the water molecule problem.
// moleculeCount holds the number of molecules formed so far.
// result string contains the sequence of "H" and "O".
// wg WaitGroup is used to wait for goroutine completion.
type Water struct {
	wg            sync.WaitGroup
	moleculeCount int
	result        string
	sum           int
	sync.Mutex
	cond *sync.Cond
}

// New initializes the water structure.
func New() *Water {
	water := &Water{}
	water.moleculeCount = 0
	water.result = ""
	water.sum = 0
	water.cond = sync.NewCond(water)
	return water
}

// releaseOxygen produces one oxygen atom if no oxygen atom is already present.
// If an oxygen atom is already present, it will block until enough hydrogen
// atoms have been produced to consume the atoms necessary to produce water.
//
// The w.wg.Done() must be called to indicate the completion of the goroutine.
func (w *Water) releaseOxygen() {
	w.cond.L.Lock()
	for !(0 <= w.sum && w.sum <= 2) {
		w.cond.Wait()
	}
	w.sum -= 2
	w.result = w.result + "O"
	w.cond.Broadcast()
	w.cond.L.Unlock()
	defer w.wg.Done()
}

// releaseHydrogen produces one hydrogen atom unless two hydrogen atoms are already present.
// If two hydrogen atoms are already present, it will block until another oxygen
// atom has been produced to consume the atoms necessary to produce water.
//
// The w.wg.Done() must be called to indicate the completion of the goroutine.
func (w *Water) releaseHydrogen() {
	w.cond.L.Lock()
	for -2 > w.sum || w.sum > 1 {
		w.cond.Wait()
	}
	w.sum += 1
	w.result = w.result + "H"
	w.cond.Broadcast()
	w.cond.L.Unlock()
	defer w.wg.Done()
}

// produceMolecule forms the water molecules.
func (w *Water) produceMolecule(done chan bool) {
	done <- true
}

func (w *Water) finish() {
}

// Molecules returns the number of water molecules that has been created.
func (w *Water) Molecules() int {
	w.moleculeCount = len(w.result) / 3
	return w.moleculeCount
}

// Make returns a sequence of water molecules derived from the input of hydrogen and oxygen atoms.
// DO NOT edit the Make method or modify the signatures of the other methods used by Make.
func (w *Water) Make(input string) string {
	done := make(chan bool)
	go w.produceMolecule(done)
	for _, ch := range input {
		w.wg.Add(1)
		switch ch {
		case 'O':
			go w.releaseOxygen()
		case 'H':
			go w.releaseHydrogen()
		}
	}
	w.wg.Wait()
	w.finish()
	<-done
	return w.result
}
