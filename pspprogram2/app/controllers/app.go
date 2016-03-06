package controllers

import (
	"github.com/revel/revel"
	"github.com/stevenyepes/pspprograms/pspprogram2/loc"
	"regexp"
)

type Modulos struct {
	Modulo []loc.Modulo `json:"modulo"`
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "PSP program 2"
	return c.Render(greeting)
}

func (c App) ListaModulos(archivo string) revel.Result {

	c.Validation.Required(archivo).Message("Ingresa el directorio de tu proyecto Go!")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	r, _ := regexp.Compile("([A-Z|a-z]:\\[^*|\"<>?\n]*)|(\\\\.*?\\.*)")

	if r.MatchString(archivo) {
		modulos, total, _ := loc.ObtenerModulos(archivo)
		return c.Render(modulos, total)
	} else {

		return c.Redirect(App.Index)
	}

}

func (c App) Help() revel.Result {

	return c.Render()
}
