package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	start := time.Now().UTC().Unix()
	end := start + 3600
	byts := bytes.NewBuffer([]byte{})
	byts.WriteString(`{"iss":"server1@api-project-31524972.iam.gserviceaccount.com","scope":"https://www.googleapis.com/auth/androidpublisher","aud":"https://www.googleapis.com/oauth2/v4/token","exp":`)
	byts.WriteString(strconv.FormatInt(end, 10))
	byts.WriteString(`,"iat":`)
	byts.WriteString(strconv.FormatInt(start, 10))
	byts.WriteString(`}`)
	dst := base64.URLEncoding.EncodeToString(byts.Bytes())
	//dst = strings.ReplaceAll(dst, "=", "")
	prikey := "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDjhSsYa+mleh5H\nPF3LmClSZ4zLsBqzcfJ9vsKLMXiM4r+e4Dr3MOzcvjtHoH1ynsUgm4jKL+FYUw5S\nV/gIgkPKJdN930goJEaHdfGD4ZRSWqH8GQe2D02JVZDahTW5jRtsrt5ApJ6eFUc7\nn8RuY++GoaNA6kIKN94/8FraMeicrKP5cLzsCd6alX0GaQCJz3fTJ1ypC2feNTV6\nJ0QPTInmriNLLX6JSDxrXZmXlParpRHSrt9kbI0PCplkH73f6fKM7vJVtp74q8I9\njczvosf/V39VjSc1pt2+xg0Gi9bpvsGod1MI8jisNPhQXarX21ppmUiHvApTkw0B\nmSMMr3VdAgMBAAECggEAEbAdrJVfIb0/s1wPEq/urnhcas1zFfZK2tnEuBuNeq56\nJTjbfLIyB+tGIohomEudma5d0RIt27cBSweJweeWq5WLPqLoMi63yPozX4RfCpP/\nOeEcR1wjNAUR8NsgVR+SPT3PC4mAx1tyIUGHfOmKCpZwYbCUl8TGI4RlG8d7hQqF\nYgFGjcBcAVjHQ9/Hh/mORuVATv6zqoGmQGhCrdEYV9Va+WULys+gUSHlsKINC0RB\n4+lVL85xEBeMH3SDd+ixR6hjTwdFM8z17243gKYvCS9OFwHS4RmFUKi+wN+XuDND\noV5dtzMZjaf0BQ+vE+haHWa2OMp2PnokTOmO0dP/QQKBgQDzveOgIBHyaR671fWw\nAtRBXkGBXn8y4SqMP2dopgtwneYBy98tNvaSF3+wPCZfCDQjmEYZDPRiDWgzGLRq\ne7AL/snkkngSW1Cp77v/hrPNFunLpTOmcjAgNDwKCSqHuDE9a6Md5JZ2J4Ct1d1n\nVIVyQ3kvMAcstZq5K3T6z7ISkQKBgQDu9m5ItbbRIDC7bbKAuZLw5Pf8I6hCy6J2\n4XztCxsa8iyH3A7MVE7LgGx3bKEBMBzpmNv7fodZQWpQo4SApC4vF4g9X9R4QwGr\n57affg1k9RK+edZ4BlH9fpPjp+U5TAhs13v4+E1zy9OCc3tITo5uhHHRHWM3G2nK\n6kv3B1ZEDQKBgHKnZCeybj7FS/u3jbaZ3hZRrCaauOLKICWQvafwU3lKDSPTLswq\nCpp2C05vPO1/Amer/W1TNrHY9Kb0fAmK3SkHVRj7/RdFdRA7AQgV6QYUPS3aLA2j\nsRe0+nkODr+A2Ui3FSe+mzhBJLqg22D71ToGmz6jLPzPAFUSKBjDElTBAoGAMlBy\n4h5YsumrORuc1Ru9w7kCOfWsDPxhZdSOgD6xY1gQZj7AYudxe8m7jN2zfNOLkufb\nkbWPfAyY/Qeg85EeJE45ImsWCohZRr/QJP7ehR5Q5wgyTy1NgClxrCKC0jCfKYOp\nl794V2RYUYRNNelMdhqu+E/OvyCngtEYU5gY4tUCgYEAlr2+J0uB2zEzLyZmt9k/\nFZsq0ShB/v2UMVZUSqTM24NivMmbiYVdTCSm2JwtMW6QhR3w7t4e+1Wt27TfBFOx\nNCxZIgbw///QoopLArsU2ytqeA0IIHVViqzm6l70/pBr00+YLuQcD48iuD3IWS3z\nmYFCCC7Tk7cRtOXJTb/aadM=\n-----END PRIVATE KEY-----\n"
	dat1 := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9." + dst
	hash := sha256.Sum256([]byte(dat1))

	//解析key
	var err error
	var block *pem.Block
	if block, _ = pem.Decode([]byte(prikey)); block == nil {
		log.Panicln(err.Error())
	}
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		log.Panicln(err.Error())
	}
	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		log.Panicln("error")
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, pkey, crypto.SHA256, hash[:])
	if err != nil {
		log.Println(err.Error())
	}
	dst = base64.URLEncoding.EncodeToString(signature)
	//dst = strings.ReplaceAll(dst, "=", "")

	tkn := dat1 + "." + dst
	reader := bytes.NewReader([]byte("grant_type=urn:ietf:params:oauth:grant-type:jwt-bearer&assertion=" + tkn))
	req, err := http.NewRequest("POST", "https://www.googleapis.com/oauth2/v4/token", reader)
	if err != nil {
		log.Println(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Host", "www.googleapis.com")
	//
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer res.Body.Close()

	ret, _ := ioutil.ReadAll(res.Body)
	log.Println(string(ret))
}
