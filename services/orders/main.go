package main

func main() {
	gRPCServer := NewGRPCServer(":3000")
	gRPCServer.Run()
}
