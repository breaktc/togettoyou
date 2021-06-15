/**
 * @Author yumy
 * @Email :yumychn@163.com
 * @Date 下午4:16 2021/5/25
 * @Description :
 **/

package controller

import "github.com/stumanageapp.com/education/service"

type Application struct {
	Setup *service.ServiceSetup
}

type User struct {
	LoginName	string
	Password	string
	IsAdmin	string
}


var users []User

func init() {

	admin := User{LoginName:"Admin", Password:"123456", IsAdmin:"T"}
	alice := User{LoginName:"ChainDesk", Password:"123456", IsAdmin:"T"}
	bob := User{LoginName:"alice", Password:"123456", IsAdmin:"F"}
	jack := User{LoginName:"bob", Password:"123456", IsAdmin:"F"}

	users = append(users, admin)
	users = append(users, alice)
	users = append(users, bob)
	users = append(users, jack)

}

func isAdmin(cuser User) bool {
	if cuser.IsAdmin == "T"{
		return true
	}
	return false
}