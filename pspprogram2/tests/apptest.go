package tests

import "github.com/revel/revel/testing"

//AppTest es una estructura que contiene todas las funciones de TestSuite
type AppTest struct {
	testing.TestSuite
}

//Before se ejecuta antes de inicializar los test
func (t *AppTest) Before() {
	println("Set up")
}

// TestThatIndexPageWorks test para verificar que se cargue correctamente Index
func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

// TestThatHelpPageWorks test para verificar que se cargue correctamente la página de ayuda
func (t *AppTest) TestThatHelpPageWorks() {
	t.Get("/App/Help")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

// After se ejecuta luego de terminar de ejecutar los tests
func (t *AppTest) After() {
	println("Tear down")
}
