package database

func CheckIfConfigExist() bool {
	var howMuchConfigs int64
	DBConn.Table("configs").Count(&howMuchConfigs)
	return howMuchConfigs > 0
}
