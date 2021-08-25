package main

import (
	"github.com/andlabs/ui"
)

func VerticalBox(children ...ui.Control) *ui.Box {
	vb := ui.NewVerticalBox()

	for _, child := range children {
		vb.Append(child, false)
	}

	return vb
}
