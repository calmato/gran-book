//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../../mock/user/domain/$GOPACKAGE/$GOFILE
package user

import "context"

// Validation - Userドメインバリデーション
type Validation interface {
	User(ctx context.Context, u *User) error
	Relationship(ctx context.Context, r *Relationship) error
}
