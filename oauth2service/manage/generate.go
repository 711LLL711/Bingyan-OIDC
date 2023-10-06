package manage

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"strings"

	"github.com/google/uuid"
)

// client_id必须唯一
// 签发client_id和client_secret
func GenerateClientIDAndSecret() (string, string) {
	randBytes := make([]byte, 32)
	rand.Read(randBytes)
	// Convert to hexadecimal encoding
	clientID := hex.EncodeToString(randBytes)

	// Convert to base64 encoding
	clientSecret := base64.StdEncoding.EncodeToString(randBytes)
	clientSecret = strings.ToUpper(strings.TrimRight(clientSecret, "=")) //去掉生成的=
	return clientID, clientSecret
}

// 签发授权码
func GenerateAhthorizationCode() string {
	randBytes := make([]byte, 32)
	rand.Read(randBytes)
	// Convert to hexadecimal encoding
	AuthorizationCode := hex.EncodeToString(randBytes)
	return AuthorizationCode
}

// 签发token
// UUIDs生成,userid作为参数
func GenerateToken(userid string) string {
	buf := bytes.NewBufferString(userid)
	token := uuid.NewMD5(uuid.Must(uuid.NewRandom()), buf.Bytes())
	code := base64.URLEncoding.EncodeToString([]byte(token.String())) //base64编码
	code = strings.ToUpper(strings.TrimRight(code, "="))              //去掉生成的=
	return code
}
