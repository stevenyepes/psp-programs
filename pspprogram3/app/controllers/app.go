package controllers

import (
	"github.com/revel/revel"
	"github.com/stevenyepes/pspprograms/pspprogram1/estadistica"
	"github.com/stevenyepes/pspprograms/pspprogram3/parejacnv"
	"strconv"
)

var pareja []parejacnv.Factor

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "PSP program 3"

	return c.Render(greeting)
}

func (c App) Estadistica(datox string, datoy string) revel.Result {
	c.Validation.Required(datox).Message("Por favor ingrese un Xi")
	c.Validation.Required(datoy).Message("Por favor ingrese un Yi")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	datoxi, _ := strconv.ParseFloat(datox, 64)
	datoyi, _ := strconv.ParseFloat(datoy, 64)
	pareja = append(pareja, parejacnv.Factor{DatoX: datoxi, DatoY: datoyi})
	return c.Render(pareja)
}

func (c App) Resultados(xParametro string) revel.Result {
	x, _ := strconv.ParseFloat(xParametro, 64)
	datos := parejacnv.ParejaCnv(pareja)
	beta1, beta0, yk, _ := estadistica.BetaParametros(datos, x)
	r, r2, _ := estadistica.CoeficienteR(datos)

	return c.Render(beta1, beta0, yk, r, r2)
}
