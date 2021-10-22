package entity

import (
	"github.com/calmato/gran-book/api/service/proto/user"
)

type Auth struct {
	*user.Auth
}

func NewAuth(a *user.Auth) *Auth {
	return &Auth{a}
}

func (a *Auth) Gender() Gender {
	return NewGender(a.GetGender())
}

type User struct {
	*user.User
}

type Users []*User

func NewUser(u *user.User) *User {
	return &User{u}
}

func NewUsers(us []*user.User) Users {
	res := make(Users, len(us))
	for i := range us {
		res[i] = NewUser(us[i])
	}
	return res
}

func (us Users) Map() map[string]*User {
	res := make(map[string]*User, len(us))
	for _, u := range us {
		res[u.Id] = u
	}
	return res
}

func (us Users) IsExists(userIDs ...string) bool {
	m := us.Map()
	for _, userID := range userIDs {
		if _, ok := m[userID]; ok {
			continue
		}

		return false
	}
	return true
}

type UserProfile struct {
	*user.UserProfile
}

func NewUserProfile(p *user.UserProfile) *UserProfile {
	return &UserProfile{p}
}

type Follow struct {
	*user.Follow
}

type Follows []*Follow

func NewFollow(f *user.Follow) *Follow {
	return &Follow{f}
}

func NewFollows(fs []*user.Follow) Follows {
	res := make(Follows, len(fs))
	for i := range fs {
		res[i] = NewFollow(fs[i])
	}
	return res
}

type Follower struct {
	*user.Follower
}

type Followers []*Follower

func NewFollower(f *user.Follower) *Follower {
	return &Follower{f}
}

func NewFollowers(fs []*user.Follower) Followers {
	res := make(Followers, len(fs))
	for i := range fs {
		res[i] = NewFollower(fs[i])
	}
	return res
}

type Admin struct {
	*user.Admin
}

type Admins []*Admin

func NewAdmin(a *user.Admin) *Admin {
	return &Admin{a}
}

func (a *Admin) Role() Role {
	return NewRole(a.GetRole())
}

func NewAdmins(as []*user.Admin) Admins {
	res := make(Admins, len(as))
	for i := range as {
		res[i] = NewAdmin(as[i])
	}
	return res
}
