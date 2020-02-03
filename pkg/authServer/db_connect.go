package authServer

import (
	"RoomStatus/config"
	"log"

	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (CAB *CreditsAuthBackend) InitDB(config *config.CfTDatabase) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	if err != nil {
		return nil, err
	}
	CAB.DBconn = db

	init_check(db)
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
	log.Println("start checking")
	if dbc == nil {
		return false, errors.New("NULL_SESSION")
	}

	log.Println("table-", UserCredMod{}.TableName())
	if dbc.HasTable(&UserCredMod{}) == false {
		log.Println("table->create", UserCredMod{}.TableName())

		dbc.Set(
			"gorm:table_options",
			"ENGINE=InnoDB",
		).CreateTable(&UserCredMod{})
	}

	log.Println("table-", CredSessionMod{}.TableName())

	if dbc.HasTable(&CredSessionMod{}) == false {
		log.Println("table->create", CredSessionMod{}.TableName())

		dbc.Set(
			"gorm:table_options",
			"ENGINE=InnoDB",
		).CreateTable(&CredSessionMod{})
		dbc.Model(&CredSessionMod{}).AddForeignKey("user_id", "user_cred(id)", "RESTRICT", "RESTRICT")
	}
	return true, nil
}
