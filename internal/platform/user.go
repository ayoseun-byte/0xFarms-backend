package platform

// User represents an individual who can register and invest in farms
type User struct {
	ID          int
	Name        string
	HasBackyard bool
	IsInvestor  bool
}
