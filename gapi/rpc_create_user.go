package gapi

import (
	"context"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	db "pocikode/simple-bank/db/sqlc"
	"pocikode/simple-bank/pb"
	"pocikode/simple-bank/util"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		Fullname:       req.GetFullname(),
		Email:          req.GetEmail(),
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "user already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	response := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return response, nil
}