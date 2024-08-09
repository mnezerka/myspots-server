package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"mnezerka/myspots-server/bootstrap"
	"mnezerka/myspots-server/entities"
	"net/http"
	"strings"
)

func Authenticate(env *bootstrap.Env) gin.HandlerFunc {

	return func(c *gin.Context) {

		// try to get token string from authorization header

		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Bearer" {
			log.Warn().Msg("invalid or missing authorization header")
			c.JSON(http.StatusForbidden, gin.H{"error": "invalid or missing authorization header"})
			c.Abort()
			return
		}

		tokenString := auth[1]

		if len(tokenString) == 0 {
			log.Warn().Msg("no authorization header provided")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no authorization header provided"})
			c.Abort()
			return
		}

		claims := &entities.JwtCustomClaims{}

		// parse the JWT string and store the result in `claims`.
		// note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(env.AccessTokenSecret), nil
		})

		if err != nil {
			log.Warn().Msgf("failed to parse token with claims %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		} else if claims, ok := token.Claims.(*entities.JwtCustomClaims); ok {
			log.Info().Str("module", "Authenticate").Msgf("user successfully validated (%s)", claims.ID.Hex())

			c.Set("user-id", claims.ID.Hex())
			c.Next()
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"": "unknown jwt auth claims type, cannot proceed"})
			c.Abort()
			return
		}
	}
}

/*
func RequireAuth(c *gin.Context) {

    // try to get auth token from authorization header
    auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
    if len(auth) != 2 || auth[0] != "Bearer" {
        return nil, "invalid or missing authorization header"
    }
    tokenString := auth[1]

    // 2. we have token string, let's validate it

    // Initialize a new instance of `Claims`
    claims := &entities.JwtCustomClaims{}

    // Parse the JWT string and store the result in `claims`.
    // Note that we are passing the key in this method as well. This method will return an error
    // if the token is invalid (if it has expired according to the expiry time we set on sign in),
    // or if the signature does not match
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        // Don't forget to validate the alg is what you expect:
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])     ■ error strings should not be capitalized
        }
        return []byte(h.cfg.JwtPassword), nil
    })

    if err != nil {
        c.AbortWithStatus(http.StatusUnauthorized)
    }

	// Decode/validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Chec k the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token Subject
		var user entities.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach the request
		c.Set("user", user)

		//Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization Header Provided")})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Next()
	}
}
*/
