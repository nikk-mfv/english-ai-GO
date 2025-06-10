package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	mysqlDriver "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDatabase() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("missing DB_DSN")
	}
	pem := os.Getenv("DB_CA_CERT")
	if pem == "" {
		log.Fatal("missing DB_CA_CERT")
	}

	// Build x509 pool và đăng ký TLS config
	rootCertPool := x509.NewCertPool()
	if !rootCertPool.AppendCertsFromPEM([]byte(pem)) {
		log.Fatal("failed to append CA cert")
	}
	mysqlDriver.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: rootCertPool,
	})

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
