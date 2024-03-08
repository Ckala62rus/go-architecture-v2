package domains

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique:var;size:150"`
	Slug string `gorm:"unique:var;size:255"`
}

func (Category) TableName() string {
	return "category"
}
