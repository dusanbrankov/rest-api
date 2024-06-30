package user

import (
	"net/http"

	"github.com/dusanbrankov/rest-api/db"
	"github.com/dusanbrankov/rest-api/types"
	_ "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/users", h.ListUsersHandler)
}

func (h *Handler) ListUsersHandler(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := db.GetDBConfig().Queries.ListUsers(ctx)
	if err != nil {
		return err
	}

	// Map database users to the custom response struct
	// var userResponses []db.User
	// for _, user := range users {
	// 	userResponses = append(userResponses, db.User{
	// 		ID:   user.ID,
	// 		Name: user.Name,
	// 		Age:  user.Age.Int16, // Extract the integer value from sql.NullInt16
	// 	})
	// }

	return c.JSON(http.StatusOK, users)
}

// func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
// 	// get JSON payload
// 	var payload types.LoginUserPayload
// 	if err := utils.ParseJSON(r, &payload); err != nil {
// 		utils.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}
//
// 	// validate the payload
// 	if err := utils.Validate.Struct(payload); err != nil {
// 		errors := err.(validator.ValidationErrors)
// 		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
// 		return
// 	}
//
// 	u, err := h.store.GetUserByEmail(payload.Email)
// 	if err != nil {
// 		log.Print(err)
// 		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("not found, invalid email or password"))
// 		return
// 	}
//
// 	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
// 		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("not found, invalid email or password"))
// 		return
// 	}
//
// 	secret := []byte(config.Envs.JWTSecret)
// 	token, err := auth.CreateJWT(secret, u.ID)
// 	if err != nil {
// 		utils.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}
//
// 	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
// }
//
// func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
// 	// get JSON payload
// 	var payload types.RegisterUserPayload
// 	if err := utils.ParseJSON(r, &payload); err != nil {
// 		utils.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}
//
// 	// validate the payload
// 	if err := utils.Validate.Struct(payload); err != nil {
// 		errors := err.(validator.ValidationErrors)
// 		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
// 		return
// 	}
//
// 	// check if the user exists
// 	_, err := h.store.GetUserByEmail(payload.Email)
// 	if err == nil {
// 		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user with email %s already exists", payload.Email))
// 		return
// 	}
//
// 	hashedPassword, err := auth.HashPassword(payload.Password)
// 	if err != nil {
// 		utils.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}
//
// 	// if user doesn't exist create new user
// 	err = h.store.CreateUser(types.User{
// 		FirstName: payload.FirstName,
// 		LastName: payload.LastName,
// 		Email: payload.Email,
// 		Password: hashedPassword,
// 	})
// 	if err != nil {
// 		utils.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}
//
// 	utils.WriteJSON(w, http.StatusCreated, nil)
// }

