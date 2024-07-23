package main

import (
	"fmt"
	"math"
)

// Listing 22.1 - latlong.go
type Latitude float64
type Longitude float64

// Listing 22.2 - location.go
type Location struct {
	Lat  Latitude
	Long Longitude
}

// Listing 22.3 - decimaldegrees.go
func (lat Latitude) Decimal() float64 {
	return float64(lat)
}

func (long Longitude) Decimal() float64 {
	return float64(long)
}

func (loc Location) String() string {
	return fmt.Sprintf("%.6f, %.6f", loc.Lat.Decimal(), loc.Long.Decimal())
}

// Listing 22.4 - distance.go
func (loc1 Location) distance(loc2 Location) float64 {
	const R = 6371 // Radius of the Earth in kilometers
	lat1 := loc1.Lat.Decimal() * math.Pi / 180
	long1 := loc1.Long.Decimal() * math.Pi / 180
	lat2 := loc2.Lat.Decimal() * math.Pi / 180
	long2 := loc2.Long.Decimal() * math.Pi / 180

	dlat := lat2 - lat1
	dlong := long2 - long1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlong/2)*math.Sin(dlong/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func main() {
	// Table 22.1 locations in decimal degrees
	locations := map[string]Location{
		"San Francisco": {Lat: 37.7749, Long: -122.4194},
		"New York":      {Lat: 40.7128, Long: -74.0060},
		"Los Angeles":   {Lat: 34.0522, Long: -118.2437},
		"London":        {Lat: 51.5074, Long: -0.1278},
		"Paris":         {Lat: 48.8566, Long: 2.3522},
		"Tokyo":         {Lat: 35.6895, Long: 139.6917},
		"Sydney":        {Lat: -33.8688, Long: 151.2093},
		"Moscow":        {Lat: 55.7558, Long: 37.6173},
		"Beijing":       {Lat: 39.9042, Long: 116.4074},
		"New Delhi":     {Lat: 28.6139, Long: 77.2090},
	}

	fmt.Println("Distances between each pair of locations:")
	var closestPair [2]string
	var farthestPair [2]string
	var minDistance = math.MaxFloat64
	var maxDistance float64

	for name1, loc1 := range locations {
		for name2, loc2 := range locations {
			if name1 != name2 {
				distance := loc1.distance(loc2)
				fmt.Printf("Distance between %s and %s: %.2f km\n", name1, name2, distance)
				if distance < minDistance {
					minDistance = distance
					closestPair = [2]string{name1, name2}
				}
				if distance > maxDistance {
					maxDistance = distance
					farthestPair = [2]string{name1, name2}
				}
			}
		}
	}

	fmt.Printf("The closest pair of locations is %s and %s with a distance of %.2f km\n", closestPair[0], closestPair[1], minDistance)
	fmt.Printf("The farthest pair of locations is %s and %s with a distance of %.2f km\n", farthestPair[0], farthestPair[1], maxDistance)

	// Calculate the specific distances required
	london := Location{Lat: 51.5074, Long: -0.1278}
	paris := Location{Lat: 48.8566, Long: 2.3522}
	distanceLondonParis := london.distance(paris)
	fmt.Printf("Distance from London to Paris: %.2f km\n", distanceLondonParis)

	// Example: Find the distance from your city to the capital of your country
	// Replace these coordinates with those of your city and capital
	myCity := Location{Lat: 0.0, Long: 0.0}  // Replace with your city's coordinates
	capital := Location{Lat: 0.0, Long: 0.0} // Replace with your capital's coordinates
	distanceMyCityToCapital := myCity.distance(capital)
	fmt.Printf("Distance from my city to the capital: %.2f km\n", distanceMyCityToCapital)

	// Mars locations
	mountSharp := Location{Lat: -5.0800, Long: 137.8500}  // Mount Sharp
	olympusMons := Location{Lat: 18.6500, Long: 226.2000} // Olympus Mons
	distanceMountSharpOlympusMons := mountSharp.distance(olympusMons)
	fmt.Printf("Distance between Mount Sharp and Olympus Mons on Mars: %.2f km\n", distanceMountSharpOlympusMons)
}
