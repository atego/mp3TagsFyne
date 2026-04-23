package main

import (
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var archivos []string

var tagsArchivo = &Tags{}

func main() {
	app := app.New()
	window := app.NewWindow("Tags MP3")

	// INPUTS TAGS (artista, título, album) ----------------------------------------------------------------------------------
	tituloInput := widget.NewEntry()
	artistaInput := widget.NewEntry()
	albumInput := widget.NewEntry()

	formInputs := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Título", Widget: tituloInput},
			{Text: "Artista", Widget: artistaInput},
			{Text: "Album", Widget: albumInput},
		},
	}

	// LISTA DE ARCHIVOS ---------------------------------------------------------------------------
	listaArchivosWidget := widget.NewList(
		func() int {
			return len(archivos)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(path.Base(archivos[lii]))
		},
	)
	listaArchivosWidget.OnSelected = func(id widget.ListItemID) {
		tagsArchivo.leerTags(archivos[id])
		tituloInput.SetText(tagsArchivo.Titulo)
		artistaInput.SetText(tagsArchivo.Artista)
		albumInput.SetText(tagsArchivo.Album)
	}

	listaScroll := container.NewVScroll(listaArchivosWidget)
	listaScroll.SetMinSize(fyne.NewSize(380, 200))

	// BOTÓN ABRIR ARCHIVOS ---------------------------------------------------------------------------
	botonAbrirArchivos := widget.NewButton(
		"Selecciona archivos",
		func() {
			archivos = AbrirArchivos()
			listaArchivosWidget.Refresh()
		},
	)

	// CONTENEDOR DE LA INTERFAZ ----------------------------------------------------------------------
	interfazContainer := container.New(
		layout.NewCustomPaddedVBoxLayout(8),
		botonAbrirArchivos,
		listaScroll,
		formInputs,
	)

	// CONTENEDOR PRINCIPAL ---------------------------------------------------------------------------
	mainContainer := container.NewPadded(interfazContainer)

	window.SetContent(mainContainer)
	window.Resize(fyne.NewSize(400, 600))
	window.ShowAndRun()
}
