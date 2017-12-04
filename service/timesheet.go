package service

import (
	"context"
	"github.com/bline/gotime-server/api/proto"
	"github.com/golang/protobuf/ptypes"
	"time"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"fmt"
	"github.com/gorilla/sessions"
)


type TimeSheetService struct {
	db *gorm.DB
	config *viper.Viper
}

func NewTimeSheetService(db *gorm.DB, cfg *viper.Viper) api.TimeSheetServer {
	tsService := TimeSheetService{db: db, config: cfg}
	return &tsService
}
func (tss *TimeSheetService) ClockIn(ctx context.Context, r *api.ClockRequest) (*api.SimpleResponse, error) {
	entryType := r.Type
	timeEntry := TimeEntry{}
	session := ctx.Value("sessionStore").(sessions.Store).Get()
	tss.db.Where("type = ?", EntryTypeWork).
		Order("modified_at DESC").
			Limit(1).
				First(&timeEntry)
	if entryType == EntryTypeBreak {
		if timeEntry.ID == 0 || !timeEntry.TimeOut.IsZero() {
			return nil, fmt.Errorf("must be clocked in to take a break")
		}
		timeEntry.ID = 0
		timeEntry.TimeIn = time.Now()
		timeEntry.TimeOut = time.Time{}
		timeEntry.TotalTime = 0
		timeEntry.UserID =
		tss.db.Create(&timeEntry)
	}
	if entryType == EntryTypeWork && !timeEntry.TimeOut.IsZero() {

	}
	return &api.SimpleResponse{IsSuccess: true, Message: "Clock-In Success"}, nil
}

func (tss *TimeSheetService) ClockOut(context.Context, *api.ClockRequest) (*api.SimpleResponse, error) {
	return &api.SimpleResponse{IsSuccess: true, Message: "Clock-Out Success"}, nil

}
func (tss *TimeSheetService) GetCurrentStatus(context.Context, *api.ClockRequest) (*api.TSStatusResponse, error) {
	ts, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		log.Fatal("Bad time")
	}
	return &api.TSStatusResponse{State: 1, Timestamp: ts}, nil
}

func (tss *TimeSheetService) GetEntries(*api.TimeSheetRequest, api.TimeSheet_GetEntriesServer) error {
	return nil
}

