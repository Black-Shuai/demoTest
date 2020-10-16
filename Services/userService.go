package Services

import (
	"demoTest/Models"
	"demoTest/Services/ServiceModels"
)

func FindUser(user Models.User)(result ServiceModels.ResultResponse)  {
	//对controller接受到的数据进行匹配
	var user1 Models.User
	user1.UserName="18275507824"
	user1.Name="Test"
	user1.UserPassword="123456"
	user1.Id="1"
	result.Code=0
	result.Message="fail"
	if user.UserName==user1.UserName&&user.UserPassword==user1.UserPassword{
		result.Code=1
		result.Message="success"
		result.Data=user1
		return result
	}
	return
}
