package main

import "testing"

func TestSoma(t *testing.T){

	esperado := 25
	total := soma(15,10)

	if (total != esperado){
		t.Errorf("Resultado da soma %d é invalido, era esperado %d", total, esperado)
	}
}
