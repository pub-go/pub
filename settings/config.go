package settings

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"code.gopub.tech/errors"
	"code.gopub.tech/logs"
	"code.gopub.tech/logs/pkg/arg"
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
	path        string
	Addr        string // 监听地址 默认为 ${defaultAddr}
	LangPath    string // 翻译文件夹 为空表示使用内置翻译
	Lang        string // 默认语言 为空表示使用系统语言
	Resource    string // 资源文件夹 为空表示使用内置资源 (可包含 static, views 文件夹)
	ViewPattern string // 模板文件名正则 为空表示 \.html$ (当有 views 文件夹时)
	Debug       bool   // 开启 Debug
	DBKey       string // 数据库密码 默认随机${defaultKeyLen}位字符
}

func (c *Config) StaticPath() string { return c.resource("static") }
func (c *Config) ViewPath() string   { return c.resource("views") }
func (c *Config) resource(sub string) string {
	resource := c.Resource
	if resource != "" { // 指定了资源文件夹
		dir := filepath.Dir(c.path)
		path := filepath.Join(dir, resource, sub)
		if _, err := os.Stat(path); err == nil {
			return path // 资源文件夹中存在指定的子文件夹
		} else {
			logs.Warn(ctx, "conf path=%v err=%+v", path, err)
		}
	}
	return ""
}

func ReadConfig(dir string) error {
	f, err := readOrCreateFile(dir)
	if err != nil {
		return errors.Wrapf(err, "readOrCreateFile")
	}
	path, err := filepath.Abs(f.Name())
	if err != nil {
		return errors.Wrapf(err, "getConfigFileAbsPath")
	}
	logs.Info(ctx, "config file abs path: %s", path)
	b, err := io.ReadAll(f)
	if err != nil {
		return errors.Wrapf(err, "failed to read config file: %s", ConfigFileName)
	}
	var conf Config
	err = json.Unmarshal(b, &conf)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal config: %s", b)
	}
	conf.path = path
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	AppConf = &conf
	logs.Info(ctx, "read config file ok: %s", arg.JSON(AppConf))
	return nil
}

func readOrCreateFile(dir string) (*os.File, error) {
	fileName := filepath.Join(dir, ConfigFileName)
	logs.Info(ctx, "read config file: %s", fileName)
	f, err := os.Open(fileName)
	if err != nil {
		logs.Notice(ctx, "config file not exist, create: %s", fileName)
		configDir := filepath.Dir(fileName)
		if err := os.MkdirAll(configDir, 0755); err != nil { // drwxr-xr-x
			return nil, errors.Wrapf(err, "can not create dir: %s", configDir)
		}
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644) // -rw-r--r--
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
		f, err = os.Open(fileName)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read config file: %s", ConfigFileName)
	}
	return f, nil
}
