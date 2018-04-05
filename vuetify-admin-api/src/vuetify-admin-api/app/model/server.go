package model

import (
	"github.com/jinzhu/gorm"
)

// Server is a server model
type Server struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null"`
	Token string `json:"token" gorm:"unique;not null"`
}

// GetAll get all servers
func (*Server) GetAll() ([]*Server, error) {
	servers := make([]*Server, 0)
	err := DB.Find(&servers).Error
	return servers, err
}

// Create create new server
func (s *Server) Create() error {
	return DB.Create(s).Error
}

// Update save server
func (s *Server) Update() error {
	return DB.Model(s).Update(s).Error
}

// Delete delete server
func (s *Server) Delete() error {
	return DB.Delete(s).Error
}
