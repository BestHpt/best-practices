package entity

type BlogType struct {
	TypeId   string `json:"typeId" gorm:"column:type_id"`
	TypeName string `json:"dataType" gorm:"column:type_name"`
}

func (s *BlogType) TableName() string {
	return "blog_type"
}