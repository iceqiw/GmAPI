package model

type User struct {
	Id        int64
	Username string `xorm:"UNIQUE NOT NULL"`
	Password      string `xorm:"UNIQUE NOT NULL"`
	Status  string
	// Email is the primary email address (to be used for communication)
	Name       string `xorm:"NOT NULL"`
}

func (u *User) Get() (*User, bool) {
	has, err := orm.Get(u)
	if err != nil {
		return u, false
	}
	return u, has
}

func (u *User) GetList(start int, num int) ([]*User, error) {
	var users = make([]*User, 0)
	start = (start - 1) * num
	err := orm.Limit(num, start).Find(&users)
	return users, err
}