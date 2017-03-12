package config

import (
	"log"
	"os"

	. "aboutcar/rest/db"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type Configs struct {
	PhantomjsFile      string
	PhantomjsTempJsDir string
	UserConfig         []map[string]string
}

type Edusite struct {
	Uuid      string `xorm:"varchar(255) index not null unique 'uuid'"`
	County    string `xorm:"varchar(255)"`
	TrainName string `xorm:"varchar(255)"`
}

type Animal struct {
	ID   int64
	Name string `gorm:"default:'galeone'"`
	Age  int64
}

var Config *Configs

func init() {
	Config = &Configs{
		PhantomjsFile:      "",
		PhantomjsTempJsDir: "",
		//   Step: []string{
		//    "西乡",
		//    "科目二训练",
		//    "杨健",
		//    "2017-02-15 13:00-14:00",
		//  },
		UserConfig: []map[string]string{
			{
				"userName": "430521199110034261",
				"passWord": "034261",
			},
		},
	}

	dburl := "host=localhost user=postgres dbname=car password=dkyz123 sslmode=disable"
	var err error
	if DB, err = xorm.NewEngine("postgres", dburl); err != nil {
		panic(err)
	}
	log.Println(err)
	log.Println(DB)

	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	DB.SetLogger(xorm.NewSimpleLogger(f))
	// DB.LogMode(false)
}

// 运行模式
const (
	UNSET int = iota - 1
	OFFLINE
	SERVER
	CLIENT
)

// 数据头部信息
const (
	// 任务请求Header
	REQTASK = iota + 1
	// 任务响应流头Header
	TASK
	// 打印Header
	LOG
)

// 运行状态
const (
	STOPPED = iota - 1
	STOP
	RUN
	PAUSE
)

var ThreadNum = 10
