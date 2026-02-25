package ldap_client

type UserInfo struct {
	Mails    []string // userPrincipalName, mail, division, url, extensionAttribute13
	Group    string   // otherHomePhone, homePhone
	FullName string   // displayName or sn + givenName + middleName
}
