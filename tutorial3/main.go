package main

import (
	"context"
	"echotest/pb"
	"echotest/utils"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	gorm.Model
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Shop struct {
	gorm.Model
	Userid   string `json:"userid"`
	Shopname string `json:"shopname"`
	ShopInfo string `json:"shop_info"`
}

type Server struct {
	DB *gorm.DB
	pb.UnimplementedGRpcAppServer
}

func (server *Server) DBConnection() {
	
	var err error

	host := "root:1111@tcp(127.0.0.1:3306)/rootdb?charset=utf8mb4&parseTime=True&loc=Local"
	server.DB , err = gorm.Open(mysql.Open(host), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	server.DB.AutoMigrate(
		&User{},
		&Shop{},
	)

	setup, err := server.DB.DB()

	if err != nil {
		log.Fatal(err)
	}

	//server.DB.Migrator().CreateView()

	setup.SetMaxIdleConns(0)
	setup.SetMaxOpenConns(5)
	setup.SetConnMaxLifetime(time.Hour)

}


func main() {

	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("server error : %v", err)
	}

	server := Server{}
	server.DBConnection()
	s := grpc.NewServer()
	pb.RegisterGRpcAppServer(s, &server)
	log.Printf("server on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed err : %v", err)
	}
}


func (s *Server) SignUp(ctx context.Context, req *pb.ReqProtoUser) (*pb.Response, error) {

	u := User{}
	u.Userid = req.Userid
	u.Username = req.Username
	hash, _ := utils.BcryptHashSync(req.Password)
	u.Password = string(hash)

	err := s.DB.Debug().Create(&u).Error

	if err != nil {
		return &pb.Response{
			Result: "fail",
		}, err
	}
	return &pb.Response{
		Result: "success",
	}, nil
}

func (s *Server) SignIn(ctx context.Context, req *pb.LoginProto) (*pb.ResToken, error) {

	u := User{}

	err := s.DB.Debug().Model(&User{}).Where("userid = ?", req.Userid).Take(&u).Error

	if err != nil {
		return &pb.ResToken{}, err
	}

	check := utils.BcryptCompareSync(u.Password, req.Password)

	if check != nil {
		return &pb.ResToken{}, check
	}

	token, token_err := utils.CreateToken(req.Userid)

	if token_err != nil {
		return &pb.ResToken{}, token_err
	}

	return &pb.ResToken{
		Token: token,
	}, nil
}

func (s *Server) UserInfo(ctx context.Context, req *pb.ResToken) (*pb.ProtoUser, error) {

	//u := User{}
	p_u := pb.ProtoUser{}

	userid := utils.VerifyToken(req.Token)

	err := s.DB.Debug().Model(&User{}).Where("userid = ?", userid).Take(&p_u).Error

	if err != nil {
		return &pb.ProtoUser{}, err
	}
	return &p_u, nil
}

func (s *Server) AllUser(ctx context.Context, req *pb.QueryPage) (*pb.ProtoUsers, error) {

	offset := 0
	page := int(req.Page)
	if page > 1 {
		offset = 10 * (page -1)
	}
	
	pus := pb.ProtoUsers{}
	pu := []pb.ProtoUser{}
	err := s.DB.Debug().Model(&User{}).Limit(10).Offset(offset).Scan(&pu).Error

	if err != nil {
		return &pb.ProtoUsers{}, err
	}

	return &pus, nil
}