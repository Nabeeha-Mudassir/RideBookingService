package main


func main() {
	// Start the gRPC serve
	grpcServer:= NewgRPCServer(":9090") //Pass a database here later!!
	grpcServer.Run()
	
}