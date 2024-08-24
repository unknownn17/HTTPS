package methods

import (
	"context"
	"database/sql"
	"fmt"
	sqlbuilder "item/internal/database/sql"
	"item/internal/models"
	"log"
)

type Database struct {
	Db *sql.DB
}

func (u *Database) ItemCreate(ctx context.Context, req *models.CreateItemRequest) (*models.DeleteResponse, error) {
	query, args, err := sqlbuilder.Create(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var id int
	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.DeleteResponse{Message: fmt.Sprintf("item has been created with this id %v", id)}, nil
}

func (u *Database) ItemGet(ctx context.Context, req *models.GetItemRequest) (*models.GeneralItem, error) {
	query, args, err := sqlbuilder.Get(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var res models.GeneralItem

	if err := u.Db.QueryRow(query, args...).Scan(&res.ID, &res.Username, &res.Name, &res.Type, &res.Amount); err != nil {
		log.Println(err)
		return nil, err
	}
	return &res, nil
}

func (u *Database) ItemGets(ctx context.Context, req *models.GetItemsRequest) ([]*models.GeneralItem, error) {
	query, args, err := sqlbuilder.Gets()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err := u.Db.Query(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var res []*models.GeneralItem
	for rows.Next() {
		var all models.GeneralItem
		if err := rows.Scan(&all.ID, &all.Username, &all.Name, &all.Type, &all.Amount); err != nil {
			log.Println(err)
			return nil, err
		}
		res = append(res, &all)
	}
	return res, nil
}

func (u *Database) ItemUpdate(ctx context.Context, req *models.GeneralItem) (*models.DeleteResponse, error) {
	query, args, err := sqlbuilder.Update(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var id int
	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.DeleteResponse{Message: fmt.Sprintf("Item has been updated with this id %v", id)}, nil
}

func (u *Database) ItemDelete(ctx context.Context, req *models.GetItemRequest) (*models.DeleteResponse, error) {
	query, args, err := sqlbuilder.Delete(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.DeleteResponse{Message: fmt.Sprintf("Item has been deleted with this id %v", req.ID)}, nil
}
