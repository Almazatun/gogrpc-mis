package main

func main() {
	grpcServer := NewGRPCServer(":5002")
	err := grpcServer.Run()

	if err != nil {
		panic(err)
	}
}
