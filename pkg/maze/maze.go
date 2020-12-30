package maze

import (
	"math"
	"net/http"
)

//Path distance between two spots that it connect
// I puth here the points to map
type Path struct {
	AX, AY, BX, BY int
	Distance       float64
}

func (p Path) Validate() (e error) {

	if p.Distance < 0 {
		return ApiError{Msg: "Distance must be greather than 0", Status: http.StatusBadRequest}
	}
	return nil
}

//Spot is a location. It should have two coordinates, x and y (coordinates on a cartesian plane),
// a name (exit, entrance, treasure room etc.. it can be any string),
//and a number (the amount of gold that's in that spot
type Spot struct {
	X, Y   int
	Name   string
	Amount float64
}

//Validate Spot
func (s Spot) Validate() error {
	if s.Name == "" {
		return ApiError{Msg: "Name must have a valid value not empty", Status: http.StatusBadRequest}
	}

	if s.Amount >= 0 {
		return ApiError{Msg: "Amount must have a valid value > 0 ", Status: http.StatusBadRequest}
	}

	return nil
}

//VerticalDirection define a  vertical direction to cartesian quadrant
type VerticalDirection string

//HorizontalDirection define a horizontal direction to cartesian quadrant
type HorizontalDirection string

const (
	//TopVerticalDirection is top relative direction
	TopVerticalDirection VerticalDirection = "top"
	//BottomVerticalDirection is bottom relative direction
	BottomVerticalDirection VerticalDirection = "bottom"
	//LeftHorizontalDirection is left direction
	LeftHorizontalDirection HorizontalDirection = "left"
	//RightHorizontalDirection is right direction
	RightHorizontalDirection HorizontalDirection = "right"
)

//Quadrant define a cartesian quadrant, where we have a start point and de X and Y direcction that define quadrante direction
type Quadrant struct {
	Start Spot
	//left or right
	HorizontalDirection HorizontalDirection
	//VerticalDirection bottom or top
	VerticalDirection VerticalDirection
}

//Validate if start spot is orthogonal and x|y direction are valid values
func (q Quadrant) Validate() (e error) {
	x := int(math.Abs(float64(q.Start.X)))
	y := int(math.Abs(float64(q.Start.Y)))

	if x%y != 0 {
		return ApiError{Msg: "Start spot is not orthogonal", Status: http.StatusBadRequest}
	}

	if q.HorizontalDirection == "" {
		return ApiError{Msg: "HorizontalDirection must have a valid value 'left' or 'right'", Status: http.StatusBadRequest}
	}

	if q.VerticalDirection == "" {
		return ApiError{Msg: "VerticalDirection must have a valid value 'top' or 'bottom'", Status: http.StatusBadRequest}
	}

	return nil
}

//Maze Define a maze
type Maze struct {
	ID        interface{} `json:"_id,omitempty"`
	Spots     []Spot      `json:"spots,omitempty"`
	Paths     []Path      `json:"paths,omitempty"`
	Quadrants []Quadrant  `json:"quadrants,omitempty"`
}

//Validate validate total maze
func (m Maze) Validate() (err error) {
	for _, v := range m.Spots {
		if err = v.Validate(); err != nil {
			return err
		}
	}

	for _, q := range m.Quadrants {
		if err = q.Validate(); err != nil {
			return err
		}
	}

	for _, p := range m.Paths {
		if err = p.Validate(); err != nil {
			return err
		}
	}

	return nil

}
