package services

import (
	"context"
	"errors"
	"fmt"
	"key-value-system/db"
	"key-value-system/enums"
	"key-value-system/helper"
	"key-value-system/models"
	"key-value-system/requests"
	"strings"
)

func StoreHead(request requests.CreateHeadRequest) error {
	_, err := db.DB.Exec(db.GetSql(db.INSERT_HEAD_SQL), request.Key)
	return err
}

func StoreNode(request requests.CreateNodeRequest, ctx context.Context) error {
	tx, err := db.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = db.DB.Exec(db.GetSql(db.INSERT_NODE_SQL), request.Key, request.Value)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec(db.GetSql(db.UPDATE_NODE_SQL), request.Key, request.Prev)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func GetHead(key string) (*models.Head, error) {
	row := db.DB.QueryRow(db.GetSql(db.GET_HEAD_SQL), key)

	var head models.Head
	err := row.Scan(&head.Key, &head.Next, &head.CreatedAt, &head.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &head, nil
}

func GetNode(key string) (*models.Node, error) {
	row := db.DB.QueryRow(db.GetSql(db.GET_NODE_SQL), key)

	var node models.Node
	err := row.Scan(&node.Key, &node.Value, &node.Next, &node.CreatedAt, &node.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func RemoveHead(key string) error {
	head, err := GetHead(key)
	if head == nil {
		return errors.New(enums.NOT_FOUND)
	}
	if err != nil {
		return err
	}
	if head.Next == nil {
		return nil
	}

	keys, err := GetAllKeys(head)
	if err != nil {
		return err
	}
	keysInterface := helper.ToInterface(keys)

	// Delete all nodes at once
	result, err := db.DB.Exec(GetDynamicParametersSQL(db.DELETE_NODES_SQL, keys), keysInterface...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected != int64(len(keys)) {
		return errors.New(enums.DELETE_FAILED)
	}

	return nil
}

/*
 * In order to give dynamic parameters, do string substitution. For example:
 * transfer from `DELETE FROM table WHERE key IN ($1)`
 * to 			`DELETE FROM table WHERE key IN ($1,$2...)` (accords to keys' length)
 */
func GetDynamicParametersSQL(originalSql string, keys []string) string {
	var parameters []string
	for index := range keys {
		parameters = append(parameters, fmt.Sprintf("$%d", index+1))
	}

	sql := db.GetSql(originalSql)
	sql = strings.Replace(sql, "$1", "%s", 1)
	sql = fmt.Sprintf(sql, strings.Join(parameters, ","))

	return sql
}

/*
 * Get all keys by the given head (except for head's)
 */
func GetAllKeys(head *models.Head) ([]string, error) {
	var keys []string

	nextNode, err := GetNode(*head.Next)
	if err != nil {
		return keys, err
	}

	for nextNode != nil {
		keys = append(keys, nextNode.Key)

		if nextNode.Next == nil {
			break
		}

		nextNode, err = GetNode(*nextNode.Next)
		if err != nil {
			return keys, err
		}
	}

	return keys, nil
}

func RemoveNode(key string) error {
	node, err := GetNode(key)
	if node == nil {
		return errors.New(enums.NOT_FOUND)
	}
	if err != nil {
		return err
	}

	result, err := db.DB.Exec(db.GetSql(db.DELETE_NODE_SQL), key)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return errors.New(enums.DELETE_FAILED)
	}

	return nil
}
