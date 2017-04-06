package main

//import (
//	"github.com/gin-gonic/gin"
//	"qiniupkg.com/api.v7/conf"
//	"qiniupkg.com/api.v7/kodo"
//
//	"api/utils"
//)
//
//func UploadToken(c *gin.Context) {
//	accessKey, err := utils.ConfigGetString("qiniu", "access_key")
//	utils.FailOnError("get access_key failed: %s", err)
//
//	secretKey, err := utils.ConfigGetString("qiniu", "secret_key")
//	utils.FailOnError("get secret_key failed: %s", err)
//
//	bucket, err := utils.ConfigGetString("qiniu", "bucket")
//	utils.FailOnError("get bucket failed: %s", err)
//
//	host, err := utils.ConfigGetString("qiniu", "host")
//	utils.FailOnError("get host failed: %s", err)
//
//	conf.ACCESS_KEY = accessKey
//	conf.SECRET_KEY = secretKey
//
//	client := kodo.New(0, nil)
//
//	policy := &kodo.PutPolicy{
//		Scope:   bucket,
//		Expires: 3600,
//	}
//	token := client.MakeUptoken(policy)
//
//	c.JSON(200,  gin.H{
//		"token": token,
//		"host":  host,
//	})
//}
