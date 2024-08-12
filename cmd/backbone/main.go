package main

import (
	"log"

	"github.com/gin-gonic/gin"
	gen "github.com/kaonmir/willofd/gen/backbone"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Person 객체 생성
	elliot := &gen.Test{}

	// 직렬화
	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 기본 포트 8080에서 서버 시작
}
