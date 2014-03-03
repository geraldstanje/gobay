package slm

import "testing"

func Test_MaxIntFK(test *testing.T) {
	m := map[float64]int{
		0.9:  9,
		0.1:  1,
		3.14: 314,
		5.1:  51,
		5.0:  51,
		1.1:  51,
	}
	if MaxIntFK(m) != 51 {
		test.Errorf("It should return 51 but %v", MaxIntFK(m))
	}
}

func Test_Top5SFK(test *testing.T) {
	m := map[float64]string{
		0.9:  "9",
		0.1:  "1",
		3.14: "314",
		5.1:  "51",
		5.0:  "51",
		1.1:  "51",
	}
	if CheckStr("1", Top5SFK(m)) {
		test.Errorf("1 shouldn't be in the result: %v", Top5SFK(m))
	}
}

func Test_UniqInts(test *testing.T) {
	s := []int{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	if len(UniqInts(s)) != 5 {
		test.Errorf("Should return 5 but %v", UniqInts(s))
	}
}
