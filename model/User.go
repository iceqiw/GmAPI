package model

type User struct {
	Id        int64 `form:"id" json:"id" xorm:"pk autoincr "`
	Username string `form:"username" json:"username" xorm:"UNIQUE NOT NULL"`
	Password      string `form:"password" json:"password" xorm:" NOT NULL"`
	Status  string `form:"status" json:"status" xorm:" NOT NULL"`
	Name       string `form:"name" json:"name" xorm:"NOT NULL"`
}

func (c *User) Get() (*User, bool) {
	has, err := orm.Get(c)
	if err != nil {
		return c, false
	}
	return c, has
}

func (c *User) GetList(start int, num int) ([]*User, error) {
	var posts = make([]*User, 0)
	start = (start - 1) * num
	err := orm.Limit(num, start).Find(&posts)
	return posts, err
}

func (c *User) All() ([]*User, error) {
	var posts = make([]*User, 0)
	err := orm.Find(&posts)
	return posts, err
}

func (c *User) Insert() (int64, error) {
	affected, err := orm.Insert(c)
	return affected, err
}

func (c *User) Update() (int64, error) {
	affected, err := orm.Id(c.Id).Update(c)
	return affected, err
}

func (c *User) Count() (int64, error) {
	affected, err := orm.Count(c)
	return affected, err
}

func (c *User) Delete() (int64, error) {
	affected, err := orm.Id(c.Id).Delete(c)
	return affected, err
}