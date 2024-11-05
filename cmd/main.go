package main

import "0xfarms-backend/internal/platform"

func main() {
	platform := platform.NewPlatform()

	// Register users
	platform.RegisterUser("Alice", true, false)
	platform.RegisterUser("Bob", false, true)

	// Show educational content
	platform.ShowEducationalContent()

	// Recommend a crop for Alice
	platform.RecommendCropPreferences(1)

	// Alice gets seedlings and plants crops
	platform.GetSeedlings(1, "Yams")
	platform.PlantCrop(1)

	// Bob invests in Alice's farm
	platform.InvestInFarm(2, 1, 1500)

	// IoT devices send crop monitoring updates
	platform.IoTDeviceUpdate(1, "Yams", 30)

	// Sell yield and distribute profits to investors
	platform.SellYieldAndDistributeProfits(1)
}
