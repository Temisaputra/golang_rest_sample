package sql

import (
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/Temisaputra/warOnk/config"
	"github.com/jmoiron/sqlx"
)

type SQLServerConnectionManager struct {
	db *sqlx.DB
}

func NewSQLServerConnectionManager(cfg config.Config) (*SQLServerConnectionManager, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBPassword, cfg.DBName)
	db, err := sqlx.Open(cfg.DBType, dsn)
	if err != nil {
		log.Fatalf("failed to connect to sql server: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping to sql server: %v", err)
		return nil, err
	}

	db.SetMaxIdleConns(cfg.DBMaxIdleConns)
	db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	db.SetConnMaxIdleTime(time.Duration(cfg.DBMaxIdleTime) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.DBMaxLifetime) * time.Second)

	log.Println("connected to sql server")

	return &SQLServerConnectionManager{
		db: db,
	}, nil
}

func (cm *SQLServerConnectionManager) Close() error {
	log.Println("closing sql server connection")
	return cm.db.Close()
}

func (cm *SQLServerConnectionManager) GetQuery() *SingleInstruction {
	return NewSingleInstruction(cm.db)
}

func (cm *SQLServerConnectionManager) GetTransaction() *MultiInstruction {
	return NewMultiInstruction(cm.db)
}
