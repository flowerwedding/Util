/**
 * @Title  七牛云图床
 * @description  #
 * @Author  沈来
 * @Update  2020/8/23 22:03
 **/
package main
/*
import (
	"context"
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

func main() {
	accessKey := "TgVGKnpCMLDI6hSS4rSWE3g-FZjMPf6Zbc******"
	secretKey := "zqZvH3fNVaggw00oc9wCrcWKgeeiV7WITF******"
	localFile := "img/1.jpg"
	bucket := "wangshubotest"
	key := "1.jpg"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}*/