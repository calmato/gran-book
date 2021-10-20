package storage

import "fmt"

const (
	userThumbnailDIR = "user_thumbnails"
)

func getUserThumbnailPath(userID string) string {
	return fmt.Sprintf("users/%s/%s", userID, userThumbnailDIR)
}
