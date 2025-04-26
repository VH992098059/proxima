package account_edit_delete

import (
	"context"
	"proxima/app/user/api/pbentity"
	"proxima/app/user/internal/dao"
	"proxima/app/user/utility"
)

func UpdateUser(ctx context.Context, in *pbentity.Users) (isOk int, err error) {
	in.Password, err = utility.Encrypt(in.Password)
	if err != nil {
		return -1, err
	}
	_, err = dao.Users.Ctx(ctx).Data(&pbentity.Users{
		Username: in.Username,
		Password: in.Password,
		Email:    in.Email,
	}).Update()
	if err != nil {
		return -1, err
	}
	return 1, nil
}
func DeleteUser(ctx context.Context, in *pbentity.Users) (isDel int, err error) {
	_, err = dao.Users.Ctx(ctx).Where("id", in.Id).Unscoped().Delete()
	if err != nil {
		return -1, err
	}
	return 1, nil

}
