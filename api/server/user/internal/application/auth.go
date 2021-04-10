package application

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/application/validation"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
)

// AuthApplication - Authアプリケーションのインターフェース
type AuthApplication interface {
	Authentication(ctx context.Context) (*user.User, error)
	Create(ctx context.Context, in *input.CreateAuth) (*user.User, error)
	UpdateEmail(ctx context.Context, in *input.UpdateAuthEmail, u *user.User) error
	UpdatePassword(ctx context.Context, in *input.UpdateAuthPassword, u *user.User) error
	UpdateProfile(ctx context.Context, in *input.UpdateAuthProfile, u *user.User) error
	UpdateAddress(ctx context.Context, in *input.UpdateAuthAddress, u *user.User) error
	UploadThumbnail(ctx context.Context, in *input.UploadAuthThumbnail, u *user.User) (string, error)
}

type authApplication struct {
	authRequestValidation validation.AuthRequestValidation
	userService           user.Service
}

// NewAuthApplication - AuthApplicationの生成
func NewAuthApplication(arv validation.AuthRequestValidation, us user.Service) AuthApplication {
	return &authApplication{
		authRequestValidation: arv,
		userService:           us,
	}
}

func (a *authApplication) Authentication(ctx context.Context) (*user.User, error) {
	uid, err := a.userService.Authentication(ctx)
	if err != nil {
		return nil, err
	}

	u, err := a.userService.Show(ctx, uid)
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return u, nil
}

func (a *authApplication) Create(ctx context.Context, in *input.CreateAuth) (*user.User, error) {
	err := a.authRequestValidation.CreateAuth(in)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		Username:  in.Username,
		Email:     strings.ToLower(in.Email),
		Password:  in.Password,
		Gender:    0,
		Role:      user.UserRole,
		Activated: true,
	}

	err = a.userService.Validation(ctx, u)
	if err != nil {
		return nil, err
	}

	err = a.userService.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a *authApplication) UpdateEmail(ctx context.Context, in *input.UpdateAuthEmail, u *user.User) error {
	err := a.authRequestValidation.UpdateAuthEmail(in)
	if err != nil {
		return err
	}

	u.Email = strings.ToLower(in.Email)

	err = a.userService.Validation(ctx, u)
	if err != nil {
		return err
	}

	return a.userService.Update(ctx, u)
}

func (a *authApplication) UpdatePassword(ctx context.Context, in *input.UpdateAuthPassword, u *user.User) error {
	err := a.authRequestValidation.UpdateAuthPassword(in)
	if err != nil {
		return err
	}

	return a.userService.UpdatePassword(ctx, u.ID, in.Password)
}

func (a *authApplication) UpdateProfile(ctx context.Context, in *input.UpdateAuthProfile, u *user.User) error {
	err := a.authRequestValidation.UpdateAuthProfile(in)
	if err != nil {
		return err
	}

	thumbnailURL, err := a.getThumbnailURL(ctx, u.ID, in.Thumbnail)
	if err != nil {
		return err
	}

	u.Username = in.Username
	u.Gender = in.Gender
	u.SelfIntroduction = in.SelfIntroduction

	// TODO: 古いサムネイルを消す処理を挟みたい
	if thumbnailURL != "" {
		u.ThumbnailURL = thumbnailURL
	}

	err = a.userService.Validation(ctx, u)
	if err != nil {
		return err
	}

	return a.userService.Update(ctx, u)
}

func (a *authApplication) UpdateAddress(ctx context.Context, in *input.UpdateAuthAddress, u *user.User) error {
	err := a.authRequestValidation.UpdateAuthAddress(in)
	if err != nil {
		return err
	}

	u.LastName = in.LastName
	u.FirstName = in.FirstName
	u.LastNameKana = in.LastNameKana
	u.FirstNameKana = in.FirstNameKana
	u.PhoneNumber = in.PhoneNumber
	u.PhoneNumber = in.PhoneNumber
	u.PostalCode = in.PostalCode
	u.Prefecture = in.Prefecture
	u.City = in.City
	u.AddressLine1 = in.AddressLine1
	u.AddressLine2 = in.AddressLine2

	err = a.userService.Validation(ctx, u)
	if err != nil {
		return err
	}

	return a.userService.Update(ctx, u)
}

func (a *authApplication) UploadThumbnail(
	ctx context.Context, in *input.UploadAuthThumbnail, u *user.User,
) (string, error) {
	err := a.authRequestValidation.UploadAuthThumbnail(in)
	if err != nil {
		return "", err
	}

	return a.userService.UploadThumbnail(ctx, u.ID, in.Thumbnail)
}

func (a *authApplication) getThumbnailURL(ctx context.Context, uid string, thumbnail string) (string, error) {
	if thumbnail == "" {
		return "", nil
	}

	// data:image/png;base64,iVBORw0KGgoAAAA... みたいなのうちの
	// `data:image/png;base64,` の部分を無くした []byte を取得
	b64data := thumbnail[strings.IndexByte(thumbnail, ',')+1:]

	data, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		ve := &exception.ValidationError{
			Field:   "thumbnail",
			Message: exception.UnableConvertBase64Massage,
		}

		return "", exception.UnableConvertBase64.New(err, ve)
	}

	return a.userService.UploadThumbnail(ctx, uid, data)
}
