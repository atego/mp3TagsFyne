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

	tituloInput := widget.NewEntryWithData(tituloBind)
	tituloInput.PlaceHolder = "Título"

	artistaInput := widget.NewEntryWithData(artistaBind)
	artistaInput.PlaceHolder = "Artista"
	artistaGeneralCheck := widget.NewCheck(
		"Toda la lista",
		func(b bool) {
			fmt.Println("Para todos 🙂‍↔️")
		},
	)
	filaArtista := container.New(
		layout.NewGridLayoutWithColumns(2),
		artistaInput,
		artistaGeneralCheck,
	)

	albumInput := widget.NewEntryWithData(albumBind)
	albumInput.PlaceHolder = "Album"
	albumGeneralCheck := widget.NewCheck(
		"Toda la lista",
		func(b bool) {
			fmt.Println("Para todos 🙂‍↔️")
		},
	)
	filaAlbum := container.New(
		layout.NewGridLayoutWithColumns(2),
		albumInput,
		albumGeneralCheck,
	)

	botonAceptar := widget.NewButton(
		"Aceptar",
		func() {
			fmt.Println(artistaInput.Text)
		},
	)

	tagsContainer = container.New(
		layout.NewCustomPaddedVBoxLayout(12),
		tituloInput,
		filaArtista,
		filaAlbum,
		botonAceptar,
	)
	return tagsContainer
}
