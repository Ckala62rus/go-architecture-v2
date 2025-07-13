package handlers

import (
	"context"
	"net/http"
	"practice/domains"
	"practice/pkg/dto"
	"practice/pkg/utils"

	"github.com/gin-gonic/gin"
)

// SignUp регистрирует нового пользователя в системе
// @Summary      Регистрация нового пользователя
// @Description  Создаёт нового пользователя с указанными email и паролем. Возвращает ID созданного пользователя.
// @Tags         Аутентификация
// @Accept       json
// @Produce      json
// @Param        input body dto.CreateAuthUser true "Данные для регистрации пользователя"
// @Success      200  {object}  StatusResponse{data=int} "Успешная регистрация"
// @Failure      400  {object}  ErrorResponse "Некорректные данные запроса"
// @Failure      500  {object}  ErrorResponse "Внутренняя ошибка сервера"
// @Router       /auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var input dto.CreateAuthUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(domains.User{
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "User success created",
		Data:    id,
	})
}

// SignIn авторизует пользователя в системе
// @Summary      Авторизация пользователя
// @Description  Авторизует пользователя по email и паролю. Возвращает JWT токен для доступа к защищённым маршрутам.
// @Tags         Аутентификация
// @Accept       json
// @Produce      json
// @Param        input body dto.SignInInput true "Данные для входа в систему"
// @Success      200  {object}  StatusResponse{data=object{token=string}} "Успешная авторизация"
// @Failure      400  {object}  ErrorResponse "Некорректные данные запроса"
// @Failure      401  {object}  ErrorResponse "Неверный email или пароль"
// @Failure      500  {object}  ErrorResponse "Внутренняя ошибка сервера"
// @Router       /auth/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	var input dto.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusOK, StatusResponse{
			Status:  false,
			Message: "authentication failed",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	ctx := context.Background()
	utils.RedisDb.SetToken(ctx, token)

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "authentication success!",
		Data: map[string]interface{}{
			"token": "Bearer " + token,
		},
	})
}

// Me возвращает информацию о текущем пользователе
// @Summary      Информация о текущем пользователе
// @Description  Возвращает информацию о пользователе на основе JWT токена. Требует авторизации.
// @Tags         Аутентификация
// @Accept       json
// @Produce      json
// @Success      200  {object}  StatusResponse{data=dto.UserOutDTO} "Информация о пользователе"
// @Failure      401  {object}  ErrorResponse "Не авторизован или токен недействителен"
// @Failure      404  {object}  ErrorResponse "Пользователь не найден"
// @Failure      500  {object}  ErrorResponse "Внутренняя ошибка сервера"
// @Router       /auth/me [get]
// @Security     BearerAuth
func (h *Handler) Me(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	user, err := h.services.Users.GetById(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userDTO := dto.MapSingleUser(user)

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "images was updated",
		Data:    userDTO,
	})
}

// Logout выход из системы
// @Summary      Выход из системы
// @Description  Выполняет выход пользователя из системы путём удаления JWT токена из Redis. Требует авторизации.
// @Tags         Аутентификация
// @Accept       json
// @Produce      json
// @Success      200  {object}  StatusResponse "Успешный выход"
// @Failure      401  {object}  ErrorResponse "Не авторизован или токен недействителен"
// @Failure      500  {object}  ErrorResponse "Внутренняя ошибка сервера"
// @Router       /auth/logout [post]
// @Security     BearerAuth
func (h *Handler) Logout(c *gin.Context) {
	token, err := getAuthenticationHeader(c)
	if err != nil {
		return
	}

	ctx := context.Background()
	err = utils.RedisDb.DeleteToken(ctx, token)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "Logout success",
		Data:    nil,
	})
}
