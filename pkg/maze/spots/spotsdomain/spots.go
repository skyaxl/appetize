package spotsdomain

import (
	"context"
	"net/http"

	"github.com/skyaxl/synack/pkg/apierrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Spot is a location. It should have two coordinates, x and y (coordinates on a cartesian plane),
// a name (exit, entrance, treasure room etc.. it can be any string),
//and a number (the amount of gold that's in that spot
type Spot struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	X      int                `json:"x,omitempty"`
	Y      int                `json:"y,omitempty"`
	Name   string             `json:"name,omitempty"`
	Amount float64            `json:"amount,omitempty"`
}

//Validate Spot
func (s Spot) Validate() error {
	if s.Name == "" {
		return apierrors.APIError{Msg: "Name must have a valid value not empty", Status: http.StatusBadRequest}
	}

	if s.Amount <= 0 {
		return apierrors.APIError{Msg: "Amount must have a valid value > 0 ", Status: http.StatusBadRequest}
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
	Spots      []Spot     `json:"spots,omitempty"`
}

//SpotService implements yor service methods.
type SpotsService interface {
	Create(ctx context.Context, spt Spot) (Spot, error)
	Get(ctx context.Context, id string) (Spot, error)
	Delete(ctx context.Context, spotID string) error
	Update(ctx context.Context, spot Spot) (Spot, error)
	GetAll(ctx context.Context, pagination Pagination) (GetAllResponse, error)
}
