package api

import (
	"encoding/json"
	"fmt"
	"my_crud/db"
	"my_crud/types"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (u *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	data, err := u.userStore.GetUsers()
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, fmt.Sprintf("Bad Getway GetUsersHandler %w", err))
	}
	json.NewEncoder(w).Encode(data)
}

func (u *UserHandler) GetUserDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры из URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Преобразуем id из строки в int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	data, err := u.userStore.GetUsersDetail(id)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, fmt.Sprintf("Bad Getway GetUsersDetail %w", err))
	}
	json.NewEncoder(w).Encode(data)
}

func (u *UserHandler) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {

	// Парсим JSON из тела запроса
	var userDetail *types.UserDetail
	err := json.NewDecoder(r.Body).Decode(&userDetail)
	if err != nil {
		http.Error(w, "Некорректный формат данных", http.StatusBadRequest)
		return
	}

	// Вызываем метод CreateUser через userStore
	id, err := u.userStore.CreateUser(userDetail)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка создания пользователя: %v", err), http.StatusInternalServerError)
		return
	}

	// Возвращаем ответ с ID нового пользователя
	userDetail.Id = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userDetail)
}

func (u *UserHandler) UpdateUsersHandler(w http.ResponseWriter, r *http.Request) {

	// Получаем параметры из URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Преобразуем id из строки в int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	// Парсим данные из тела запроса
	var userDetail types.UserDetail
	err = json.NewDecoder(r.Body).Decode(&userDetail)
	if err != nil {
		http.Error(w, fmt.Sprintf("Некорректный формат данных: %v", err), http.StatusBadRequest)
		return
	}

	// Обновляем пользователя
	updatedUser, err := u.userStore.UpdateUser(id, &userDetail)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка обновления пользователя: %v", err), http.StatusInternalServerError)
		return
	}

	// Возвращаем обновлённого пользователя
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

func (u *UserHandler) DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {

	// Получаем параметры из URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Преобразуем id из строки в int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	err = u.userStore.DeleteUser(id)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, fmt.Sprintf("Bad Getway DeleteUsersHandler %w", err))
	}
	json.NewEncoder(w).Encode(err)
}
