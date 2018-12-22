package main

import "fmt"

func main() {
	s1 := Student{
		Name: "zhang",
		Id:   200000225,
		M_Score: Score{
			Chinese: 85.368,
			English: 87.4564,
		},
	}

	fmt.Println(s1)
	fmt.Println(s1.toString())
}

type Student struct {
	Name    string
	Id      uint64
	M_Score Score
}
type Score struct {
	Chinese float64
	English float64
}

func (s *Student) toString() string {
	return fmt.Sprintf("student: %s, chinese: %f, english: %f",
		s.Name, s.M_Score.Chinese, s.M_Score.English)
}
