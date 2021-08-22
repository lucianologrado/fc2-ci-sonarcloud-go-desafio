package main

import "testing"

func TestSoma(t *testing.T){

	esperado := 25
	total := soma(15,10)

	if (total != esperado){
		t.Errorf("Resultado da soma %d é invalido, era esperado %d", total, esperado)
	}
}

func TestDividir(t *testing.T){

	esperado := 4
	total := divida(8,2)

	if (total != esperado){
		t.Errorf("Resultado da operacao %d é invalido, era esperado %d", total, esperado)
	}
}

func TestMultiplica(t *testing.T){

	esperado := 16
	total := multiplica(8,2)

	if (total != esperado){
		t.Errorf("Resultado da operacao %d é invalido, era esperado %d", total, esperado)
	}
}

func TestSubtrai(t *testing.T){

	esperado := 6
	total := subtrai(8,2)

	if (total != esperado){
		t.Errorf("Resultado da operacao %d é invalido, era esperado %d", total, esperado)
	}
}

