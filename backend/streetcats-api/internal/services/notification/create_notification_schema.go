package notification

func (s *Service) CreateNotificationSchemas() error {
	return s.notificationRepo.CreateNotificationSchemas()
}
