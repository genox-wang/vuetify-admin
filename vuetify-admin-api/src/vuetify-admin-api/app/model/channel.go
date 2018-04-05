package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"fmt"
)

// Channel is a channel server
type Channel struct {
	gorm.Model
	UUID    string   `json:"uuid" gorm:"unique;not null"`
	Name    string   `json:"name" gorm:"not null"`
	Servers []Server `json:"servers" gorm:"many2many:channel_servers;ASSOCIATION_SAVE_REFERENCE:true;ASSOCIATION_AUTOUPDATE:false;"`
}

// BeforeCreate add uuid before user created
func (ch *Channel) BeforeCreate(scope *gorm.Scope) error {
	UUID := fmt.Sprintf("%s", uuid.NewV1())
	scope.SetColumn("UUID", UUID)
	logrus.Warnf("[ Channel Create ] => uuid = %s", UUID)
	return nil
}

// GetAll get all channels
func (*Channel) GetAll() ([]*Channel, error) {
	channels := make([]*Channel, 0)
	err := DB.Preload("Servers").Find(&channels).Error
	return channels, err
}

// GetByToken get channel by token
func (ch *Channel) GetByToken() error {
	return DB.Preload("Servers").First(ch, "uuid = ?", ch.UUID).Error
}

// Create create new channel
func (ch *Channel) Create() error {
	if err := DB.Create(ch).Error; err != nil {
		return err
	}
	return DB.Model(ch).Association("Servers").Append(ch.Servers).Error
}

//Update save channel
func (ch *Channel) Update() error {
	if err := DB.Model(ch).Update(ch).Error; err != nil {
		return err
	}
	return DB.Model(ch).Association("Servers").Replace(ch.Servers).Error
}

// Delete delete channel
func (ch *Channel) Delete() error {
	return DB.Delete(ch).Error
}
