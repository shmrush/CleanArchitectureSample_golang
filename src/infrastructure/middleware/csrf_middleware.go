package middleware

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	sr "github.com/tuvistavie/securerandom"
)

const csrfTokenLength = 64

var context *gin.Context
var session sessions.Session

func CSRFProtection(ctx *gin.Context) {
	context = ctx
	session = sessions.Default(ctx)
	if !isVerifiedRequest() {
		ctx.AbortWithStatusJSON(500, gin.H{"status": "Internal server error", "message": "Invalid CSRF token."})
	} else {
		ctx.SetCookie("_csrf_token", maskedToken(), 3600, "/", "localhost", false, false)
		ctx.Next()
	}
}

func isVerifiedRequest() bool {
	method := context.Request.Method
	return method == http.MethodGet || method == http.MethodHead || (isCSRFTokenValid() && isRequestOriginValid())
}

func isRequestOriginValid() bool {
	origin := context.Request.Header.Get("Origin")
	return origin == "" || origin == context.Request.RequestURI
}

func isCSRFTokenValid() bool {
	encodedMaskedToken := context.Request.Header.Get("X-CSRF-Token")
	if encodedMaskedToken == "" {
		return false
	}
	maskedToken, err := base64.StdEncoding.DecodeString(encodedMaskedToken)
	if err != nil {
		return false
	}
	if len(maskedToken) == csrfTokenLength*2 {
		return unmaskToken(maskedToken) == realCSRFToken()
	}
	return false
}

func realCSRFToken() string {
	token := session.Get("realCSRFToken")
	if token == nil {
		newToken, err := sr.RandomBytes(csrfTokenLength)
		if err != nil {
			log.Fatal("Failed to create new realCSRFtoken.")
		}
		encodedNewToken := base64.StdEncoding.EncodeToString(newToken)
		session.Set("realCSRFToken", encodedNewToken)
		session.Save()
		return encodedNewToken
	}
	return token.(string)
}

func maskedToken() string {
	oneTimePad, err := sr.RandomBytes(csrfTokenLength)
	if err != nil {
		log.Fatal("Failed to create new oneTimePad.")
	}
	encryptedCSRFToken := []byte{}
	decodedRealCSRFToken, err := base64.StdEncoding.DecodeString(realCSRFToken())
	if err != nil {
		log.Fatal("Failed to create new oneTimePad.")
	}
	for i, p := range oneTimePad {
		encryptedCSRFToken = append(encryptedCSRFToken, p^decodedRealCSRFToken[i])
	}
	maskedCSRFToken := append(oneTimePad, encryptedCSRFToken...)
	return base64.StdEncoding.EncodeToString(maskedCSRFToken)
}

func unmaskToken(maskedToken []byte) string {
	oneTimePad := maskedToken[:csrfTokenLength]
	encryptedCSRFToken := maskedToken[csrfTokenLength:]
	unmaskedToken := []byte{}
	for i, p := range oneTimePad {
		unmaskedToken = append(unmaskedToken, p^encryptedCSRFToken[i])
	}
	return base64.StdEncoding.EncodeToString(unmaskedToken)
}
