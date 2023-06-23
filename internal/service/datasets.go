package service

import (
	"AutoVision/internal/model"
	"AutoVision/pkg/app"
	"archive/zip"
	"fmt"
	_ "github.com/gin-gonic/gin/binding"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type DatasetTrainListRequest struct {
	UserID uint `form:"user_id" binding:"gte=0"`
}

type DatasetPredListRequest struct {
	UserID uint `form:"user_id" binding:"gte=0"`
}

type DeleteDatasetRequest struct {
	DatasetID uint `form:"dataset_id" binding:"required,gte=0"`
}

type CreateTrainDatasetRequest struct {
	UserID      uint   `form:"user_id" binding:"gte=0"`
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
}

type CreatePredDatasetRequest struct {
	UserID      uint   `form:"user_id" binding:"gte=0"`
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
}

func (svc *Service) DatasetTrainList(param *DatasetTrainListRequest, pager *app.Pager) ([]*model.Datasets, error) {
	return svc.dao.GetTrainDataset(param.UserID, pager.Page, pager.PageSize)
}

func (svc *Service) DatasetPredList(param *DatasetPredListRequest, pager *app.Pager) ([]*model.Datasets, error) {
	return svc.dao.GetPredDataset(param.UserID, pager.Page, pager.PageSize)
}

func (svc *Service) DeleteDataset(param *DeleteDatasetRequest) error {
	return svc.dao.DeleteDataset(param.DatasetID)
}

func (svc *Service) CreateTrainDataset(param *CreateTrainDatasetRequest, fileSize int) error {
	return svc.dao.CreateTrainDataset(param.Name, param.Description, time.Now(), param.UserID, fileSize)
}

func (svc *Service) CreatePredDataset(param *CreatePredDatasetRequest, fileSize int) error {
	return svc.dao.CreatePredDataset(param.Name, param.Description, time.Now(), param.UserID, fileSize)
}

// Unzip 解压 ZIP 文件
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func(r *zip.ReadCloser) {
		err := r.Close()
		if err != nil {

		}
	}(r)

	for _, f := range r.File {
		// 获取解压后的文件路径
		fpath := filepath.Join(dest, f.Name)

		// 检查文件路径是否在解压目录中
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("非法文件路径：%s", fpath)
		}

		if f.FileInfo().IsDir() {
			// 创建目录
			err := os.MkdirAll(fpath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		// 创建文件
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		err = outFile.Close()
		if err != nil {
			return err
		}
		err = rc.Close()
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
	}

	return nil
}
