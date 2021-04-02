package leetcode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Polynomial [2]int

func SolveEquation(equation string) (result string) {
	sides := strings.Split(equation, "=")
	left, right := sides[0], sides[1]

	leftPolynomial := addCoheficients(left)
	rightPolynomial := addCoheficients(right)

	p := [2]int{leftPolynomial[0] - rightPolynomial[0], rightPolynomial[1] - leftPolynomial[1]}

	if p[0] == 0 && p[1] != 0 {
		result = "No solution"
	} else if p[0] == 0 && p[1] == 0 {
		result = "Infinite solutions"
	} else {
		result = fmt.Sprintf("x=%v", p[1]/p[0])
	}

	return
}

func addCoheficients(equation string) (coheficients Polynomial) {
	reader := []int{0,0}
	for {
		indexes := regexp.MustCompile(`\+|-`).FindStringIndex(equation[reader[1]:])

		if indexes == nil {
			parseStatement(equation[reader[0]:], &coheficients)
			break
		} else if indexes[0] == 0 {
			reader[1] = 1
		} else {
			parseStatement(equation[reader[0]:reader[1]+indexes[0]], &coheficients)
			reader = []int{reader[1] + indexes[0], reader[1] + indexes[1]}
		}
	}

	return
}

func parseStatement(sub string, p *Polynomial) {
	if strings.Contains(sub, "x") {
		sub := strings.ReplaceAll(sub, "x", "")

		var incr int
		if sub == "-" {
			incr = -1
		} else if sub == "" || sub == "+" {
			incr = 1
		} else {
			incr, _ = strconv.Atoi(sub)
		}

		p[0] += incr
	} else {
		incr, _ := strconv.Atoi(sub)
		p[1] += incr
	}
}