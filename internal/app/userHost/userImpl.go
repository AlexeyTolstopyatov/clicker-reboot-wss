package userHost

import (
	"clicker/internal/def/models"
	"golang.org/x/net/context"
)

func (u *UserHost) Register(ctx context.Context, user models.RegistryUserModel) error {
	teamModel, err := u.teamApi.SelectTeamByName(ctx, user.TeamName)
	if err != nil {
		return err
	}

	_, err = u.userApi.InsertUser(ctx, models.CreateUserModel{
		TelegramId: user.TelegramId,
		TeamId:     int64(teamModel.Id),
	})

	return err
}

func (u *UserHost) Login(ctx context.Context, tgId int64) (bool, error) {
	_, err := u.userApi.SelectUserByTelegramId(ctx, tgId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *UserHost) Click(ctx context.Context, tgId int64) error {
	user, err := u.userApi.SelectUserByTelegramId(ctx, tgId)
	if err != nil {
		return err
	}

	totalClicks := user.TotalClicks + 1

	if err := u.thostApi.WithinTransaction(ctx, func(txCtx context.Context) error {
		if err := u.userApi.UpdateTotalClicksByTelegramId(ctx, user.TelegramId, totalClicks); err != nil {
			return err
		}

		// FIXME: implicit int32 <==> int64 convertion.
		if err := u.teamApi.UpdateTotalClicksByTeamId(txCtx, int(user.TeamId), totalClicks); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
