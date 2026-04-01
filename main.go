package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ncruces/zenity"
)

var (
	listaContainer  *widget.List
	listaScroll     *container.Scroll
	ListaArchivos   []string
	interfaz        *fyne.Container
	vistaDatos      *fyne.Container
	btnAbrirCarpeta *widget.Button
	//variables bind ---------------------
	artistaBind  = binding.NewString()
	tituloBind   = binding.NewString()
	albumBind    = binding.NewString()
	caratulaBind = binding.NewString()
)

func main() {
	//LeerTags("styles.mp4")
	a := app.NewWithID("mp3TagsFyne")
	w := a.NewWindow("mp3TagsFyne")

	ListaArchivos = []string{}
	listaContainer = CrearLista()

	listaScroll = container.NewVScroll(listaContainer)
	listaScroll.SetMinSize(fyne.NewSize(380, 200))

	vistaDatos = VistaTags()

	btnAbrirCarpeta = widget.NewButton(
		"Abrir carpeta",
		func() {
			ListaArchivos, _ = zenity.SelectFileMultiple(
				zenity.FileFilters{
					zenity.FileFilter{
						Name:     "Archivos de audio",
						Patterns: []string{"*.mp3", "*.mp4", "*.m4a"},
						CaseFold: false,
					},
				},
			)
			if len(ListaArchivos) > 0 {
				listaContainer.Refresh()
				LeerTags(ListaArchivos[0])
				listaContainer.Select(0)
			}
		},
	)

	listaContainer.OnSelected = func(id widget.ListItemID) {
		LeerTags(ListaArchivos[id])
	}

	interfaz = container.New(
		layout.NewCustomPaddedVBoxLayout(12),
		btnAbrirCarpeta,
		listaScroll,
		vistaDatos,
	)

	mainContainer := container.NewPadded(interfaz)

	w.SetContent(mainContainer)
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}
