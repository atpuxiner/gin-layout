## 安装及更新：
  - `go get -u`
  - `go mod tidy`

## 详情如下：
- gin
  - go get -u github.com/gin-gonic/gin
- viper
  - go get -u github.com/spf13/viper
- zap
  - go get -u go.uber.org/zap
  - 日志分割：
    - go get -u gopkg.in/natefinch/lumberjack.v2
- gorm
  - go get -u gorm.io/gorm
  - 不同的driver：
    - go get -u gorm.io/driver/postgres
    - go get -u gorm.io/driver/mysql
    - go get -u gorm.io/driver/sqlite
- swagger
  - go get -u github.com/swaggo/swag/cmd/swag
  - go install github.com/swaggo/swag/cmd/swag@latest
  - gin-swagger
    - go get -u github.com/swaggo/gin-swagger
    - go get -u github.com/swaggo/files
- cobra
  - go get -u github.com/spf13/cobra@latest
- jwt
  - go get github.com/golang-jwt/jwt/v4
- redis
  - go get -u github.com/go-redis/redis/v8
