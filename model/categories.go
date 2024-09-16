package model

type Category struct {
	Name string
}

func GetCategories() []Category {
	return []Category{{Name: "wine"}, {Name: "pickles"}, {Name: "kraut"}}
}
