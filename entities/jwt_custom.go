package entities

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtCustomClaims struct {
	ID primitive.ObjectID `bson:"id"`
	jwt.RegisteredClaims
}
