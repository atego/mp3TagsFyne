package main

import (
	"go.senan.xyz/taglib"
)

type Tags struct {
	Artista string
	Titulo  string
	Album   string
}

func (t Tags) escribirTags(archivo string) {
	taglib.WriteTags(archivo, map[string][]string{
		taglib.Artist: {t.Artista, t.Artista},
		taglib.Title:  {t.Titulo},
		taglib.Album:  {t.Album},
	}, 0)
}

func (t *Tags) leerTags(archivo string) error {
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
