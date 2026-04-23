package main

import (
	"github.com/ncruces/zenity"
)

func AbrirArchivos() []string {
	listaArchivos, _ := zenity.SelectFileMultiple(
		zenity.FileFilters{
			zenity.FileFilter{
				Name:     "Archivos de audio",
				Patterns: []string{"*.mp3", "*.mp4", "*.m4a"},
				CaseFold: false,
			},
		},
	)
	return listaArchivos
}
