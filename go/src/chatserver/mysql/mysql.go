/*
 * @Author: dzw
 * @Date: 2020-03-15 13:45:00
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-15 14:03:19
 */

package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

// InitDB init db
func InitDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// 尝试与数据库建立链接，检验dataSourceName的正确性
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// QueryRows ...
func QueryRows(db *sql.DB, dest interface{}, query string, args ...interface{}) error {
	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	// 关闭 rows 释放持有的数据库链接
	defer rows.Close()

	// 判断dest类型
	value := reflect.ValueOf(dest)
	if value.Kind() != reflect.Ptr {
		return errors.New("must pass a point, not a value")
	}

	if value.IsNil() {
		return errors.New("nil pointer passed")
	}

	if value.Elem().Kind() != reflect.Slice {
		return errors.New("type is not slice")
	}

	// return value point to
	direct := reflect.Indirect(value)

	// slice类型
	slice := value.Type().Elem()
	// 获取slice中元素的类型
	base := slice.Elem()
	// 查看slice中元素是否是Ptr
	isPtr := slice.Elem().Kind() == reflect.Ptr
	if isPtr {
		base = slice.Elem().Elem()
	}

	//
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	// isScanable
	isScanable := false
	if base.Kind() != reflect.Struct { // not struct
		isScanable = true
	}

	// 不是结构体，但是查的字段大于1
	if isScanable && len(columns) > 1 {
		return fmt.Errorf("non-struct dest type %s with > 1 columns (%d)", base.Kind(), len(columns))
	}

	// 循环读取结果
	if isScanable {
		for rows.Next() {
			vp := reflect.New(base)
			err := rows.Scan(vp.Interface())
			if err != nil {
				return err
			}

			if isPtr {
				direct.Set(reflect.Append(direct, vp))
			} else {
				direct.Set(reflect.Append(direct, reflect.Indirect(vp)))
			}
		}
	} else {
		for rows.Next() {
			vp := reflect.New(base)
			v := reflect.Indirect(vp)

			num := len(columns)
			values := make([]interface{}, num)
			for i := 0; i < num; i++ {
				values[i] = v.Field(i).Addr().Interface()
			}

			err := rows.Scan(values...)
			if err != nil {
				return err
			}

			if isPtr {
				direct.Set(reflect.Append(direct, vp))
			} else {
				direct.Set(reflect.Append(direct, v))
			}
		}
	}

	return nil
}

// InsertInfo insert info
func InsertInfo(db *sql.DB, query string, args ...interface{}) error {
	ret, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	_, err = ret.LastInsertId() // 新插入数据的Id
	if err != nil {
		return err
	}

	return nil
}

// UpdateInfo update or delete info
func UpdateInfo(db *sql.DB, query string, args ...interface{}) error {
	ret, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	_, err = ret.RowsAffected() // 操作影响的行数
	if err != nil {
		return err
	}

	return nil
}
