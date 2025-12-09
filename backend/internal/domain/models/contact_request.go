package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// ContactRequest mirrors the Rolston HVAC contact form payload
type ContactRequest struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Reason    string    `json:"reason" gorm:"type:text;not null;default:''"`
	Phone     *string   `json:"phone" gorm:"type:varchar(20)"`
	Message   string    `json:"message" gorm:"type:text;not null"`
	IPAddress *string   `json:"ip_address" gorm:"column:ip_address;type:varchar(45)"`
	UserAgent *string   `json:"user_agent" gorm:"column:user_agent;type:text"`
	Metadata  JSONB     `json:"metadata" gorm:"type:jsonb;default:'{}'"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName explicitly sets the table name to match the migration
func (ContactRequest) TableName() string {
	return "contact_requests"
}

// JSONB supports PostgreSQL JSONB fields
type JSONB map[string]interface{}

// Value satisfies the driver.Valuer interface
func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(j)
}

// Scan satisfies the sql.Scanner interface
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = JSONB{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}

	return json.Unmarshal(bytes, j)
}
