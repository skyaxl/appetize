package quadrantsdomain

import (
	"context"
	"math"
	"net/http"

	"github.com/skyaxl/synack/pkg/apierrors"
	"github.com/skyaxl/synack/pkg/maze/spots/spotsdomain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Start spotsdomain.Spot   `json:"start,omitempty"`
	//left or right
	HorizontalDirection HorizontalDirection `json:"horizontal_direction,omitempty"`
	//VerticalDirection bottom or top
	VerticalDirection VerticalDirection `json:"vertical_direction,omitempty"`
}

//Validate if start spot is orthogonal and x|y direction are valid values
func (q Quadrant) Validate() (e error) {
	x := int(math.Abs(float64(q.Start.X)))
	y := int(math.Abs(float64(q.Start.Y)))

	if x%y != 0 {
		return apierrors.APIError{Msg: "Start spot is not orthogonal a valid value its (0,0) (2,2)", Status: http.StatusBadRequest}
	}

	if q.HorizontalDirection != LeftHorizontalDirection && q.HorizontalDirection != RightHorizontalDirection {
		return apierrors.APIError{Msg: "horizontal_direction must have a valid value 'left' or 'right'", Status: http.StatusBadRequest}
	}

	if q.VerticalDirection != TopVerticalDirection && q.VerticalDirection != BottomVerticalDirection {
		return apierrors.APIError{Msg: "vertical_direction must have a valid value 'top' or 'bottom'", Status: http.StatusBadRequest}
	}

	return nil
}

type Pagination struct {
	Limit int64 `json:"limit,omitempty" schema:"limit,required"`
	Page  int64 `json:"page,omitempty" schema:"page,required"`
	Total int64 `json:"total,omitempty"`
}

type GetAllResponse struct {
	Pagination Pagination `json:"pagination,omitempty"`
	Quadrants  []Quadrant `json:"quadrants,omitempty"`
}

//QuadrantsService implements yor service methods.
type QuadrantsService interface {
	Create(ctx context.Context, spt Quadrant) (Quadrant, error)
	Get(ctx context.Context, id string) (Quadrant, error)
	Delete(ctx context.Context, quadrantID string) error
	Update(ctx context.Context, quadrant Quadrant) (Quadrant, error)
	GetAll(ctx context.Context, pagination Pagination) (GetAllResponse, error)
}
