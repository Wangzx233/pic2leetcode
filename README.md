# 批量替换图片来源的小工具
> 只支持替换为 leetcode.com.cn 来源

## 使用说明
- 在编译好的文件同一目录下新增一个 `resource` 文件夹和一个 `cookie.txt` 文件
- 将需要处理的文章放入 `resource` 文件夹（文件路径随意，只要在 `resource` 中即可
- 将自己的 leetcode cookie 粘贴进 `cookie.txt`
- 运行

## 其他
如果需要自定义替换可以更改 `load/load.go` 中的 `InitReplaceMap` 函数，并取消