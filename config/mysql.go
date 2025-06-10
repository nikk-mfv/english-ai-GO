package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDatabase() {
	// Load .env if available
	_ = godotenv.Load()

	caCert := os.Getenv("CA_CERT")
	if caCert == "" {
		log.Fatal("missing CA_CERT env")
	}
	caPath := "aiven-ca.pem"
	err := os.WriteFile(caPath, []byte(caCert), 0644)
	if err != nil {
		log.Fatal("failed to write CA cert file:", err)
	}

	// Load CA cert
	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile(caPath)
	if err != nil {
		log.Fatal(err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}

	// Register custom TLS config
	err = mysqlDriver.RegisterTLSConfig("aiven", &tls.Config{
		RootCAs: rootCertPool,
	})
	if err != nil {
		log.Fatal("Failed to register TLS config: ", err)
	}

	dsn := os.Getenv("DB_DSN") + "&tls=aiven"
	if dsn == "" {
		log.Fatal("missing database connection string")
	}

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}

	db = database
	fmt.Println("Connected to database!")
}

func GetDatabase() *gorm.DB {
	if db == nil {
		log.Fatal("Database not initialized, call ConnectDatabase() first")
	}
	return db
}
