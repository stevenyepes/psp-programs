package listaligada

import ()

type Node struct {
	data       interface{}
	next, prev *Node
}

type ListaLigada struct {
	firstNode, lastNode *Node
}

func (l *ListaLigada) Push(data interface{}) {
	last := new(Node)
	last.data = data

	if l.firstNode == nil {
		l.firstNode = last
		l.lastNode = last
	} else {

		l.lastNode.next = last
		last.prev = l.lastNode
		l.lastNode = last
	}

}

func (l *ListaLigada) InsertBegining(data interface{}) {

	first := new(Node)
	first.data = data

	if l.firstNode == nil {
		l.firstNode = first
		l.lastNode = first
	} else {

		l.firstNode.prev = first
		first.next = l.firstNode
		l.firstNode = first
	}
}

func (l *ListaLigada) Get(index int) interface{} {

	nodoAux := new(Node)
	nodoAux = l.firstNode

	count := 0
	for count < index {

		nodoAux = nodoAux.next
		count++
	}

	return nodoAux.data
}

func (l *ListaLigada) Len() int {

	if l.firstNode == nil {
		return 0
	}

	nodoAux := new(Node)
	nodoAux = l.firstNode
	count := 1
	for nodoAux.next != nil {

		nodoAux = nodoAux.next
		count++
	}

	return count
}

// Elimina el ultimo elemento de la lista
func (l *ListaLigada) Delete() {
	nodoAux := new(Node)

	if l.firstNode == nil {

		l.firstNode = nil
		l.lastNode = nil
	} else {
		nodoAux = l.lastNode.prev
		nodoAux.next = nil
		l.lastNode = nodoAux
	}

}
