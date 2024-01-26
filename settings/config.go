package settings

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"code.gopub.tech/logs"
	"code.gopub.tech/logs/pkg/arg"
	"code.gopub.tech/errors"
	"code.gopub.tech/pub/util"
	"github.com/gin-gonic/gin"
)

var AppConf *Config
var ctx = context.Background()

const (
	ConfigFileName = "conf/app.json"
	defaultAddr    = ":8765"
	defaultKeyLen  = 7
)

type Config struct {
	Addr        string // 监听地址 默认为 ${defaultAddr}
	LangPath    string // 翻译文件夹 为空表示使用内置翻译
	Lang        string // 默认语言 为空表示使用系统语言
	ViewPath    string // 模板文件夹 为空表示使用内置模板
	ViewPattern string // 模板文件名正则 为空表示 \.html$
	Debug       bool   // 开启 Debug
	DBKey       string // 数据库密码 默认随机${defaultKeyLen}位字符
}

func ReadConfig(dir string) error {
	f, err := readOrCreateFile(dir)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return errors.Wrapf(err, "failed to read config file: %s", ConfigFileName)
	}
	var conf Config
	err = json.Unmarshal(b, &conf)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal config: %s", b)
	}
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	AppConf = &conf
	logs.Info(ctx, "read config file ok: %s", arg.JSON(AppConf))
	return nil
}

func readOrCreateFile(dir string) (io.Reader, error) {
	fileName := filepath.Join(dir, ConfigFileName)
	f, err := os.Open(fileName)
	if err != nil {
		logs.Info(ctx, "config file not exist, create: %s", fileName)
		configDir := filepath.Dir(fileName)
		if err := os.MkdirAll(configDir, 0755); err != nil { // drwxr-xr-x
			return nil, errors.Wrapf(err, "can not create dir: %s", configDir)
		}
		f, err = os.OpenFile(ConfigFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644) // -rw-r--r--
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create config file: %s", ConfigFileName)
		}
		var config = Config{Addr: defaultAddr, DBKey: util.RandStr(defaultKeyLen)}
		b, _ := json.MarshalIndent(config, "", "\t")
		logs.Info(ctx, "write defaut config: %s", b)
		_, err = f.Write(b)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to write config file: %s", ConfigFileName)
		}
		f, err = os.Open(ConfigFileName)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read config file: %s", ConfigFileName)
	}
	return f, nil
}
