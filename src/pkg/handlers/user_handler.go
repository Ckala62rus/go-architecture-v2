package handlers

import (
	"fmt"
	"net/http"
	"practice/domains"
	"practice/pkg/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUser struct {
	Name string `uri:"name" binding:"required"`
}

// Hello тестовый endpoint
// @Summary      Тестовый маршрут
// @Description  Возвращает тестовый JSON ответ с переданным именем. Используется для проверки работоспособности API.
// @Tags         Тестирование
// @Accept       json
// @Produce      json
// @Param        name path string true "Имя пользователя для приветствия"
// @Success      200  {object}  StatusResponse{data=string} "Успешный ответ"
// @Router       /hello/{name} [get]
func (h *Handler) Hello(c *gin.Context) {
	//ctx := context.Background()
	//res := utils.RedisDb.SetToken(ctx, "lorem ipsum dollar sit amet")
	//fmt.Println(res)

	//logger.MainLogger.Info("hello handler")
	h.log.Debug("#Debug hello handler")
	h.log.Info("#Info hello handler")
	fmt.Println("test")

	var getUser GetUser
	err := c.ShouldBindUri(&getUser)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSONP(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "Hello handler",
		Data:    getUser.Name,
	})
}

// GetAllUsers возвращает всех пользователей
// @Summary      Получить всех пользователей
// @Description  Возвращает список всех пользователей в системе. Требует авторизации.
// @Tags         Пользователи
// @Accept       json
// @Produce      json
// @Success      200  {object}  StatusResponse{data=[]dto.UserOutDTO} "Список пользователей"
// @Failure      401  {object}  ErrorResponse "Не авторизован или токен недействителен"
// @Failure      500  {object}  ErrorResponse "Внутренняя ошибка сервера"
// @Router       /users/ [get]
// @Security     BearerAuth
func (h *Handler) GetAllUsers(c *gin.Context) {
	users := h.services.Users.GetAllUsers()
	//var usersDTO []dto.AllUsersOutDTO

	//for _, user := range users {
	//	dtoUserMap := dto.AllUsersOutDTO{
	//		Id:        user.Id,
	//		Name:      user.Name,
	//		Email:     user.Email,
	//		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	//		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	//	}
	//	usersDTO = append(usersDTO, dtoUserMap)
	//}

	usersDTO := dto.MapAllUser(users)

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "all users",
		Data:    usersDTO,
	})
}

// GetUserByName получает пользователя по имени
// @Summary      Получить пользователя по имени
// @Description  Возвращает информацию о пользователе по его имени. Требует авторизации.
// @Tags         Пользователи
// @Accept       json
// @Produce      json
// @Param        name path string true "Имя пользователя для поиска"
// @Success      200  {object}  StatusResponse{data=dto.UserOutDTO} "Информация о пользователе"
// @Failure      401  {object}  ErrorResponse "Не авторизован или токен недействителен"
// @Failure      404  {object}  ErrorResponse "Пользователь не найден"
// @Failure      500  {object}  ErrorResponse "Внутренняя ошибка сервера"
// @Router       /users/user/{name} [get]
// @Security     BearerAuth
func (h *Handler) GetUserByName(c *gin.Context) {
	name := c.Param("name")
	user, err := h.services.Users.GetUserByName(name)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	userMap := dto.MapSingleUser(user)
	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "one user",
		Data:    userMap,
	})
}

// GetById получает пользователя по ID
// @Summary      Получить пользователя по ID
// @Description  Возвращает информацию о пользователе по его уникальному идентификатору. Требует авторизации.
// @Tags         Пользователи
// @Accept       json
// @Produce      json
// @Param        id path int true "Уникальный идентификатор пользователя"
// @Success      200  {object}  StatusResponse{data=dto.UserOutDTO} "Информация о пользователе"
// @Failure      400  {object}  ErrorResponse "Некорректный ID пользователя"
// @Failure      401  {object}  ErrorResponse "Не авторизован или токен недействителен"
// @Failure      404  {object}  ErrorResponse "Пользователь не найден"
// @Failure      500  {object}  ErrorResponse "Внутренняя ошибка сервера"
// @Router       /users/{id} [get]
// @Security     BearerAuth
func (h *Handler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	user, err := h.services.Users.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	userMap := dto.MapSingleUser(user)

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "one user",
		Data:    userMap,
	})
}

// DeleteUserById удаляет пользователя по ID
// @Summary      Удалить пользователя по ID
// @Description  Удаляет пользователя из системы по его уникальному идентификатору. Требует авторизации.
// @Tags         Пользователи
// @Accept       json
// @Produce      json
// @Param        id path int true "Уникальный идентификатор пользователя для удаления"
// @Success      200  {object}  StatusResponse{data=string} "Пользователь успешно удалён"
// @Failure      400  {object}  ErrorResponse "Некорректный ID пользователя"
// @Failure      401  {object}  ErrorResponse "Не авторизован или токен недействителен"
// @Failure      404  {object}  ErrorResponse "Пользователь не найден"
// @Failure      500  {object}  ErrorResponse "Внутренняя ошибка сервера"
// @Router       /users/{id} [delete]
// @Security     BearerAuth
func (h *Handler) DeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	isDelete, err := h.services.Users.DeleteUserById(id)
	if !isDelete {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: fmt.Sprintf("User was delete with id:%d", id),
	})
}

// UpdateUser обновляет данные пользователя
// @Summary      Обновить пользователя по ID
// @Description  Обновляет информацию о пользователе по его уникальному идентификатору. Требует авторизации.
// @Tags         Пользователи
// @Accept       json
// @Produce      json
// @Param        id path int true "Уникальный идентификатор пользователя"
// @Param        input body dto.UpdateUserInDTO true "Данные для обновления пользователя"
// @Success      200  {object}  StatusResponse{data=dto.UserOutDTO} "Пользователь успешно обновлён"
// @Failure      400  {object}  ErrorResponse "Некорректный ID пользователя или данные запроса"
// @Failure      401  {object}  ErrorResponse "Не авторизован или токен недействителен"
// @Failure      404  {object}  ErrorResponse "Пользователь не найден"
// @Failure      500  {object}  ErrorResponse "Внутренняя ошибка сервера"
// @Router       /users/{id} [put]
// @Security     BearerAuth
func (h *Handler) UpdateUser(c *gin.Context) {
	var user dto.UpdateUserInDTO

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	err = c.BindJSON(&user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	updateUser, err := h.services.Users.UpdateUser(
		domains.User{
			Id:   id,
			Name: user.Name,
		},
	)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status:  true,
		Message: "updated user",
		Data:    updateUser,
	})
}
