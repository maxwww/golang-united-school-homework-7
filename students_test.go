package coverage

import (
	"os"
	"testing"
	"time"
)

const (
	defaultMatrix      = "1 2 3\n4 5 6"
	matrixIsNotCreated = "matrix is not created"
	errExpected        = "should be error"
	rowsError          = "Rows doesn't work properly"
	setError           = "Set doesn't work properly"
	lessError          = "Less doesn't work properly"
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
func TestNewMatrix(t *testing.T) {
	matrix, err := New(defaultMatrix)
	if err != nil {
		t.Errorf(matrixIsNotCreated)
	}
	if matrix.cols != 3 {
		t.Errorf("cols is wrong")
	}
	if matrix.rows != 2 {
		t.Errorf("rows is wrong")
	}
}

func TestNewEmptyMatrix(t *testing.T) {
	_, err := New("")
	if err == nil {
		t.Errorf(errExpected)
	}
}

func TestNewWrongMatrix(t *testing.T) {
	_, err := New("1 2 3\n4 5")
	if err == nil {
		t.Errorf(errExpected)
	}
}

func TestMatrixRows(t *testing.T) {
	matrix, err := New(defaultMatrix)
	if err != nil {
		t.Errorf(matrixIsNotCreated)
	}

	expectedRows := [][]int{{1, 2, 3}, {4, 5, 6}}
	rows := matrix.Rows()
	for rowIndex, rowValue := range rows {
		for columnIndex, columnValue := range rowValue {
			if columnValue != expectedRows[rowIndex][columnIndex] {
				t.Errorf(rowsError)
			}
		}
	}
}

func TestMatrixCols(t *testing.T) {
	matrix, err := New(defaultMatrix)
	if err != nil {
		t.Errorf(matrixIsNotCreated)
	}

	expectedCols := [][]int{{1, 4}, {2, 5}, {3, 6}}
	cols := matrix.Cols()
	for columnIndex, columnValue := range cols {
		for rowIndex, rowValue := range columnValue {
			if rowValue != expectedCols[columnIndex][rowIndex] {
				t.Errorf(rowsError)
			}
		}
	}
}

func TestMatrixSet(t *testing.T) {
	matrix, err := New(defaultMatrix)
	if err != nil {
		t.Errorf(matrixIsNotCreated)
	}

	result := matrix.Set(0, 2, 55)

	if !result || matrix.data[2] != 55 {
		t.Errorf(setError)
	}
}

func TestMatrixSetWrongRow(t *testing.T) {
	matrix, err := New(defaultMatrix)

	if err != nil {
		t.Errorf(matrixIsNotCreated)
	}
	result := matrix.Set(-1, 2, 55)

	if result {
		t.Errorf(setError)
	}
}

func TestMatrixSetBigRow(t *testing.T) {
	matrix, err := New(defaultMatrix)

	if err != nil {
		t.Errorf(matrixIsNotCreated)
	}
	result := matrix.Set(150, 2, 55)

	if result {
		t.Errorf(setError)
	}
}

func TestMatrixSetWrongCol(t *testing.T) {
	matrix, err := New(defaultMatrix)

	if err != nil {
		t.Errorf(matrixIsNotCreated)
	}
	result := matrix.Set(0, -10, 55)

	if result {
		t.Errorf(setError)
	}
}

func TestMatrixSetBigCol(t *testing.T) {
	matrix, err := New(defaultMatrix)

	if err != nil {
		t.Errorf(matrixIsNotCreated)
	}
	result := matrix.Set(1, 140, 55)

	if result {
		t.Errorf(setError)
	}
}

func TestPeopleLen(t *testing.T) {
	now := time.Now()
	people := People{{
		firstName: "AA",
		lastName:  "BB",
		birthDay:  now,
	}, {
		firstName: "AA",
		lastName:  "BB",
		birthDay:  now,
	}}
	length := people.Len()
	if length != 2 {
		t.Errorf("Len doesn't work properly")
	}
}

func TestPeopleLessFirstName(t *testing.T) {
	now := time.Now()
	people := People{{
		firstName: "AA",
		lastName:  "BB",
		birthDay:  now,
	}, {
		firstName: "BB",
		lastName:  "BB",
		birthDay:  now,
	}}

	result := people.Less(0, 1)
	if !result {
		t.Errorf(lessError)
	}
}

func TestPeopleLessLastName(t *testing.T) {
	now := time.Now()
	people := People{{
		firstName: "AA",
		lastName:  "AA",
		birthDay:  now,
	}, {
		firstName: "AA",
		lastName:  "BB",
		birthDay:  now,
	}}

	result := people.Less(0, 1)
	if !result {
		t.Errorf(lessError)
	}
}

func TestPeopleLessLastBirthDay(t *testing.T) {
	now := time.Now()
	people := People{{
		firstName: "AA",
		lastName:  "BB",
		birthDay:  now,
	}, {
		firstName: "AA",
		lastName:  "BB",
		birthDay:  now.Add(1 * time.Second),
	}}

	result := people.Less(0, 1)
	if result {
		t.Errorf(lessError)
	}
}

func TestPeopleSwap(t *testing.T) {
	now := time.Now()
	people := People{{
		firstName: "AA",
		lastName:  "BB",
		birthDay:  now,
	}, {
		firstName: "CC",
		lastName:  "DD",
		birthDay:  now.Add(1 * time.Second),
	}}

	people.Swap(0, 1)

	if people[0].firstName != "CC" || people[0].lastName != "DD" || people[1].firstName != "AA" || people[1].lastName != "BB" {
		t.Errorf("Swap doesn't work properly")
	}
}
