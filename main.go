package main

import (
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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

	// CARÁTULA ---------------------------------------------------------------------------------------------------------
	var caratulaImagen = canvas.NewImageFromResource(nil)
	caratulaImagen.Refresh()
	caratulaImagen.CornerRadius = 8

	caratulaWrap := container.NewGridWrap(fyne.NewSize(150, 150), caratulaImagen)
	caratulaContainer := container.NewHBox(
		layout.NewSpacer(),
		caratulaWrap,
		layout.NewSpacer(),
	)

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
		caratulaImagen.Resource = &tagsArchivo.Caratula
		caratulaImagen.Refresh()
	}

	listaScroll := container.NewVScroll(listaArchivosWidget)
	listaScroll.SetMinSize(fyne.NewSize(380, 200))

	// BOTÓN ABRIR ARCHIVOS ---------------------------------------------------------------------------
	botonAbrirArchivos := widget.NewButton(
		"Selecciona archivos",
		func() {
			archivos = AbrirArchivos()
			listaArchivosWidget.UnselectAll()
			listaArchivosWidget.Refresh()
			artistaInput.SetText("")
			artistaInput.Refresh()
			tituloInput.SetText("")
			tituloInput.Refresh()
			albumInput.SetText("")
			albumInput.Refresh()
			caratulaImagen.Resource = nil
			caratulaImagen.Refresh()
		},
	)

	// CONTENEDOR DE LA INTERFAZ ----------------------------------------------------------------------
	interfazContainer := container.New(
		layout.NewCustomPaddedVBoxLayout(8),
		botonAbrirArchivos,
		listaScroll,
		formInputs,
		caratulaContainer,
	)

	// CONTENEDOR PRINCIPAL ---------------------------------------------------------------------------
	mainContainer := container.NewPadded(interfazContainer)

	window.SetContent(mainContainer)
	window.Resize(fyne.NewSize(400, 600))
	window.ShowAndRun()
}
