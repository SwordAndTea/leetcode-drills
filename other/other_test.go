package other

import (
	"fmt"
	"testing"
	"time"
)

func TestCalculateDate(t *testing.T) {
	bDate := time.Date(1965, time.July, 8, 0, 0, 0, 0, time.UTC)
	eDate := time.Date(2065, time.July, 8, 0, 0, 0, 0, time.UTC)
	//t.Log(bDate.Weekday())
	//t.Log(eDate.Weekday())
	numDays := CalculateDate(bDate, eDate)
	t.Log(numDays)
	t.Log(numDays % 7)
}

func TestGenerateP(t *testing.T) {
	result := GenerateP([]int{1, 2, 3})

	fmt.Printf("result is %v", result)
}

func TestSelectMostInterval(t *testing.T) {
	intervals := []*Interval{
		{
			Start: 1,
			End:   3,
		},
		{
			Start: 2,
			End:   4,
		},
		{
			Start: 3,
			End:   5,
		},
		{
			Start: 6,
			End:   7,
		},
	}

	result := SelectMostIntervals(intervals)

	for _, v := range result {
		fmt.Printf("%v", v)
	}
}

func TestBinarySearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	if BinarySearch(a, 2) != 1 {
		t.Fatalf("search error")
	}

	if BinarySearch(a, 6) != -1 {
		t.Fatalf("search error")
	}
}

func TestBinaryPow(t *testing.T) {
	a, b, m := uint64(2), uint64(11), uint64(3)

	ret := BinaryPow(a, b, m)

	print(ret)
}

func TestGCD(t *testing.T) {
	a, b := 49, 14

	if GCD(a, b) != 7 {
		t.Fatalf("gcd fail")
	}
}

func TestLCM(t *testing.T) {
	a, b := 49, 14

	if LCM(a, b) != 98 {
		t.Fatalf("lcm fail")
	}
}

func TestListPrime(t *testing.T) {
	res := ListPrime(15)

	fmt.Printf("res is %+v", res)
}

func TestCombinationNumber(t *testing.T) {
	if CombinationNumber(5, 3) != 10 {
		t.Fatalf("calculate combination number fail")
	}
}

func TestSimpleCalculator(t *testing.T) {
	//res, err := SimpleCalculator("(1 + 2) / 3")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if res != 1 {
	//	t.Fatalf("expect res to be 1, but get %v", res)
	//}
	//
	//res, err = SimpleCalculator("30/90-26+97-5-6-13/88*6+51/29+79*87+57*92")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Logf("res is %v", res)

	res, err := SimpleCalculator("(1+(4+5+2)-3)+(6+8)")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("res is %v", res)
}

func TestKMP(t *testing.T) {
	count := KMP("BBC ABCDAB ABCDABCDABDE", "ABCDABD")
	t.Log(count)
}
