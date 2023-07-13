package rollee

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
	
  for _, ch := range chs {

	go func(ch chan List) {

		for myList := range ch {
			result := Fold(initialValue, f, myList)

			for id, value := range result {
				m[id] = f(m[id], value)
			}
		}
	}(ch)
  }
  return m
}

