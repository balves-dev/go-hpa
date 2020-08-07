package main

import "testing"

func TestRaizQuadrada(t *testing.T) {
	resultado := raizQuadrada(9)
	if resultado != 3 {
		t.Error("O resultado deve ser 3", resultado)
	}

}
