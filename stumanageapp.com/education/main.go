/**
 * @Author yumy
 * @Email :yumychn@163.com
 * @Date 下午4:16 2021/5/25
 * @Description :
 **/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/stumanageapp.com/education/sdkInit"
	"github.com/stumanageapp.com/education/service"
	"github.com/stumanageapp.com/education/web"
	"github.com/stumanageapp.com/education/web/controller"
	"os"
)

const (
	configFile  = "config.yaml"
	initialized = false
	SimpleCC    = "educc"
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID:     "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/stumanageapp.com/education/fixtures/artifacts/channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID:     SimpleCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/stumanageapp.com/education/chaincode/",
		UserName:        "User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	serviceSetup := service.ServiceSetup{
		ChaincodeID: SimpleCC,
		Client:      channelClient,
	}

	edu := service.Education{
		Name:           "张三",
		Gender:         "男",
		Nation:         "汉",
		EntityID:       "101",
		Place:          "北京",
		BirthDay:       "1999年01月01日",
		EnrollDate:     "2017年9月",
		GraduationDate: "2021年7月",
		SchoolName:     "中北大学",
		Major:          "计算机",
		QuaType:        "普通",
		Length:         "四年",
		Mode:           "普通全日制",
		Level:          "本科",
		Graduation:     "毕业",
		CertNo:         "111",
		Photo:          "/static/phone/11.png",
	}

	edu2 := service.Education{
		Name:           "李四",
		Gender:         "男",
		Nation:         "汉",
		EntityID:       "102",
		Place:          "上海",
		BirthDay:       "1999年02月01日",
		EnrollDate:     "2017年9月",
		GraduationDate: "2021年7月",
		SchoolName:     "中北大学",
		Major:          "物联网",
		QuaType:        "普通",
		Length:         "四年",
		Mode:           "普通全日制",
		Level:          "本科",
		Graduation:     "毕业",
		CertNo:         "222",
		Photo:          "/static/phone/22.png",
	}

	msg, err := serviceSetup.SaveEdu(edu)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}

	msg, err = serviceSetup.SaveEdu(edu2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}

	// 根据证书编号与名称查询信息
	result, err := serviceSetup.FindEduByCertNoAndName("222", "李四")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		fmt.Println("根据证书编号与姓名查询信息成功：")
		fmt.Println(edu)
	}

	//===========================================//

	// 根据身份证号码查询信息
	result, err = serviceSetup.FindEduInfoByEntityID("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		fmt.Println("根据身份证号码查询信息成功：")
		fmt.Println(edu)
	}

	//===========================================//

	// 修改/添加信息
	info := service.Education{
		Name: "张三",
		Gender: "男",
		Nation: "汉",
		EntityID: "101",
		Place: "北京",
		BirthDay: "1999年01月01日",
		EnrollDate: "2017年9月",
		GraduationDate: "2019年7月",
		SchoolName: "中北大学",
		Major: "软件工程",
		QuaType: "普通",
		Length: "两年",
		Mode: "普通全日制",
		Level: "研究生",
		Graduation: "毕业",
		CertNo: "333",
		Photo: "/static/phone/11.png",
	}
	msg, err = serviceSetup.ModifyEdu(info)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息操作成功, 交易编号为: " + msg)
	}

	//===========================================//
	// 根据证书编号与名称查询信息
	result, err = serviceSetup.FindEduByCertNoAndName("333","张三")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		fmt.Println("根据证书编号与姓名查询信息成功：")
		fmt.Println(edu)
	}

	//===========================================//

	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)
}
