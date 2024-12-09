package dbhost

import (
	"clicker/internal/def/models"
	"clicker/internal/util"
	"context"
)

// InsertUser
// Inserts new model's statistics
// in database.
func (c *Client) InsertUser(ctx context.Context, user models.CreateUserModel) (models.UserModel, error) {
	q := `INSERT INTO users ([telegram_id], [team_id]) VALUES($1, $2) 
			RETURNING [telegram_id], [team_id], [total_clicks], [battery_capacity], [battery_charge], [battery_recharge_rate]`

	newUser := models.UserModel{}

	if err := c.Pool.SelectRow(
		ctx,
		q,
		user.TelegramId,
		user.TeamId).Scan(
		&newUser.TelegramId,
		&newUser.TeamId,
		&newUser.TotalClicks,
		&newUser.BatteryCapacity,
		&newUser.BatteryCharge,
		&newUser.BatteryRecharge,
	); err != nil {
		if err := util.SendPostgresError(err); err != nil {
			return newUser, err
		}
		return newUser, err
	}

	return newUser, nil
}

// SelectUserByTelegramId
// Runs SELECT query by known Telegram ID
// Returns models.UserModel
func (c *Client) SelectUserByTelegramId(ctx context.Context, tgId int64) (models.UserModel, error) {
	q := `SELECT [telegram_id], [team_id], [total_clicks], [battery_capacity], [battery_charge], [battery_recharge_rate] 
			FROM [users] WHERE [telegram_id] = $1`

	user := models.UserModel{}

	if err := c.Pool.SelectRow(ctx, q, tgId).Scan(
		&user.TelegramId,
		&user.TeamId,
		&user.TotalClicks,
		&user.BatteryCapacity,
		&user.BatteryCharge,
		&user.BatteryRecharge,
	); err != nil {
		if err := util.SendPostgresError(err); err != nil {
			// Postgre Error
			return user, err
		}
		return user, err
	}

	return user, nil
}

// UpdateTotalClicksByTelegramId
// Runs UPDATE of All clicks registered by server into Database
// Updates previous value
func (c *Client) UpdateTotalClicksByTelegramId(ctx context.Context, tgId int64, totalClicks int64) error {
	q := `UPDATE [users] SET [total_clicks] = $1 WHERE [telegram_id] = $2`

	if _, err := c.Pool.Execute(ctx, q, totalClicks, tgId); err != nil {
		if err := util.SendPostgresError(err); err != nil {
			// Postgre Error
			return err
		}
		// logger call
		return err
	}

	return nil
}

func (c *Client) SelectTeamByName(ctx context.Context, name string) (models.Team, error) {
	q := `SELECT [id], [name], [total_clicks] FROM teams WHERE [name] = $1`

	team := models.Team{}

	if err := c.Pool.SelectRow(ctx, q, name).Scan(
		&team.Id,
		&team.Name,
		&team.TotalClicks,
	); err != nil {
		if err := util.SendPostgresError(err); err != nil {
			return team, err
		}

		// query error
		return team, err
	}

	return team, nil
}

func (c *Client) UpdateTotalClicksByTeamId(ctx context.Context, teamId int, totalClicks int64) error {
	q := `UPDATE [teams] SET [total_clicks] = $1 WHERE [id] = $2`

	if _, err := c.Pool.Execute(ctx, q, totalClicks, teamId); err != nil {
		if err := util.SendPostgresError(err); err != nil {
			return err
		}
		// call error
		return err
	}

	return nil
}
