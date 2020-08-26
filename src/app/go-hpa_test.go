package main

import "testing"

func TestSoma(t *testing.T)  {

	got := calcular_raiz()
	want := 2499999414254808.5

	if got != want {
		t.Errorf("funcao calcular_raiz não esta certa.\n devolveu: %v \n esperado: %v", got, want)
	}
	
}