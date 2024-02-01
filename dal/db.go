package dal

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	"code.gopub.tech/errors"
	"code.gopub.tech/logs"
	"code.gopub.tech/pub/dal/model"
	"code.gopub.tech/pub/dal/query"
	"code.gopub.tech/pub/settings"
	"code.gopub.tech/pub/util"
	"code.gopub.tech/pub/webs"
	driverHook "github.com/youthlin/driver"
	sqlite3 "github.com/youthlin/go-sqlcipher"
	"github.com/youthlin/sqlcipher"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const dbName = "data.db"
const driverName = "sqlite3_hook"

var DB *gorm.DB
var sqlLogger logs.Logger
var ctx = context.Background()

func MustInit(dir string) {
	// sql 日志
	sqlLogger = logs.NewLogger(logs.CombineHandlers(
		logs.NewHandler(),
		logs.NewHandler(
			logs.WithFile(filepath.Join(dir, "logs", "sql.log")),
		),
	))
	register(dir) // 注册驱动
	open(dir)     // 打开
	migrate()     // 自动更新表结构
}

func register(dir string) {
	// 注册包装后的驱动 用于 sql 计数
	driverHook.Register(driverName, &sqlite3.SQLiteDriver{
		OnOpenHook: sqlite3.SimpleOpenHook, // for _pragma_xxx=yyy
	}, driverHook.NewHook(
		func(ctx context.Context, method driverHook.Method, query string, args any) context.Context {
			return ctx
		},
		func(ctx context.Context, method driverHook.Method, query string, args, result any, err error) (any, error) {
			webs.AddSqlCount(ctx)
			return result, err
		},
	))
}

// calculateDepth 忽略 gorm 调用栈 定位到项目调用的代码行
func calculateDepth() (skip int) {
	pc := make([]uintptr, 30)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	exclude := []string{"code.gopub.tech/pub/dal/query"}
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
		util.Panic(ctx, errors.Wrapf(err, "mkdir %s", dbDir))
	}
	dsn := fmt.Sprintf("%s?%s", filepath.Join(dbDir, dbName), strings.Join(params, "&"))

	db, err := gorm.Open(&sqlcipher.Dialector{DriverName: driverName, DSN: dsn}, &gorm.Config{
		Logger: &dbLogger{
			Logger:   sqlLogger,
			LogLevel: logger.Info,
		},
	})
	if err != nil {
		util.Panic(ctx, errors.Wrapf(err, "can not open database `%s`, is key correct?", dsn))
	}
	logs.Info(ctx, "init database success")
	DB = db
	query.SetDefault(db)
}

func migrate() {
	db := DB
	db.AutoMigrate(&model.User{}, &model.Option{})
}

var _ logger.Interface = (*dbLogger)(nil)

type dbLogger struct {
	logs.Logger
	logger.LogLevel
	SlowThreshold time.Duration
}

// LogMode implements logger.Interface.
func (l *dbLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

// Info implements logger.Interface.
func (l *dbLogger) Info(ctx context.Context, format string, args ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Logger.Log(ctx, calculateDepth(), logs.LevelInfo, format, args...)
	}
}

// Warn implements logger.Interface.
func (l *dbLogger) Warn(ctx context.Context, format string, args ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Logger.Log(ctx, calculateDepth(), logs.LevelWarn, format, args...)
	}
}

// Error implements logger.Interface.
func (l *dbLogger) Error(ctx context.Context, format string, args ...interface{}) {
	if l.LogLevel >= logger.Error {
		args = append(args, debug.Stack())
		l.Logger.Log(ctx, calculateDepth(), logs.LevelError, format+"\n%s", args...)
	}
}

// Trace implements logger.Interface.
func (l *dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error:
		var (
			sql, rows     = fc()
			rowsPrint any = rows
			skip          = calculateDepth()
		)
		if rows == -1 {
			rowsPrint = "-"
		}
		l.Logger.Log(ctx, skip, logs.LevelError, "err=[%s] [%.3fms] [rows:%v] %s",
			err, float64(elapsed.Nanoseconds())/1e6, rowsPrint, sql)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		var (
			sql, rows     = fc()
			rowsPrint any = rows
			skip          = calculateDepth()
		)
		if rows == -1 {
			rowsPrint = "-"
		}
		l.Logger.Log(ctx, skip, logs.LevelWarn, "[SLOW SQL >= %v] [%.3fms] [rows:%v] %s",
			l.SlowThreshold, float64(elapsed.Nanoseconds())/1e6, rowsPrint, sql)
	case l.LogLevel == logger.Info:
		var (
			sql, rows     = fc()
			rowsPrint any = rows
			skip          = calculateDepth()
		)
		if rows == -1 {
			rowsPrint = "-"
		}
		l.Logger.Log(ctx, skip, logs.LevelInfo, "[%.3fms] [rows:%v] %s",
			float64(elapsed.Nanoseconds())/1e6, rowsPrint, sql)
	}
}
