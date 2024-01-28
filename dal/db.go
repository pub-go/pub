package dal

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"code.gopub.tech/logs"
	"code.gopub.tech/pub/dal/model"
	"code.gopub.tech/pub/dal/query"
	"code.gopub.tech/pub/settings"
	driverHook "github.com/youthlin/driver"
	sqlite3 "github.com/youthlin/go-sqlcipher"
	"github.com/youthlin/sqlcipher"
	"gorm.io/gorm"
)

const dbName = "data.db"
const driverName = "sqlite3_hook"

var logSkip int
var DB *gorm.DB
var ctx = context.Background()

func MustInit(dir string) {
	register(dir) // 注册驱动
	open(dir)     // 打开
	migrate()     // 自动插入表结构
}

func register(dir string) {
	log := logs.NewLogger(logs.CombineHandlers(
		logs.NewHandler(),
		logs.NewHandler(
			logs.WithFile(filepath.Join(dir, "logs", "sql.log")),
		),
	))
	driverHook.Register(driverName, &sqlite3.SQLiteDriver{
		OnOpenHook: sqlite3.SimpleOpenHook, // for _pragma_xxx=yyy
	}, driverHook.NewHook(
		func(ctx context.Context, method driverHook.Method, query string, args any) context.Context {
			skip := calculateDepth([]string{"code.gopub.tech/pub/dal/query"})
			ctx = context.WithValue(ctx, &logSkip, skip)
			// before sql execute
			log.Log(ctx, skip, logs.LevelNotice, "[sql] method=%v, sql=%v, args=%v",
				method, query, args)
			return ctx
		},
		func(ctx context.Context, method driverHook.Method, query string, args, result any, err error) (any, error) {
			// after sql execute
			if skip, ok := ctx.Value(&logSkip).(int); ok {
				log.Log(ctx, skip, logs.LevelNotice, "[sql] method=%v, cost=%v, sql=%v, args=%v, result=%#v, err=%+v",
					method, driverHook.Cost(ctx), query, args, result, err)
			} else {
				log.Info(ctx, "[sql] method=%v, cost=%v, sql=%v, args=%v, result=%#v, err=%+v",
					method, driverHook.Cost(ctx), query, args, result, err)
			}
			return result, err
		},
	))
}

func calculateDepth(exclude []string) (skip int) {
	pc := make([]uintptr, 30)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	skip++
	for {
		frame, more := frames.Next()
		file := frame.File
		if strings.Contains(file, "code.gopub.tech/pub") {
			var ok = true
			for _, s := range exclude {
				if strings.Contains(file, s) {
					ok = false
				}
			}
			if ok {
				return
			}
		}
		skip++
		if !more {
			break
		}
	}
	return
}

func open(dir string) {
	params := []string{
		fmt.Sprintf(`_pragma_key=%q`, settings.AppConf.DBKey),
		// 注意顺序 _pragma_key 最先；注意引号
		`_pragma_cipher_page_size=1024`,
		`_pragma_kdf_iter=4000`,
		`_pragma_cipher_hmac_algorithm="HMAC_SHA1"`,
		`_pragma_cipher_kdf_algorithm="PBKDF2_HMAC_SHA1"`,
		`_pragma_cipher_use_hmac="OFF"`,
	}
	dbDir := filepath.Join(dir, "data")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s?%s", filepath.Join(dbDir, dbName), strings.Join(params, "&"))

	db, err := gorm.Open(&sqlcipher.Dialector{DriverName: driverName, DSN: dsn}, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	query.SetDefault(db)
}

func migrate() {
	db := DB
	db.AutoMigrate(&model.User{})
	users, err := query.User.WithContext(ctx).Find()
	logs.Info(ctx, "users: %v, err=%+v", users, err)
}
