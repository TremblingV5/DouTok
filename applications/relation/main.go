package main

import (
	"github.com/TremblingV5/DouTok/applications/relation/handler"
	"github.com/TremblingV5/DouTok/kitex_gen/relation/relationservice"
)

func main() {
	s := relationservice.NewServer(&handler.RelationServiceImpl{})
	s.Run()
}
