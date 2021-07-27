package utils_test

import (
	"encoder-video/framework/utils"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestIsJson(t *testing.T) {
	json := `{
		"id": "234234324",
		"file_path": "convite.mp4",
		"status": "pending"
	}`
	err := utils.IsJson(json)
	require.Nil(t, err)

	json = `was`

	err = utils.IsJson(json)
	require.Error(t, err)
}