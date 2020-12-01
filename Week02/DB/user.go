package DB

import "github.com/pkg/errors"
type User struct {
	Id int64
	Name string
	Age int64
}

func FindUser(id int64) (*User, error) {
	u := new(User)
	if u.Id <= 0 {
		return u, errors.New("do not find user")
	}
	return u, nil
}
