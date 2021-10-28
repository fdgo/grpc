package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var db *gorm.DB
var err error

func init() {
	GoodsInfo = append(GoodsInfo, GoodsInvInfo{
		GoodsId: 12345,
		Num:     5,
	})
	GoodsInfo = append(GoodsInfo, GoodsInvInfo{
		GoodsId: 56789,
		Num:     20,
	})
}
func main() {
	dsn := "root:000000@tcp(192.168.164.128:3306)/gormlock?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	// 全局模式
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	}, &gorm.Config{NamingStrategy: schema.NamingStrategy{ /*TablePrefix: preTag,*/ SingularTable: true}})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Inventory{})
	db.AutoMigrate(&InventoryNew{})
	db.AutoMigrate(&Delivery{})
	db.AutoMigrate(&StockSellDetail{})

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetConnMaxIdleTime(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		tx := db.Begin()
		sellDetail := StockSellDetail{
			OrderSn: uuid.NewString(),
			Status:  1,
		}
		var details []GoodsDetail
		for _, goodInfo := range GoodsInfo {
			details = append(details, GoodsDetail{
				Goods: goodInfo.GoodsId,
				Num:   goodInfo.Num,
			})
			var inv Inventory
			if result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
				tx.Rollback() //回滚之前的操作
				return
			}
			if result := tx.Where(&Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
				tx.Rollback() //回滚之前的操作
				return
			}
			//判断库存是否充足
			if inv.Stocks < goodInfo.Num {
				tx.Rollback() //回滚之前的操作
				return
			}
			//扣减， 会出现数据不一致的问题 - 锁，分布式锁
			inv.Stocks -= goodInfo.Num
			tx.Save(&inv)
		}
		sellDetail.Detail = details
		//写selldetail表
		if result := tx.Create(&sellDetail); result.RowsAffected == 0 {
			tx.Rollback()
			return
		}
		tx.Commit()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
