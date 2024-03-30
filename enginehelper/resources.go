package main

import (
	"dream-play/resources"
	"fmt"
)

func (a *App) Tmpl() {
	path, err := resources.Embed.ReadDir("div-avatar")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
}
