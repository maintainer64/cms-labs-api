package ldap_client

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

type LDAPClient struct {
	URL string
	DN  string
}

func (c *LDAPClient) GetInfo(username, password string) (*UserInfo, error) {
	// Connect to LDAP over TLS
	l, err := ldap.DialURL(c.URL, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true})) // #nosec G402
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}
	defer l.Close()

	// Bind with the user's credentials (using UPN, e.g., user@domain)
	// The example used -D "Vadim.Asatulin@urfu.me", so we use the provided username as is.
	err = l.Bind(username, password)
	if err != nil {
		return nil, fmt.Errorf("bind failed: %w", err)
	}

	// Determine the sAMAccountName from the username.
	// If the username is a UPN (contains '@'), we extract the part before '@'.
	// Otherwise, assume it's already the sAMAccountName.
	samAccountName := username
	if idx := strings.Index(username, "@"); idx != -1 {
		samAccountName = username[:idx]
	}

	// Search for the user's entry.
	// Base DN for the search (from ldapsearch example)
	// Filter by sAMAccountName (as in the example)
	filter := fmt.Sprintf("(sAMAccountName=%s)", samAccountName)
	// Attributes we want to retrieve
	attributes := []string{
		"userPrincipalName",
		"mail",
		"division",
		"url",
		"extensionAttribute13",
		"otherHomePhone",
		"homePhone",
		"displayName",
		"sn",
		"givenName",
		"middleName",
	}

	searchReq := ldap.NewSearchRequest(
		c.DN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		attributes,
		nil,
	)

	sr, err := l.Search(searchReq)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	if len(sr.Entries) == 0 {
		return nil, fmt.Errorf("user not found")
	}
	if len(sr.Entries) > 1 {
		return nil, fmt.Errorf("multiple entries found for user")
	}

	entry := sr.Entries[0]

	// Extract the required fields.
	info := &UserInfo{
		Mails: collectMailAttributes(entry),
		Group: collectGroupAttributes(entry),
	}

	// Build full name: try displayName first, else combine sn, givenName, middleName.
	if displayName := entry.GetAttributeValue("displayName"); displayName != "" {
		info.FullName = displayName
	} else {
		parts := []string{
			entry.GetAttributeValue("sn"),
			entry.GetAttributeValue("givenName"),
			entry.GetAttributeValue("middleName"),
		}
		var nonEmpty []string
		for _, p := range parts {
			if p != "" {
				nonEmpty = append(nonEmpty, p)
			}
		}
		info.FullName = strings.Join(nonEmpty, " ")
	}

	return info, nil
}

// collectMailAttributes gathers all email-like attributes into a slice.
func collectMailAttributes(entry *ldap.Entry) []string {
	var mails []string
	// Attributes that may contain email addresses
	attrs := []string{"extensionAttribute13", "mail", "division", "url", "userPrincipalName"}
	for _, attr := range attrs {
		if val := entry.GetAttributeValue(attr); val != "" {
			mails = append(mails, val)
		}
	}
	return mails
}

// collectGroupAttributes gathers group-like attributes (otherHomePhone, homePhone).
func collectGroupAttributes(entry *ldap.Entry) string {
	if val := entry.GetAttributeValue("otherHomePhone"); val != "" {
		return val
	}
	if val := entry.GetAttributeValue("homePhone"); val != "" {
		return val
	}
	return ""
}
