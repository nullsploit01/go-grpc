package main

func main() {
	httpServer := NewHttpServer(":8000")
	go httpServer.Run()

	gRPCServer := NewGRPCServer(":3000")
	gRPCServer.Run()
}
