package service

import (
	"time"

	"github.com/bline/gotime-server/api/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/golang/protobuf/ptypes"
	"log"
)


const (
	EntryTypeWork  = 0
	EntryTypeBreak = 1
)

type TimeEntry struct {
	gorm.Model
	UserID     UserID          `json:"user_id"`
	Timestamp  time.Time       `gorm:"type:bigint"`
	Type       uint8
	TimeIn     time.Time
	TimeOut    time.Time
	TotalTime  uint64
}
