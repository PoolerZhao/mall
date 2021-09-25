package models

type CategoryOption struct {
	Value    uint             `json:"value"`
	Label    string           `json:"label"`
	Children []CategoryOption `json:"children"`
}
