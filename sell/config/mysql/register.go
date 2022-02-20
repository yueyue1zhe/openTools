package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
	"yueyue-sell/config"
)

var (
	DB    *gorm.DB
	DbErr error
)

func Test(dbHost, dbPort, dbDataName, dbUsername, dbPassword, dbTablePreFix string) error {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbDataName, "utf8mb4",
	)
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbTablePreFix,
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	return err
}

func InitMysql() {
	dbHost := config.Env.Mysql.Host
	dbPort := config.Env.Mysql.Port
	dbCharset := config.Env.Mysql.Charset
	dbDataName := config.Env.Mysql.Database
	dbUsername := config.Env.Mysql.Username
	dbPassword := config.Env.Mysql.Password
	dbTablePreFix := config.Env.Mysql.Pre
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbDataName, dbCharset,
	)
	DB, DbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbTablePreFix,
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if DbErr != nil {
		logrus.Errorf("数据库链接异常:%v", DbErr.Error())
		return
	}

	sqlDb, err := DB.DB()
	if err != nil {
		logrus.Errorf("链接池配置异常:%v", err.Error())
		return
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
}
