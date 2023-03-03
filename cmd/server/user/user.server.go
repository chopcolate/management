package user

import (
	"context"
	"log"
	"os"

	"github.com/chopcolate/management/internal/common/database"
	userPack "github.com/chopcolate/management/internal/user"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	// u "github.com/chopcolate/management/pkg/proto/user/user.proto"
)

type UserServer struct {
	Service *userPack.UserService
}

func Start(ctx context.Context) {
	server := UserServer{}
	err := godotenv.Load("internal/user/config/.env")
	if err != nil {
		panic(err)
	}
	db, err := database.NewDatabase(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Fatal("Successfully connected to database")

	server.Service = userPack.NewUserService(db.DB)

	grpcServer := grpc.NewServer()

	u.RegisterUserServiceServer(grpcServer, server.Service)

	defer db.DB.Close()
}
