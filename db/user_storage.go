package db

import (
	"database/sql"
	"fmt"
	"my_crud/types"
)

type UserStore interface {
	GetUsers() ([]*types.User, error)
	GetUsersDetail(id int) (*types.UserDetail, error)
	CreateUser(user *types.UserDetail) (int, error)
	UpdateUser(id int, user *types.UserDetail) (*types.UserDetail, error)
	DeleteUser(id int) error
}

type MySqlUserStore struct {
	db *sql.DB
}

func NewMysqlUserSotre(db *sql.DB) *MySqlUserStore {
	return &MySqlUserStore{
		db: db,
	}
}

func (s MySqlUserStore) GetUsers() ([]*types.User, error) {
	query := `SELECT u.id,u.name,u.msisdn FROM users u`
	rows, err := s.db.Query(query)
	if err != nil {
		fmt.Printf("Query erro %w", err)
	}
	defer rows.Close()

	users := []*types.User{}
	for rows.Next() {
		user := new(types.User)
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Msisdn,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s MySqlUserStore) GetUsersDetail(id int) (*types.UserDetail, error) {
	query := `SELECT u.id,u.name,u.msisdn, u.status_id, u.trpl_id, u.trpl_name, u.birthday 
	FROM users u WHERE u.id = ?`
	row := s.db.QueryRow(query, id)

	user := new(types.UserDetail)
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Msisdn,
		&user.StatusID,
		&user.Trpl_id,
		&user.Trpl_name,
		&user.Birthday,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (s *MySqlUserStore) CreateUser(user *types.UserDetail) (int, error) {
	query := `
        INSERT INTO users (name, msisdn, status_id, trpl_id, trpl_name, birthday)
        VALUES (?, ?, ?, ?, ?, ?)
    `
	_, err := s.db.Exec(query, user.Name, user.Msisdn, user.StatusID, user.Trpl_id, user.Trpl_name, user.Birthday)
	if err != nil {
		return 0, err
	}

	if err != nil {
		return 1, err
	} else {
		return 0, err
	}
}

func (s MySqlUserStore) UpdateUser(id int, user *types.UserDetail) (*types.UserDetail, error) {
	query := `
			UPDATE users 
			SET name = ?, msisdn = ?, status_id = ?, trpl_id = ?, trpl_name = ?, birthday = ?
			WHERE id = ?
		`

	// Выполняем запрос
	_, err := s.db.Exec(
		query,
		user.Name,
		user.Msisdn,
		user.StatusID,
		user.Trpl_id,
		user.Trpl_name,
		user.Birthday,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("Ошибка обновления пользователя: %w", err)
	}

	// Возвращаем обновлённого пользователя
	return s.GetUsersDetail(id)
}

func (s MySqlUserStore) DeleteUser(id int) error {
	query := `DELETE FROM  users  WHERE id = ?`

	// Выполняем запрос
	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления пользователя: %w", err)
	}

	// Проверяем, сколько строк было затронуто
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка проверки удаления: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("пользователь с id %d не найден", id)
	}

	fmt.Println("Пользователь удалён успешно")
	return nil
}
