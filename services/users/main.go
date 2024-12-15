package main
import (

	"log"
	
)
func main() {
	
	//Initialise database
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	
	
	// Start the gRPC serve
	grpcServer:= NewgRPCServer(":9090",db) //Pass a database here later!!
	grpcServer.Run()
}