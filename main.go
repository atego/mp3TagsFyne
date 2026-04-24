package main

import (
	"io"
	"os"
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

var archivoSeleccionado string

func main() {
	app := app.New()
	window := app.NewWindow("Tags MP3")

	// INPUTS TAGS (artista, título, album) ----------------------------------------------------------------------------------
	tituloInput := widget.NewEntry()
	artistaInput := widget.NewEntry()
	artistaCheck := widget.NewCheck(
		"",
		func(b bool) {
		},
	)
	artistaFila := container.NewBorder(nil, nil, nil, artistaCheck, artistaInput)

	albumInput := widget.NewEntry()
	albumCheck := widget.NewCheck(
		"",
		func(b bool) {
		},
	)
	albumFila := container.NewBorder(nil, nil, nil, albumCheck, albumInput)

	formInputs := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Título", Widget: tituloInput},
			{Text: "Artista", Widget: artistaFila},
			{Text: "Album", Widget: albumFila},
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
		archivoSeleccionado = archivos[id]
		tagsArchivo = &Tags{}
		if artistaCheck.Checked {
			tagsArchivo.Artista = artistaInput.Text
		}
		if albumCheck.Checked {
			tagsArchivo.Album = albumInput.Text
		}
		tagsArchivo.leerTags(archivoSeleccionado)
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
			archivos = AbrirArchivos([]string{"*.mp3", "*.mp4", "*.m4a"})
			listaArchivosWidget.UnselectAll()
			listaArchivosWidget.Refresh()
			artistaInput.SetText("")
			tituloInput.SetText("")
			albumInput.SetText("")
			caratulaImagen.Resource = nil
			caratulaImagen.Refresh()
		},
	)

	// BOTÓN CAMBIAR CARATULA ------------------------------------------------------------------------
	botonCambiarCaratula := widget.NewButton(
		"Elegir carátula",
		func() {
			caratulaPath := AbrirArchivo([]string{"*.png", "*.jpg", "*.jpeg"})
			if caratulaPath != "" {
				archivo, _ := os.Open(caratulaPath)
				imagenBytes, _ := io.ReadAll(archivo)
				tagsArchivo.Caratula = *fyne.NewStaticResource("", imagenBytes)
				caratulaImagen.Resource = &tagsArchivo.Caratula
				caratulaImagen.Refresh()
			}
		},
	)
	botonCambiarCaratula.Importance = widget.LowImportance

	// BOTON GUARDAR DATOS ----------------------------------------------------------------------------
	botonGuardarDatos := widget.NewButton(
		"Guardar tags",
		func() {
			tagsArchivo.Artista = artistaInput.Text
			tagsArchivo.Titulo = tituloInput.Text
			tagsArchivo.Album = albumInput.Text
			tagsArchivo.escribirTags(archivoSeleccionado)
			notificacionExito := fyne.NewNotification(
				"Aviso!!",
				"Tags guardadas con éxito 🙂",
			)
			app.SendNotification(notificacionExito)
		},
	)
	botonGuardarDatos.Importance = widget.WarningImportance

	// CONTENEDOR DE LA INTERFAZ ----------------------------------------------------------------------
	interfazContainer := container.New(
		layout.NewCustomPaddedVBoxLayout(8),
		botonAbrirArchivos,
		listaScroll,
		layout.NewSpacer(),
		formInputs,
		caratulaContainer,
		botonCambiarCaratula,
		botonGuardarDatos,
	)

	// CONTENEDOR PRINCIPAL ---------------------------------------------------------------------------
	mainContainer := container.NewPadded(interfazContainer)

	window.SetContent(mainContainer)
	window.Resize(fyne.NewSize(400, 600))
	window.ShowAndRun()
}
