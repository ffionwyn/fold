package rollee

import "sync"

type ID = int

// We suppose L is always valid with len (l.Values) >= 1).
type List struct {
	ID     ID
	Values []int
}

// This function takes 3 arguments, initialValue(int), f(takes 2 ints returns int) and l(struct).
// Created an empty map(m) to store the results.
// Check the length of the list(l.Values) is less than 1. If it is, it means the list is empty, so it assigns the initialValue to the map with the ID(l.ID) as the key and then returns the map. 
// If the list is not empty, it takes the firstValue from the list(l.Values).
// Creates a new list(newValues) by excluding the first value from the original list(l.Values).
// This recursively calls the Fold function with the updated params. The result of this is returned as the final result of the Fold operation. 

func Fold(initialValue int, f func(int, int) int, l List) map[ID]int {
	m := make(map[ID]int)
	
	 if (len(l.Values) < 1) {
	    m[l.ID] = initialValue
	    return m
	 }
		 firstValue:= l.Values[0]

    	 newValues := l.Values[1:]
	
	 return Fold(f(initialValue, firstValue), f, List{l.ID, newValues})
}

// This function takes 3 arguments, initialValue(int), f(takes 2 ints, returns int) and ch(channel which receives list vals).
// Created an empty map(m) to store the results.
// Loops over the channel(range), this keeps going until the channel is closed. 
// Inside the loop, myList is received from the channel. It calls the Fold function with myList, initialValue and f to compete the fold operation on the list. 
// Check if the ID of myList already exists in the map(m), If it does exist it applies the function f to the value and the folded value from Fold and then updates the map with the new value. If it does not exist, it assigns the folded value to the map. 
// When the loop is finished (no more values in the channel) the function will return the final map that contains the folded results for each ID. 

func FoldChan(initialValue int, f func(int, int) int, ch chan List) map[ID]int {
	m := make(map[ID]int)	
	
	for myList := range ch {
		result := Fold(initialValue, f, myList)
		
		if val, ok := m[myList.ID] ; ok {
			m[myList.ID] = f(val,result[myList.ID])
		} else { 
		m[myList.ID] = result[myList.ID]
	}
}
	return m
}

// This function takes 3 arguments, initialValue(int), f(takes 2 ints, returns int) and ch(channel which receives list vals). 
// Created an empty map(m) to store the results. 
// Use mutexes to synchronise access to the map(m) when the map is being updated concurrently. 
// Use waitgroups to synchronise the completion of the goroutines.
// Loop over the chs slice(range) to iterate over each channel.
// Inside the loop, the waitgroup counter is incrementing by wg.Add(1), using this to track the number of go routines being used.
// Using a new go routine for each channel by using the anonymous function with the channel as the parameter. Each go routine – uses defer to make sure that wg.Done is called at the end, starts a loop over the channel(ch) and keeps going until the channel is closed, inside the loop receives myList from channel, calls the Fold function with myList and initialvalue and f to complete the fold operation on the list, iterate over the results(map) from the Fold function and updates the map to merge the folded values with the existing values for each ID. The map is protected with the mutex which is locked and unlocked. 
// Once the loop is finished the go routine will call wg.Done. After the loop over the channels, it waits for all go routines to be finished with wg.Wait. 
// Returns the final map that contains the folded results for each ID. 

func FoldChanX(initialValue int, f func(int, int) int, chs ...chan List) map[ID]int {
	m := make(map[ID]int)
	var mutex sync.Mutex	
	
	wg := sync.WaitGroup{}
    	for _, ch := range chs {
	wg.Add(1)
	go func(ch chan List) {
		defer wg.Done()
		for myList := range ch {
			result := Fold(initialValue, f, myList)
		for id, value := range result {
				mutex.Lock()
				m[id] = f(m[id], value)
				mutex.Unlock()
			}
		}
	}(ch)
  }
  wg.Wait()
  return m
}