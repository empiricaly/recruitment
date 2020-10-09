package admin

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/empiricaly/recruitment/internal/ent"
	adminT "github.com/empiricaly/recruitment/internal/ent/admin"
	"github.com/empiricaly/recruitment/internal/storage"
	"github.com/o1egl/paseto/v2"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

// User will store the credentials of an admin.
type User struct {
	Name     string
	Username string
	Password string
}

func Init(ctx context.Context, admins []User, store *storage.Conn) error {
	for _, a := range admins {
		ad, err := store.Admin.Query().Where(adminT.UsernameEQ(a.Username)).First(ctx)
		if err != nil {
			_, err := store.Admin.Create().
				SetID(xid.New().String()).
				SetName(a.Name).
				SetUsername(a.Username).
				Save(ctx)
			if err != nil {
				return errors.Wrap(err, "creating admin")
			}
		} else {
			if ad.Name != a.Name {
				_, err := ad.Update().SetName(a.Name).Save(ctx)
				if err != nil {
					log.Error().Msgf("could not update name for %s", a.Username)
				}
			}
		}
	}

	return nil
}

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var adminCtxKey = struct{}{}

// AuthMiddleware decodes the share session cookie and packs the session into context
func Middleware(conn *storage.Conn, key []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if auth == "" {
				next.ServeHTTP(w, r)
				return
			}

			userID, err := validateAndGetUserID(key, auth)
			if err != nil {
				log.Error().Err(err).Msg("failed validate user for ID")
				http.Error(w, "Invalid auth", http.StatusForbidden)
				return
			}

			// Get the user from the database
			admin, err := getUserByID(r.Context(), conn, userID)
			if err != nil {
				log.Error().Err(err).Msg("failed get user for ID")
				http.Error(w, "Invalid auth", http.StatusForbidden)
				return
			}

			// Put it in context
			ctx := context.WithValue(r.Context(), adminCtxKey, admin)

			// And call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func getUserByID(ctx context.Context, conn *storage.Conn, userID string) (*ent.Admin, error) {
	fmt.Println("GET userID", userID)
	return conn.Admin.Get(ctx, userID)
}

// CreateUserIDToken ForContext finds the user from the context. REQUIRES Middleware to have run.
func CreateUserIDToken(key []byte, userID string) (string, error) {
	now := time.Now()
	exp := now.Add(90 * 24 * time.Hour)

	fmt.Println("ADD userID", userID)

	jsonToken := paseto.JSONToken{
		Audience:   "recruitment.empirica.app",
		Issuer:     "recruitment.empirica.app",
		Jti:        xid.New().String(),
		Subject:    userID,
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  now,
	}

	token, err := paseto.Encrypt(key, jsonToken, "")
	// token = "v2.local.E42A2iMY9SaZVzt-WkCi45_aebky4vbSUJsfG45OcanamwXwieieMjSjUkgsyZzlbYt82miN1xD-X0zEIhLK_RhWUPLZc9nC0shmkkkHS5Exj2zTpdNWhrC5KJRyUrI0cupc5qrctuREFLAvdCgwZBjh1QSgBX74V631fzl1IErGBgnt2LV1aij5W3hw9cXv4gtm_jSwsfee9HZcCE0sgUgAvklJCDO__8v_fTY7i_Regp5ZPa7h0X0m3yf0n4OXY9PRplunUpD9uEsXJ_MTF5gSFR3qE29eCHbJtRt0FFl81x-GCsQ9H9701TzEjGehCC6Bhw.c29tZSBmb290ZXI"

	return token, err
}

// validateAndGetUserID decrypts the user authentication token.
func validateAndGetUserID(key []byte, token string) (string, error) {
	t := strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))

	var jsonToken paseto.JSONToken
	err := paseto.Decrypt(t, key, &jsonToken, nil)
	if err != nil {
		return "", errors.Wrap(err, "paseto json key decrypt")
	}

	return jsonToken.Subject, nil
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *ent.Admin {
	raw, _ := ctx.Value(adminCtxKey).(*ent.Admin)
	return raw
}
