package estadistica

import (
	"github.com/stevenyepes/pspprograms/pspprogram1/listaligada"
	"testing"
)

func TestMedia1(t *testing.T) {

	lista := new(listaligada.ListaLigada)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)

	if Media(lista) != 1 {

		t.Errorf("Error, se esperaba %q al obtener la media, se obtuvo %q", 1, Media(lista))
	}
}

func TestMedia(t *testing.T) {

	lista := new(listaligada.ListaLigada)
	lista.Push(1.0)
	lista.Push(2.0)
	lista.Push(3.0)
	lista.Push(1.0)
	lista.Push(2.0)
	lista.Push(3.0)

	if Media(lista) != 2 {

		t.Errorf("Error, se esperaba %q al obtener la media, se obtuvo %q", 2, Media(lista))
	}
}

func TestMediaPSPData1(t *testing.T) {
	lista := new(listaligada.ListaLigada)
	lista.Push(160.0)
	lista.Push(591.0)
	lista.Push(114.0)
	lista.Push(229.0)
	lista.Push(230.0)
	lista.Push(270.0)
	lista.Push(128.0)
	lista.Push(1657.0)
	lista.Push(624.0)
	lista.Push(1503.0)

	if Media(lista) != 550.6 {
		t.Errorf("Error, se esperaba %f al obtener la media, se obtuvo %f", 550.6, Media(lista))
	}
}

func TestMediaPSPData2(t *testing.T) {

	lista := new(listaligada.ListaLigada)
	lista.Push(15.0)
	lista.Push(69.9)
	lista.Push(6.5)
	lista.Push(22.4)
	lista.Push(28.4)
	lista.Push(65.9)
	lista.Push(19.4)
	lista.Push(198.7)
	lista.Push(38.8)
	lista.Push(138.2)

	if Media(lista) < 60.32 || Media(lista) > 60.33 {
		t.Errorf("Error, se esperaba %f al obtener la media, se obtuvo %f", 60.32, Media(lista))
	}
}

func TestDesvEstandar(t *testing.T) {
	lista := new(listaligada.ListaLigada)
	lista.Push(1.0)
	lista.Push(2.0)
	lista.Push(3.0)
	lista.Push(1.0)
	lista.Push(2.0)
	lista.Push(3.0)

	if Desviacion(lista) < 0.88 || Desviacion(lista) > 0.90 {
		t.Errorf("Error, se esperaba %f al obtener la desviación estandar, se obtuvo %f", 0.89, Desviacion(lista))
	}
}

func TestDesvEstandarPSPData1(t *testing.T) {

	lista := new(listaligada.ListaLigada)
	lista.Push(160.0)
	lista.Push(591.0)
	lista.Push(114.0)
	lista.Push(229.0)
	lista.Push(230.0)
	lista.Push(270.0)
	lista.Push(128.0)
	lista.Push(1657.0)
	lista.Push(624.0)
	lista.Push(1503.0)

	if Desviacion(lista) < 572.02 || Desviacion(lista) > 572.03 {
		t.Errorf("Error, se esperaba %f al obtener la desviación estandar, se obtuvo %f", 572.026845, Desviacion(lista))
	}
}

func TestDesvEstandarPSPData2(t *testing.T) {

	lista := new(listaligada.ListaLigada)
	lista.Push(15.0)
	lista.Push(69.9)
	lista.Push(6.5)
	lista.Push(22.4)
	lista.Push(28.4)
	lista.Push(65.9)
	lista.Push(19.4)
	lista.Push(198.7)
	lista.Push(38.8)
	lista.Push(138.2)

	if Desviacion(lista) < 62.24 || Desviacion(lista) > 62.26 {
		t.Errorf("Error, se esperaba %f al obtener la desviación estandar, se obtuvo %f", 62.255831, Desviacion(lista))
	}
}

func TestSumatoria(t *testing.T) {
	lista := new(listaligada.ListaLigada)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	suma, _ := Sumatoria(lista)

	if suma != 10.0 {
		t.Errorf("Error, se esperaba %f al obtener la sumatria, se obtuvo %f", 10.0, suma)
	}
}

func TestXporY(t *testing.T) {
	lista := new(listaligada.ListaLigada)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(2.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)

	lista2 := new(listaligada.ListaLigada)
	lista2.Push(1.0)
	lista2.Push(1.0)
	lista2.Push(1.0)
	lista2.Push(2.0)
	lista2.Push(1.0)
	lista2.Push(1.0)
	lista2.Push(1.0)
	lista2.Push(1.0)
	lista2.Push(1.0)
	lista2.Push(1.0)
	pareja := Pareja{lista, lista2}

	xpory, _ := Xpory(pareja)
	suma, _ := Sumatoria(xpory)
	if suma != 13.0 {
		t.Errorf("Error, se esperaba %f al obtener la sumatria, se obtuvo %f", 10.0, suma)
	}
}

func TestListaCuadrado(t *testing.T) {
	lista := new(listaligada.ListaLigada)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(2.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)
	lista.Push(1.0)

	listacuadrado, _ := ListaCuadrado(lista)
	suma, _ := Sumatoria(listacuadrado)
	if suma != 13.0 {
		t.Errorf("Error, se esperaba %f al obtener la sumatria, se obtuvo %f", 13.0, suma)
	}
}
