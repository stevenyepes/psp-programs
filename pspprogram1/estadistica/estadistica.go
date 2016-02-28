package estadistica

import (
	"github.com/stevenyepes/pspprograms/pspprogram1/listaligada"
	"math"
)

func Media(lista *listaligada.ListaLigada) float64 {

	suma := float64(0)
	for i := 0; i < lista.Len(); i++ {

		suma = suma + lista.Get(i).(float64)
	}

	return float64(suma / float64(lista.Len()))
}

func Desviacion(lista *listaligada.ListaLigada) float64 {

	suma := float64(0)
	media := Media(lista)
	for i := 0; i < lista.Len(); i++ {
		suma = suma + math.Pow(lista.Get(i).(float64)-media, 2)
	}

	raiz := suma / float64(lista.Len()-1)

	desv := math.Sqrt(raiz)

	return desv
}
