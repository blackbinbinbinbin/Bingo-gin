package data

import (
	"context"
	"fmt"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/ent"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/ent/user"
)


func CreateUser(ctx context.Context, client *ent.Client, age int, name string) (*ent.User, error) {
	u, err := client.User.Create().
		SetAge(age).SetName(name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return u, nil
}


func UpdateUserById(ctx context.Context, client *ent.Client, id int, age int, name string) (*ent.User, error) {
	u , err := client.User.
		UpdateOneID(id).
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}
	return u, nil
}



func QueryUserById(ctx context.Context, client *ent.Client, id int) ([]*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.ID(id)).
	// `Only` fails if no user found,
	// or more than 1 user returned.
	// `All` return all data
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return u, nil
}