package platform

// Farm represents a farm on the platform
type Farm struct {
	ID        int
	OwnerID   int
	CropType  string
	CropCycle string
	IsPlanted bool
	Yield     float64
	Location  string
}
