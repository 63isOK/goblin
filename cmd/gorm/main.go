package main

// import (
// 	"fmt"
//
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/schema"
// )
//
// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }
//
// func main() {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s&timeout=%ds&readTimeout=%ds&writeTimeout=%ds",
// 		"main",
// 		"tttc-2016",
// 		"dev.ttpark.cn:3306",
// 		"iot-core",
// 		"charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
// 		10, 10, 10)
//
// 	dialector := mysql.New(mysql.Config{
// 		DSN:                       dsn,   // DSN配置信息
// 		DefaultStringSize:         256,   // string类型字段的默认长度
// 		DisableDatetimePrecision:  true,  // 禁用datetime精度，因为MySQL5.6之前的数据库不支持
// 		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并创建的方式，因为MySQL5.7之前的数据库和MariaDB不支持重命名索引
// 		DontSupportRenameColumn:   true,  // 用`change`重命名列，因为MySQL8之前的数据库和MariaDB不支持重命名列
// 		SkipInitializeWithVersion: false, // 根据当前MySQL版本自动配置
// 	})
//
// 	db, err := gorm.Open(dialector,
// 		&gorm.Config{
// 			DisableForeignKeyConstraintWhenMigrating: true, // 禁用AutoMigrate自动创建数据库外键约束
// 			NamingStrategy: schema.NamingStrategy{
// 				TablePrefix:   "iot_", // 表名前缀
// 				SingularTable: true,   // 使用单数表名
// 			},
// 		})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
//
// 	db.Table("iot_product_model").Select("").
// }
