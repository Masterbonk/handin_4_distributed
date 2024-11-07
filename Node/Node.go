package main

import (
	cc "Client"
)

type server struct {
	cc.UnimplementedClientServer
	msg string
	key int32
}
