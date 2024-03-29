//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../mock/user/$GOPACKAGE/$GOFILE
package application

import (
	"context"

	"github.com/calmato/gran-book/api/internal/user/domain/chat"
	"github.com/calmato/gran-book/api/internal/user/domain/user"
	"github.com/calmato/gran-book/api/pkg/database"
	"github.com/calmato/gran-book/api/pkg/firebase/firestore"
)

type UserApplication interface {
	Authentication(ctx context.Context) (*user.User, error)
	List(ctx context.Context, q *database.ListQuery) (user.Users, int, error)
	ListAdmin(ctx context.Context, q *database.ListQuery) (user.Users, int, error)
	ListFollow(ctx context.Context, userID, targetID string, limit, offset int) (user.Follows, int, error)
	ListFollower(ctx context.Context, userID, targetID string, limit, offset int) (user.Followers, int, error)
	MultiGet(ctx context.Context, userIDs []string) (user.Users, error)
	Get(ctx context.Context, userID string) (*user.User, error)
	GetAdmin(ctx context.Context, userID string) (*user.User, error)
	GetUserProfile(ctx context.Context, userID, targetID string) (*user.User, error)
	Create(ctx context.Context, u *user.User) error
	Update(ctx context.Context, u *user.User) error
	UpdatePassword(ctx context.Context, u *user.User) error
	Delete(ctx context.Context, u *user.User) error
	DeleteAdmin(ctx context.Context, u *user.User) error
	Follow(ctx context.Context, userID, followerID string) (*user.User, error)
	Unfollow(ctx context.Context, userID, followerID string) (*user.User, error)
	UploadThumbnail(ctx context.Context, userID string, thumbnail []byte) (string, error)
}

type ChatApplication interface {
	ListRoom(ctx context.Context, userID string, p *firestore.Params) (chat.Rooms, error)
	GetRoom(ctx context.Context, roomID string, userID string) (*chat.Room, error)
	CreateRoom(ctx context.Context, cr *chat.Room) error
	CreateMessage(ctx context.Context, cr *chat.Room, cm *chat.Message) error
	UploadImage(ctx context.Context, cr *chat.Room, image []byte) (string, error)
}
