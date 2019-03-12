package utils

import (
	"fmt"
	"strconv"
)

// user define validation function to be passed in struct for user input validation
type validate func(string) error

// user defeind filter function to process the user defined input
type filter func(string) (error, string)

// Query struct type to handle user input and Question
type Query struct {
	Name         string
	Question     string
	DefaultValue string
	AnswerType   string
	Answer       string
}

// Validate Handle user input validation
func (q Query) Validate(fn validate) error {
	return fn(q.Answer)
}

// Filter hanlde filteration on user data
func (q *Query) Filter(fn filter) error {
	err, Answer := fn(q.Answer)
	q.Answer = Answer
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
