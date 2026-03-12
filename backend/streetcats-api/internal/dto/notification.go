package dto

type NotificationDTO struct {
	ID uint64 `json:"id"`
	//LicensePlate *string `json:"licensePlate,omitempty"`
	//CustomerName *string `json:"customerName,omitempty"`
	Subject     *string `json:"subject,omitempty"`
	Data        any     `json:"data" swaggertype:"object"`
	Ts          int64   `json:"ts"`
	Category    string  `json:"category"`
	Severity    string  `json:"severity"`
	AlertCode   string  `json:"alertCode"`
	MarkedState bool    `json:"markedState"`
} //	@name	NotificationResult

type NotificationEventDTO struct {
	ID         uint64 `json:"id"`
	Source     string `json:"source"`
	SourceType uint16 `json:"source_type"`
	Category   string `json:"category"`
	Severity   string `json:"severity"`
	Data       any    `json:"data" swaggertype:"object"`
	Ts         int64  `json:"ts"`
	AlertCode  string `json:"alert_code"`
} //@name NotificationEventResponse

type NotificationMarkerDTO struct {
	IDs    []uint64 `json:"ids"`
	Marked bool     `json:"marked"`
} //@name NotificationMarkedResponse

type UserInfoDTO struct {
	ID         uint64  `json:"id"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Email      *string `json:"email"`
	Username   *string `json:"username,omitempty"`
	IdUserRole *uint16 `json:"iduserrole,omitempty"`
	UsVhId     *uint64 `json:"usvhid,omitempty"`
	Language   *string `json:"language,omitempty"`
}

type EventUserDTO struct {
	EventId uint64 `json:"event_id"`
	UserId  uint64 `json:"user_id"`
	Marked  bool   `json:"marked"`
}

type Rule struct {
	Customer               uint64 `json:"customer"`
	SeverityCode           string `json:"severity_code"`
	ID                     uint64 `json:"id"`
	AlertCapability        bool   `json:"alert_capability"`
	AppCapability          bool   `json:"app_capability"`
	EmailCapability        bool   `json:"email_capability"`
	CustomerId             uint64 `json:"customer_id"`
	EmailToRecipients      bool   `json:"email_to_recipients"`
	FilterUserEnabled      bool   `json:"filter_user_enabled"`
	FilterRecipientEnabled bool   `json:"filter_recipient_enabled"`
	CreatedAt              int64  `json:"created_at"`
	UpdatedAt              int64  `json:"updated_at"`
	IsActive               bool   `json:"is_active"`
	Category               string `json:"category"`
}
