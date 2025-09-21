package driver

import "go.uber.org/zap"

type DriverService struct {
	driverRepository DriverRepository
	logger           *zap.Logger
}

func NewDriverService(driverRepository DriverRepository, logger *zap.Logger) *DriverService {
	return &DriverService{
		driverRepository: driverRepository,
		logger:           logger,
	}
}

func (s *DriverService) Get(id int) (Driver, error) {
	return s.driverRepository.FindById(id)
}
