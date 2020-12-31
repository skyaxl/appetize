package pathsdomain

import (
	"context"
	"net/http"

	"github.com/skyaxl/synack/pkg/apierrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Path distance between two spots that it connect
// I puth here the points to map
type Path struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AX       int                `json:"ax,omitempty"`
	AY       int                `json:"ay,omitempty"`
	BX       int                `json:"bx,omitempty"`
	BY       int                `json:"by,omitempty"`
	Distance float64            `json:"distance,omitempty"`
}

//Validate Path
func (p Path) Validate() (e error) {
	if p.Distance < 0 {
		return apierrors.APIError{Msg: "Distance must be greather than 0", Status: http.StatusBadRequest}
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
	Paths      []Path     `json:"paths,omitempty"`
}

//PathService implements yor service methods.
type PathsService interface {
	Create(ctx context.Context, spt Path) (Path, error)
	Get(ctx context.Context, id string) (Path, error)
	Delete(ctx context.Context, pathID string) error
	Update(ctx context.Context, path Path) (Path, error)
	GetAll(ctx context.Context, pagination Pagination) (GetAllResponse, error)
}
