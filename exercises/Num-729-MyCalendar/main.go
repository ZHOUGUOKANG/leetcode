package Num_729_MyCalendar

type MyCalendar struct {
	list [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{list: [][2]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, v := range this.list {
		if !(start >= v[1] || end <= v[0]) {
			return false
		}
	}
	this.list = append(this.list, [2]int{start, end})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
