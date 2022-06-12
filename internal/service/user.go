package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) Authentication(ctx context.Context, req *v1.AuthenticationRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{User: &v1.UserReply_User{Username: "Xiaomin"}}, nil
}

func (s *RealWorldService) Registration(ctx context.Context, req *v1.RegistrationRequest) (replay *v1.UserReply, err error) {
	return &v1.UserReply{User: &v1.UserReply_User{Username: "Xiaomin"}}, nil
}

