package main

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func main() {
	//// The username and password we want to check
	//username := "admin"
	//password := "admin"
	//
	//bindusername := "readonly"
	//bindpassword := "readonly"
	//
	//l, err := ldap.DialURL("ldap://127.0.0.1:389")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer l.Close()
	//
	//// Reconnect with TLS
	//err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// First bind with a read only user
	//err = l.Bind(bindusername, bindpassword)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Search for the given username
	//searchRequest := ldap.NewSearchRequest(
	//	"dc=example,dc=org",
	//	ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
	//	fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", username),
	//	[]string{"dn"},
	//	nil,
	//)
	//
	//sr, err := l.Search(searchRequest)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if len(sr.Entries) != 1 {
	//	log.Fatal("User does not exist or too many entries returned")
	//}
	//
	//userdn := sr.Entries[0].DN
	//
	//// Bind as the user to verify their password
	//err = l.Bind(userdn, password)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Rebind as the read only user for any further queries
	//err = l.Bind(bindusername, bindpassword)
	//if err != nil {
	//	log.Fatal(err)
	//}
	con, err := ldap.Dial("tcp", "127.0.0.1:3898")
	if err != nil {
		fmt.Println(err)
	}
	err = con.Bind("cn=admin,dc=softnix,dc=co,dc=th", "P1234@1234")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connection success")

	//mod, err := con.PasswordModify(&ldap.PasswordModifyRequest{
	//	UserIdentity: "cn=admin,dc=softnix,dc=co,dc=th",
	//	OldPassword:  "",
	//	NewPassword:  "P1234@1234",
	//})
	//if err != nil {
	//	//e, ok := err.(*ldap.Error)
	//	//if ok {
	//	//	fmt.Println(e.ResultCode == ldap.LDAPResultCodeMap[53])
	//	//}
	//	fmt.Println("Error", err)
	//}
	//fmt.Println(mod)
}
