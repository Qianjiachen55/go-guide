package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image
var loadIconsOnce sync.Once

func loadIcons()  {
	icons = map[string]image.Image{
		"left": loadIcon(""),
	}
}

func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
}