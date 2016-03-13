package estadistica

import (
	"errors"
	"github.com/stevenyepes/pspprograms/pspprogram1/listaligada"
	"math"
)

// Pareja es una estructura que contiene los datos tanto en x como en y
type Pareja struct {
	DatosX *listaligada.ListaLigada
	DatosY *listaligada.ListaLigada
}

// Media es una función que retorna la media de una lista de datos
func Media(lista *listaligada.ListaLigada) float64 {

	suma := float64(0)
	for i := 0; i < lista.Len(); i++ {

		suma = suma + lista.Get(i).(float64)
	}

	return float64(suma / float64(lista.Len()))
}

// Desviacion es una función que retorna la desviación estándar de
// Una lista de datos
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

// Sumatoria recibe una lista de float y devuelve su suma
func Sumatoria(lista *listaligada.ListaLigada) (float64, error) {

	suma := 0.0

	if lista.Len() == 0 {
		return 0.0, errors.New("La lista debe contener al menos un dato")
	}

	for i := 0; i < lista.Len(); i++ {

		suma = suma + lista.Get(i).(float64)

	}

	return suma, nil
}

// Xpory recibe una pareja y devuelve una lista con la multiplicación de cada
// dato con su correspondiente relación
func Xpory(pareja *Pareja) (*listaligada.ListaLigada, error) {

	if pareja.DatosX.Len() == 0 || pareja.DatosY.Len() == 0 {
		return nil, errors.New("Las listas no pueden estar vacías")
	}

	if pareja.DatosX.Len() != pareja.DatosY.Len() {
		return nil, errors.New("El numero de datos de ambas listas debe ser igual")
	}

	lista := new(listaligada.ListaLigada)
	for i := 0; i < pareja.DatosX.Len(); i++ {
		xy := pareja.DatosX.Get(i).(float64) * pareja.DatosY.Get(i).(float64)
		lista.Push(xy)
	}

	return lista, nil

}

// ListaCuadrado recibe una lista y devuelve una lista con los datos elevados al
// cuadrado
func ListaCuadrado(lista *listaligada.ListaLigada) (*listaligada.ListaLigada, error) {

	if lista.Len() == 0 {

		return nil, errors.New("La lista debe contener al menos un dato")
	}

	listaCuadrado := new(listaligada.ListaLigada)
	for i := 0; i < lista.Len(); i++ {
		xcuadrado := math.Pow(lista.Get(i).(float64), 2)
		lista.Push(xcuadrado)
	}

	return listaCuadrado, nil

}

// BetaParametros recibe una pareja y calcula el parámetro b1 y b0
// para calcular la regresión lineal
func BetaParametros(pareja *Pareja) (float64, float64, error) {

	if pareja.DatosX.Len() == 0 || pareja.DatosY.Len() == 0 {
		return 0.0, 0.0, errors.New("Las listas no pueden estar vacías")
	}

	if pareja.DatosX.Len() != pareja.DatosY.Len() {
		return 0.0, 0.0, errors.New("El numero de datos de ambas listas debe ser igual")
	}

	n := float64(pareja.DatosX.Len())

	xpory, err1 := Xpory(pareja)
	if err1 != nil {

		return 0.0, 0.0, err1
	}

	sumaXporY, err2 := Sumatoria(xpory)
	if err2 != nil {
		return 0.0, 0.0, err2
	}

	mediax := Media(pareja.DatosX)
	mediay := Media(pareja.DatosY)

	xcuadrado, _ := ListaCuadrado(pareja.DatosX)
	sumaxcuadrado, _ := Sumatoria(xcuadrado)

	beta1 := (sumaXporY - n*(mediax*mediay)/(sumaxcuadrado-(n*math.Pow(mediax, 2))))

	beta0 := mediay - beta1*mediax
	return beta1, beta0, nil

}

// CoeficienteR recibe una pareja y devuelve el coeficiente de correlación 'r'
// y 'r^2'
func CoeficienteR(pareja *Pareja) (float64, float64, error) {

	if pareja.DatosX.Len() == 0 || pareja.DatosY.Len() == 0 {
		return 0.0, 0.0, errors.New("Las listas no pueden estar vacías")
	}

	if pareja.DatosX.Len() != pareja.DatosY.Len() {
		return 0.0, 0.0, errors.New("El numero de datos de ambas listas debe ser igual")
	}

	n := float64(pareja.DatosX.Len())

	xpory, _ := Xpory(pareja)
	sumaXporY, _ := Sumatoria(xpory)

	sumax, _ := Sumatoria(pareja.DatosX)
	sumay, _ := Sumatoria(pareja.DatosY)

	xcuadrado, _ := ListaCuadrado(pareja.DatosX)
	sumaxcuadrado, _ := Sumatoria(xcuadrado)

	ycuadrado, _ := ListaCuadrado(pareja.DatosY)
	sumaycuadrado, _ := Sumatoria(ycuadrado)

	sumatoriaXcuadrado := math.Pow(sumax, 2)
	sumatoriaYcuadrado := math.Pow(sumay, 2)

	raiz := math.Sqrt((n*sumaxcuadrado - sumatoriaXcuadrado) * (n*sumaycuadrado - sumatoriaYcuadrado))

	r := (n*sumaXporY - sumax*sumay) / (raiz)
	rcuadrado := math.Pow(r, 2)

	return r, rcuadrado, nil
}
