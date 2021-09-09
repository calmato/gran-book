package entity

import (
	pb "github.com/calmato/gran-book/api/gateway/admin/proto"
)

type Auth struct {
	*pb.Auth
}

func NewAuth(a *pb.Auth) *Auth {
	return &Auth{a}
}

func (a *Auth) Role() Role {
	return NewRole(a.GetRole())
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

type Admin struct {
	*pb.Admin
}

type Admins []*Admin

func NewAdmin(a *pb.Admin) *Admin {
	return &Admin{a}
}

func (a *Admin) Role() Role {
	return NewRole(a.GetRole())
}

func NewAdmins(as []*pb.Admin) Admins {
	res := make(Admins, len(as))
	for i := range as {
		res[i] = NewAdmin(as[i])
	}
	return res
}
