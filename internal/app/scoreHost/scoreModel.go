package scoreHost

import "clicker/internal/app/dbhost"

type ScoreHost struct {
	dbHostApi dbhost.TeamQueryable
}

func Instance(dbhApi dbhost.TeamQueryable) *ScoreHost {
	return &ScoreHost{
		dbHostApi: dbhApi,
	}
}

// Other ScoreHost API implemented in ...Impl.go
