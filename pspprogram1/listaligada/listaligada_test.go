package listaligada

import (
	"testing"
)

func TestGetDataNode4(t *testing.T) {

	nodo := &Node{4, new(Node), new(Node)}

	if nodo.data != 4 {

		t.Errorf("Error obteniendo dato, se esperaba %q, se obtuvo %q", 4, nodo.data)
	}

}

func TestGetNodeNext(t *testing.T) {

	nodo := &Node{4, new(Node), new(Node)}
	next := &Node{5, new(Node), new(Node)}
	nodo.next = next

	if nodo.next.data != 5 {
		t.Errorf("Error obteniendo dato Next, se esperaba %q, se obtuvo %q", 5, nodo.next.data)
	}
}

func TestGetNodePrev(t *testing.T) {

	nodo := &Node{4, new(Node), new(Node)}
	prev := &Node{5, new(Node), new(Node)}
	nodo.prev = prev

	if nodo.prev.data != 5 {
		t.Errorf("Error obteniendo dato Next, se esperaba %q, se obtuvo %q", 5, nodo.next.data)
	}
}

func TestPushDataListaLigada(t *testing.T) {

	lista := &ListaLigada{nil, nil}

	lista.Push(4)
	lista.Push(6)

	if lista.lastNode.data != 6 {

		t.Errorf("Error, se esperaba %q al final de la lista, se obtuvo %q", 6, lista.lastNode.data)
	}

}

func TestInsertBeginingListaLigada(t *testing.T) {

	lista := &ListaLigada{nil, nil}

	lista.InsertBegining(4)
	lista.InsertBegining(6)

	if lista.firstNode.data != 6 {

		t.Errorf("Error, se esperaba %q al final de la lista, se obtuvo %q", 6, lista.lastNode.data)
	}
}

func TestGetListaLigada(t *testing.T) {
	lista := new(ListaLigada)
	lista.Push(3)
	lista.Push(4)
	lista.Push(6)
	lista.Push(7)
	lista.Push(9)

	if lista.Get(3) != 7 {

		t.Errorf("Error, se esperaba %q al obtener el dato en la pos(3), se obtuvo %q", 7, lista.Get(3))
	}
}

func TestLenListaLigada(t *testing.T) {
	lista := new(ListaLigada)
	lista.Push(3)
	lista.Push(4)
	lista.Push(6)
	lista.Push(7)
	lista.Push(9)
	lista.Push(9)
	lista.Push(9)

	if lista.Len() != 7 {

		t.Errorf("Error, se esperaba %q al obtener el tamaño de la lista, se obtuvo %q", 7, lista.Len())
	}

}

func TestLenListaLigadaVacia(t *testing.T) {
	lista := new(ListaLigada)

	if lista.Len() != 0 {

		t.Errorf("Error, se esperaba %q al obtener el tamaño de la lista, se obtuvo %q", 0, lista.Len())
	}

}

func TestDeleteListaLigada(t *testing.T) {
	lista := new(ListaLigada)
	lista.Push(3)
	lista.Push(4)
	lista.Push(6)
	lista.Push(7)
	lista.Push(9)

	lenActual := lista.Len()
	lista.Delete()

	if lenActual == lista.Len() {

		t.Errorf("Error, se esperaba %q al obtener el tamaño de la lista, se obtuvo %q", lista.Len()-1, lista.Len())
	}
}
