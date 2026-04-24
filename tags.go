package main

import (
	"io"
	"os"

	"fyne.io/fyne/v2"
	"go.senan.xyz/taglib"
)

type Tags struct {
	Artista  string
	Titulo   string
	Album    string
	Caratula fyne.StaticResource
}

func (t Tags) escribirTags(archivo string) {
	taglib.WriteTags(archivo, map[string][]string{
		taglib.Artist: {t.Artista, t.Artista},
		taglib.Title:  {t.Titulo},
		taglib.Album:  {t.Album},
	}, 0)
	taglib.WriteImage(archivo, t.Caratula.Content())
}

func (t *Tags) leerTags(archivo string) error {
	var imagen []byte
	imagen, _ = taglib.ReadImage(archivo)
	if imagen == nil {
		archivo, _ := os.Open("sinCaratula.png")
		imagen, _ = io.ReadAll(archivo)
	}
	t.Caratula.StaticContent = imagen
	tags, err := taglib.ReadTags(archivo)
	if err != nil {
		return taglib.ErrInvalidFile
	}
	if len(tags[taglib.Artist]) > 0 {
		t.Artista = tags[taglib.Artist][0]
	}
	if len(tags[taglib.Title]) > 0 {
		t.Titulo = tags[taglib.Title][0]
	}
	if len(tags[taglib.Album]) > 0 {
		t.Album = tags[taglib.Album][0]
	}
	return nil
}
