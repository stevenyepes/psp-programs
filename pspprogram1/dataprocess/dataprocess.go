package dataprocess

import (
	"github.com/stevenyepes/pspprograms/pspprogram1/listaligada"
	"regexp"
	"strconv"
)

func LoadData(data []byte) *listaligada.ListaLigada {

	if data == nil {
		return nil
	}

	lista := new(listaligada.ListaLigada)

	text := string(data)

	nums := ""

	r, _ := regexp.Compile("([a-z]+)")

	for _, c := range text {

		if r.MatchString(string(c)) {
			break
		}

		if c == ',' {

			if num, err := strconv.ParseFloat(nums, 64); err == nil {
				lista.Push(num)
				nums = ""
			}

		} else {

			nums = nums + string(c)
		}
	}

	return lista

}
