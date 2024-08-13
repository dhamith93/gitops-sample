package maths_test

import (
	"testing"

	"gitlab.local/dhamith93/devops-playground/app/internal/maths"
)

func TestIsOp(t *testing.T) {
	exp := maths.Expression{}
	got := exp.IsOp("+")
	if got != true {
		t.Errorf("exp.IsOp(+) = %v; want true", got)
	}
}

func TestIsOp2(t *testing.T) {
	exp := maths.Expression{}
	got := exp.IsOp("-")
	if got != true {
		t.Errorf("exp.IsOp(-) = %v; want true", got)
	}
}

func TestIsOp3(t *testing.T) {
	exp := maths.Expression{}
	got := exp.IsOp("*")
	if got != true {
		t.Errorf("exp.IsOp(*) = %v; want true", got)
	}
}

func TestIsOp4(t *testing.T) {
	exp := maths.Expression{}
	got := exp.IsOp("/")
	if got != true {
		t.Errorf("exp.IsOp(/) = %v; want true", got)
	}
}

func TestIsOp5(t *testing.T) {
	exp := maths.Expression{}
	got := exp.IsOp("2")
	if got != false {
		t.Errorf("exp.IsOp(2) = %v; want false", got)
	}
}

func TestOrder(t *testing.T) {
	exp := maths.Expression{}
	got := exp.Order("+")
	if got != 1 {
		t.Errorf("exp.Order(+) = %v; want 1", got)
	}
}

func TestOrder2(t *testing.T) {
	exp := maths.Expression{}
	got := exp.Order("-")
	if got != 1 {
		t.Errorf("exp.Order(-) = %v; want 1", got)
	}
}

func TestOrder3(t *testing.T) {
	exp := maths.Expression{}
	got := exp.Order("*")
	if got != 2 {
		t.Errorf("exp.Order(*) = %v; want 2", got)
	}
}

func TestOrder4(t *testing.T) {
	exp := maths.Expression{}
	got := exp.Order("/")
	if got != 2 {
		t.Errorf("exp.Order(/) = %v; want 2", got)
	}
}

func TestParse(t *testing.T) {
	exp := maths.Expression{}
	exp.Parse("1+1+1")
	exp2 := maths.Expression{}
	exp2.Elements = []maths.Node{
		{Value: "1", IsOperation: false},
		{Value: "1", IsOperation: false},
		{Value: "+", IsOperation: true},
		{Value: "1", IsOperation: false},
		{Value: "+", IsOperation: true},
	}
	for i, e := range exp.Elements {
		if e.Value != exp2.Elements[i].Value {
			t.Errorf("exp.Parse(1+1+1) = %v; want %v", exp.Elements, exp2.Elements)
		}
	}
}

func TestSolve(t *testing.T) {
	exp := maths.Expression{}
	exp.Parse("1+1+1")
	exp.Solve()
	if exp.Error != nil {
		t.Errorf("exp.Solve(1+1+1) got error %v", exp.Error)
	}
	got := exp.Result

	if got != "3" {
		t.Errorf("exp.Solve(1+1+1) = %v; want %v", got, "3")
	}
}
