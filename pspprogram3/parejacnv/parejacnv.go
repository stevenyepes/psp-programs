package parejacnv

import (
	"github.com/stevenyepes/pspprograms/pspprogram1/estadistica"
	"github.com/stevenyepes/pspprograms/pspprogram1/listaligada"
)

// Factor es una estructura que representa un par de datos ordenados
type Factor struct {
	DatoX float64
	DatoY float64
}

// ParejaCnv recibe un arreglo de fac
func ParejaCnv(array []Factor) estadistica.Pareja {

	pareja := estadistica.Pareja{DatosX: new(listaligada.ListaLigada), DatosY: new(listaligada.ListaLigada)}
	for _, data := range array {

		pareja.DatosX.Push(data.DatoX)
		pareja.DatosY.Push(data.DatoY)
	}

	return pareja

}
