```bash
# 安装抽取工具
go install -v code.gopub.tech/tpl/cmd/xtpl@latest

# 抽取翻译字符串
xtpl -path ../resource/views -output tpl.pot

# tree
#    -I 忽略指定的 pattern
#    -P 指定文件 pattern
#    -F 在文件夹后添加斜线结尾
#    -f 给每个文件添加路径前缀
#    -i 不要输入树形 直接每行一个文件
#    --noreport 不要输入统计信息: x directories, y files
# grep
#    -v 排除以斜线结尾的
tree -I "cmd"  -P "*.go" -F -f -i --noreport .. | grep -v /$ | xargs xgettext -C --from-code=UTF-8 -o go.pot  -kT -kN:1,2 -kN64:1,2 -kX:2,1c -kXN:2,3,1c -kXN64:2,3,1c

msgcat tpl.pot go.pot -o messages.pot


# 初始化翻译
GETTEXTCLDRDIR=./cldr msginit -i messages.pot -l zh_CN -o zh_CN.po

# 更新模板
msgmerge -U  zh_CN.po messages.pot
```
