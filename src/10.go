
package main

import "math"

func GenerateHomework(grade int) []string {
	var homework []string
	for i := 0; i < grade; i++ {
		x1 := math.Floor(math.Rand()*10) + 1
		y1 := math.Floor(math.Rand()*10) + 1
		operation := math.Floor(math.Rand()*2)
		if operation == 0 {
			homework = append(homework, x1+"+"+y1)
		} else if operation == 1 {
			homework = append(homework, x1+"-"+y1)
		} else if operation == 2 {
			homework = append(homework, x1+"*"+y1)
		} else if operation == 3 {
			homework = append(homework, x1+"/"+y1)
		}
	}
	return homework
}