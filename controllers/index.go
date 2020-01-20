package controllers

import "github.com/amanjuman/gocontact/views"

func NewIndex() *Index {
	return &Index{
		Homepage: views.NewView("bootstrap", "index/homepage"),
	}
}

type Index struct {
	Homepage *views.View
}
