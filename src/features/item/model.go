package item

type ItemModel struct {
	Id        int    `json:"id" gorm:"primary_key, not null"`
	Title     string `json:"title" gorm:"not null"`
	CreatedAt string `json:"created_at" sql:"DEFAULT:current_timestamp"`
}
