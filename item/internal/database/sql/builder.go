package sqlbuilder

import (
	"item/internal/models"
	"log"

	"github.com/Masterminds/squirrel"
)

func Create(req *models.CreateItemRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("items").
		Columns("username", "name", "type", "amount").
		Values(req.Username, req.Name, req.Type, req.Amount).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func Get(req *models.GetItemRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("items").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func Gets() (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("items").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func Update(req *models.GeneralItem) (string, []interface{}, error) {
	query, args, err := squirrel.Update("items").SetMap(map[string]interface{}{
		"username": req.Username,
		"name":     req.Name,
		"type":     req.Type,
		"amount":   req.Amount,
	}).
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func Delete(req *models.GetItemRequest) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("items").
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}
