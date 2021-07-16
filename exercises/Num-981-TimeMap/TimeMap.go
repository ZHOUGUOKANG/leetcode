package Num_981_TimeMap

import "sort"

type TimeMap struct {
	mapping map[string][]*Item
}
type Item struct {
	value     string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{mapping: map[string][]*Item{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.mapping[key] = append(this.mapping[key], &Item{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if items, ok := this.mapping[key]; ok {
		i := sort.Search(len(items), func(i int) bool {
			return items[i].timestamp > timestamp
		})
		if i > 0 {
			return items[i-1].value
		}
	}
	return ""
}
