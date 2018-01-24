package service

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"log"
	"wxapi/models"
	"strconv"
)

func initDb() {
	conf, err := config.NewConfig("ini", "conf/db.conf")
	if err != nil {
		log.Fatal(err)
	}
	dbtype := conf.String("dbtype")
	maxidle, _ := strconv.Atoi(conf.String("maxidle"))
	maxconn, _ := strconv.Atoi(conf.String("maxconn"))
	switch dbtype {
	case "postgresql":postgresql(conf, maxidle, maxconn)
	case "mysql":mysql(conf, maxidle, maxconn)
	case "sqlite3":sqlite(conf)
	default:log.Fatal("数据库不支持或配置文件填写有误")
	}
	orm.RegisterModel(new(models.User), new(models.Reply), new(models.Log))
	orm.RunSyncdb("default", false, true)
}

func postgresql(conf config.Configer, maxidle, maxconn int) {
	//注册Driver
	orm.RegisterDriver("postgres", orm.DRPostgres)
	//获取配置
	db, _ := conf.GetSection("postgresql")
	//构造conn连接
	conn := "user=" + db["user"] + " password=" + db["password"] +
		" dbname=" + db["dbname"] + " host=" + db["host"] +
		" port=" + db["port"] + " sslmode=" + db["sslmode"]
	//注册数据库连接
	if err := orm.RegisterDataBase("default", "postgres", conn, maxidle, maxconn); err != nil {
		log.Fatal("postgresql连接失败，请检查配置文件填写是否有误")
	}
}

func mysql(conf config.Configer, maxidle, maxconn int) {
	//注册Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//获取配置
	db, _ := conf.GetSection("mysql")
	//构造conn连接
	conn := db["user"] + ":" + db["password"] +
		"@tcp(" + db["host"] + ":" + db["port"] + ")/" +
		db["dbname"] + "?charset=utf8"
	//注册数据库连接
	if err := orm.RegisterDataBase("default", "mysql", conn, maxidle, maxconn); err != nil {
		log.Fatal("mysql连接失败，请检查配置文件填写是否有误")
	}

}

func sqlite(conf config.Configer) {
	//注册Driver
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	//获取配置
	db, _ := conf.GetSection("sqlite3")
	//构造conn连接
	if err := orm.RegisterDataBase("default", "sqlite3", db["dbname"] + ".db"); err != nil {
		log.Fatal("sqlite连接失败，请检查配置文件填写是否有误")
	}
}
