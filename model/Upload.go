package model

import (
	"context"
	"goblog/utils"
	"goblog/utils/errmsg"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var accessKey = utils.AccessKey
var secretKey = utils.SecretKey
var bucket = utils.Bucket
var imgUrl = utils.QiniuServer

func UploadFile(file multipart.File, fileName string, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuabei,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}

	err := formUploader.Put(context.Background(), &ret, upToken, fileName, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := imgUrl + ret.Key
	return url, errmsg.SUCCSE
}
