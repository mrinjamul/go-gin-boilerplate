package models

// Quote is a struct for quote model
type Quote struct {
	Base
	Text   string `json:"text"`
	Author string `json:"author"`
}
