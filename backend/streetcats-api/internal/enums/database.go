package enums

type platform int

// User role
const (
	RootUser uint16 = iota + 1
	SuperUser
	CRMAdmin
	CustomerAdmin
	CustomerUser
	CustomerManager
)

// Platforms

const (
	Crm platform = iota + 1
	Fdm
)

// Source Type

const (
	Imei uint16 = iota + 1
	VehicleId
	CustomerId
	DistributorId
	GeneralAppId
)

// Category

const (
	Tpms      = "TPMS"
	Refueling = "REFUELING"
	Tacho     = "TACHO"
	Poi       = "POI"
)

// Severity

const (
	Error        = "ERROR"
	Warning      = "WARNING"
	Info         = "INFO"
	PriorityInfo = "PRIORITY_INFO"
)

// AlertCode

const (
	GenericError             = "GENERIC_ERROR"
	WarningPressure          = "WARNING_PRESSURE"
	AlertPressure            = "ALERT_PRESSURE"
	WarningTemperature       = "WARNING_TEMPERATURE"
	AlertTemperature         = "ALERT_TEMPERATURE"
	MissingSensor            = "MISSING_SENSOR"
	LowBatterySensor         = "LOW_BATTERY_SENSOR"
	TachoDrivingTimeExceeded = "TACHO_DRIVING_TIME_EXCEEDED"
	TachoWorkTimeExceeded    = "TACHO_WORK_TIME_EXCEEDED"
	TachoInsufficientBreak   = "TACHO_INSUFFICIENT_BREAK"
	WarningTireControl       = "WARNING_TIRE_CONTROL"
	AlertTireControl         = "ALERT_TIRE_CONTROL"
)
