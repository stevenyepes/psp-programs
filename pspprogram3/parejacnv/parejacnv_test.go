package parejacnv

import (
	"testing"
)

// TestParejaConv es un test que verifica que los datos en la estructura Pareja
// Si correspondan con los datos de el array de Factores
func TestParejaConv(t *testing.T) {

	factores := []Factor{}
	factores = append(factores, Factor{DatoX: 3.8, DatoY: 17.6})
	factores = append(factores, Factor{DatoX: 3.7, DatoY: 27.6})
	factores = append(factores, Factor{DatoX: 5.9, DatoY: 37.6})
	factores = append(factores, Factor{DatoX: 37.8, DatoY: 47.6})

	pareja := ParejaCnv(factores)

	cases := []struct {
		in, want float64
	}{
		{3.8, 17.6},
		{3.7, 27.6},
		{5.9, 37.6},
		{37.8, 47.6},
	}
	for i, c := range cases {

		got := pareja.DatosY.Get(i)
		if got != c.want {
			t.Errorf("Se esperaba la pareja(%f,%f), se obtuvo(%f,%f)", c.in, c.want, c.in, got)
		}
	}

}
