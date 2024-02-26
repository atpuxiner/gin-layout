# gin-layout

## What is this?
- by: axiner
- gin-layout
- This is a gin layout.

## 脚手架使用
- go get -u github.com/atpuxiner/gtools/gtcli
- go install github.com/atpuxiner/gtools/gtcli
- gtcli gin -p <项目名称> -m <模块名称> -d <目录(不指定则默认当前)>

## 项目启动
- 1）cd 到项目根目录
- 2）初始化相关
  - 第三方模块
    - `go get -u`
    - `go mod tidy`
  - swagger
    - swag init
- 3）编译启动
  - win:
    - `go build -o gin-layout.exe main.go`
    - `./gin-layout.exe runserver`
  - linux:
    - `go build -o gin-layout main.go`
    - `./gin-layout runserver`

## 项目结构
- ABD：ABD模式
  - A   api
  - B   business
  - D   datatype
- 调用过程：main.go(initializer) - router(middleware) - api - business - (datatype)
- 项目结构详情（命名经过多次修改敲定，简洁易懂，ABD目录贴合避免杂乱无章）
  ```
  └── gin-layout
      ├── app                         (应用)
      │   ├── api                     ├── (api)
      │   │   └── v1                  │   └── (v1版本)
      │   ├── business                ├── (业务)
      │   ├── datatype                ├── (数据类型)
      │   │   ├── entity              │   ├── (实体)
      │   │   └── model               │   └── (模型)
      │   ├── initializer             ├── (初始化)
      │   │   ├── conf                │   ├── (配置)
      │   │   ├── db                  │   ├── (数据库)
      │   │   ├── logger              │   ├── (日志)
      │   │   └── redis               │   └── (redis)
      │   ├── middleware              ├── (中间件)
      │   ├── router                  ├── (路由)
      │   └── utils                   └── (utils)
      ├── cmd                         (命令目录)
      ├── config                      (配置目录)
      ├── deploy                      (部署目录)
      ├── docs                        (文档目录)
      ├── log                         (日志目录)
      ├── .gitignore
      ├── go.mod
      ├── LICENSE
      ├── main.go
      └── README.md
  ```

## LICENSE
This project is released under the MIT License (MIT). See [LICENSE](LICENSE)
