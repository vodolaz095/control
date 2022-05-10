package server

import (
	"gitflic.ru/vodolaz095/control/pb"
	"log"
)

type RPC struct {
	pb.UnimplementedControlServer
	Logger *log.Logger
}
