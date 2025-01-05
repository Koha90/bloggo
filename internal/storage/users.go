package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"

	"github.com/koha90/bloggo/internal/types"
)

// UserStore - interface of users store.
type UserStorage interface {
	CreateUser(*types.User) error
	Users() ([]*types.User, error)
	UserByID(id uint) (*types.User, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, updates map[string]interface{}) error
}

var ErrUserAlreadyExists = errors.New("user with this username already exists")

// CreateUser - creates a user.
func (s *Storage) CreateUser(user *types.User) error {
	const op = "storage.CreateUser"

	query := `INSERT INTO users 
  (username, first_name, last_name, role,
  encrypted_password, created_at, updated_at)
  VALUES ($1, $2, $3, $4, $5, $6, $7);`
	_, err := s.db.Exec(
		query,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Role,
		user.EncryptedPassword,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		// проверяем ошибку уникальным ограничением
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.Code == sqlite3.ErrConstraint &&
			sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return fmt.Errorf("%s: %w", op, ErrUserAlreadyExists)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// UserByID - return one user by his id.
func (s *Storage) UserByID(id uint) (*types.User, error) {
	const op = "storage.UserByID"
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	if rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("%s: user not found %w", op, err)
}

// utilitars functions
func scanIntoUser(rows *sql.Rows) (*types.User, error) {
	const op = "storage.scanIntoUser"

	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.EncryptedPassword,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
