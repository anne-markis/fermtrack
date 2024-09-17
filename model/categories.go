package model

import (
	"github.com/charmbracelet/bubbles/list"
)

type Category struct {
	Name string
	Desc string
}

func getCategories() []Category {
	return []Category{
		{Name: "wine", Desc: "🍷 Old fruit"},
		{Name: "pickles", Desc: "🥒 Old hard vegetables"},
		{Name: "kraut", Desc: "🥬 Old leafy vegetables"},
	}
}

func ToCategoryStateList() []list.Item {
	items := []list.Item{}
	for _, category := range getCategories() {
		items = append(items, ListItem{Name: category.Name, Desc: category.Desc})
	}

	return items
}
