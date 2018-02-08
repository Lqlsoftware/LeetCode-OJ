/*
Insert Interval
Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
You may assume that the intervals were initially sorted according to their start times.

Example 1:
Given intervals [1,3],[6,9], insert and merge [2,5] in as [1,5],[6,9].
Example 2:
Given [1,2],[3,5],[6,7],[8,10],[12,16], insert and merge [4,9] in as [1,2],[3,10],[12,16].
This is because the new interval [4,9] overlaps with [3,5],[6,7],[8,10].
*/

package main

func insert(intervals []Interval, newInterval Interval) []Interval {
	var result []Interval
	for i := 0;i < len(intervals);i++ {
		temp := intervals[i]
		if newInterval.Start > temp.End {
			result = append(result, temp)
		} else {
			if newInterval.Start > temp.Start {
				newInterval.Start = temp.Start
			}
			for ; i < len(intervals) && newInterval.End >= intervals[i].Start; i++ {
				if intervals[i].End > newInterval.End {
					newInterval.End = intervals[i].End
				}
			}
			result = append(result, newInterval)
			for ; i < len(intervals); i++ {
				result = append(result, intervals[i])
			}
			return result
		}
	}
	return append(result, newInterval)
}