package postgres

import (
	"context"
	"errors"
	"fmt"

	"user-service/internal/postgres/database"
	"user-service/internal/users"

	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	db *database.Queries
}

func NewUserRepository(db *database.Queries) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) FindByID(ctx context.Context, id int64) (*users.User, error) {
	u, err := repository.db.FindUser(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, users.ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return &users.User{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
	}, nil
}

func (repository *UserRepository) CountByUsername(ctx context.Context, username string) (int64, error) {
	count, err := repository.db.CountUsersByUsername(ctx, username)
	if err != nil {
		return 0, fmt.Errorf("failed to count users by username: %w", err)
	}
	return count, nil
}

func (repository *UserRepository) CountByEmail(ctx context.Context, email string) (int64, error) {
	count, err := repository.db.CountUsersByEmail(ctx, email)
	if err != nil {
		return 0, fmt.Errorf("failed to count users by email: %w", err)
	}
	return count, nil
}

func (repository *UserRepository) Save(ctx context.Context, user *users.User) error {
	var u database.User
	var err error

	if user.ID == 0 {
		u, err = repository.db.CreateUser(ctx, database.CreateUserParams{
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
		})
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		user.ID = u.ID
	} else {
		u, err = repository.db.UpdateUser(ctx, database.UpdateUserParams{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
		})
		if err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}
	}

	user.Username = u.Username
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Email = u.Email
	user.Phone = u.Phone

	return nil
}

func (repository *UserRepository) Delete(ctx context.Context, id int64) error {
	err := repository.db.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
