package models

type Quote struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Text      string `json:"text"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}

func init() {

}
