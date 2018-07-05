package godap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGodap_HappyPath(t *testing.T) {
	client := New(Options{
		BaseDN:   "dc=example,dc=com",
		Filter:   "uid",
		Password: "password",
		Username: "cn=read-only-admin,dc=example,dc=com",
		URL:      "ldap.forumsys.com:389",
	})
	assert.Nil(t, client.Authenticate("tesla", "password"),
		"should authenticate successfully")
}

func TestGodap_BindFail(t *testing.T) {
	client := New(Options{
		BaseDN:   "dc=example,dc=com",
		Filter:   "uid",
		Password: "badpassword",
		Username: "cn=read-only-admin,dc=example,dc=com",
		URL:      "ldap.forumsys.com:389",
	})
	err := client.Authenticate("tesla", "password")
	assert.Equal(t, "godap: could not perform initial bind", err.Error(),
		"should fail initial bind")
}

func TestGodap_EmptyPassword(t *testing.T) {
	client := New(Options{
		BaseDN:   "dc=example,dc=com",
		Filter:   "uid",
		Password: "password",
		Username: "cn=read-only-admin,dc=example,dc=com",
		URL:      "ldap.forumsys.com:389",
	})
	err := client.Authenticate("tesla", "")
	assert.Equal(t, "godap: empty password", err.Error(),
		"should fail because of empty password")
}
func TestGodap_EmptyUsername(t *testing.T) {
	client := New(Options{
		BaseDN:   "dc=example,dc=com",
		Filter:   "uid",
		Password: "password",
		Username: "cn=read-only-admin,dc=example,dc=com",
		URL:      "ldap.forumsys.com:389",
	})
	err := client.Authenticate("", "password")
	assert.Equal(t, "godap: empty username", err.Error(),
		"should fail because of empty password")
}

func TestGodap_FailSearch(t *testing.T) {
	client := New(Options{
		BaseDN:   "dc=example,dc=com",
		Filter:   "(",
		Password: "password",
		Username: "cn=read-only-admin,dc=example,dc=com",
		URL:      "ldap.forumsys.com:389",
	})
	err := client.Authenticate("tesla", "password")
	assert.Equal(t, "godap: error performing search", err.Error(),
		"should fail to perform search")
}

func TestGodap_NotFound(t *testing.T) {
	client := New(Options{
		BaseDN:   "dc=example,dc=com",
		Filter:   "uid",
		Password: "password",
		Username: "cn=read-only-admin,dc=example,dc=com",
		URL:      "ldap.forumsys.com:389",
	})
	err := client.Authenticate("daddy", "password")
	assert.Equal(t, "godap: user not found in directory", err.Error(),
		"should fail to find user")
}

func TestGodap_NoAuth(t *testing.T) {
	client := New(Options{
		BaseDN:   "dc=example,dc=com",
		Filter:   "uid",
		Password: "password",
		Username: "cn=read-only-admin,dc=example,dc=com",
		URL:      "ldap.forumsys.com:389",
	})
	err := client.Authenticate("tesla", "wrongpassword")
	assert.Equal(t, "godap: could not authenticate user", err.Error(),
		"should fail to authenticate user")
}

func TestGodap_NoConnection(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			if err.Error() != "godap: could not establish connection" {
				panic(err)
			}
		} else {
			panic("should panic")
		}
	}()
	New(Options{
		BaseDN:   "dc=example,dc=com",
		Filter:   "uid",
		Password: "password",
		Username: "cn=read-only-admin,dc=example,dc=com",
		URL:      "ldap.isnotarealthing.com:389",
	})
}
