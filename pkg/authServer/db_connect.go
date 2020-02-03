package authServer

import (
	"RoomStatus/config"

	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (CAB *CreditsAuthBackend) InitDB(config *config.CfTDatabase) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	if err != nil {
		return nil, err
	}

	return db, nil

}

func (CAB *CreditsAuthBackend) CloseDB() (bool, error) {
	CAB.mu.Lock()
	defer CAB.mu.Unlock()

	err := CAB.DBconn.Close()
	if err != nil {
		return false, err
	}
	return true, nil
}

type UserCredMod struct {
	gorm.Model
	Username *string `gorm:"type:varchar(100)"`
	Password *string `gorm:"type:varchar(100)"`
	KeyPem   *[]byte
}

func (UserCredMod) TableName() string {
	return "user_cred"
}

type CredSessionMod struct {
	gorm.Model
	UserId     *string `gorm:"column:user_id,type:varchar(100)"`
	DeviceName *string
}

func (CredSessionMod) TableName() string {
	return "cred_session"
}

func init_check(dbc *gorm.DB) (bool, error) {
	if dbc == nil {
		return false, errors.New("NULL_SESSION")
	}

	if dbc.HasTable(&UserCredMod{}) == false {
		dbc.Set(
			"gorm:table_options",
			"ENGINE=InnoDB",
		).CreateTable(&UserCredMod{})
	}

	if dbc.HasTable(&CredSessionMod{}) == false {
		dbc.Set(
			"gorm:table_options",
			"ENGINE=InnoDB",
		).CreateTable(&CredSessionMod{})
		dbc.Model(&CredSessionMod{}).AddForeignKey("user_id", "user_cred(id)", "RESTRICT", "RESTRICT")
	}
	return true, nil
}
