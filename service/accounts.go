package service

import (
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/bline/gotime-server/api/proto"
	"github.com/futurenda/google-auth-id-token-verifier"

	"context"
	"strings"
	"time"
	"fmt"
	"log"
	"github.com/davecgh/go-spew/spew"
)

type UserID uint

type User struct {
	gorm.Model
	GoogleID    string      `sql:"type:char(40)" json:"google_id"`
	Email       string      `sql:"type:char(255)" json:"email"`
	DisplayName string      `sql:"type: char(150)" json:"display_name"`
	GivenName   string      `sql:"type: char(100)" json:"given_name"`
	FamilyName  string      `sql:"type: char(150)" json:"family_name"`
	Picture     string      `sql:"type: char(200)" json:"picture"`
	LastLogin   time.Time   `sql:"type:bigint" json:"last_login"`
	IsAdmin     bool        `json:"is_admin"`
	TimeEntries []TimeEntry `gorm:"ForeignKey:UserID" json:"-"`
}


type AccountsService struct {
	db *gorm.DB
	config *viper.Viper
}

func NewAccountsService(db *gorm.DB, config *viper.Viper) api.AccountsServer {
	acctService := AccountsService{db: db, config: config}
	return &acctService
}

func (*AccountsService) GetUser(ctx context.Context, r *api.GetUserRequest) (*api.User, error) {
	return &api.User{}, nil
}
func (*AccountsService) GetUsers(r *api.GetUsersRequest, server api.Accounts_GetUsersServer) error {
	return nil
}
func (*AccountsService) DisableUser(ctx context.Context, r *api.DisableUserRequest) (*api.SimpleResponse, error) {
	return &api.SimpleResponse{IsSuccess: false, Message: "disabling a user is unsupported"}, nil
}
func (*AccountsService) DeleteUser(ctx context.Context, r *api.DeleteUserRequest) (*api.SimpleResponse, error) {
	return &api.SimpleResponse{IsSuccess: true, Message: "User Deleted"}, nil
}
func (*AccountsService) LockUser(ctx context.Context, r *api.LockUserRequest) (*api.SimpleResponse, error) {
	return &api.SimpleResponse{IsSuccess: true, Message: "User Locked"}, nil
}

func UserFromIDToken(db *gorm.DB, config *viper.Viper, idtoken string) (*User, error) {
	var curUser User
	set, err := verifyToken(config, idtoken)
	if err != nil {
		return nil, err
	}
	db.Where("google_id = ? AND email = ?", set.Sub, set.Email).First(&curUser)
	if curUser.ID != 0 {
		return &curUser, nil
	}

	curUser.DisplayName = set.Name
	curUser.Email = set.Email
	curUser.FamilyName = set.FamilyName
	curUser.GivenName = set.GivenName
	curUser.GoogleID = set.Sub
	curUser.IsAdmin = false
	curUser.LastLogin = time.Now()
	curUser.Picture = set.Picture
	db.Save(&curUser) // save sets ID
	return &curUser, nil
}

func verifyToken(config *viper.Viper, token string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	v := googleAuthIDTokenVerifier.Verifier{}
	aud := config.GetString("google_client_id")
	suffix := config.GetString("oauth2_email_suffix")
	err := v.VerifyIDToken(token, []string{
		aud,
	})
	if err != nil {
		return nil, err
	}
	claimSet, err := googleAuthIDTokenVerifier.Decode(token)
	// claimSet.Iss,claimSet.Email ... (See claimset.go)
	// XXX get hd from PrivateClaim?
	log.Printf("claimSet: %v", spew.Sdump(claimSet))
	if !strings.HasSuffix(claimSet.Email, suffix) {
		return nil, fmt.Errorf("must use a %v email address", suffix)
	}
	return claimSet, nil
}