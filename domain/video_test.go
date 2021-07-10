package domain_test

import (
	"encoder-video/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotAUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "abc"
	video.ResourseID = "a"
	video.FilePath = "path"
	video.CreateAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}
func TestVideoValidate(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourseID = "a"
	video.FilePath = "path"
	video.CreateAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}
