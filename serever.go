package main

import (
	"awesomeProject/controller"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
)

var (
	//비디오 컨트롤러 인터페이스를 New로 생성 후 service를 매핑
	videoController controller.VideoController = controller.New(videoService)
	//비디오 서비스를 구현한 인터페이스를 New 로 생성
	videoService service.VideoService = service.New()
)

func main() {

	server := gin.Default()

	server.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})

	server.Run(":8080")

}
