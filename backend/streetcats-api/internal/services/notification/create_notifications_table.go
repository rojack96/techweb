package notification

// CreateEventTables - Setup all tables useful for service
func (s *Service) CreateEventTables() error {
	return s.notificationRepo.CreateNotificationsTables()
}
