package controllers

import (
	"github.com/revel/revel"
	"github.com/stevenyepes/pspprograms/pspprogram1/dataprocess"
	"github.com/stevenyepes/pspprograms/pspprogram1/estadistica"
	"github.com/stevenyepes/pspprograms/pspprogram1/listaligada"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "PSP program 1"
	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("El nombre es requerido!")
	c.Validation.MinSize(myName, 3).Message("El nombre debe ser mas largo")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}

func (c App) Estadistica(archivo []byte) revel.Result {

	c.Validation.Required(archivo).Message("Selecciona un archivo!")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	lista := new(listaligada.ListaLigada)
	lista = dataprocess.LoadData(archivo)
	media := estadistica.Media(lista)
	desv := estadistica.Desviacion(lista)
	return c.Render(media, desv)

}

func (c App) Help() revel.Result {

	return c.Render()
}
