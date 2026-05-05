package main

import (
	"github.com/ncruces/zenity"
)

func AbrirArchivos(filtros []string) []string {
	listaArchivos, _ := zenity.SelectFileMultiple(
		zenity.FileFilters{
			zenity.FileFilter{
				Name:     "Archivos de audio",
				Patterns: filtros,
				CaseFold: false,
			},
		},
	)
	return listaArchivos
}

func AbrirArchivo(filtros []string) string {
	archivo, _ := zenity.SelectFile(
		zenity.FileFilters{
			zenity.FileFilter{
				Name:     "Archivos de audio",
				Patterns: filtros,
				CaseFold: false,
			},
		},
	)
	return archivo
}
