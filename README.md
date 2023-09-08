# Golang-practice
Task given in 2022 by UiS in Operating Systems and System Programming

Contains four different tasks:
* FizzBizz
  This program uses goroutines to print a list of outputs depending on the integer input.
- `fizz()` prints the word `Fizz`, if the number is divisible by 3 and not 5.
- `bizz()` prints the word `Bizz`, if the number is divisible by 5 and not 3.
- `number()` prints the `number` if it is not divisible by 3 and 5.
- `fizzBizz()` prints the word `FizzBizz`, if the number is divisible by both 3 and 5.

* Stack
  This is a stack data structure which is accessed concurrently by several goroutines.
  The stack handles several functions:
- `Size()` returns the current number of items on the stack,
- `Pop() interface{}` pops an item of the stack (`nil` if empty), while
- `Push(value interface{})` pushes an item onto the stack.

* Water
  This program gets input of H and O (atoms) and forms them in groups of HHO, HOH, OHH (water molecules)
  It includes two main functions that handle the input:
- `releaseOxygen`: can produce only one oxygen atom and should block until that atom is consumed to form the water molecule before releasing another oxygen atom.
- `releaseHydrogen`: can produce at most two hydrogen atoms and should block until these atoms are consumed to form the water molecule before releasing additional hydrogen atoms.

  It also includes a Make() function which was pre-implemented by the teacher.

* Wordcount
  This program counts the words in wordcount/mobydick.txt and returns every word and how many times they were used.
  It uses parallel execution e.g. it counts the words using multiple goroutines, typically as many as there are CPU cores in your machine.