package dbhost

import (
	"clicker/internal/def/models"
	"clicker/internal/driver/db"
	"context"
)

//
// Here is containing the model of DataBase "models-side"
// modeling.
// Remember the structure example:
// Windows Terminal works because his 2 parts are connecting.
// ConDrv.sys -- Driver of terminals (represents itself API) or Server-side of project
// ConHost.exe -- Controls director (if I remember) Client-side of project. Connects with ConDrv.sys
// for enable of making terminals in every FrameworkContent-element what you want (ex. all IDEs has internal terminal)
//
// ./driver/db -- Driver package, which represents PostgreSQL driver's API
// ./app/dbhost -- Application's service to manipulate data inside Postgre Database.
//

type Client struct {
	Pool db.Pool
}

type UserQueryable interface {
	InsertUser(ctx context.Context, user models.CreateUserModel) (models.UserModel, error)
	SelectUserByTelegramId(ctx context.Context, tgId int64) (models.UserModel, error)
	UpdateTotalClicksByTelegramId(ctx context.Context, tgId int64, totalClicks int64) error
}

type TeamQueryable interface {
	UpdateTotalClicksByTeamId(ctx context.Context, teamId int, totalClicks int64) error
	SelectTeamByName(ctx context.Context, name string) (models.Team, error)
}

// Instance
// Creates DB-Host instance, provides database host's API
// for manipulate data
func Instance(pool db.Pool) *Client {
	return &Client{
		Pool: pool,
	}
}
