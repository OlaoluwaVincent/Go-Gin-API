package interfaces

import (
	"go/tutorial/entity"
)

type VideoInterface interface {
	FindAll() []entity.Video
	Save(v entity.Video) entity.Video
	Delete(id string) bool
}
