package main

type Config struct {
	ServerAddr    string `env:"SERVER_ADDR"`     // ex: "0.0.0.0:8080"
	DbConnString  string `env:"DB_CONN_STRING"`  // ex: "postgres://myusername:mypassword@localhost:5432/mydb?sslmode=disable"
	ThezosApiAddr string `env:"THEZOS_API_ADDR"` // ex "https://api.tzkt.io"
}
