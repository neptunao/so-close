package geo

import (
	"fmt"
	"math"
)

//Coord is geo coordinates representation with latitude and longitude
type Coord struct {
	Name string
	Lat  float64
	Lon  float64
}

// RelativeCoord is Coord with Center point and Distance in KM from it
type RelativeCoord struct {
	Coord
	Center   Coord
	Distance float64
}

func deg2rad(deg float64) float64 {
	return (deg * math.Pi / 180.0)
}

func rad2deg(rad float64) float64 {
	return (rad / math.Pi * 180.0)
}

// algorithm from http://www.geodatasource.com/developers/c-sharp
func distance(coord1, coord2 Coord) float64 {
	theta := coord1.Lon - coord2.Lon
	dist := math.Sin(deg2rad(coord1.Lat))*math.Sin(deg2rad(coord2.Lat)) +
		math.Cos(deg2rad(coord1.Lat))*math.Cos(deg2rad(coord2.Lat))*math.Cos(deg2rad(theta))
	dist = rad2deg(math.Acos(dist))
	dist = dist * 60 * 1.1515 * 1.609344 //last num converts to KM
	return dist
}

// IsValidCoord checks if geo coordinates in deegrees is valid
func IsValidCoord(c Coord) bool {
	return c.Lat >= -90 && c.Lat <= 90 && c.Lon >= -180 && c.Lon <= 180
}

func (c Coord) String() string {
	return fmt.Sprintf("%s (%f, %f)", c.Name, c.Lat, c.Lon)
}

func (c RelativeCoord) String() string {
	return fmt.Sprintf("%s (%f, %f) distance from %s: %f km",
		c.Name,
		c.Lat,
		c.Lon,
		c.Center,
		c.Distance)
}
