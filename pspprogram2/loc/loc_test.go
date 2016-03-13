package loc

import (
	"io/ioutil"
	"regexp"
	"testing"
)

// Test para verificar que se listen los archivos de un directorio
func TestAddfiles(t *testing.T) {

	_, error := addFiles("C:/Users/Steven/go_projects/src/github.com/stevenyepes/pspprograms/pspprogram1")

	if error != nil {

		t.Errorf("Error, se esperaba err nul en la variable error, se obtuvo %q", error)
	}
}

// Test para verificar que se listen solo los archivos .go
func TestAddfilesJustGo(t *testing.T) {

	files, error := addFiles("C:/Users/Steven/go_projects/src/github.com/stevenyepes/pspprograms/pspprogram1")

	r, _ := regexp.Compile("\\.go")

	for _, file := range files {

		if error != nil && !r.MatchString(file) {

			t.Errorf("Error, existen archivos con extendion diferente a .go, se obtuvo %q", r.MatchString(file))
			break
		}
	}
}

//Test para verificar la función contar lineas
func TestContarLineas(t *testing.T) {

	data, _ := ioutil.ReadFile("C:/Users/Steven/Downloads/UdeA-PS-Prog2/UdeA-PS-Prog2/archivo.go")
	count, funcion := contarLineas(data)
	if count != 27 || funcion != 4 {

		t.Errorf("Error, se esperaba %d numero de lineas y %d en el numero de funciones, se obtuvo %d y %d", 21, 3, count, funcion)
	}

}

// Test para verificar que se cuenten todos los archivos .go de un proyecto
func TestObtenerModulos(t *testing.T) {

	_, _, err := ObtenerModulos("C:/Users/Steven/Downloads/UdeA-PS-Prog2/UdeA-PS-Prog2/project")

	if err != nil {

		t.Errorf("Error, se obtuvo un error al obtener los modulos del proyecto, error: %q", err)
	}

}
