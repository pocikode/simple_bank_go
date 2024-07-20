package gapi

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	db "pocikode/simple-bank/db/sqlc"
	"pocikode/simple-bank/pb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		Fullname:          user.Fullname,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
