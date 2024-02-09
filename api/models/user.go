package models

import "time"

// User represents a user.
// swagger:response userResponse
//
//	type User struct {
//		ID         uint32     `gorm:"primary_key;auto_increment" json:"id" swaggerignore:"true"`
//		UserName   string     `gorm:"default:user_name size:20;not null;unique_index" json:"userName,omitempty" example:"piter"`
//		Email      string     `gorm:"size:35;not null;unique_index" json:"email,omitempty" example:"piter.teste@example.com"`
//		Password   string     `gorm:"size:60;not null" json:"password,omitempty" example:"123"`
//		CreatedAt  *time.Time `gorm:"default:current_timestamp()" json:"created_at" swaggerignore:"true"`
//		UpdatedAt  *time.Time `gorm:"default:current_timestamp()" json:"updated_at" swaggerignore:"true"`
//		AuditLogin *time.Time `gorm:"<-:false" json:"auditLogin" swaggerignore:"true"`
//
// User represents a user.
// swagger:response userResponse
type User struct {
	ID         uint32     `gorm:"primary_key;auto_increment" json:"id" swaggerignore:"true"`
	UserName   string     `gorm:"size:20;not null;unique_index" json:"userName,omitempty" example:"piter"`
	Email      string     `gorm:"size:35;not null;unique_index" json:"email,omitempty" example:"piter.teste@example.com"`
	Password   string     `gorm:"size:60;not null" json:"password,omitempty" example:"123"`
	CreatedAt  *time.Time `gorm:"default:current_timestamp()" json:"created_at" swaggerignore:"true"`
	UpdatedAt  *time.Time `gorm:"default:current_timestamp()" json:"updated_at" swaggerignore:"true"`
	AuditLogin *time.Time `gorm:"<-:false" json:"auditLogin" swaggerignore:"true"`
	Rule       string     `gorm:"not null" json:"rule,omitempty" example:"ADMIN"`
}

// UserWithoutPassword represents a user.
// swagger:response userResponse
type UserWithoutPassword struct {
	ID        uint32     `json:"id"`
	UserName  string     `json:"userName,omitempty"`
	Email     string     `json:"email,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
