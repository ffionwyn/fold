package rollee

type ID = int

// We suppose L is always valid with len (l.Values) >= 1).
type List struct {
	ID     ID
	Values []int
}

func Fold(initialValue int, f func(int, int) int, l List) map[ID]int {
	panic("not implemented")
}

func FoldChan(initialValue int, f func(int, int) int, ch chan List) map[ID]int {
	panic("not implemented")
}

func FoldChanX(initialValue int, f func(int, int) int, chs ...chan List) map[ID]int {
	panic("not implemented")
}
