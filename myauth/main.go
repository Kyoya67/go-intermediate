package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
)

func main() {
	idToken := `eyJhbGciOiJSUzI1NiIsImtpZCI6ImYxMGY4NzQwNWE5NzljMWRmMzZkZjI2NjA2NzM0ZjMzY2Q4NWMyNzEiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI0ODI2MTA2NTk4NTgtaW1mczc5a2w5Y3NzcWRjOWE0NmM4ZHMyajB1YW5lOGYuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI0ODI2MTA2NTk4NTgtaW1mczc5a2w5Y3NzcWRjOWE0NmM4ZHMyajB1YW5lOGYuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTYxMDg1Njc2MzUxMTczMzE1MDciLCJub25jZSI6ImVzbiIsIm5iZiI6MTc4NzA0NTM0OSwibmFtZSI6IuS6rOiwt-WHjOaxsCIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS9BQ2c4b2NLSmNtZTZuZmtaektpZWtncEJzUk9RRTZ5cGVVM1FrV0NOSlozMkF4WkNZbGVTZW5nVD1zOTYtYyIsImdpdmVuX25hbWUiOiLlh4zmsbAiLCJmYW1pbHlfbmFtZSI6IuS6rOiwtyIsImlhdCI6MTc4NzA0NTY0OSwiZXhwIjoxNzg3MDQ5MjQ5LCJqdGkiOiJkYmIzZDU1OWNmODg0YTBiM2FjNWRhYmE2ZWFjOTA2YzRmZGRiY2VlIn0.GGfkeMHcoN8aCLMYFwfrz-PynbpILpgSBndTxcd3pJRh61xnIiPkFceaVl4XEuhLBCtgiJ7aZ2i4PtWUTw_s6trteH-X3fwPo1auizw7lk5jlwmp9WoXz-zOggWCE4UY8qtRrIk02QIpWE065j4h9XeiZvcd9mQUlpxzv8Sdr4P-stqWx5VkKNHANhG9wkwviLOV9QR-dnkUOoX9cK-5MbUyYCJerYlDFRslqPVwzA_pm9o-EIMk3im1fhhSIv3N1PrMlLQtxqGjz45n3YHAwPJWEUTgaOfJw6dEWIwExYXVEO8AfDnb4_l7VdFtp021UL98gPR1H_jAwpufVBaaEw`

	dataArray := strings.Split(idToken, ".")
	header, payload, sig := dataArray[0], dataArray[1], dataArray[2]

	// headerをbase64 decodeする
	headerData, err := base64.RawURLEncoding.DecodeString(header)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// payloadをbase64 decodeする
	payloadData, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// fmt.Println("header: ", string(headerData))
	// fmt.Println("payload: ", string(payloadData))

	// 公開鍵を使えるようにする
	E := "AQAB"
	N := "4rY5uwZK1dQ-UVgB5s4NLyC-u5LC2MT7b8GWZztiNgMsp0Nnqx0pM7Ofx0ws32N2aZcx10-J8ydQxnNb9uAcf-7LyhyOIcv_WEyzaSbUAMOgoF-nQmJetckxNg6ekhNfaFcTQS0T-29ql2_CBLIML6CvSh-r0fgWRsqN2ayB7wCl74Gv6OOVbvagUWhj5z2L6o_plmsPDwLVuvA7o3WDEDjoq-IXafRQowj92kQUenrOKD4YCopuLIBhel6VH8doFRNZ6KISQhMcOivWaLU_UtKKAMloGJieTf_3r-_nErs2h5wB7T7FrMCScmO7mvFQXKh8_4P-MlbfgS9CUvQksw"
	dn, _ := base64.RawURLEncoding.DecodeString(N)
	de, _ := base64.RawURLEncoding.DecodeString(E)
	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(dn),
		E: int(new(big.Int).SetBytes(de).Int64()),
	}
	// 検証するデータ
	message := sha256.Sum256([]byte(header + "." + payload))

	// 署名を base64 decode する
	sigData, err := base64.RawURLEncoding.DecodeString(sig)
	if err != nil {
		fmt.Println("署名を base64 decode error:", err)
		return
	}

	if err := rsa.VerifyPKCS1v15(pk, crypto.SHA256, message[:], sigData); err != nil {
		fmt.Println("invalid token")
	} else {
		fmt.Println("valid token")
		fmt.Println("header: ", string(headerData))
		fmt.Println("payload: ", string(payloadData))
	}
}
