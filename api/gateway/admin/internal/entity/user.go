package entity

import (
	"github.com/calmato/gran-book/api/gateway/admin/proto/service/user"
)

type Auth struct {
	*user.Auth
}

func NewAuth(a *user.Auth) *Auth {
	return &Auth{a}
}

func (a *Auth) Role() Role {
	return NewRole(a.GetRole())
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
