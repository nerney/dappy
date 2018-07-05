package godap

import (
	"fmt"

	"gopkg.in/ldap.v2"
)

//Client interface for performing LDAP authentication
type Client interface {
	Authenticate(username, password string) error
}

//Options required for GoDAP client
type Options struct {
	BaseDN   string
	Filter   string
	Password string
	Username string
	URL      string
}

type client struct {
	options Options
	conn    connection
}

//New Godap client with options
func New(options Options) Client {
	return client{
		options: options,
		conn:    connect(options.URL),
	}
}

func (client client) Authenticate(username, password string) error {
	if len(password) < 1 {
		return errEmptyPassword
	}
	if len(username) < 1 {
		return errEmptyUsername
	}
	client.conn = connect(client.options.URL)
	if client.conn.Bind(client.options.Username, client.options.Password) != nil {
		return errCouldNotBind
	}
	searchRequest := ldap.NewSearchRequest(
		client.options.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0, false,
		fmt.Sprintf("(%v=%v)", client.options.Filter, username),
		[]string{}, nil,
	)
	searchResult, err := client.conn.Search(searchRequest)
	if err != nil {
		return errSearch
	}
	if len(searchResult.Entries) < 1 {
		return errNotFound
	}
	user := searchResult.Entries[0]
	if client.conn.Bind(user.DN, password) != nil {
		return errCouldNotAuth
	}
	return nil
}
