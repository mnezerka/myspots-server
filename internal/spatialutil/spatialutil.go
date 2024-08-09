package spatialutil

import (
	"fmt"
	"mnezerka/myspots-server/entities"
)

/*
Valid longitude values are between -180 and 180, both inclusive.
Valid latitude values are between -90 and 90, both inclusive.
*/
func ValidateCoordinates(coords entities.Coordinates) error {

	if len(coords) != 2 {
		return fmt.Errorf("Invalid coordinates format (%v). It shall be an array of exactly two float numbers - [lng, lat]", coords)
	}

	if coords[0] < -180 || coords[0] > 180 {
		return fmt.Errorf("Invalid longitude (%f). It shall be a float number between -180 and 180.", coords[0])
	}

	if coords[1] < -90 || coords[1] > 90 {
		return fmt.Errorf("Invalid latitude (%f). It shall be a float number between -90 and 90.", coords[1])
	}

	return nil
}
