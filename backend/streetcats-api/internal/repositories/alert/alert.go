package alert

type Repository interface {
	CreateAlertSchemas() error
}
