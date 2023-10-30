package mapper

import (
	"NewThread/src/pojo"
)

type GroupMysql struct{}

func NewGroupMysql() *GroupMysql {
	return &GroupMysql{}
}

func (c *GroupMysql) GroupListMysql() ([]pojo.Group, error) {
	var m []pojo.Group
	if err := Db.Raw("select id,name,description from t_group").Scan(&m).Error; err != nil {
		return []pojo.Group{}, err
	}
	return m, nil
}

/*
查询 所有老师 和 老师负责的小组
返回 老师的用户ID、Name、Group、头像URL
*/
func (c *GroupMysql) GroupTeacherListAndGroupMysql() ([]pojo.Teacher, error) {
	var m []pojo.Teacher
	err := Db.Raw("SELECT t.user_id,t.`name`,g.name `group`,img.url FROM t_teacher t LEFT JOIN t_group g ON t.id = g.teacher_id LEFT JOIN t_imageUser img ON img.user_id = t.user_id ").Scan(&m).Error
	if err != nil {
		return []pojo.Teacher{}, err
	}
	return m, nil
}

/*
查询 所有老师 和 老师负责的小组
返回 老师的用户ID、Name、Group、头像URL
*/
func (c *GroupMysql) GroupStudentListAndWishesMysql() ([]pojo.Student, error) {
	var m []pojo.Student
	err := Db.Raw("SELECT " +
		"s.id," +
		"s.`name`," +
		"g.`name` `group`," +
		"s.user_id, " +
		"img.url " +
		"FROM " +
		"t_position p," +
		"t_student s " +
		"LEFT JOIN t_group g ON g.id = s.group_id " +
		"LEFT JOIN t_imageUser img ON img.user_id = s.user_id " +
		"WHERE " +
		"s.user_id = p.user_id " +
		"GROUP BY id").Scan(&m).Error
	if err != nil {
		return []pojo.Student{}, err
	}
	return m, nil
}

func (c *GroupMysql) PositionByUserIdMysql(userid int) ([]string, error) {
	var m []string
	if err := Db.Raw("SELECT name FROM t_position where user_id = ?", &userid).Scan(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}
