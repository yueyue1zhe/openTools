package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"openTools/y/sys/env"
	"time"
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
	dbHost := env.Conf.Mysql.Host
	dbPort := env.Conf.Mysql.Port
	dbCharset := env.Conf.Mysql.Charset
	dbDataName := env.Conf.Mysql.Database
	dbUsername := env.Conf.Mysql.Username
	dbPassword := env.Conf.Mysql.Password
	dbTablePreFix := env.Conf.Mysql.Pre
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
