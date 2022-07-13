package entity

// UserEntity
type UserEntity struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

// TableName
func (p *UserEntity) TableName() string {
	return "user"
}
