package main

import (
	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/docker/docker-credential-helpers/envvarauth"
)

func main() {
	credentials.Serve(envvarauth.Envvarauth{})
}
