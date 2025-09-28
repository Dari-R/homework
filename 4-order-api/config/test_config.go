package configs

import (
	"os"
)

type TestConfig struct {
	Db   DbConfig
	Auth AuthConfig
}

func LoadTestConfig() *TestConfig {
	// Используем отдельную тестовую базу
	testDSN := os.Getenv("TEST_DSN")
	if testDSN == "" {
		testDSN = "host=localhost user=postgres password=postgres dbname=order_api_test port=5432 sslmode=disable"
	}

	return &TestConfig{
		Db: DbConfig{
			Dsn: testDSN,
		},
		Auth: AuthConfig{
			Secret: "test_secret_key",
		},
	}
}