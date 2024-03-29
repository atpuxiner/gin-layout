package api

import (
	"errors"
	"github.com/atpuxiner/gin-layout/app/utils/errs"
	"github.com/atpuxiner/gin-layout/app/utils/status"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type BaseApi struct{}

func (r BaseApi) Success(c *gin.Context, data any) {
	if data == nil {
		data = gin.H{}
	}
	res := ResponseJson{
		Time: time.Now(),
		Code: status.CodeSuccess,
		Msg:  status.CodeMsg(status.CodeSuccess),
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}

func (r BaseApi) Failure(c *gin.Context, code int, err error) {
	// 特殊错误处理
	if errors.Is(err, gorm.ErrRecordNotFound) {
		code = status.CodeRecordNotFound
	} else if errors.Is(err, errs.ErrRecordExisted) {
		code = status.CodeRecordExisted
	}
	res := ResponseJson{
		Time:  time.Now(),
		Code:  code,
		Msg:   status.CodeMsg(code),
		Error: err.Error(),
		Data:  gin.H{},
	}
	c.JSON(http.StatusOK, res)
}

func (r BaseApi) FailureOnlyErr(c *gin.Context, err error) {
	res := ResponseJson{
		Time:  time.Now(),
		Code:  status.CodeFailure,
		Msg:   status.CodeMsg(status.CodeFailure),
		Error: err.Error(),
		Data:  gin.H{},
	}
	c.JSON(http.StatusOK, res)
}
