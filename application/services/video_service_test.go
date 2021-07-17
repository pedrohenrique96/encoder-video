package services_test

import (
	"encoder-video/application/repositories"
	"encoder-video/application/services"
	"encoder-video/domain"
	"encoder-video/framework/database"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {

	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}



func prepare() (*domain.Video, repositories.VideoRepositoryDb) {
	db := database.NewDbTest()

	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "VID_20210603_114822.mp4"
	video.CreateAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	return video, repo
	
}

func TestVideoserviceDownload(t *testing.T) {
	video, repo := prepare()
	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err :=videoService.Download("testencoder")

	require.Nil(t, err)
	err = videoService.Fragment()
	require.Nil(t, err)
}