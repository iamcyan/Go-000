package Service

import "train/Week02/DB"

func UserInfo(id int64) (*DB.User, error) {
	return DB.FindUser(id)
}
