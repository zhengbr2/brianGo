package myorm

/*
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT ,
  `age` smallint(10) unsigned NOT NULL DEFAULT 0 ,
  `first_name` varchar(45) NOT NULL DEFAULT '' ,
  `last_name` varchar(45) NOT NULL DEFAULT '' ,
  `email` varchar(45) NOT NULL DEFAULT '' ,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
  PRIMARY KEY (`id`),
  KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='user table';
*/

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"strings"
	"time"
)

func Connect(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//设置连接池
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(10 * time.Minute)
	return conn, conn.Ping()
}

type Query struct {
	db    *sql.DB
	table string
}

//Table bind db and table
func Table(db *sql.DB, tableName string) func() *Query {
	return func() *Query {
		return &Query{
			db:    db,
			table: tableName,
		}
	}
}

//Insert in can be *User, []*User, map[string]interface{}
func (q *Query) Insert(in interface{}) (int64, error) {
	var keys, values []string
	v := reflect.ValueOf(in)
	//剥离指针
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct:
		keys, values = sKV(v)
	case reflect.Map:
		keys, values = mKV(v)
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			//Kind是切片时，可以用Index()方法遍历
			sv := v.Index(i)
			for sv.Kind() == reflect.Ptr || sv.Kind() == reflect.Interface {
				sv = sv.Elem()
			}
			//切片元素不是struct或者指针，报错
			if sv.Kind() != reflect.Struct {
				return 0, errors.New("method Insert error: in slice is not structs")
			}
			//keys只保存一次就行，因为后面的都一样了
			if len(keys) == 0 {
				keys, values = sKV(sv)
				continue
			}
			_, val := sKV(sv)
			values = append(values, val...)
		}
	default:
		return 0, errors.New("method Insert error: type error")
	}

	kl := len(keys)
	vl := len(values)
	if kl == 0 || vl == 0 {
		return 0, errors.New("method Insert error: no data")
	}
	var insertValue string
	//插入多条记录时需要用","拼接一下values
	if kl < vl {
		var tmpValues []string
		for kl <= vl {
			if kl%(len(keys)) == 0 {
				tmpValues = append(tmpValues, fmt.Sprintf("(%s)", strings.Join(values[kl-len(keys):kl], ",")))
			}
			kl++
		}
		insertValue = strings.Join(tmpValues, ",")
	} else {
		insertValue = fmt.Sprintf("(%s)", strings.Join(values, ","))
	}
	query := fmt.Sprintf(`insert into %s (%s) values %s`, q.table, strings.Join(keys, ","), insertValue)
	log.Printf("insert sql: %s", query)
	st, err := q.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	result, err := st.Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func sKV(v reflect.Value) ([]string, []string) {
	var keys, values []string
	t := v.Type()
	for n := 0; n < t.NumField(); n++ {
		tf := t.Field(n)
		vf := v.Field(n)
		//忽略非导出字段
		if tf.Anonymous {
			continue
		}
		//忽略无效、零值字段
		if !vf.IsValid() || reflect.DeepEqual(vf.Interface(), reflect.Zero(vf.Type()).Interface()) {
			continue
		}
		for vf.Type().Kind() == reflect.Ptr {
			vf = vf.Elem()
		}
		//有时候根据需求会组合struct，这里处理下，支持获取嵌套的struct tag和value
		//如果字段值是time类型之外的struct，递归获取keys和values
		if vf.Kind() == reflect.Struct && tf.Type.Name() != "Time" {
			cKeys, cValues := sKV(vf)
			keys = append(keys, cKeys...)
			values = append(values, cValues...)
			continue
		}
		//根据字段的json tag获取key，忽略无tag字段
		key := strings.Split(tf.Tag.Get("json"), ",")[0]
		if key == "" {
			continue
		}
		value := format(vf)
		if value != "" {
			keys = append(keys, key)
			values = append(values, value)
		}
	}
	return keys, values
}

func mKV(v reflect.Value) ([]string, []string) {
	var keys, values []string
	//获取map的key组成的切片
	mapKeys := v.MapKeys()
	for _, key := range mapKeys {
		value := format(v.MapIndex(key))
		if value != "" {
			values = append(values, value)
			keys = append(keys, key.Interface().(string))
		}
	}
	return keys, values
}

func format(v reflect.Value) string {
	//断言出time类型直接转unix时间戳
	if t, ok := v.Interface().(time.Time); ok {
		return fmt.Sprintf("FROM_UNIXTIME(%d)", t.Unix())
	}
	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf(`'%s'`, v.Interface())
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return fmt.Sprintf(`%d`, v.Interface())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf(`%f`, v.Interface())
		//如果是切片类型，遍历元素，递归格式化成"(, , , )"形式
	case reflect.Slice:
		var values []string
		for i := 0; i < v.Len(); i++ {
			values = append(values, format(v.Index(i)))
		}
		return fmt.Sprintf(`(%s)`, strings.Join(values, ","))
		//接口类型剥一层递归
	case reflect.Interface:
		return format(v.Elem())
	}
	return ""
}
