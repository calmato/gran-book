package application

import (
	"context"
	"strings"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/application/output"
	"github.com/calmato/gran-book/api/server/user/internal/application/validation"
	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
)

// AdminApplication - Adminアプリケーションのインターフェース
type AdminApplication interface {
	List(ctx context.Context, in *input.ListAdmin) ([]*user.User, *output.ListQuery, error)
	Search(ctx context.Context, in *input.SearchAdmin) ([]*user.User, *output.ListQuery, error)
	Show(ctx context.Context, uid string) (*user.User, error)
	Create(ctx context.Context, in *input.CreateAdmin) (*user.User, error)
	UpdateRole(ctx context.Context, in *input.UpdateAdminRole, uid string) (*user.User, error)
	UpdatePassword(ctx context.Context, in *input.UpdateAdminPassword, uid string) (*user.User, error)
	UpdateProfile(ctx context.Context, in *input.UpdateAdminProfile, uid string) (*user.User, error)
	UploadThumbnail(ctx context.Context, in *input.UploadAdminThumbnail, uid string) (string, error)
}

type adminApplication struct {
	adminRequestValidation validation.AdminRequestValidation
	userService            user.Service
}

// NewAdminApplication - AdminApplicationの生成
func NewAdminApplication(arv validation.AdminRequestValidation, us user.Service) AdminApplication {
	return &adminApplication{
		adminRequestValidation: arv,
		userService:            us,
	}
}

func (a *adminApplication) List(ctx context.Context, in *input.ListAdmin) ([]*user.User, *output.ListQuery, error) {
	err := a.adminRequestValidation.ListAdmin(in)
	if err != nil {
		return nil, nil, err
	}

	query := &domain.ListQuery{
		Limit:  in.Limit,
		Offset: in.Offset,
		Conditions: []*domain.QueryCondition{
			{
				Field:    "role",
				Operator: "BETWEEN",
				Value:    []int{1, 3},
			},
		},
	}

	if in.By != "" {
		o := &domain.QueryOrder{
			By:        in.By,
			Direction: in.Direction,
		}

		query.Order = o
	}

	us, err := a.userService.List(ctx, query)
	if err != nil {
		return nil, nil, err
	}

	total, err := a.userService.ListCount(ctx, query)
	if err != nil {
		return nil, nil, err
	}

	out := &output.ListQuery{
		Limit:  query.Limit,
		Offset: query.Offset,
		Total:  total,
	}

	if query.Order != nil {
		o := &output.QueryOrder{
			By:        query.Order.By,
			Direction: query.Order.Direction,
		}

		out.Order = o
	}

	return us, out, nil
}

func (a *adminApplication) Search(ctx context.Context, in *input.SearchAdmin) ([]*user.User, *output.ListQuery, error) {
	err := a.adminRequestValidation.SearchAdmin(in)
	if err != nil {
		return nil, nil, err
	}

	query := &domain.ListQuery{
		Limit:  in.Limit,
		Offset: in.Offset,
		Conditions: []*domain.QueryCondition{
			{
				Field:    "role",
				Operator: "BETWEEN",
				Value:    []int{1, 3},
			},
			{
				Field:    in.Field,
				Operator: "LIKE",
				Value:    in.Value,
			},
		},
	}

	if in.By != "" {
		o := &domain.QueryOrder{
			By:        in.By,
			Direction: in.Direction,
		}

		query.Order = o
	}

	us, err := a.userService.List(ctx, query)
	if err != nil {
		return nil, nil, err
	}

	total, err := a.userService.ListCount(ctx, query)
	if err != nil {
		return nil, nil, err
	}

	out := &output.ListQuery{
		Limit:  query.Limit,
		Offset: query.Offset,
		Total:  total,
	}

	if query.Order != nil {
		o := &output.QueryOrder{
			By:        query.Order.By,
			Direction: query.Order.Direction,
		}

		out.Order = o
	}

	return us, out, nil
}

func (a *adminApplication) Show(ctx context.Context, uid string) (*user.User, error) {
	u, err := a.userService.Show(ctx, uid)
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return u, nil
}

func (a *adminApplication) Create(ctx context.Context, in *input.CreateAdmin) (*user.User, error) {
	err := a.adminRequestValidation.CreateAdmin(in)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		Username:      in.Username,
		Email:         strings.ToLower(in.Email),
		Password:      in.Password,
		Gender:        0,
		Role:          in.Role,
		LastName:      in.LastName,
		FirstName:     in.FirstName,
		LastNameKana:  in.LastNameKana,
		FirstNameKana: in.FirstNameKana,
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

func (a *adminApplication) UpdateRole(ctx context.Context, in *input.UpdateAdminRole, uid string) (*user.User, error) {
	err := a.adminRequestValidation.UpdateAdminRole(in)
	if err != nil {
		return nil, err
	}

	u, err := a.userService.Show(ctx, uid)
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	u.Role = in.Role

	err = a.userService.Validation(ctx, u)
	if err != nil {
		return nil, err
	}

	err = a.userService.Update(ctx, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a *adminApplication) UpdatePassword(
	ctx context.Context, in *input.UpdateAdminPassword, uid string,
) (*user.User, error) {
	err := a.adminRequestValidation.UpdateAdminPassword(in)
	if err != nil {
		return nil, err
	}

	u, err := a.userService.Show(ctx, uid)
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	err = a.userService.UpdatePassword(ctx, uid, in.Password)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a *adminApplication) UpdateProfile(
	ctx context.Context, in *input.UpdateAdminProfile, uid string,
) (*user.User, error) {
	err := a.adminRequestValidation.UpdateAdminProfile(in)
	if err != nil {
		return nil, err
	}

	u, err := a.userService.Show(ctx, uid)
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	u.Username = in.Username
	u.Email = in.Email
	u.LastName = in.LastName
	u.FirstName = in.FirstName
	u.LastNameKana = in.LastNameKana
	u.FirstNameKana = in.FirstNameKana

	err = a.userService.Validation(ctx, u)
	if err != nil {
		return nil, err
	}

	err = a.userService.Update(ctx, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a *adminApplication) UploadThumbnail(
	ctx context.Context, in *input.UploadAdminThumbnail, uid string,
) (string, error) {
	err := a.adminRequestValidation.UploadAdminThumbnail(in)
	if err != nil {
		return "", err
	}

	return a.userService.UploadThumbnail(ctx, uid, in.Thumbnail)
}
