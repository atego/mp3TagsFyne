package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func VistaTags() *fyne.Container {
	var tagsContainer *fyne.Container

	a, _ := artistaBind.Get()
	fmt.Println(a)
	artistaInput := widget.NewEntryWithData(artistaBind)
	artistaInput.PlaceHolder = "Artista"

	tituloInput := widget.NewEntryWithData(tituloBind)
	tituloInput.PlaceHolder = "Título"

	albumInput := widget.NewEntryWithData(albumBind)
	albumInput.PlaceHolder = "Album"

	botonAceptar := widget.NewButton(
		"Aceptar",
		func() {
			fmt.Println(artistaInput.Text)
		},
	)

	tagsContainer = container.New(
		layout.NewCustomPaddedVBoxLayout(12),
		artistaInput,
		tituloInput,
		albumInput,
		botonAceptar,
	)
	return tagsContainer
}
