```bash
# 安装抽取工具
go install -v code.gopub.tech/tpl/cmd/xtpl@latest

# 抽取翻译字符串
xtpl -path ../views -output messages.pot


# 初始化翻译
GETTEXTCLDRDIR=./cldr msginit -i messages.pot -l zh_CN -o zh_CN.po

# 更新模板
msgmerge -U zh_CN.po messages.pot
```
