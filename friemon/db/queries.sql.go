// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const deleteEverything = `-- name: DeleteEverything :exec
TRUNCATE TABLE characters
`

func (q *Queries) DeleteEverything(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteEverything)
	return err
}

const createCharacter = `-- name: createCharacter :one
INSERT INTO characters (id, owner_id, claimed_timestamp, idx, character_id, level, xp, personality, shiny, iv_hp, iv_atk, iv_def, iv_sp_atk, iv_sp_def, iv_spd, iv_total, nickname, favourite, held_item, moves, color)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
RETURNING id, owner_id, claimed_timestamp, idx, character_id, level, xp, personality, shiny, iv_hp, iv_atk, iv_def, iv_sp_atk, iv_sp_def, iv_spd, iv_total, nickname, favourite, held_item, moves, color
`

type createCharacterParams struct {
	OwnerID          string    `json:"owner_id"`
	ClaimedTimestamp time.Time `json:"claimed_timestamp"`
	Idx              int32     `json:"idx"`
	CharacterID      int32     `json:"character_id"`
	Level            int32     `json:"level"`
	Xp               int32     `json:"xp"`
	Personality      string    `json:"personality"`
	Shiny            bool      `json:"shiny"`
	IvHp             int32     `json:"iv_hp"`
	IvAtk            int32     `json:"iv_atk"`
	IvDef            int32     `json:"iv_def"`
	IvSpAtk          int32     `json:"iv_sp_atk"`
	IvSpDef          int32     `json:"iv_sp_def"`
	IvSpd            int32     `json:"iv_spd"`
	IvTotal          float64   `json:"iv_total"`
	Nickname         string    `json:"nickname"`
	Favourite        bool      `json:"favourite"`
	HeldItem         int32     `json:"held_item"`
	Moves            []int32   `json:"moves"`
	Color            int32     `json:"color"`
}

func (q *Queries) createCharacter(ctx context.Context, arg createCharacterParams) (Character, error) {
	row := q.db.QueryRow(ctx, createCharacter,
		arg.OwnerID,
		arg.ClaimedTimestamp,
		arg.Idx,
		arg.CharacterID,
		arg.Level,
		arg.Xp,
		arg.Personality,
		arg.Shiny,
		arg.IvHp,
		arg.IvAtk,
		arg.IvDef,
		arg.IvSpAtk,
		arg.IvSpDef,
		arg.IvSpd,
		arg.IvTotal,
		arg.Nickname,
		arg.Favourite,
		arg.HeldItem,
		arg.Moves,
		arg.Color,
	)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.ClaimedTimestamp,
		&i.Idx,
		&i.CharacterID,
		&i.Level,
		&i.Xp,
		&i.Personality,
		&i.Shiny,
		&i.IvHp,
		&i.IvAtk,
		&i.IvDef,
		&i.IvSpAtk,
		&i.IvSpDef,
		&i.IvSpd,
		&i.IvTotal,
		&i.Nickname,
		&i.Favourite,
		&i.HeldItem,
		&i.Moves,
		&i.Color,
	)
	return i, err
}

const createUser = `-- name: createUser :one
INSERT INTO users (id) VALUES ($1) RETURNING id, balance, selected_id, order_by, order_desc, shinies_caught
`

func (q *Queries) createUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, createUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.SelectedID,
		&i.OrderBy,
		&i.OrderDesc,
		&i.ShiniesCaught,
	)
	return i, err
}

const deleteCharacter = `-- name: deleteCharacter :exec
DELETE FROM characters WHERE id = $1
`

func (q *Queries) deleteCharacter(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteCharacter, id)
	return err
}

const getCharacter = `-- name: getCharacter :one
SELECT id, owner_id, claimed_timestamp, idx, character_id, level, xp, personality, shiny, iv_hp, iv_atk, iv_def, iv_sp_atk, iv_sp_def, iv_spd, iv_total, nickname, favourite, held_item, moves, color FROM characters WHERE id = $1
`

func (q *Queries) getCharacter(ctx context.Context, id uuid.UUID) (Character, error) {
	row := q.db.QueryRow(ctx, getCharacter, id)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.ClaimedTimestamp,
		&i.Idx,
		&i.CharacterID,
		&i.Level,
		&i.Xp,
		&i.Personality,
		&i.Shiny,
		&i.IvHp,
		&i.IvAtk,
		&i.IvDef,
		&i.IvSpAtk,
		&i.IvSpDef,
		&i.IvSpd,
		&i.IvTotal,
		&i.Nickname,
		&i.Favourite,
		&i.HeldItem,
		&i.Moves,
		&i.Color,
	)
	return i, err
}

const getCharactersForUser = `-- name: getCharactersForUser :many
SELECT id, owner_id, claimed_timestamp, idx, character_id, level, xp, personality, shiny, iv_hp, iv_atk, iv_def, iv_sp_atk, iv_sp_def, iv_spd, iv_total, nickname, favourite, held_item, moves, color FROM characters WHERE owner_id = $1
`

func (q *Queries) getCharactersForUser(ctx context.Context, ownerID string) ([]Character, error) {
	rows, err := q.db.Query(ctx, getCharactersForUser, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Character
	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.ClaimedTimestamp,
			&i.Idx,
			&i.CharacterID,
			&i.Level,
			&i.Xp,
			&i.Personality,
			&i.Shiny,
			&i.IvHp,
			&i.IvAtk,
			&i.IvDef,
			&i.IvSpAtk,
			&i.IvSpDef,
			&i.IvSpd,
			&i.IvTotal,
			&i.Nickname,
			&i.Favourite,
			&i.HeldItem,
			&i.Moves,
			&i.Color,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: getUser :one
SELECT id, balance, selected_id, order_by, order_desc, shinies_caught FROM users WHERE id = $1
`

func (q *Queries) getUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.SelectedID,
		&i.OrderBy,
		&i.OrderDesc,
		&i.ShiniesCaught,
	)
	return i, err
}

const updateCharacter = `-- name: updateCharacter :one
UPDATE characters SET owner_id = $2, claimed_timestamp = $3, idx = $4, character_id = $5, level = $6, xp = $7, personality = $8, shiny = $9, iv_hp = $10, iv_atk = $11, iv_def = $12, iv_sp_atk = $13, iv_sp_def = $14, iv_spd = $15, iv_total = $16, nickname = $17, favourite = $18, held_item = $19, moves = $20, color = $21 WHERE id = $1 RETURNING id, owner_id, claimed_timestamp, idx, character_id, level, xp, personality, shiny, iv_hp, iv_atk, iv_def, iv_sp_atk, iv_sp_def, iv_spd, iv_total, nickname, favourite, held_item, moves, color
`

type updateCharacterParams struct {
	ID               uuid.UUID `json:"id"`
	OwnerID          string    `json:"owner_id"`
	ClaimedTimestamp time.Time `json:"claimed_timestamp"`
	Idx              int32     `json:"idx"`
	CharacterID      int32     `json:"character_id"`
	Level            int32     `json:"level"`
	Xp               int32     `json:"xp"`
	Personality      string    `json:"personality"`
	Shiny            bool      `json:"shiny"`
	IvHp             int32     `json:"iv_hp"`
	IvAtk            int32     `json:"iv_atk"`
	IvDef            int32     `json:"iv_def"`
	IvSpAtk          int32     `json:"iv_sp_atk"`
	IvSpDef          int32     `json:"iv_sp_def"`
	IvSpd            int32     `json:"iv_spd"`
	IvTotal          float64   `json:"iv_total"`
	Nickname         string    `json:"nickname"`
	Favourite        bool      `json:"favourite"`
	HeldItem         int32     `json:"held_item"`
	Moves            []int32   `json:"moves"`
	Color            int32     `json:"color"`
}

func (q *Queries) updateCharacter(ctx context.Context, arg updateCharacterParams) (Character, error) {
	row := q.db.QueryRow(ctx, updateCharacter,
		arg.ID,
		arg.OwnerID,
		arg.ClaimedTimestamp,
		arg.Idx,
		arg.CharacterID,
		arg.Level,
		arg.Xp,
		arg.Personality,
		arg.Shiny,
		arg.IvHp,
		arg.IvAtk,
		arg.IvDef,
		arg.IvSpAtk,
		arg.IvSpDef,
		arg.IvSpd,
		arg.IvTotal,
		arg.Nickname,
		arg.Favourite,
		arg.HeldItem,
		arg.Moves,
		arg.Color,
	)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.ClaimedTimestamp,
		&i.Idx,
		&i.CharacterID,
		&i.Level,
		&i.Xp,
		&i.Personality,
		&i.Shiny,
		&i.IvHp,
		&i.IvAtk,
		&i.IvDef,
		&i.IvSpAtk,
		&i.IvSpDef,
		&i.IvSpd,
		&i.IvTotal,
		&i.Nickname,
		&i.Favourite,
		&i.HeldItem,
		&i.Moves,
		&i.Color,
	)
	return i, err
}
