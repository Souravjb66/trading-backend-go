package connect

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	// _ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// type SqlDbInstance struct {
// 	MySql *sql.DB
// 	PostgreSQL *sql.DB
// }
type PostgreSQLDbInstance struct{
	PostgreSQL *sql.DB


}
type MySqlRequired struct {
	Url      string
	UserName string
	Password string
	Port     string
	DbName   string
}
type PostgreSQLRequired struct {
	Url      string
	UserName string
	Password string
	Port     string
	DbName   string
}

// var MySqlInstance MySqlDbInstance

// func ConnectMysql(requiredInput MySqlRequired) *MySqlDbInstance {
// 	// "root:@tcp(127.0.0.1:3306)/mydb"
// 	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", requiredInput.UserName, requiredInput.Password, requiredInput.Url, requiredInput.Port, requiredInput.DbName)

// 	// DSN for migrate (notice mysql:// prefix)
// 	dsnMigrate := fmt.Sprintf("mysql://%v:%v@tcp(%v:%v)/%v?parseTime=true",
// 		requiredInput.UserName, requiredInput.Password,
// 		requiredInput.Url, requiredInput.Port, requiredInput.DbName)

//     absPath, _ := filepath.Abs("../migrations")
//     absPath = filepath.ToSlash(absPath) // ✅ Convert \ to /
//     m, err := migrate.New("file://" + absPath, dsnMigrate)
//     if err != nil {
//         log.Fatalf("Migration setup failed: %v", err)
//     }

//     if err := m.Up(); err != nil && err != migrate.ErrNoChange {
//         log.Fatalf("Migration failed: %v", err)
//     }
// 	// defer db.Close()
// 	// Open the database connection
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Fatal("Error connecting to the database:", err)
// 	}

// 	// Test the connection
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("Error pinging the database:", err)
// 	}

// 	fmt.Println("Successfully connected to MySQL!")
// 	return &MySqlDbInstance{MySql: db}

// }

func ConnectMysql(requiredInput PostgreSQLRequired) *PostgreSQLDbInstance {
	// Normal DB connection DSN
	dsn := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		requiredInput.Url,
		requiredInput.Port,
		requiredInput.UserName,
		requiredInput.Password,
		requiredInput.DbName,
	)

	// DSN for migrate
	dsnMigrate := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		requiredInput.UserName,
		requiredInput.Password,
		requiredInput.Url,
		requiredInput.Port,
		requiredInput.DbName,
	)

	// Migration path
	absPath, _ := filepath.Abs("../migrations")
	absPath = filepath.ToSlash(absPath)

	m, err := migrate.New("file://"+absPath, dsnMigrate)
	if err != nil {
		log.Fatalf("Migration setup failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	// Open DB connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	log.Println("✅ PostgreSQL connected successfully")
	return &PostgreSQLDbInstance{PostgreSQL: db}
}
