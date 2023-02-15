package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name       string `valid:"required"`
	Email      string
	CustomerID string `valid:"matches(^[L|M|H]([0-9]{7})$)~CustomerID ไม่ถูกต้อง"`
}
