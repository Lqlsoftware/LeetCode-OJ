/*
Merge Intervals
Given a collection of intervals, merge all overlapping intervals.

For example,
Given [1,3],[2,6],[8,10],[15,18],
return [1,6],[8,10],[15,18].
*/

package main

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	var result []Interval
	for i := 0;i < len(intervals); {
		temp := intervals[i]
		for i++;i < len(intervals) && temp.End >= intervals[i].Start;i++ {
			if intervals[i].End > temp.End {
				temp.End = intervals[i].End
			}
		}
		result = append(result, temp)
	}
	return result
}