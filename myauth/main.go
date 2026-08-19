package main

import (
	"context"
	"fmt"

	"google.golang.org/api/idtoken"
)

func main() {
	googleClientID := "482610659858-imfs79kl9cssqdc9a46c8ds2j0uane8f.apps.googleusercontent.com"
	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImYxMGY4NzQwNWE5NzljMWRmMzZkZjI2NjA2NzM0ZjMzY2Q4NWMyNzEiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI0ODI2MTA2NTk4NTgtaW1mczc5a2w5Y3NzcWRjOWE0NmM4ZHMyajB1YW5lOGYuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI0ODI2MTA2NTk4NTgtaW1mczc5a2w5Y3NzcWRjOWE0NmM4ZHMyajB1YW5lOGYuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTYxMDg1Njc2MzUxMTczMzE1MDciLCJub25jZSI6ImVzbiIsIm5iZiI6MTc4NzA0NTM0OSwibmFtZSI6IuS6rOiwt-WHjOaxsCIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS9BQ2c4b2NLSmNtZTZuZmtaektpZWtncEJzUk9RRTZ5cGVVM1FrV0NOSlozMkF4WkNZbGVTZW5nVD1zOTYtYyIsImdpdmVuX25hbWUiOiLlh4zmsbAiLCJmYW1pbHlfbmFtZSI6IuS6rOiwtyIsImlhdCI6MTc4NzA0NTY0OSwiZXhwIjoxNzg3MDQ5MjQ5LCJqdGkiOiJkYmIzZDU1OWNmODg0YTBiM2FjNWRhYmE2ZWFjOTA2YzRmZGRiY2VlIn0.GGfkeMHcoN8aCLMYFwfrz-PynbpILpgSBndTxcd3pJRh61xnIiPkFceaVl4XEuhLBCtgiJ7aZ2i4PtWUTw_s6trteH-X3fwPo1auizw7lk5jlwmp9WoXz-zOggWCE4UY8qtRrIk02QIpWE065j4h9XeiZvcd9mQUlpxzv8Sdr4P-stqWx5VkKNHANhG9wkwviLOV9QR-dnkUOoX9cK-5MbUyYCJerYlDFRslqPVwzA_pm9o-EIMk3im1fhhSIv3N1PrMlLQtxqGjz45n3YHAwPJWEUTgaOfJw6dEWIwExYXVEO8AfDnb4_l7VdFtp021UL98gPR1H_jAwpufVBaaEw"

	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		fmt.Println("failed to create token validator:", err)
		return
	}

	payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
	if err != nil {
		fmt.Println("failed to validate ID token:", err)
		return
	}

	fmt.Println("ID token is valid.")
	fmt.Println("Payload:", payload.Claims["name"])
}
