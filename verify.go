// 验证逻辑
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

// 计算Hash
func CountHash(data []byte) (hashString string) {
	hashByte := sha256.Sum256(data)
	hashString = hex.EncodeToString(hashByte[:])
	return hashString
}

// 检查Hash是否真确
func CheckingHash() (success []FileInfo, fail []FileInfo) {
	// 保存解析sha256.json的数据
	shaSaveData := []FileInfo{}

	// 读取Hash文件
	data, err := os.ReadFile(ScanDirPath + "sha256.json")
	ErrprDisplay(err)

	// 开始解析
	err = json.Unmarshal(data, &shaSaveData)
	ErrprDisplay(err)

	// 开始验证
	success, fail = shaVerify(shaSaveData)

	return success, fail
}

// 从sha256.json验证是否被篡改
func shaVerify(shaSaveData []FileInfo) (success []FileInfo, fail []FileInfo) {
	for _, v := range shaSaveData {
		data, err := os.ReadFile(v.Path)
		if err != nil {
			v.Path = err.Error()
			fail = append(fail, v)
		}

		if v.Hash == CountHash(data) {
			success = append(success, v)
		} else {
			// 无法同时读取和写入sha256.json文件Hash值，允许不同
			if v.Name != "sha256.json" {
				fail = append(fail, v)
			}
		}
	}

	return success, fail
}

// 判断文件是否存在于sha256.json
func fileChecking(dirName string) {
	// 读取文件目录
	if dirName == "" {
		dirName = "./"
	}

	dir, err := os.ReadDir(dirName)
	ErrprDisplay(err)
	fail := []string{}
	for _, v := range dir {
		is := 1
		if v.IsDir() {
			fileChecking(dirName + v.Name() + "/")
		} else {
			for _, t := range GetHash() {
				if v.Name() == t.Name {
					is = 0
				}
			}
			if is == 1 {
				fail = append(fail, v.Name())
			}
		}
	}

	fmt.Println(fail, "：不存在于hash表中")
}
