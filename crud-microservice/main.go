package main

import (
	"log"
	"net"

	"crud-microservice/crud-microservice/pb"
	"crud-microservice/store"

	usergRPC "crud-microservice/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// // RESTful API
	// userStore := store.NewUserStore()
	// userHandler := handlers.NewUserHandler(userStore)

	// router := mux.NewRouter()

	// // Routes
	// router.HandleFunc("/createUser", userHandler.CreateUser).Methods("POST")
	// router.HandleFunc("/getAllUsers", userHandler.GetAllUsers).Methods("GET")
	// router.HandleFunc("/getUser/{id}", userHandler.GetUser).Methods("GET")
	// router.HandleFunc("/updateUser/{id}", userHandler.UpdateUser).Methods("PUT")
	// router.HandleFunc("/deleteUser/{id}", userHandler.DeleteUser).Methods("DELETE")

	// log.Println("Server listening on port :8080")
	// log.Fatal(http.ListenAndServe(":8080", router))

	// gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &usergRPC.UserServiceServer{
		Store: *store.NewUserStore(), // or postgres later
	})

	// Enable reflection for grpcurl and other tools
	reflection.Register(s)

	log.Println("gRCP server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
