package storage

import "fmt"

const (
	chatImageDIR     = "chat_images"
	userThumbnailDIR = "user_thumbnails"
)

func getUserThumbnailPath(userID string) string {
	return fmt.Sprintf("users/%s/%s", userID, userThumbnailDIR)
}

func getChatMessagePath(roomID string) string {
	return fmt.Sprintf("rooms/%s/%s", roomID, chatImageDIR)
}
