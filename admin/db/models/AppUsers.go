package models

import (
	"time"

	"gorm.io/gorm"
)

// Enums for UserType and UserProfile
type UserType string
type UserProfile string

const (
	UserTypeCallCenter UserType = "call_center"
	UserTypeBoxOffice  UserType = "box_office"

	UserProfileOperator UserProfile = "operator"
	UserProfileAdmin    UserProfile = "admin"
)

// PointsOfSaleSet model
type PointsOfSale struct {
	PointOfSaleId uint           `gorm:"primaryKey" json:"pointsOfSaleId"`
	Name          string         `json:"name"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

// ReportSet model
type Report struct {
	ReportId  uint           `gorm:"primaryKey" json:"reportId"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Users model
type User struct {
	AppUserID      uint           `gorm:"primaryKey" json:"appUserId"`
	Status         bool           `json:"status"`
	Type           UserType       `gorm:"type:text" json:"type"` // ya que UserType es un string alias
	Profile        UserProfile    `gorm:"type:text" json:"profile"`
	Name           string         `json:"name"`
	LastName       string         `json:"lastName"`
	UserName       string         `json:"userName"`
	Email          string         `json:"email"`
	APIUser        bool           `json:"apiUser"`
	Password       string         `json:"password"`
	PointOfSaleSet []PointsOfSale `gorm:"many2many:user_points_of_sale_sets;" json:"pointOfSaleSet"`
	ReportSet      []Report       `gorm:"many2many:user_report_sets;" json:"reportSet"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}
