package platform

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Platform manages users, farms, investments, and educational content
type Platform struct {
	Users       map[int]*User
	Farms       map[int]*Farm
	Investments []Investment
	Blogs       []Blog
	mu          sync.Mutex
}

// NewPlatform creates a new instance of the platform
func NewPlatform() *Platform {
	return &Platform{
		Users: map[int]*User{},
		Farms: map[int]*Farm{
			1: {ID: 1, OwnerID: 1, CropType: "Yams", CropCycle: "6 months", IsPlanted: false, Yield: 500, Location: "Backyard"},
			2: {ID: 2, OwnerID: 2, CropType: "Tomatoes", CropCycle: "3 months", IsPlanted: false, Yield: 300, Location: "Community Garden"},
		},
		Blogs: []Blog{
			{Title: "Starting Your Backyard Farm", Content: "Learn the basics of starting your own farm..."},
			{Title: "Choosing the Best Crops", Content: "A guide to selecting the right crops for your climate..."},
		},
	}
}

// RegisterUser registers a new user on the platform
func (p *Platform) RegisterUser(name string, hasBackyard bool, isInvestor bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	userID := len(p.Users) + 1
	p.Users[userID] = &User{ID: userID, Name: name, HasBackyard: hasBackyard, IsInvestor: isInvestor}
	fmt.Printf("User registered: %s (Has Backyard: %t, Investor: %t)\n", name, hasBackyard, isInvestor)
}

// ShowEducationalContent displays available blog posts and forums
func (p *Platform) ShowEducationalContent() {
	fmt.Println("Educational Blogs & Forums on 0xFarms:")
	for _, blog := range p.Blogs {
		fmt.Printf("Title: %s\nContent: %s\n", blog.Title, blog.Content)
	}
}

// RecommendCropPreferences recommends a crop based on user preferences
func (p *Platform) RecommendCropPreferences(userID int) {
	crops := []string{"Yams", "Tomatoes", "Lettuce"}
	rand.Seed(time.Now().UnixNano())
	recommendedCrop := crops[rand.Intn(len(crops))]
	fmt.Printf("Recommended crop for %s: %s\n", p.Users[userID].Name, recommendedCrop)
}

// GetSeedlings provides crop seedlings to eligible users
func (p *Platform) GetSeedlings(userID int, cropType string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	user, exists := p.Users[userID]
	if !exists || !user.HasBackyard {
		fmt.Println("User not eligible for receiving seedlings.")
		return
	}
	fmt.Printf("Seedlings for %s provided to %s\n", cropType, user.Name)
}

// PlantCrop allows users to plant crops in their backyard
func (p *Platform) PlantCrop(farmID int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	farm, exists := p.Farms[farmID]
	if !exists || farm.IsPlanted {
		fmt.Println("Farm not available or crop already planted.")
		return
	}
	farm.IsPlanted = true
	fmt.Printf("Crops planted in farm %s located in %s\n", farm.CropType, farm.Location)
}

// InvestInFarm allows users to invest in a farm project
func (p *Platform) InvestInFarm(investorID, farmID int, amount float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	user, userExists := p.Users[investorID]
	farm, farmExists := p.Farms[farmID]
	if !userExists || !farmExists || !user.IsInvestor {
		fmt.Println("Invalid investor or farm for investment.")
		return
	}
	p.Investments = append(p.Investments, Investment{InvestorID: investorID, FarmID: farmID, Amount: amount})
	fmt.Printf("Investor %s invested %.2f in %s farm\n", user.Name, amount, farm.CropType)
}

// IoTDeviceUpdate simulates real-time crop monitoring updates
func (p *Platform) IoTDeviceUpdate(farmID int, cropType string, growthStage int) {
	farm, exists := p.Farms[farmID]
	if !exists {
		fmt.Println("Farm not found.")
		return
	}
	fmt.Printf("IoT update for farm %s: Crop '%s' at growth stage %d%%\n", farm.CropType, cropType, growthStage)
}

// SellYieldAndDistributeProfits simulates selling crop yields and distributing profits to investors
func (p *Platform) SellYieldAndDistributeProfits(farmID int) {
	farm, exists := p.Farms[farmID]
	if !exists || !farm.IsPlanted {
		fmt.Println("Farm not available or crop not planted.")
		return
	}
	fmt.Printf("Yield from %s farm sold at proximity market: %.2f units\n", farm.CropType, farm.Yield)
	for _, inv := range p.Investments {
		if inv.FarmID == farmID {
			profitShare := farm.Yield * 0.1 // Investors get 10% of the yield
			fmt.Printf("Investor %d receives a profit share of: %.2f units\n", inv.InvestorID, profitShare)
		}
	}
}
