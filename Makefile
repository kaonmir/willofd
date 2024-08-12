

generate:
	protoc --go_out=gen proto/backbone/*
	protoc --go_out=gen proto/virtualmachine/*