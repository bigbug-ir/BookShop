package util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt/v5"
)

/*****************************************************************/

// create private_key.pem
func Token() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	err = os.WriteFile("private_key.pem", privateKeyPEM, 0644)
	if err != nil {
		panic(err)
	}
	println("Successfull ECDSA private key generate.")
}

/*****************************************************************/

// genrate jwt token
func GenerateTokenJwt(c *fiber.Ctx, id uint) string {
	privateKeyPEM, err := ioutil.ReadFile("private_key.pem")
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error loading private key"})
	}
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "EC PRIVATE KEY" {
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid generate token"})
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid generate token"})
	}
	return tokenString
}

/*****************************************************************/
