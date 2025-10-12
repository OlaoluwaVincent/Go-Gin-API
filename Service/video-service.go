package service

import (
	"go/tutorial/entity"
)

type VideoService struct {
	videos []entity.Video
}

func (s *VideoService) FindAll() []entity.Video {
	return s.videos
}

func (s *VideoService) Save(v entity.Video) entity.Video {
	s.videos = append(s.videos, v)
	return v
}

func (s *VideoService) Delete(id string) bool {
	for i, v := range s.videos {
		if v.Id == id {
			s.videos = append(s.videos[:i], s.videos[i+1:]...)
			return true
		}
	}
	return false
}

// Call this function to Instantiate the VideoService
// After instantiate, you can call the methods on the returned object
func NewVideoService() *VideoService {
	return &VideoService{}
}
