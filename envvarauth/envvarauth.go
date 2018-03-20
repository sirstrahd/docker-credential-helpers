package envvarauth

/*
#cgo CFLAGS: -x objective-c -mmacosx-version-min=10.10
#cgo LDFLAGS: -framework Security -framework Foundation -mmacosx-version-min=10.10

#include "envvarauth.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"os"
	"strings"

	"github.com/docker/docker-credential-helpers/credentials"
)

// errCredentialsNotFound is the specific error message returned by OS X
// when the credentials are not in the keychain.
const errCredentialsNotFound = "The specified item could not be found in the keychain."

// Envvarauth gets credentials from an environment variable
type Envvarauth struct{}

// Add adds new credentials to the keychain.
func (h Envvarauth) Add(creds *credentials.Credentials) error {
	return nil
}

// Delete removes credentials from the keychain.
func (h Envvarauth) Delete(serverURL string) error {
	return nil
}

// Get returns the username and secret to use for a given registry server URL.
func (h Envvarauth) Get(serverURL string) (string, string, error) {
	if serverURL == "" {
		return "", "", errors.New("missing server url")
	}
	credential := os.Getenv(strings.Replace(serverURL, ".", "_", -1) + "_CREDENTIALS")
	if credential == "" {
		return "", "", errors.New("Could not read environment variable for server")
	}
	userPassArray := strings.Split(credential, ":")
	user := userPassArray[0]
	pass := userPassArray[1]
	return user, pass, nil
}

// List returns the stored URLs and corresponding usernames.
func (h Envvarauth) List() (map[string]string, error) {
	resp := make(map[string]string)
	return resp, nil
}
