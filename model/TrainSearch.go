package model

type TrainSearch struct {
	Id        int64 `form:"id" json:"id" xorm:"pk autoincr "`
	TrainNo string `form:"trainNo" json:"trainNo" xorm:"UNIQUE NOT NULL"`
	Date     string `form:"date" json:"date" xorm:" NOT NULL"`
	StartStation  string `form:"startStation" json:"startStation" xorm:" NOT NULL"`
	EndStation       string `form:"endStation" json:"endStation" xorm:"NOT NULL"`
}

func (c *TrainSearch) Get() (*TrainSearch, bool) {
	has, err := orm.Get(c)
	if err != nil {
		return c, false
	}
	return c, has
}

func (c *TrainSearch) GetList(start int, num int) ([]*TrainSearch, error) {
	var posts = make([]*TrainSearch, 0)
	start = (start - 1) * num
	err := orm.Limit(num, start).Find(&posts)
	return posts, err
}

func (c *TrainSearch) All() ([]*TrainSearch, error) {
	var posts = make([]*TrainSearch, 0)
	err := orm.Find(&posts)
	return posts, err
}

func (c *TrainSearch) Insert() (int64, error) {
	affected, err := orm.Insert(c)
	return affected, err
}

func (c *TrainSearch) Update() (int64, error) {
	affected, err := orm.Id(c.Id).Update(c)
	return affected, err
}

func (c *TrainSearch) Count() (int64, error) {
	affected, err := orm.Count(c)
	return affected, err
}

func (c *TrainSearch) Delete() (int64, error) {
	affected, err := orm.Id(c.Id).Delete(c)
	return affected, err
}