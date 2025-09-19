package driver

type DriverService struct {
	driverRepository DriverRepository //dependency injection and not exportable
}

func NewDriverService(driverRepository DriverRepository) *DriverService {
	return &DriverService{driverRepository: driverRepository}
}

func (s *DriverService) Get(id int) (Driver, error) {
	return s.driverRepository.FindById(id)
}
