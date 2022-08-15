/**
    @author: huchao
    @since: 2022/8/2
    @desc: //TODO sync.Once
**/
package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image

var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

func loadIcon(s string) image.Image {
	return icons[s]
}

// Icon 被多个goroutine调用时不是并发安全的
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

// Icon2 是并发安全的
func Icon2(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
