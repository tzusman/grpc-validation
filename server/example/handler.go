package example

import (
	context "context"

	"github.com/google/uuid"
	p "github.com/journeyai/grpc-validation/protocols/example"
	"github.com/journeyai/grpc-validation/server/util"
)

// Server handles the Example protocol
type Server struct {
	store util.UserStore
}

// NewServer creates a new server
func NewServer() (*Server, error) {
	return &Server{
		store: util.NewUserStore(),
	}, nil
}

// CreateWidget handles the CreateWidget request
func (s Server) CreateWidget(ctx context.Context, in *p.CreateWidgetRequest) (*p.CreateWidgetReply, error) {

	err := validateCreateWidget(*in)
	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	s.store.AddUser(id, util.User{
		Name:        in.Name,
		Age:         in.Age,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
		Attributes:  in.Attributes,
	})

	out := p.CreateWidgetReply{
		Id: id,
	}

	return &out, nil
}

// GetWidget handles the GetWidget request
func (s Server) GetWidget(ctx context.Context, in *p.GetWidgetRequest) (*p.GetWidgetReply, error) {

	err := validateGetWidget(*in)
	if err != nil {
		return nil, err
	}

	user := s.store.RetrieveUser(in.Id)

	out := p.GetWidgetReply{
		Id:          in.Id,
		Name:        user.Name,
		Age:         user.Age,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Attributes:  user.Attributes,
	}

	return &out, nil
}
