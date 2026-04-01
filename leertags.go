package main

import (
	"fmt"

	"go.senan.xyz/taglib"
)

func LeerTags(archivo string) {
	fmt.Println(archivo)
	/* 		archivoPortada, err := os.Open("./portada.jpg")
	if err != nil {
		log.Fatal(err)
	}
	portadaBytes, _ := io.ReadAll(archivoPortada) */
	//taglib.WriteImage(archivo, portadaBytes)
	/* 	taglib.WriteTags(archivo, map[string][]string{
		taglib.Artist: {"Harry Styles"},
		taglib.Album:  {"Grandes éxitos"},
		taglib.Title:  {"As it was"},
	}, 0) */
	tags, err := taglib.ReadTags(archivo)
	if err != nil {
		fmt.Println("error")
		return
	}
	if len(tags[taglib.Artist]) > 0 {
		artistaBind.Set(tags[taglib.Artist][0])
	}
	if len(tags[taglib.Title]) > 0 {
		tituloBind.Set(tags[taglib.Title][0])

	}
	if len(tags[taglib.Album]) > 0 {
		tituloBind.Set(tags[taglib.Album][0])
	}
}
