package godap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGodapAuth_HappyPath(t *testing.T) {
	client := New(Options{
		BaseDN:       "dc=example,dc=com",
		Filter:       "uid",
		BasePassword: "password",
		BaseUser:     "cn=read-only-admin,dc=example,dc=com",
		URL:          "ldap.forumsys.com:389",
	})
	assert.Nil(t, client.Authenticate("tesla", "password"),
		"should authenticate successfully")
}

func TestGodapAuth_BindFail(t *testing.T) {
	client := New(Options{
		BaseDN:       "dc=example,dc=com",
		Filter:       "uid",
		BasePassword: "badpassword",
		BaseUser:     "cn=read-only-admin,dc=example,dc=com",
		URL:          "ldap.forumsys.com:389",
	})
	err := client.Authenticate("tesla", "password")
	assert.Equal(t, "godap: could not perform initial bind", err.Error(),
		"should fail initial bind")
}

func TestGodapAuth_EmptyPassword(t *testing.T) {
	client := New(Options{
		BaseDN:       "dc=example,dc=com",
		Filter:       "uid",
		BasePassword: "password",
		BaseUser:     "cn=read-only-admin,dc=example,dc=com",
		URL:          "ldap.forumsys.com:389",
	})
	err := client.Authenticate("tesla", "")
	assert.Equal(t, "godap: empty password", err.Error(),
		"should fail because of empty password")
}
func TestGodapAuth_EmptyUsername(t *testing.T) {
	client := New(Options{
		BaseDN:       "dc=example,dc=com",
		Filter:       "uid",
		BasePassword: "password",
		BaseUser:     "cn=read-only-admin,dc=example,dc=com",
		URL:          "ldap.forumsys.com:389",
	})
	err := client.Authenticate("", "password")
	assert.Equal(t, "godap: empty username", err.Error(),
		"should fail because of empty username")
}

func TestGodapAuth_FailSearch(t *testing.T) {
	client := New(Options{
		BaseDN:       "dc=example,dc=com",
		Filter:       "(",
		BasePassword: "password",
		BaseUser:     "cn=read-only-admin,dc=example,dc=com",
		URL:          "ldap.forumsys.com:389",
	})
	err := client.Authenticate("tesla", "password")
	assert.Equal(t, "godap: error performing search", err.Error(),
		"should fail to perform search")
}

func TestGodapAuth_NotFound(t *testing.T) {
	client := New(Options{
		BaseDN:       "dc=example,dc=com",
		Filter:       "uid",
		BasePassword: "password",
		BaseUser:     "cn=read-only-admin,dc=example,dc=com",
		URL:          "ldap.forumsys.com:389",
	})
	err := client.Authenticate("daddy", "password")
	assert.Equal(t, "godap: user not found in directory", err.Error(),
		"should fail to find user")
}

func TestGodapAuth_NoAuth(t *testing.T) {
	client := New(Options{
		BaseDN:       "dc=example,dc=com",
		Filter:       "uid",
		BasePassword: "password",
		BaseUser:     "cn=read-only-admin,dc=example,dc=com",
		URL:          "ldap.forumsys.com:389",
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
		BaseDN:       "dc=example,dc=com",
		Filter:       "uid",
		BasePassword: "password",
		BaseUser:     "cn=read-only-admin,dc=example,dc=com",
		URL:          "ldap.isnotarealthing.com:389",
	})
}

func TestGodapGetEntry(t *testing.T) {
	client := New(Options{
		BaseDN:       "dc=example,dc=com",
		Filter:       "uid",
		BasePassword: "password",
		BaseUser:     "cn=read-only-admin,dc=example,dc=com",
		URL:          "ldap.forumsys.com:389",
	})
	_, err := client.GetUserEntry("tesla")
	assert.Nil(t, err, "should get entry")
	_, err = client.GetUserEntry("daddy")
	assert.Equal(t, "godap: user not found in directory", err.Error(),
		"should fail to find user")
	_, err = client.GetUserEntry("")
	assert.Equal(t, "godap: empty username", err.Error(),
		"should fail because of empty username")
	client = New(Options{
		BaseDN:       "dc=example,dc=com",
		Filter:       "(",
		BasePassword: "password",
		BaseUser:     "cn=read-only-admin,dc=example,dc=com",
		URL:          "ldap.forumsys.com:389",
	})
	_, err = client.GetUserEntry("tesla")
	assert.Equal(t, "godap: error performing search", err.Error(),
		"should fail to perform search")
	client = New(Options{
		BaseDN:       "dc=example,dc=com",
		Filter:       "uid",
		BasePassword: "badpassword",
		BaseUser:     "cn=read-only-admin,dc=example,dc=com",
		URL:          "ldap.forumsys.com:389",
	})
	_, err = client.GetUserEntry("tesla")
	assert.Equal(t, "godap: could not perform initial bind", err.Error(),
		"should fail initial bind")
}
