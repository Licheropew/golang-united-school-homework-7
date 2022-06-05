package coverage

import (
	"errors"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
func TestMain(m *testing.M) {
	log.Println("START TESTING")
	code := m.Run()
	log.Println("TEST END")
	os.Exit(code)
}

func TestLenPeople(t *testing.T) {
	tests := map[string]struct {
		input  People
		output int
	}{
		"Empty": {People{}, 0},
		"One": {People{
			{firstName: "Anton", lastName: "A"},
		}, 1},
		"Two": {People{
			{firstName: "Anton", lastName: "A"},
			{firstName: "Anton", lastName: "B"},
		}, 2},
		"Three": {People{
			{firstName: "Anton", lastName: "A"},
			{firstName: "Anton", lastName: "B"},
			{firstName: "Anton", lastName: "C"},
		}, 3},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			res := testCase.input.Len()
			if res != testCase.output {
				t.Errorf("got %d, expect %d", res, testCase.output)
			}
		})
	}
}

func TestLessPeople(t *testing.T) {
	birthday := time.Date(1990, 3, 22, 0, 0, 0, 0, time.UTC)
	tests := map[string]struct {
		input  People
		output bool
		i, j   int
	}{
		"Birthday and first name even, last name i < j": {People{
			{firstName: "Anton", lastName: "A", birthDay: birthday},
			{firstName: "Anton", lastName: "B", birthDay: birthday},
		}, true, 0, 1},
		"Birthday and first name even, last name i > j": {People{
			{firstName: "Anton", lastName: "A", birthDay: birthday},
			{firstName: "Anton", lastName: "B", birthDay: birthday},
		}, false, 1, 0},
		"Birthday, first and last name even (i==j)": {People{
			{firstName: "Anton", lastName: "A", birthDay: birthday},
		}, false, 0, 0},
		"Birthday even, first name i < j": {People{
			{firstName: "Anton", birthDay: birthday},
			{firstName: "Notna", birthDay: birthday},
		}, true, 0, 1},
		"Birthday even, first name i > j": {People{
			{firstName: "Anton", birthDay: birthday},
			{firstName: "Notna", birthDay: birthday},
		}, false, 1, 0},
		"Birthday i > j": {People{
			{firstName: "Anton", birthDay: birthday.AddDate(0, 1, 0)},
			{firstName: "Notna", birthDay: birthday},
		}, true, 0, 1},
		"Birthday i < j": {People{
			{firstName: "Anton", birthDay: birthday.AddDate(0, 1, 0)},
			{firstName: "Notna", birthDay: birthday},
		}, false, 1, 0},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			res := testCase.input.Less(testCase.i, testCase.j)
			if res != testCase.output {
				t.Errorf("got %v, expect %v", res, testCase.output)
			}
		})
	}
}

func TestSwapFunc(t *testing.T) {
	tests := map[string]struct {
		input  People
		output People
		i, j   int
	}{
		"Positive case": {
			People{
				{firstName: "Anton", lastName: "A"},
				{firstName: "Anton", lastName: "B"},
			},
			People{
				{firstName: "Anton", lastName: "B"},
				{firstName: "Anton", lastName: "A"},
			}, 0, 1},
		"Case where i==j": {
			People{
				{firstName: "Anton", lastName: "A"},
				{firstName: "Anton", lastName: "B"},
			},
			People{
				{firstName: "Anton", lastName: "A"},
				{firstName: "Anton", lastName: "B"},
			}, 0, 0},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			testCase.input.Swap(testCase.i, testCase.j)
			if testCase.i != testCase.j {
				switch {
				case testCase.input[testCase.i] == testCase.output[testCase.j]:
					t.Errorf("swap func goes wrong, expect %v, got %v", testCase.output[testCase.j], testCase.input[testCase.i])
				case testCase.input[testCase.j] == testCase.output[testCase.i]:
					t.Errorf("swap func goes wrong, expect %v, got %v", testCase.output[testCase.i], testCase.input[testCase.j])
				}
			} else {
				if testCase.input[testCase.i] != testCase.output[testCase.j] {
					t.Errorf("swap func goes wrong, expect %v, got %v", testCase.output[testCase.j], testCase.input[testCase.i])
				}
			}
		})
	}
}

func TestNewMatrix(t *testing.T) {
	var tErr error = errors.New("Rows need to be the same length")
	var numError *strconv.NumError
	tests := map[string]struct {
		input  string
		output *Matrix
		err    error
		atoiErr bool
	}{
		"Positive case, one row": {input: "1 2 3 4 5", output: &Matrix{rows: 1, cols: 5, data: []int{1, 2, 3, 4, 5}}, err: nil},
		"Positive case, three rows": {input: "1 2 3\n1 2 3\n1 2 3", output: &Matrix{rows: 3, cols: 3, data: []int{1, 2, 3, 1, 2, 3, 1, 2, 3}}, err: nil},
		"Positive case, five rows": {input: "1\n1\n1\n1\n1", output: &Matrix{rows: 5, cols: 1, data: []int{1, 1, 1, 1, 1}}, err: nil},
		"Empty string": {input: "", output: nil, err: numError, atoiErr: true},
		"String with letters": {input: "1 2 3 4 a", output: nil, err: numError, atoiErr: true},
		"Different cols": {input: "1 2 3\n1 2", output: nil, err: tErr},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			res, err := New(testCase.input)
			if testCase.err != nil {
				if testCase.atoiErr {
					if !errors.As(err, &numError) {
						t.Errorf("returned wrong error type: got %T, want %T", err, testCase.err)
					}
				} else {
					if err.Error() != testCase.err.Error() {
						t.Errorf("returned wrong error: got %s, want %s", err, testCase.err)
					}
				}
			} else {
				if len(res.data) != len(testCase.output.data) {
					t.Errorf("rows: got: %v with len %d, want %v with len %d", res.data, len(res.data), testCase.output.data, len(testCase.output.data))
				}
				if err != nil {
					t.Errorf("error returned while not awaited: %s", err.Error())
				}
				if res.rows != testCase.output.rows {
					t.Errorf("rows: got: %v, want %v", res.rows, testCase.output.rows)
				}
				if res.cols != testCase.output.cols {
					t.Errorf("cols: got: %v, want %v", res.cols, testCase.output.cols)
				}
			}			
		})
	}
}
