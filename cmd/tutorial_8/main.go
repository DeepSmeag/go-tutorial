package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8", "id9", "id10"}
var wg = sync.WaitGroup{}
var results = []string{}
var m = sync.Mutex{}
var rwm = sync.RWMutex{}

func main() {

	// var t0 = time.Now()
	// for i := 0; i < len(dbData); i++ {
	// 	dbCall(i)
	// }
	// fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	var t1 = time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t1))
	// fmt.Printf("The results are: %v\n", results)

}

func dbCall(i int) {
	// simulate db call delay
	var delay float32 = 2000 //rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// fmt.Println("The result from the database is: ", dbData[i])
	m.Lock()
	// results = append(results, dbData[i])
	save(dbData[i])
	log()
	m.Unlock()
	wg.Done()
}

func save(result string) {
	rwm.Lock()
	// fmt.Println("Write lock acquired")
	results = append(results, result)
	// fmt.Println("Write lock releasing")
	rwm.Unlock()
}

func log() {
	rwm.RLock()
	// fmt.Println("Read lock acquired")
	fmt.Println("The results are: ", results)
	// fmt.Println("Read lock releasing")
	rwm.RUnlock()
}
