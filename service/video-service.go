package service

import "awesomeProject/entity"

//비디오 서비스 인터페이스 구현
type VideoService interface {
	//저장용 video entity를 한개 받아옴
	Save(video entity.Video) entity.Video
	// 전부 찾아서 배열로 반환
	FindAll() []entity.Video
}

//비디오 서비스 구조체 값을 반환함
type videoService struct {
	videos []entity.Video
}

//새롭게 생성하면 비디오 서비스 배열의 구조체를 반환

func New() VideoService {
	return &videoService{}
}

//메서드 비디오 서비스의 Save로 엔티디의 값을 가져와서 추가함  그리고 video를 반환
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

//비디오를 전부 반환
func (service *videoService) FindAll() []entity.Video {
	return service.videos

}
