每一步后都应该先自行记录前置种子文件，以及最终补充 交互和数据 的确认逻辑
- New: 新建
- Duplicate: 复制，再次复制
- VersionDialog: 打开版本列表，切换选择，切换tab，切换选择，关键词A，关键词B，选中当前显示并确认选择，选中非当前显示并确认选择
- DeleteVersion: 打开版本列表，删除非当前选中也非当前显示，删除当前选中，删除当前显示，删除所有版本
- PublishAndUnPulish: 发布，取消发布
- Schedule: 草稿态的 start < end < now，now < start < end，end < start < now ，其他态可代码调整


```
# gen command
sh $GOPATH/src/github.com/qor5/admin/utils/testflow/gentool/gen.sh ./sample ./patch ./_backup .
```
