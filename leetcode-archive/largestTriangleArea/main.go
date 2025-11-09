package main

import (
	"fmt"
	"math"
)

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	areaOfTriangle := 0.0
	maxArea := 0.0
	for i:=0; i<n; i++ {
		for j:=i+1; j<n; j++ {
			for k:=j+1; k<n; k++ {
				x1 := points[i][0]
				y1 := points[i][1]
				x2 := points[j][0]
				y2 := points[j][1]
				x3 := points[k][0]
				y3 := points[k][1]

				y2MinusY3 := y2 - y3
				y3MinusY1 := y3 - y1
				y1MinusY2 := y1 - y2

				areaOfTriangle = float64((math.Abs(float64((x1 * y2MinusY3) + (x2 * y3MinusY1) + (x3 * y1MinusY2)))) / 2)
				maxArea = math.Max(areaOfTriangle, maxArea)
			}
		}
	}
	return maxArea
}

func main() {
	points := [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
	fmt.Println("largestTriangleArea -> ", largestTriangleArea(points))
}
