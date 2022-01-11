package controller

import (
	"awesomeProject/entity"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
)

//뷰단에서 사용할 비디오 컨트롤러
type VideoController interface {
	FindAll() []entity.Video
	Save(context *gin.Context) entity.Video
}

//컨트롤러의 구조체 service를 매핑함
type controller struct {
	service service.VideoService
}

//뷰단에서 사용
//1. service를 New로 생성 뒤 받아옴
//2. 리턴값으로 VideoController 인터페이스를 리턴
//3. 리턴값으로 컨트롤러 구조체에 서비스를 매핑
func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

//뷰단에서 사용할 Video FindAll 메서드 service에서
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

//Save시 Context를 받아오고 Entity를 생성 뒤  context Body JSON의 값을  video주소에 값을 넣고 service를 호출하여 Save시킴
func (c *controller) Save(context *gin.Context) entity.Video {
	var video entity.Video
	context.BindJSON(&video)
	c.service.Save(video)
	return video
}
