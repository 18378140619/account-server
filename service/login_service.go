package service

import (
	"account-server/dao"
)

var loginSerive *LoginService

type LoginService struct {
	BaseService
	Dao *dao.LoginDao
}

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "wx64c12a1645d44164"
	appSecret       = "f87abde966e34da5c906dfeaef0ba859"
)

func NewLoginService() *LoginService {
	if loginSerive == nil {
		loginSerive = &LoginService{
			Dao: dao.NewLoginDao(),
		}
	}
	return loginSerive
}

func GetOpenidByCode() string {
	//url := fmt.Sprintf(code2sessionURL, appID, appSecret, code)
	//resp, err := http.DefaultClient.Get(url)
	//if err != nil {
	//	return "", err
	//}
	//var wxMap map[string]string
	//err = json.NewDecoder(resp.Body).Decode(&wxMap)
	//if err != nil {
	//	return "", err
	//}
	//defer resp.Body.Close()
	//
	//return wxMap["openid"], nil
	return "123"
}
