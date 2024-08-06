package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/thechotinun/authentication/pkg/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlConnection func for connection to Mysql database.
func MysqlConnection() (*gorm.DB, error) {
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Build Mysql connection URL.
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}

	// Define database connection for PostgreSQL.
	db, err := gorm.Open(mysql.Open(mysqlConnURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting generic database object: %w", err)
	}

	// Set database connection settings:
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	// Try to ping database.
	if sqlDB, err := db.DB(); err != nil {
		defer sqlDB.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	} else if err := sqlDB.Ping(); err != nil {
		defer sqlDB.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
