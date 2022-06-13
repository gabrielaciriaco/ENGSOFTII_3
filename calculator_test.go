package main

import "testing"

func TestPlus(t *testing.T) {
	expectedTotal := 10
	plus := plus(5, 5)
	if plus != expectedTotal {
		t.Errorf("Soma incorreta, obtido: %d, esperado: %d", plus, expectedTotal)
	}
}

func TestSub(t *testing.T) {
	expectedTotal := 10
	sub := sub(15, 5)
	if sub != expectedTotal {
		t.Errorf("Soma incorreta, obtido: %d, esperado: %d", sub, expectedTotal)
	}
}

func TestMult(t *testing.T) {
	expectedTotal := 25
	mult := mult(5, 5)
	if mult != expectedTotal {
		t.Errorf("Soma incorreta, obtido: %d, esperado: %d", mult, expectedTotal)
	}
}
