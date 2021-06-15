/**
 * @Author yumy
 * @Email :yumychn@163.com
 * @Date 下午4:03 2021/5/25
 * @Description :
 **/
package main

/**
姓名：张小三，性别：男，

民族：汉，身份证号：1011010101010101

籍贯：XXX，出生日期：1999年01月01日，			照片，

入学日期：2017年9月，毕（结）业日期：2021年7月，

学校名称：中北大学，专业：计算机，

学历类别：普通，学制：四年，

学习形式：普通全日制，层次：本科，

毕（结）业：毕业，证书编号：11111111111111

 */
type Education struct {
	ObjectType	string	`json:"docType"`
	Name	string	`json:"Name"`		// 姓名
	Gender	string	`json:"Gender"`		// 性别
	Nation	string	`json:"Nation"`		// 民族
	EntityID	string	`json:"EntityID"`		// 身份证号
	Place	string	`json:"Place"`		// 籍贯
	BirthDay	string	`json:"BirthDay"`		// 出生日期

	EnrollDate	string	`json:"EnrollDate"`		// 入学日期
	GraduationDate	string	`json:"GraduationDate"`	// 毕（结）业日期
	SchoolName	string	`json:"SchoolName"`	// 学校名称
	Major	string	`json:"Major"`	// 专业
	QuaType	string	`json:"QuaType"`	// 学历类别
	Length	string	`json:"Length"`	// 学制
	Mode	string	`json:"Mode"`	// 学习形式
	Level	string	`json:"Level"`	// 层次
	Graduation	string	`json:"Graduation"`	// 毕（结）业
	CertNo	string	`json:"CertNo"`	// 证书编号

	Photo	string	`json:"Photo"`	// 照片

	Historys	[]HistoryItem	// 当前edu的历史记录
}

type HistoryItem struct {
	TxId	string
	Education	Education
}
