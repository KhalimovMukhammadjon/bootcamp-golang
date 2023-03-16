package hour

import "fmt"

func FindHour() {

	var d, h, m, s int

	s = 150005 

	d = s / (24 * 3600)

	s = s % (24 * 3600)
	h = s / 3600

	s %= 3600
	m = s / 60

	s %= 60
	sec := s

	fmt.Println(d, "days", h, "hours",
		m, "minutes",
		sec, "seconds")
	// fmt.Println(d)
	// fmt.Println(h)
	// fmt.Println(m)
	// fmt.Println(sec)
	//fmt.Println(d)
	//fmt.Printf("")
}
