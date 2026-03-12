package customer

type Repository interface {
	GetCustomerInfoById(id uint64) (string, error)
}
