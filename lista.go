package main

import (
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CrearLista() *widget.List {
	listaObjeto := widget.NewList(
		func() int {
			return len(ListaArchivos)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(path.Base(ListaArchivos[lii]))
		},
	)
	return listaObjeto
}
