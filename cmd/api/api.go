package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/dusanbrankov/rest-api/config"
	"github.com/dusanbrankov/rest-api/service/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr, db}
}

func (s *APIServer) Run() error {
	e := echo.New()

	// productStore := product.NewStore(s.db)
	// productHandler := product.NewHandler(productStore)
	// productHandler.RegisterRoutes(subrouter)
	//
	// log.Println("Listening on", s.addr)

	// return http.ListenAndServe(s.addr, router)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		HSTSMaxAge: 3600,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "form:_csrf",
		CookiePath:     "/",
		CookieDomain:   config.GetAppConfig().PublicHost,
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(e)

	// Static assets
	e.Static("/static", "static")

	srv := &http.Server{
		Addr:         s.addr,
		Handler:      e,
		ReadTimeout:  5*time.Second,
		WriteTimeout: 10*time.Second,
	}

	return srv.ListenAndServe()
}

func Routes() http.Handler {
	e := echo.New()

	// Read the session key from the environment variable
	// encodedKey := os.Getenv("SESSION_KEY")
	// if encodedKey == "" {
	// 	log.Fatal("SESSION_KEY environment variable is not set")
	// }

	// Decode the base64 key
	// key, err := base64.StdEncoding.DecodeString(encodedKey)
	// if err != nil {
	// 	log.Fatalf("error decoding the key: %v", err)
	// }

	// Create a new cookie store with the decoded key
	// store := sessions.NewCookieStore(key)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		HSTSMaxAge: 3600,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "form:_csrf",
		CookiePath:     "/",
		CookieDomain:   config.GetAppConfig().PublicHost,
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))

	// Routes
	// e.GET("/", handlers.Home)
	// e.GET("/users", handlers.ListUsersHandler)
	// e.GET("/users/add", handlers.CreateUserHandler)

	// Static assets
	e.Static("/static", "static")

	return e
}

