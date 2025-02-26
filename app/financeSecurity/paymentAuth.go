package financeSecurity

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	// 生成随机数作为示例内容
	randNum, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}

	// 生成数字签名
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%d", randNum)))
	digest := hash.Sum(nil)
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digest)

	// 验证数字签名
	hash2 := sha256.New()
	hash2.Write([]byte(fmt.Sprintf("%d", randNum)))
	digest2 := hash2.Sum(nil)
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, digest2, signature)
	if err != nil {
		panic(err)
	}

	fmt.Println("支付安全算法实现成功")
}
