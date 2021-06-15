/**
 * @Author yumy
 * @Email :yumychn@163.com
 * @Date 下午4:10 2021/5/25
 * @Description :
 **/

package sdkInit

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)

type InitInfo struct {
	ChannelID     string
	ChannelConfig string
	OrgAdmin      string
	OrgName       string
	OrdererOrgName	string
	OrgResMgmt *resmgmt.Client

	ChaincodeID	string
	ChaincodeGoPath	string
	ChaincodePath	string
	UserName	string
}
