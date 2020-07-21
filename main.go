package calc

import (
	"errors"
)

type nums map[bool]int64

func findNums(e string) (nums nums, sign string, err error) {
	nums = make(map[bool]int64, 2)
	var currentStep bool

	var elem rune

	for _, elem = range e {

		switch elem {
		case ' ':
			continue
		case '*':
			sign = "*"
		case '/':
			sign = "/"
		case '+':
			sign = "+"
		case '-':
			sign = "-"
		default:
			err = nums.numberCheck(currentStep, elem, false)
		}

		if err != nil {
			return
		}

		if sign != "" && currentStep == false {
			currentStep = true
		}

		//log.Println(elem)
		//log.Println(currentStep)
		//log.Println(nums)
		//log.Println(sign)
	}

	// Check the last elem in the string if it is a number
	err = nums.numberCheck(currentStep, elem, true)
	if err != nil {
		return
	}

	if sign == "" {
		err = errors.New("no supported sign is in the expression")
	}

	return
}

func (nums nums) numberCheck(currentStep bool, input rune, onlyCheck bool) (err error) {
	if input < 48 || input > 57 {
		err = errors.New("expression contains incorrect symbols")
	} else if !onlyCheck {
		nums[currentStep] = nums[currentStep]*10 + int64(input) - 48
	}
	return
}

func Calc(e string) (result int64, err error) {
	var nums nums
	var sign string

	nums, sign, err = findNums(e)
	if err != nil {
		return
	}
	//log.Println(nums)

	switch sign {
	case "*":
		result = nums[false] * nums[true]
	case "/":
		if nums[true] == 0 {
			return 0, errors.New("you tried to divide by zero")
		}
		result = nums[false] / nums[true]
	case "+":
		result = nums[false] + nums[true]
	case "-":
		result = nums[false] - nums[true]
	}

	return
}
