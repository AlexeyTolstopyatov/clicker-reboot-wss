package models

type UserModel struct {
	TelegramId      int64
	TeamId          int64
	TotalClicks     int64
	BatteryCharge   int
	BatteryCapacity int
	BatteryRecharge int
}

type CreateUserModel struct {
	TelegramId int64
	TeamId     int64
}

type RegistryUserModel struct {
	TelegramId int64
	TeamName   string
}

type LoginUserModel struct {
	TelegramId int64
}
