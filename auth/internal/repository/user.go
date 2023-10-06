package repository

import (
	"auth/internal/core"
	"auth/internal/database"
	"auth/internal/models"
	"context"
	"log"
)

type userRepo struct {
	pg *database.PG
}

// DeleteByEmail implements core.UserRepo.
func (ur *userRepo) DeleteByEmail(ctx context.Context, email string) error {
	_, err := ur.pg.DB.ExecContext(ctx, DELETE_USER_BY_EMAIL_QUERY, email)
	if err != nil {
		log.Printf("[repository layer] | error while deleting user with email = %v : %v \n", email, err)
		return err
	}
	return nil
}

// DeleteByID implements core.UserRepo.
func (ur *userRepo) DeleteByID(ctx context.Context, id int64) error {
	_, err := ur.pg.DB.ExecContext(ctx, DELETE_USER_QUERY, id)
	if err != nil {
		log.Printf("[repository layer] | error while deleting user with id = %v : %v \n", id, err)
		return err
	}
	return nil
}

// GetByEmail implements core.UserRepo.
func (ur *userRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	u := new(models.User)
	err := ur.pg.DB.QueryRowContext(ctx, GET_BY_EMAIL_QUERY,
		email,
	).Scan(
		&u.ID,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
	)
	if err != nil {
		log.Printf("[repository layer] | error while retrieving user with email = %v : %v \n", email, err)
		return nil, err
	}

	return u, nil
}

// GetByID implements core.UserRepo.
func (ur *userRepo) GetByID(ctx context.Context, id int64) (*models.User, error) {
	u := new(models.User)
	err := ur.pg.DB.QueryRowContext(ctx, GET_BY_ID_QUERY,
		id,
	).Scan(
		&u.ID,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
	)
	if err != nil {
		log.Printf("[repository layer] | error while retrieving user with id = %v : %v \n", id, err)
		return nil, err
	}
	return u, nil
}

// Insert implements core.UserRepo.
func (ur *userRepo) Insert(ctx context.Context, user *models.User) (*models.User, error) {
	u := new(models.User)
	err := ur.pg.DB.QueryRowContext(ctx, INSERT_USER_QUERY,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&u.ID,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.Email,
	)
	if err != nil {
		log.Printf("[repository layer] | error while inserting new user : %v \n", err)
		return nil, err
	}

	return u, nil
}

// Update implements core.UserRepo.
func (ur *userRepo) Update(ctx context.Context, user *models.User) (*models.User, error) {
	u := new(models.User)
	// the service layer already fetched the user and knew that this user exists
	// so the recieved user is filled with all the fields, and we can start updating now
	err := ur.pg.DB.QueryRowContext(ctx, UPDATE_USER_QUERY,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	).Scan(
		&u.ID,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.Email,
	)
	if err != nil {
		log.Printf("[repository layer] | error while updating user with id = %v : %v \n", user.ID, err)
		return nil, err
	}

	return u, nil
}

func New(pg database.DBTX) core.UserRepo {
	return &userRepo{
		pg: &database.PG{
			DB: pg,
		},
	}
}

const (
	INSERT_USER_QUERY = `
		INSERT INTO users 
		(first_name, last_name, username, email, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING 
		id, username, first_name, last_name, email
	`

	GET_BY_EMAIL_QUERY = `
		SELECT id, username, first_name, last_name, email, password
		FROM users 
		WHERE email = $1
	`

	GET_BY_ID_QUERY = `
		SELECT id, username, first_name, last_name, email, password
		FROM users 
		WHERE id = $1
	`

	UPDATE_USER_QUERY = `
		UPDATE users 
		SET username = $1, first_name = $2, last_name = $3, email = $4, password = $5
		WHERE id = $6
		RETURNING 
		id, username, first_name, last_name, email
	`

	DELETE_USER_QUERY = `
		DELETE FROM users 
		WHERE id = $1
	`

	DELETE_USER_BY_EMAIL_QUERY = `
		DELETE FROM users 
		WHERE email = $1
	`
)
