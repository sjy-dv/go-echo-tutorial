package main

import (
	"context"
	"echotest/pb"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := pb.NewGRpcAppClient(conn)
/*
	resp, err1 := c.SignUp(context.Background(), &pb.ReqProtoUser{
		Userid: "aaa",
		Username: "proto_테스터",
		Password: "1111",
	})

	if err1 != nil {
		fmt.Printf("result : %s, error : %v", resp.Result, err1)
	}

	fmt.Printf("result : %s", resp.Result)


    resp, err1 := c.SignIn(context.Background(), &pb.LoginProto{
		Userid: "aaa",
		Password: "1111",
	})

	if err1 != nil {
		fmt.Printf("result : %s, error : %v", resp.Token, err1)
	}

	fmt.Printf("result : %s", resp.Token)
	

	resp, err1 := c.UserInfo(context.Background(), &pb.ResToken{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjQ5MzU1MTQsInVzZXJpZCI6ImFhYSJ9.x0j4HkcTNpfdDE-Gq7ZI-RMD6lIGM_8YbqmrR-UOVas",
	})

	if err != nil {
		fmt.Println(resp, err1)
	}
	fmt.Println(resp)
	*/

	resp, err1 := c.AllUser(context.Background(), &pb.QueryPage{
		Page: 1,
	})

	if err != nil {
		fmt.Println(resp, err1)
	}
	fmt.Println(resp.ProtoUser)
}