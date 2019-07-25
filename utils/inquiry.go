package utils

import (
	"fmt"
	"strconv"
)

// Validation defines validation function to be passed in struct for user input validation
type Validation func(string) error

// Filter defines function for running filter on the data
type Filter func(string) (error, string)

// Query struct type to handle user input and Question
type Query struct {
	Name         string
	Question     string
	DefaultValue string
	AnswerType   string
	Answer       string
}

// Validate Handle user input validation
func (q Query) Validate(fn Validation) error {
	return fn(q.Answer)
}

// Filter hanlde filteration on user data
func (q *Query) Filter(fn Filter) error {
	err, Answer := fn(q.Answer)
	if err == nil {
		q.Answer = Answer
		return nil
	}
	return err
}

// InType Convert user Input into the format user wants
func (q *Query) InType() interface{} {
	switch q.AnswerType {
	case "int":
		result, _ := strconv.Atoi(q.Answer)
		return result
	case "string":
		return q.Answer
	case "float":
		result, _ := strconv.ParseFloat(q.Answer, 32)
		return result
	default:
		return q.Answer
	}
}

// Prompt triggers event to Ask Question to users
func (q *Query) Prompt() {
	fmt.Printf("\n%s [%s]: ", q.Question, q.DefaultValue)
	var response string
	fmt.Scanln(&response)
	q.Answer = response

}
