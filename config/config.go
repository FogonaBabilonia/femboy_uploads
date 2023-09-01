package config

import "os"

var Secret_key string

func Load_key() {
	Secret_key = os.Getenv("JWT_SECRET_KEY")

	if Secret_key == "" {
		panic("could not get jwt secret key")
	}
}
