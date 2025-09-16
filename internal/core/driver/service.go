package driver

type DriverService struct {
	driverRepository DriverRepository //dependency injection and not exportable
}

func NewDriverService(driverRepository DriverRepository) *DriverService {
	return &DriverService{driverRepository: driverRepository}
}

func (s *DriverService) Get(id int) (Driver, error) {
	driver, err := s.driverRepository.FindById(id)
	if err != nil {
		return Driver{}, err //I want to handle error here with a custom error
	}
	return driver, nil
}
