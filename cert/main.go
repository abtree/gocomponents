package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

/*
	生成TLS证书

RSA的公钥与私钥
*/
func main() {
	generateTLS()
}

func generateTLS() {
	// sn为一个随机数(大int类型)
	sn, _ := rand.Int(rand.Reader, big.NewInt(time.Now().Unix()))
	// 证书模板
	temp := x509.Certificate{
		SerialNumber: sn,
		Subject:      pkix.Name{CommonName: "localhost"},
		NotBefore:    time.Now(),                           //生效时间
		NotAfter:     time.Now().Add(365 * 24 * time.Hour), //过期时间
	}
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 生成证书
	certDER, err := x509.CreateCertificate(rand.Reader, &temp, &temp, &key.PublicKey, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 私钥写入文件
	file, err := os.OpenFile("private.pem", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	keyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	err = pem.Encode(file, keyBlock)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 公钥写入文件
	file1, err := os.OpenFile("cert.pem", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file1.Close()
	certBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	}
	err = pem.Encode(file1, certBlock)
	if err != nil {
		fmt.Println(err)
		return
	}
}
