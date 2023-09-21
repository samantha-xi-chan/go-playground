package play301_x509

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

func Play() {
	// 1. 读取 basic.json 文件
	jsonFilePath := "test/basic.json"
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatalf("无法读取 JSON 文件：%v", err)
	}

	// 2. 计算 SHA256 摘要
	hash := sha256.Sum256(jsonData)

	// 3. 读取 id_rsa_pkcs1 的私钥文件
	privateKeyPath := "test/id_rsa_pkcs1"
	privateKeyData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("无法读取私钥文件：%v", err)
	}

	// 4. 解析私钥
	privateKey, err := parsePrivateKey(privateKeyData)
	if err != nil {
		log.Fatalf("无法解析私钥：%v", err)
	}

	// 5. 使用私钥进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		log.Fatalf("签名失败：%v", err)
	}

	// 6. 对签名结果进行 BASE64 编码
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
	fmt.Printf("签名结果（BASE64 编码）：%s\n", signatureBase64)
}

func parsePrivateKey(privateKeyData []byte) (*rsa.PrivateKey, error) {
	// 解析 PEM 格式的私钥
	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return nil, fmt.Errorf("无法解析私钥")
	}

	// 解析 RSA 私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
