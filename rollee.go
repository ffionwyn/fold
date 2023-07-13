package rollee

import "sync"

type ID = int

// We suppose L is always valid with len (l.Values) >= 1).
type List struct {
	ID     ID
	Values []int
}

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

