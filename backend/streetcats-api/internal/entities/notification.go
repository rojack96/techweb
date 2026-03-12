package entities

import "gorm.io/datatypes"

const (
	Geofence  = "geofence"
	Sightings = "sightings"
	Topology  = "topology"
	Users     = "users"
)

type Event struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	Source     string
	SourceType int32
	Category   string
	Severity   string
	AlertCode  string
	Data       datatypes.JSON `gorm:"type:jsonb"`
	Ts         int64
	CreateDate int64
}

type EventUser struct {
	IdEvent     uint64 `gorm:"primaryKey"`
	IdUser      uint64 `gorm:"primaryKey"`
	MarkedState bool

	Event   Event   `gorm:"foreignKey:IdEvent;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UrUsers UrUsers `gorm:"foreignKey:IdUser;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UrUsers struct {
	ID            uint64 `gorm:"primaryKey"`
	Name          string `gorm:"column:name"`
	Surname       string `gorm:"column:surname"`
	Username      string `gorm:"column:username;unique;"`
	Email         string `gorm:"column:email;unique;"`
	IdUserType    uint64 `gorm:"column:idusertype"`
	IdUserRole    uint16 `gorm:"column:iduserrole"`
	IdLanguage    uint64 `gorm:"column:idlanguage"`
	IdTimezone    uint64 `gorm:"column:idtimezone"`
	IdClient      uint64 `gorm:"column:idclient"`
	IdDistributor uint64 `gorm:"column:iddistributor"`
	IdTemplate    uint64 `gorm:"column:idtemplate"`
	IsActive      bool   `gorm:"column:isactive"`
	IdSystemType  uint64 `gorm:"column:idsystemtype"`
}

func (Event) TableName() string {
	return NotificationSchema + ".event"
}

func (EventUser) TableName() string {
	return NotificationSchema + ".event_user"
}

func (Rule) TableName() string { return NotificationSchema + ".rule" }

func (UrUsers) TableName() string {
	return CrmdbSchema + ".ur_users"
}

func (EmailReceipt) TableName() string {
	return NotificationSchema + ".email_receipt"
}
