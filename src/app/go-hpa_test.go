package main

import "testing"

func TestSoma(t *testing.T)  {

	got := calcular_raiz()
	want := 2.4999528710966443e+11

			

	if got != want {
		t.Errorf("funcao calcular_raiz não esta certa.\n devolveu: %v \n esperado: %v", got, want)
	}
	
}