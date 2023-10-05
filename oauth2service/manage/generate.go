package manage

import (
	"OIDC/oauth2service/model"
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
	return clientID, clientSecret
}

// 签发授权码
func GenerateAhthorizationCode() {

}

// 签发token
// UUIDs生成
func GenerateToken(ClientInfo model.ClientInfo) string {
	buf := bytes.NewBufferString(ClientInfo.ClientID)
	token := uuid.NewMD5(uuid.Must(uuid.NewRandom()), buf.Bytes())
	code := base64.URLEncoding.EncodeToString([]byte(token.String()))
	code = strings.ToUpper(strings.TrimRight(code, "="))
	return code
}
