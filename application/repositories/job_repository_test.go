package repositories_test

import (
	"encoder-video/application/repositories"
	"encoder-video/domain"
	"encoder-video/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()

	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreateAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)

	require.Nil(t, err)

	repojob := repositories.JobRepositoryDb{Db: db}

	repojob.Insert(job)

	j, err := repojob.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)

	require.Equal(t, j.VideoID, job.VideoID)

}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()

	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreateAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)

	require.Nil(t, err)

	repojob := repositories.JobRepositoryDb{Db: db}

	repojob.Insert(job)

	job.Status = "Complete"

	repojob.Update(job)

	j, err := repojob.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}
