package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

func (*UserInfoAtomicService) Save(u *UserInfo) error {
	session := eng.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Insert(u) //插入一条记录
	if err != nil {
		session.Rollback()
		return err
	}
	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (*UserInfoAtomicService) FindAll() []UserInfo {
	allr := make([]UserInfo, 0)
	eng.Find(&allr) //从数据库中查询多条记录
	return allr
}

func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	var u UserInfo
	eng.Id(id).Get(&u)
	return &u
}
