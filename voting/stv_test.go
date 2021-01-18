package voting

import (
	"sort"
	"testing"
)

// TestSimple tests a really basic STV example
func TestSimple(t *testing.T) {
	actual := STVElection(
		2,
		[]string{"a", "b", "c"},
		[][]string{
			[]string{"a", "b"},
			[]string{"a", "b"},
			[]string{"b", "a"},
			[]string{"b", "a"},
			[]string{"b", "a"}})

	sort.Strings(actual)
	expected := []string{"a", "b"}
	if len(expected) != len(actual) {
		t.Errorf("STV giving wrong number of winners. Expected %v, Got %v", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("STV giving different winners. Expected %v, Got %v", expected, actual)
		}
	}
}

// TestWikipedia tests the example in the wikipedia article
func TestWikipedia(t *testing.T) {
	actual := STVElection(
		3,
		[]string{"Oranges", "Pears", "Chocolate", "Strawberries", "Hamburgers"},
		[][]string{
			[]string{"Oranges"},
			[]string{"Oranges"},
			[]string{"Oranges"},
			[]string{"Oranges"},
			[]string{"Pears", "Oranges"},
			[]string{"Pears", "Oranges"},
			[]string{"Chocolate", "Strawberries"},
			[]string{"Chocolate", "Strawberries"},
			[]string{"Chocolate", "Strawberries"},
			[]string{"Chocolate", "Strawberries"},
			[]string{"Chocolate", "Strawberries"},
			[]string{"Chocolate", "Strawberries"},
			[]string{"Chocolate", "Strawberries"},
			[]string{"Chocolate", "Strawberries"},
			[]string{"Chocolate", "Hamburgers"},
			[]string{"Chocolate", "Hamburgers"},
			[]string{"Chocolate", "Hamburgers"},
			[]string{"Chocolate", "Hamburgers"},
			[]string{"Strawberries"},
			[]string{"Hamburgers"}})

	sort.Strings(actual)
	expected := []string{"Chocolate", "Oranges", "Strawberries"}
	if len(expected) != len(actual) {
		t.Errorf("STV giving wrong number of winners. Expected %v, Got %v", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("STV giving different winners. Expected %v, Got %v", expected, actual)
		}
	}
}

// TestCGPSimple tests the simple example in the CGP Grey Video
func TestCGPSimple(t *testing.T) {
	actual := STVElection(
		3,
		[]string{"Lynx", "Monkey", "Tarcia", "Gorilla", "Tiger"},
		[][]string{
			[]string{"Tarcia", "Gorilla"},
			[]string{"Tarcia", "Gorilla"},
			[]string{"Tarcia", "Gorilla"},
			[]string{"Tarcia", "Gorilla"},
			[]string{"Tarcia", "Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Gorilla"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Monkey"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"}})

	sort.Strings(actual)
	expected := []string{"Gorilla", "Monkey", "Tiger"}
	if len(expected) != len(actual) {
		t.Errorf("STV giving wrong number of winners. Expected %v, Got %v", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("STV giving different winners. Expected %v, Got %v", expected, actual)
		}
	}
}

// TestCGPExtended tests the extended CGP Grey video
func TestCGPExtended(t *testing.T) {
	actual := STVElection(
		5,
		[]string{
			"Tarcia",
			"Gorilla",
			"Silverback",
			"Owl",
			"Turtle",
			"Snake",
			"Tiger",
			"Lynx",
			"Jackalope",
			"Buffalo"},
		[][]string{
			[]string{"Tarcia", "Silverback"},
			[]string{"Tarcia", "Silverback"},
			[]string{"Tarcia", "Silverback"},
			[]string{"Tarcia", "Silverback"},
			[]string{"Tarcia", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Silverback"},
			[]string{"Gorilla", "Tarcia", "Silverback"},
			[]string{"Gorilla", "Tarcia", "Silverback"},
			[]string{"Gorilla", "Tarcia", "Silverback"},
			[]string{"Gorilla", "Tarcia", "Silverback"},
			[]string{"Gorilla", "Tarcia", "Silverback"},
			[]string{"Gorilla", "Tarcia", "Silverback"},
			[]string{"Gorilla", "Tarcia", "Silverback"},
			[]string{"Gorilla", "Tarcia", "Silverback"},
			[]string{"Gorilla", "Tarcia", "Silverback"},
			[]string{"Gorilla", "Tarcia", "Silverback"},
			[]string{"Silverback"},
			[]string{"Silverback"},
			[]string{"Silverback"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Owl", "Turtle"},
			[]string{"Turtle"},
			[]string{"Snake", "Turtle"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Lynx", "Tiger"},
			[]string{"Jackalope"},
			[]string{"Jackalope"},
			[]string{"Buffalo", "Jackalope"},
			[]string{"Buffalo", "Jackalope"},
			[]string{"Buffalo", "Jackalope", "Turtle"}})

	sort.Strings(actual)
	expected := []string{"Gorilla", "Owl", "Silverback", "Tiger", "Turtle"}
	if len(expected) != len(actual) {
		t.Errorf("STV giving wrong number of winners. Expected %v, Got %v", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("STV giving different winners. Expected %v, Got %v", expected, actual)
		}
	}
}
