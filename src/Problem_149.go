/*
Max Points on a Line
Given n points on a 2D plane, find the maximum number of points that lie on the same straight line.

Example 1:
Input: [[1,1],[2,2],[3,3]]
Output: 3

Explanation:
^
|
|        o
|     o
|  o
+------------->
0  1  2  3  4

Example 2:
Input: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
Output: 4
Explanation:
^
|
|  o
|     o        o
|        o
|  o        o
+------------------->
0  1  2  3  4  5  6
*/

package main

// Calculate every pair of points to check the same line expression.
// If it is the same point, only add the number(not line)
func maxPoints(points []Point) int {
	if len(points) <= 2 {
		return len(points)
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].X != points[j].X {
			return points[i].X < points[j].X
		}
		return points[i].Y < points[j].Y
	})
	max := 0
	for i := range points {
		mem, number := map[Line]int{}, 0
		for j := range points {
			if points[i] == points[j] {
				number++
			} else {
				mem[GetLine(points[i],points[j])]++
			}
		}
		max = Max(number, max)
		for j := range mem {
			max = Max(mem[j] + number, max)
		}
	}
	return max
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// define line like "au/ad x + by + cu/cd = 0"
type Line struct {
	Au int
	Ad int
	B  int
	Cu int
	Cd int
}

// au and ad need to divide GCD(au, ad) to make sure same line have same form.
func GetLine(point1, point2 Point) Line {
	diffY, diffX := point1.Y - point2.Y, point1.X - point2.X
	if diffX == 0 {
		return Line{1,1,0,-point1.X,1}
	}
	p := GCD(diffX,diffY)
	diffX /= p
	diffY /= p
	return Line{diffY, diffX, -1,-diffY * point1.X + diffX * point1.Y, diffX}
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	} else if a % b == 0 {
		return b
	}
	return GCD(b, a % b)
}