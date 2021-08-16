package models

import (
	"time"

	"gorm.io/gorm"
)

type Tenant struct {
	ID        string `gorm:"PrimaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Name   string `gorm:"size:255"`
	Type   int8   `gorm:"size:255"`
	Status int8   `gorm:"size:255"`
}


func (e *Tenant) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = GenerateID()
	e.Status = 0
	return nil
}

func AddTenant(name string, tenantType, status int8) (tenant Tenant, err error) {
	tenant.Status = tenantType
	tenant.Type = tenantType
	tenant.Name = name
	MetaDB.Create(&tenant)
	// 返回ID
	return
}

func FindTenantById(tenantId string) (tenant Tenant, err error) {
	MetaDB.First(&tenant, tenantId)
	return
}

func FindTenantByName(name string) (tenant Tenant, err error) {
	MetaDB.Where("name = ?", name).First(&tenant)
	return
}