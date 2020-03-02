package server

import (
	"github.com/gofrs/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"promotioner/cmd/authService/manager"
	auth "promotioner/protobuf"
)

type AuthService struct {
	SessionManger manager.SessionManagerInterface
}

func NewAuthService(manager manager.SessionManagerInterface) *AuthService {
	return &AuthService{SessionManger: manager}
}

func (a AuthService) Create(ctx context.Context, in *auth.UserID) (*auth.Token, error) {
	u := uuid.Must(uuid.NewV4())
	err := a.SessionManger.SetTokenForUserID(u.String(), in.UserID)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, "cannot create token for user %v", err.Error())
	}
	return &auth.Token{ID: u.String()}, nil
}

func (a AuthService) Check(ctx context.Context, in *auth.Token) (*auth.UserID, error) {
	userID, err := a.SessionManger.GetUserIDByToken(in.ID)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "user not authenticated")
	}
	return &auth.UserID{UserID: userID}, nil
}

func (a AuthService) Delete(ctx context.Context, in *auth.Token) (*auth.Nothing, error) {
	err := a.SessionManger.DeleteToken(in.ID)
	if err != nil {
		return &auth.Nothing{Dummy: true}, status.Errorf(codes.NotFound, "cannot delete token", err.Error())
	}
	return &auth.Nothing{Dummy: true}, nil
}
