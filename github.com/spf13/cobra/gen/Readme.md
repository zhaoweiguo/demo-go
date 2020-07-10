
## 生成cobra文件
```
$ go get -u github.com/spf13/cobra/cobra
```

## 使用
```
$ cobra init --pkg-name <name>
如:
$ cobra init --pkg-name github.com/zhaoweiguo/demo-go/github.com/spf13/cobra/gen
Your Cobra applicaton is ready at
/Users/zhaoweiguo/9tool/go/src/github.com/zhaoweiguo/demo-go/github.com/spf13/cobra/gen/init
```

## 查看生成结果
```
$ tree
.
├── cmd
│   └── root.go
└── main.go
```

## 添加一个子命令
```
$ cobra add <subName>
如:
$ cobra add version
version created at /Users/zhaoweiguo/9tool/go/src/github.com/zhaoweiguo/demo-go/github.com/spf13/cobra/gen
$ tree
├── LICENSE
├── cmd
│   ├── root.go
│   └── version.go
└── main.go

```


