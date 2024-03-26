package config

const (
	// The port on which the application is running
	ServerPort = ":1323"
	// Address and parameters for accessing the database
	DatabaseUrl = "host=localhost user=vk_admin password=password dbname=appdb port=5432 sslmode=disable"
	// The secret key for JWT encryption
	SecretKeyJWT = "RandomX100][weqwr]"
	// The size of one page in the ad feed
	PageSize = 24
)
