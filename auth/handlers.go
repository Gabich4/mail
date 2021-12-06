package main

import (
	"auth/utils"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	"net/http"
	"net/http/pprof"
	"time"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func login(w http.ResponseWriter, r *http.Request) {
	authorized := false
	username, password, ok := r.BasicAuth()
	if ok {
		if hash, ok := cfg.Users[username]; ok {
			authorized = utils.HashPasswordValid(hash, password)

		}
	}

	if !authorized {
		utils.Logger.Print("User unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	createTokenAndSetCookie(1*time.Minute, username, "accessToken", w)
	createTokenAndSetCookie(1*time.Hour, username, "refreshToken", w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("accessToken")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		tknStr := c.Value
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecretKey), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
			}
			w.WriteHeader(http.StatusBadRequest)
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}

	expirationTime := time.Now().Add(-1 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "accessToken",
		Value:   "",
		Expires: expirationTime,
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "refreshToken",
		Value:   "",
		Expires: expirationTime,
	})
}

func info(w http.ResponseWriter, r *http.Request) {
	tokenInfo := r.Context().Value("tokenInfo").(struct {
		token    *jwt.Token
		username string
	})
	username := tokenInfo.username

	if tokenInfo.token.Valid {
		createTokenAndSetCookie(1*time.Minute, username, "accessToken", w)
		createTokenAndSetCookie(1*time.Hour, username, "refreshToken", w)

		result := fmt.Sprintf("{\"username\": \"%s\"}", username)
		_, err := w.Write([]byte(result))
		if err != nil {
			utils.Logger.Error(err)
		}
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	return
}

// ---------------------------------------------------- MIDDLEWARE -----------------------------------------------------

func addContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ct := w.Header().Get("content-type"); ct == "" {
			w.Header().Set("content-type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}

func redirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		redirectIfNeeded(w, r)
	})
}

func checkAccessAndRefreshCookiesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tknStr := getTokenFromCookieWithoutResponse(r, "accessToken")
		if len(tknStr) != 0 {
			claims := &Claims{}
			tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.JWTSecretKey), nil
			})
			if err == nil && tkn.Valid {
				result := fmt.Sprintf("{\"username\": \"%s\"}", claims.Username)
				_, err := w.Write([]byte(result))
				if err != nil {
					utils.Logger.Error(err)
				}
				return
			}
			if err != nil && err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		tknStr = getTokenFromCookie(r, w, "refreshToken")
		if len(tknStr) == 0 {
			return
		}
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecretKey), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(context.TODO(), "tokenInfo", struct {
			token    *jwt.Token
			username string
		}{tkn, claims.Username})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func checkEnableDebug(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if enableDebug := r.URL.Query().Get("enable"); enableDebug == "true" {
			next.ServeHTTP(w, r)
		} else {
			err := fmt.Errorf("debug is disabled")
			utils.Logger.Print(err)
			http.Error(w, err.Error(), http.StatusNotAcceptable)
		}
	})
}

// ---------------------------------------------------- SUPPORTING -----------------------------------------------------
func createTokenAndSetCookie(tokenLifetime time.Duration, username, tokenName string, w http.ResponseWriter) {
	expirationTime := time.Now().Add(tokenLifetime)
	tokenString, err := createToken(expirationTime, username)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    tokenName,
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func createToken(expirationTime time.Time, username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecretKey))
}

func getTokenFromCookie(r *http.Request, w http.ResponseWriter, tokenName string) string {
	c, err := r.Cookie(tokenName)
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return ""
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return ""
	}

	// Get the JWT string from the cookie
	return c.Value
}

func getTokenFromCookieWithoutResponse(r *http.Request, tokenName string) string {
	c, err := r.Cookie(tokenName)
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return ""
		}
		// For any other type of error, return a bad request status
		return ""
	}

	// Get the JWT string from the cookie
	return c.Value
}

func redirectIfNeeded(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["redirect_uri"]

	if !ok || len(keys[0]) < 1 {
		utils.Logger.Println("Url Param 'redirect_uri' is missing")
		return
	}

	key := keys[0]
	http.Redirect(w, r, key, 301)

	utils.Logger.Println("Url Param 'redirect_uri' is: " + string(key))
}

// ----------------------------------------------------- PROFILER ------------------------------------------------------

// go tool pprof http://localhost:3000/debug/pprof/profile?seconds=5
// go tool pprof http://localhost:3000/debug/pprof/heap
func Profiler() http.Handler {
	r := chi.NewRouter()

	r.Use(checkEnableDebug)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI+"/pprof/", http.StatusMovedPermanently)
	})
	r.HandleFunc("/pprof", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI+"/", http.StatusMovedPermanently)
	})

	// Получение списка всех профилей
	r.HandleFunc("/pprof/*", pprof.Index)

	// Отображение строки запуска (например: /go-observability-course/examples/caching/redis/__debug_bin)
	r.HandleFunc("/pprof/cmdline", pprof.Cmdline)

	// профиль ЦПУ, в query-параметрах можно указать seconds со значением времени в секундах для снимка (по-умолчанию 30с)
	r.HandleFunc("/pprof/profile", pprof.Profile)
	r.HandleFunc("/pprof/symbol", pprof.Symbol)

	// профиль для получения трассировки (последовательности инструкций)
	// выполнения приложения за время seconds из query-параметров ( по-умолчанию 1с)
	r.HandleFunc("/pprof/trace", pprof.Trace)

	return r
}
