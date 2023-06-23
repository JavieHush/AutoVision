package service

import (
	"AutoVision/global"
	"AutoVision/internal/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"os"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{
		ctx: ctx,
	}
	svc.dao = dao.New(global.DBEngine)

	return svc
}

func DownloadHandler(c *gin.Context) {
	filePath := "storage/res/result.csv"
	fileName := "result.csv"

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		c.String(http.StatusNotFound, "File not found")
		return
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, _ := file.Stat()

	// 设置响应头信息
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// 将文件内容写入响应体中
	buf := make([]byte, 8192)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			break
		}
		if n == 0 {
			break
		}

		c.Writer.Write(buf[:n])
	}

	c.Status(http.StatusOK)
}
