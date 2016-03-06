package loc

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/*
	Representa un módulo del programa
*/
type Modulo struct {
	Nombre    string `json:"nombre"`
	Size      int    `json:"size"`
	Funciones int    `json:"funciones"`
}

/*
	Esta funcion recibe un path root (ruta raiz del proyecto) y devuelve un arreglo de
	string con las rutas a todos los archivos .go dentro del proyecto
*/
func addFiles(root string) ([]string, error) {

	r, _ := regexp.Compile("\\.go")

	fileList := []string{}
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {

		if !f.IsDir() && r.MatchString(path) {
			fileList = append(fileList, path)
			return nil
		}

		return err
	})

	if err != nil {

		return nil, err
	}

	return fileList, err
}

/*
	Esta funcion devuelve un arreglo de modulos(partes del programa), cada módulo contiene su
	nombre, el numero de funciones y su tamaño
*/
func ObtenerModulos(path string) ([]Modulo, int, error) {

	modulos := []Modulo{}
	total := 0
	filelist, err := addFiles(path)

	if err != nil {
		return nil, 0, err
	}

	for _, file := range filelist {

		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, 0, err
		}

		lineas, funcion := contarLineas(data)
		modulo := Modulo{Nombre: file, Size: lineas, Funciones: funcion}
		total = total + lineas
		modulos = append(modulos, modulo)
	}

	return modulos, total, nil
}

/*
	Esta función devuelve el numero de lineas y el numero de funciones de un arreglo de bytes
*/
func contarLineas(data []byte) (lineas int, nfunciones int) {

	regFunc, _ := regexp.Compile("(func) .")                                          // expresión regular para una funcion
	regvar1, _ := regexp.Compile("([a-z]+[0-9]* *,* *):=")                            // expresión regular para una declaracion de variable
	regvar2, _ := regexp.Compile("(var +[a-z]+[0-9]* *)[a-z]+")                       // expresión regular para una declaracion de variable
	regasig, _ := regexp.Compile("([a-z]+[0-9]* *)= *.")                              // expresión regular para una asignación
	regProp, _ := regexp.Compile("[a-z]+\\.")                                         // expresión regular para una propiedad, ejem : fmt.Println()
	regCall, _ := regexp.Compile("[a-z]+\\(")                                         // expresión regular para la llamada a una función, ejem : contarPalabras()
	regCom, _ := regexp.Compile("return|if|for|type|package|switch|case|else|\\_|\"") // expresión regular para los tipos mencionados

	modulo := string(data)
	lineasCodigo := 0
	funciones := 0
	linea := ""
	for _, c := range modulo {

		if c == '\n' {
			if strings.Trim(linea, " \r)}") != "" {

				switch true {

				case regFunc.MatchString(linea):
					funciones++
					lineasCodigo++

				case regvar1.MatchString(linea):
					lineasCodigo++

				case regvar2.MatchString(linea):
					lineasCodigo++

				case regasig.MatchString(linea):
					lineasCodigo++

				case regCom.MatchString(linea):
					lineasCodigo++

				case regProp.MatchString(linea):
					lineasCodigo++

				case regCall.MatchString(linea):
					lineasCodigo++
				}

			}

			linea = ""

		} else {

			linea = linea + string(c)
		}
	}

	return lineasCodigo, funciones

}
