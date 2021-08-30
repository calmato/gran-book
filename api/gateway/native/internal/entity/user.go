package entity

import (
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
)

type Auth struct {
	*pb.Auth
}

func NewAuth(a *pb.Auth) *Auth {
	return &Auth{a}
}

func (a *Auth) Gender() Gender {
	return NewGender(a.GetGender())
}

type User struct {
	*pb.User
}

type Users []*User

func NewUser(u *pb.User) *User {
	return &User{u}
}

func NewUsers(us []*pb.User) Users {
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

type UserProfile struct {
	*pb.UserProfile
}

func NewUserProfile(p *pb.UserProfile) *UserProfile {
	return &UserProfile{p}
}

type Follow struct {
	*pb.Follow
}

type Follows []*Follow

func NewFollow(f *pb.Follow) *Follow {
	return &Follow{f}
}

func NewFollows(fs []*pb.Follow) Follows {
	res := make(Follows, len(fs))
	for i := range fs {
		res[i] = NewFollow(fs[i])
	}
	return res
}

type Follower struct {
	*pb.Follower
}

type Followers []*Follower

func NewFollower(f *pb.Follower) *Follower {
	return &Follower{f}
}

func NewFollowers(fs []*pb.Follower) Followers {
	res := make(Followers, len(fs))
	for i := range fs {
		res[i] = NewFollower(fs[i])
	}
	return res
}
