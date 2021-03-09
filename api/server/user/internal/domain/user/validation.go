package user

import "context"

// Validation - Userドメインバリデーション
type Validation interface {
	User(ctx context.Context, u *User) error
	Relationship(ctx context.Context, r *Relationship) error
}
