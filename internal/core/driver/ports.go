package driver

type DriverRepository interface {
	FindById(id int) (Driver, error)
}
type IDriverService interface {
	Get(id int) (Driver, error)
}
