package example

import (
	"gorm.io/gorm"
	"time"
)

//go:generate gormgen -structs User,Admin -input . -imports gorm.io/gorm -transformErr true
type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique"`
	Age   int
	Email string
}

type Admin struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique"`
	Age   int
	Email string
}

type CodeccDefectDetailCcnResult struct {
	ID              int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	TaskID          int       `gorm:"column:task_id" json:"task_id"`
	FuncationAmount int       `gorm:"column:funcation_amount" json:"funcation_amount"`
	OriCcn          int       `gorm:"column:ori_ccn" json:"ori_ccn"`
	ComputeCcn      int       `gorm:"column:compute_ccn" json:"compute_ccn"`
	DataDate        int       `gorm:"column:data_date;default:0" json:"data_date"`
	UpdateTime      time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
}

func (m *CodeccDefectDetailCcnResult) TableName() string {
	return "codecc_defect_detail_ccn_result"
}
