package model

type ListItem struct {
	Name string
	Desc string
}

func (i ListItem) FilterValue() string { return i.Name }
func (i ListItem) Description() string { return i.Desc }
func (i ListItem) Title() string       { return i.Name }
