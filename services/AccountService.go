package services

import (
	"errors"

	"github.com/astaxie/beego/logs"
	"github.com/myxzjie/able-admin/models"
	"github.com/myxzjie/able-admin/utils"
)

type AccountService struct {
	BaseService
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (m *AccountService) FindAccount(name string) (res *models.Account, err error) {
	res = &models.Account{}
	err = getOrm().QueryTable(new(models.Account)).Filter("name", name).One(res)
	return
}

func (m *AccountService) AccountLogin(name string, password string) (res *models.Account, e error) {

	data, err := m.FindAccount(name)
	// logs.Debug(">>", data)
	if err != nil {
		logs.Error(">>", err)
		// e = errors.New("data error")
		return nil, err
	}
	// p, s := passwordEncryption(name, password)
	// logs.Info(">>p,s:", p, s)
	if (verify(data.Name, data.Salt, data.Password, password)) == false {
		return nil, errors.New("password error")
	}
	// res = data
	// return
	return data, nil
}

func passwordEncryption(name string, password string) (pass string, salt string) {
	salt = utils.SecureRandomGenerator()
	nameSalt := name + salt
	pass = utils.SimpleHash(utils.AlgorithmName, password, nameSalt, utils.HashIterations)
	return
}

func verify(name string, salt string, password string, sourcePass string) bool {
	nameSalt := name + salt
	// logs.Debug(">>salt:", salt, pass)
	npass := utils.SimpleHash(utils.AlgorithmName, sourcePass, nameSalt, utils.HashIterations)
	logs.Debug(">> verify password:", sourcePass, password)

	if npass == password {
		return true
	}
	return false
}

func (m *AccountService) GetDataList() (res []*models.Account, err error) {
	logs.Debug(">>table name", TableName(new(models.Account)))
	_, err = getOrm().QueryTable(new(models.Account)).All(&res)
	return
}

//分页查询指定用户的项目
//按照最新的进行排序
func (m *AccountService) GetDataListPage(page *models.Page, PrivatelyOwned ...int) (res []*models.Account, total int, err error) {

	// relationship := NewRelationship()
	sql := "SELECT t.* FROM " + TableName(new(models.Account)) + " t WHERE 1=1 "

	// sql1 := "SELECT COUNT(book.book_id) AS total_count FROM " + m.TableNameWithPrefix() + " AS book LEFT JOIN " +
	// 	relationship.TableNameWithPrefix() + " AS rel ON book.book_id=rel.book_id AND rel.member_id = ? WHERE rel.relationship_id > 0 "
	// if len(PrivatelyOwned) > 0 {
	// 	sql1 = sql1 + " and book.privately_owned=" + strconv.Itoa(PrivatelyOwned[0])
	// }
	querySQL, conutSQL := getPageSQLString(sql, page)
	err = getOrm().Raw(conutSQL).QueryRow(&total)
	// if err != nil {
	// 	return
	// }

	// offset := (page.CurrentPage - 1) * page.ShowCount
	// querySQL := "SELECT t.* FROM " + TableName(new(models.Account)) + " t WHERE 1=1 LIMIT " + fmt.Sprintf("%d,%d", offset, pageSize)
	// sql2 := "SELECT t.*,rel.member_id,rel.role_id,m.account as create_name FROM " + m.TableNameWithPrefix() + " AS book" +
	// 	" LEFT JOIN " + relationship.TableNameWithPrefix() + " AS rel ON book.book_id=rel.book_id AND rel.member_id = ? AND rel.role_id=0" +
	// 	" LEFT JOIN " + NewMember().TableNameWithPrefix() + " AS m ON rel.member_id=m.member_id " +
	// 	" WHERE rel.relationship_id > 0 %v ORDER BY book.book_id DESC LIMIT " + fmt.Sprintf("%d,%d", offset, pageSize)

	// if len(PrivatelyOwned) > 0 {
	// 	sql2 = fmt.Sprintf(sql2, " and book.privately_owned="+strconv.Itoa(PrivatelyOwned[0]))
	// }
	_, err = getOrm().Raw(querySQL).QueryRows(&res)
	if err != nil {
		logs.Error("分页查询项目列表 => ", err)
		return
	}

	// if err == nil && len(books) > 0 {
	// 	sql := "SELECT m.account,doc.modify_time FROM md_documents AS doc LEFT JOIN md_members AS m ON doc.modify_at=m.member_id WHERE book_id = ? ORDER BY doc.modify_time DESC LIMIT 1 "

	// 	for index, book := range books {
	// 		var text struct {
	// 			Account    string
	// 			ModifyTime time.Time
	// 		}

	// 		err = o.Raw(sql, book.BookId).QueryRow(&text)
	// 		if err == nil {
	// 			books[index].LastModifyText = text.Account + " 于 " + text.ModifyTime.Format("2006-01-02 15:04:05")
	// 		}

	// 		if book.RoleId == 0 {
	// 			book.RoleName = "创始人"
	// 		} else if book.RoleId == 1 {
	// 			book.RoleName = "管理员"
	// 		} else if book.RoleId == 2 {
	// 			book.RoleName = "编辑者"
	// 		} else if book.RoleId == 3 {
	// 			book.RoleName = "观察者"
	// 		}
	// 	}
	// }
	return
}
