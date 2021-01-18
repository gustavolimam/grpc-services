package main

import (
	"database/sql"
	"grpc-services/internal/proto/messagepb"
	"grpc-services/internal/services/client"
	"grpc-services/internal/services/portdomain"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	go startClient()
	startgRPC()
}

func startClient() {
	godotenv.Load()
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := messagepb.NewMessageServiceClient(conn)
	s := client.NewService(c)

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/add", s.AddDataFromJson)
	myRouter.HandleFunc("/data", s.GetAllData)

	log.Println("Server start ", os.Getenv("CLIENT_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("CLIENT_PORT"), myRouter))
}

func startgRPC() {
	godotenv.Load()
	listen, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}
	log.Println("Listening gRPC Server at port: ", os.Getenv("GRPC_PORT"))

	grpcServer := grpc.NewServer()

	db, err := sql.Open("postgres", os.Getenv("DB_CONN"))
	if err != nil {
		log.Fatalln("Error database - ", err)
	}

	service := portdomain.Service{
		DB: db,
	}

	messagepb.RegisterMessageServiceServer(grpcServer, &service)

	defer db.Close()

	log.Fatal(grpcServer.Serve(listen))
}
