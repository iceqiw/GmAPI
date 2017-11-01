package model

type Student struct {
	Id        int64 `xorm:"pk autoincr "`
	Name string `xorm:"NOT NULL"`
	Num      int `xorm:"NOT NULL"`
	Sex  string `xorm:"NOT NULL"`
	Tel       string `xorm:"NOT NULL"`
	Adress string `xorm:"NOT NULL"`
}

func (c *Student) Get() (*Student, bool) {
	has, err := orm.Get(c)
	if err != nil {
		return c, false
	}
	return c, has
}

func (c *Student) GetList(start int, num int) ([]*Student, error) {
	var posts = make([]*Student, 0)
	start = (start - 1) * num
	err := orm.Limit(num, start).Find(&posts)
	return posts, err
}

func (c *Student) All() ([]*Student, error) {
	var posts = make([]*Student, 0)
	err := orm.Find(&posts)
	return posts, err
}

func (c *Student) Insert() (int64, error) {
	affected, err := orm.Insert(c)
	return affected, err
}

func (c *Student) Update() (int64, error) {
	affected, err := orm.Id(c.Id).Update(c)
	return affected, err
}

func (c *Student) Count() (int64, error) {
	affected, err := orm.Count(c)
	return affected, err
}

func (c *Student) Delete() (int64, error) {
	affected, err := orm.Id(c.Id).Delete(c)
	return affected, err
}