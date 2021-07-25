package storage

import "fmt"

const (
	chatImageDIR     = "chat_images"
	userThumbnailDIR = "user_thumbnails"
)

func getChatMessagePath(roomID string) string {
	return fmt.Sprintf("rooms/%s/%s", roomID, chatImageDIR)
}
