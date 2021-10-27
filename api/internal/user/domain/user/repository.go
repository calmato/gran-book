//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../../mock/user/domain/$GOPACKAGE/$GOFILE
package user

import (
	"context"

	"github.com/calmato/gran-book/api/pkg/database"
)

// Repository - Userレポジトリ
type Repository interface {
	Authentication(ctx context.Context) (string, error)
	List(ctx context.Context, q *database.ListQuery) (Users, error)
	ListFollow(ctx context.Context, q *database.ListQuery) (Follows, error)
	ListFollower(ctx context.Context, q *database.ListQuery) (Followers, error)
	ListInstanceID(ctx context.Context, q *database.ListQuery) ([]string, error)
	ListFollowID(ctx context.Context, userID string, userIDs ...string) ([]string, error)
	ListFollowerID(ctx context.Context, userID string, userIDs ...string) ([]string, error)
	Count(ctx context.Context, q *database.ListQuery) (int, error)
	CountRelationship(ctx context.Context, q *database.ListQuery) (int, error)
	MultiGet(ctx context.Context, userIDs []string) (Users, error)
	Get(ctx context.Context, userID string) (*User, error)
	GetAdmin(ctx context.Context, userID string) (*User, error)
	GetRelationship(ctx context.Context, followID string, followerID string) (*Relationship, error)
	GetUserIDByEmail(ctx context.Context, email string) (string, error)
	GetRelationshipIDByUserID(ctx context.Context, followID string, followerID string) (int, error)
	Create(ctx context.Context, u *User) error
	CreateWithOAuth(ctx context.Context, u *User) error
	CreateRelationship(ctx context.Context, r *Relationship) error
	Update(ctx context.Context, u *User) error
	UpdatePassword(ctx context.Context, userID string, password string) error
	Delete(ctx context.Context, userID string) error
	DeleteRelationship(ctx context.Context, relationshipID int) error
}
