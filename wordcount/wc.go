package wordcount

import (
	"log"
	"os"
	"runtime"
	"sync"
	"unicode"
)

var (
	mutex sync.Mutex
)

func loadMoby() []byte {
	moby, err := os.ReadFile("mobydick.txt")
	if err != nil {
		log.Fatal(err)
	}
	return moby
}

func wordCount(b []byte) (words int) {
	inWord := false
	for _, v := range b {
		r := rune(v)
		if unicode.IsSpace(r) && inWord {
			words++
		}
		inWord = unicode.IsLetter(r)
	}
	return
}

func shardSlice(input []byte, numShards int) (shards [][]byte) {
	shards = make([][]byte, numShards)
	if numShards < 2 {
		shards[0] = input[:]
		return
	}
	shardSize := len(input) / numShards
	start, end := 0, shardSize
	for i := 0; i < numShards; i++ {
		for j := end; j < len(input); j++ {
			char := rune(input[j])
			if unicode.IsSpace(char) {
				// split slice at position j, where there is a space
				// note: need to include the space in the shard to get accurate count
				end = j + 1
				shards[i] = input[start:end]
				start = end
				end += shardSize
				break
			}
		}
	}
	shards[numShards-1] = input[start:]
	return
}

func parallelWordCount(input []byte) (words int) {
	return doParallelWordCount(input, runtime.NumCPU())
}

func doParallelWordCount(input []byte, numShards int) (words int) {
	shard := shardSlice(input, numShards)
	var wg sync.WaitGroup
	for i := 0; i < numShards; i++ {
		wg.Add(1)
		go func(i int) {
			count := wordCount(shard[i])
			mutex.Lock()
			words += count
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	return
}
