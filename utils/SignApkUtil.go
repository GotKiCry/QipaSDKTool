package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const signPath = "E:\\tool\\tool\\qpabox.keystore"
const signAlias = "7pabox"
const keyAliasPassword = "androidgamebox"
const keyPassword = "7paboxandroid"

// SignApkV1 执行V1签名
func SignApkV1(apkFile string, signApkFile string) {
	apkFile = strings.Replace(apkFile, "\"", "", -1)
	signApkFile = strings.Replace(signApkFile, "\"", "", -1)

	signError := RunCmd("jarsigner", "-keystore", signPath, "-signedjar", signApkFile, apkFile, signAlias, "-keypass", keyPassword, "-storepass", keyAliasPassword)
	if signError == io.EOF {
		fmt.Println("签名V1成功")
		os.Rename(signApkFile, apkFile)
		RunCmd("zipalign", "-v", "4", apkFile, apkFile+"zipalign")
		os.RemoveAll(apkFile)
		os.Rename(apkFile+"zipalign", apkFile)
	} else {
		fmt.Println("签名失败")
	}
}
