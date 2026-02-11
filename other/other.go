package other

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func getDayInMonth(month int, year int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	}
	return 0
}

func CalculateDate(begin, end time.Time) int {
	bYear, bMonth, bDay := begin.Date()
	eYear, eMonth, eDay := end.Date()
	numberOfDays := 0
	for bDay < eDay || bMonth < eMonth || bYear < eYear {
		bDay++
		numberOfDays++
		if bDay == getDayInMonth(int(bMonth), bYear)+1 {
			bDay = 1
			bMonth++
		}
		if bMonth == 13 {
			bYear++
			bMonth = 1
		}
	}
	return numberOfDays
}

func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func generatePImpl(input []int, result *[][]int, useInfo map[int]bool, current []int) {
	if len(current) == len(input) {
		*result = append(*result, current)
		return
	}
	for _, v := range input {
		if !useInfo[v] {
			newRes := make([]int, len(current), len(input))
			copy(newRes, current)
			newRes = append(newRes, v)
			useInfo[v] = true
			generatePImpl(input, result, useInfo, newRes)
			useInfo[v] = false
		}
	}
}

func GenerateP(input []int) [][]int {
	useInfo := make(map[int]bool)
	for _, v := range input {
		useInfo[v] = false
	}

	result := make([][]int, 0, factorial(len(input)))
	current := make([]int, 0, len(input))
	generatePImpl(input, &result, useInfo, current)

	return result
}

type Interval struct {
	Start int
	End   int
}

// from select most intervals, there is a significant difference of
//choosing from the smallest start or the biggest start

func SelectMostIntervals(intervals []*Interval) []*Interval {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start > intervals[j].Start
	})

	result := make([]*Interval, 0, len(intervals))
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i].End <= result[len(result)-1].Start {
			result = append(result, intervals[i])
		}
	}

	return result
}

func BinarySearch(input []int, target int) int {
	left, right := 0, len(input)-1
	for left <= right {
		mid := left + (right-left)/2 // not using (left+right) / 2 to avoid overflow
		if input[mid] == target {
			return mid
		} else if target < input[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func BinaryPow(a, b, m uint64) uint64 {
	if b == 0 {
		return 1
	}

	if b%2 == 1 {
		return a * BinaryPow(a, b-1, m) % m
	} else {
		r := BinaryPow(a, b/2, m) % m
		return r * r % m
	}
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

func LCM(a, b int) int {
	gcd := GCD(a, b)

	return a / gcd * b
}

func ListPrime(a int) []int {
	if a < 2 {
		return nil
	}

	flags := make(map[int]bool)
	for i := 3; i <= a; i += 2 {
		flags[i] = true
	}

	res := make([]int, 0, 4)
	res = append(res, 2)

	for i := 3; i <= a; i += 2 {
		if flags[i] {
			res = append(res, i)
			for j := i; j <= a; j += i {
				flags[j] = false
			}
		}
	}

	return res
}

// CombinationNumber number of choices of choosing m elements from n
// Note: CombinationNumber(n, m) equals CombinationNumber(n, n-m)
func CombinationNumber(n, m int) int {
	res := 1
	for i := 1; i <= m; i++ {
		res = res * (n - m + i) / i
	}
	return res
}

var operatorPriority = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
}

type expressionNode struct {
	Number float64
	OP     string
	IsNum  bool
}

func SimpleCalculator(expression string) (float64, error) {
	operatorStack := make([]string, 0, len(expression))
	expressionStack := make([]*expressionNode, 0, len(expression)) //postfix expression stack

	for i := 0; i < len(expression); i++ {
		c := expression[i]
		if c >= 48 && c <= 57 /*0-9*/ {
			j := i + 1
			for ; j < len(expression) && expression[j] >= 48 && expression[j] <= 57; j++ {
			}
			v, err := strconv.ParseFloat(expression[i:j], 64)
			if err != nil {
				return 0, err
			}
			expressionStack = append(expressionStack, &expressionNode{
				Number: v,
				IsNum:  true,
			})
			i = j - 1
		} else if c == 40 /*(*/ {
			tmpStack := make([]byte, 0, len(expression))
			tmpStack = append(tmpStack, c)
			for j := i + 1; j < len(expression); j++ {
				if expression[j] == '(' {
					tmpStack = append(tmpStack, ')')
				} else if expression[j] == ')' {
					tmpStack = tmpStack[:len(tmpStack)-1]
					if len(tmpStack) == 0 {
						v, err := SimpleCalculator(expression[i+1 : j])
						if err != nil {
							return 0, err
						}
						expressionStack = append(expressionStack, &expressionNode{
							Number: v,
							IsNum:  true,
						})
						i = j
						break
					}
				}
			}
			if len(tmpStack) != 0 {
				return 0, fmt.Errorf("invalid ( and ) pairs")
			}
		} else if c == 32 { // blank space
			continue
		} else {
			op := string([]byte{c})
			if len(operatorStack) == 0 {
				if op == "-" && len(expressionStack) == 0 {
					expressionStack = append(expressionStack, &expressionNode{
						Number: 0,
						IsNum:  true,
					})
				}
				operatorStack = append(operatorStack, op)
			} else {
				opPriority := operatorPriority[op]
				j := len(operatorStack) - 1
				for ; j >= 0; j-- {
					topOp := operatorStack[j]
					if opPriority <= operatorPriority[topOp] {
						expressionStack = append(expressionStack, &expressionNode{
							OP: topOp,
						})
					} else {
						break
					}
				}
				operatorStack = append(operatorStack[0:j+1], op)
			}
		}
	}
	for i := len(operatorStack) - 1; i >= 0; i-- {
		expressionStack = append(expressionStack, &expressionNode{
			OP: operatorStack[i],
		})
	}

	calculationStack := make([]float64, 0, len(expressionStack))

	for _, v := range expressionStack {
		if v.IsNum {
			fmt.Printf("%v ", v.Number)
			calculationStack = append(calculationStack, v.Number)
		} else {
			fmt.Printf("%v ", v.OP)
			right := calculationStack[len(calculationStack)-1]
			if len(calculationStack) < 2 && v.OP != "-" {
				if v.OP != "-" {
					return 0, fmt.Errorf("invalid expression")
				}
				calculationStack[len(calculationStack)-1] = 0 - right
			} else {
				left := calculationStack[len(calculationStack)-2]
				calculationStack = calculationStack[0 : len(calculationStack)-2]
				switch v.OP {
				case "+":
					calculationStack = append(calculationStack, left+right)
					break
				case "-":
					calculationStack = append(calculationStack, left-right)
					break
				case "*":
					calculationStack = append(calculationStack, left*right)
					break
				case "/":
					calculationStack = append(calculationStack, left/right)
					break
				}
			}
		}
	}

	return calculationStack[0], nil
}

type UnionFindSet struct {
	Pre map[int]int
}

func NewUnionFindSet(elements []int) *UnionFindSet {
	pre := make(map[int]int)

	for _, v := range elements {
		pre[v] = v // in initialization, each element is a set, which root in itself
	}

	return &UnionFindSet{
		Pre: pre,
	}
}

// FindTop find the root of the set that the element belongs to,
// meanwhile, during finding process, it will zip the path
// which all element directly point to the root element
func (ufs *UnionFindSet) FindTop(element int) int {
	i := element
	for element != ufs.Pre[element] { // if is not root element
		element = ufs.Pre[element] // go forward
	}

	// now the element is the root element of current set
	// zip path
	for i != ufs.Pre[i] && ufs.Pre[i] != element {
		z := ufs.Pre[i] // temporary hold the previous element
		ufs.Pre[i] = element
		i = z
	}

	return element
}

func (ufs *UnionFindSet) Union(a, b int) bool {
	topA, topB := ufs.FindTop(a), ufs.FindTop(b)
	if topA != topB {
		ufs.Pre[topB] = topA
		return true
	}
	return false
}

func getNext(s string) []int {
	if len(s) == 0 {
		return nil
	}
	next := make([]int, len(s))
	next[0] = 0
	for i := 1; i < len(s); i++ {
		j := next[i-1]
		for j != 0 && s[i] != s[j] {
			j = next[j-1]
		}

		if s[i] == s[j] {
			next[i] = j + 1
		} else {
			next[i] = j // j is 0
		}
	}
	return next
}

func KMP(text, pattern string) int {
	patternNext := getNext(pattern)
	i, j := 0, 0
	count := 0
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
			if j == len(pattern) {
				count++
				j = 0
			}
		} else if j == 0 {
			i++
		} else {
			j = patternNext[j-1]
		}
	}
	return count
}
