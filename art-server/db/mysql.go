package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/humgal/art-server/util"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", util.Config.Mysqlhost)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DB = db
}

func CloseDB() error {
	return DB.Close()
}

func Migrate() {
	if err := DB.Ping(); err != nil {
		util.Logger.Fatal(err)
	}
	driver, _ := mysql.WithInstance(DB, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations/mysql",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		util.Logger.Fatal(err)
	}

}

func GenInsertSql(s interface{}, tablename string) string {

	sqlstr := "insert into " + tablename + "($field) values" + "($value)"
	fileds := ""
	values := ""

	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return ""
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// only accept struct param
	if v.Kind() != reflect.Struct {
		return ""
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf(CamelCaseToUdnderscore(t.Field(i).Name))
		fmt.Printf(":")
		fmt.Printf("%v", v.Field(i).Interface())
		fmt.Println("")
		switch v.Field(i).Kind() {
		case reflect.Int:
			fileds += CamelCaseToUdnderscore(t.Field(i).Name) + ","
			values += strconv.Itoa(int(v.Field(i).Int())) + ","
		case reflect.String:
			if v.Field(i).String() != "" {
				fileds += CamelCaseToUdnderscore(t.Field(i).Name) + ","
				values += "'" + v.Field(i).String() + "',"
			}
		case reflect.Struct:
			if v.Field(i).CanFloat() {
				fileds += CamelCaseToUdnderscore(t.Field(i).Name) + ","
				values += strconv.FormatFloat(v.Field(i).Float(), 'g', 5, 32) + ","
			}
		case reflect.Ptr:
			if !v.Field(i).IsNil() {
				fileds += CamelCaseToUdnderscore(t.Field(i).Name) + ","
				if v.Field(i).Elem().CanFloat() {
					values += strconv.FormatFloat(v.Field(i).Elem().Float(), 'g', 5, 32) + ","
				}
				if v.Field(i).Elem().CanInt() {
					values += strconv.Itoa(int(v.Field(i).Elem().Int())) + ","
				}

				values += "'" + v.Field(i).Elem().String() + "',"
			}

		}
	}

	sqlstr = strings.Replace(sqlstr, "$field", fileds[0:len(fileds)-1], 1)
	sqlstr = strings.Replace(sqlstr, "$value", values[0:len(values)-1], 1)
	return sqlstr
}

func GenUpdateSql(s interface{}, tablename string) string {
	sqlstr := "update " + tablename + " set $value"
	values := ""
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return ""
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// only accept struct param
	if v.Kind() != reflect.Struct {
		return ""
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf(CamelCaseToUdnderscore(t.Field(i).Name))
		fmt.Printf(":")
		fmt.Printf("%v", v.Field(i).Interface())
		fmt.Println("")
		switch v.Field(i).Kind() {
		case reflect.Int:

			values += CamelCaseToUdnderscore(t.Field(i).Name) + "=" + strconv.Itoa(int(v.Field(i).Int())) + ","
		case reflect.String:
			if v.Field(i).String() != "" {
				values += CamelCaseToUdnderscore(t.Field(i).Name) + "='" + v.Field(i).String() + "',"
			}
		case reflect.Struct:
			if v.Field(i).CanFloat() {

				values += CamelCaseToUdnderscore(t.Field(i).Name) + "=" + strconv.FormatFloat(v.Field(i).Float(), 'g', 5, 32) + ","
			}
		case reflect.Ptr:
			if !v.Field(i).IsNil() {

				if v.Field(i).Elem().CanFloat() {
					values += CamelCaseToUdnderscore(t.Field(i).Name) + "=" + strconv.FormatFloat(v.Field(i).Elem().Float(), 'g', 5, 32) + ","
				}
				if v.Field(i).Elem().CanInt() {
					values += CamelCaseToUdnderscore(t.Field(i).Name) + "=" + strconv.Itoa(int(v.Field(i).Elem().Int())) + ","
				}

				values += CamelCaseToUdnderscore(t.Field(i).Name) + "='" + v.Field(i).Elem().String() + "',"
			}

		}
	}

	sqlstr = strings.Replace(sqlstr, "$value", values, 1)

	return sqlstr[0 : len(sqlstr)-1]

}
func CamelCaseToUdnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		} else {
			if unicode.IsUpper(r) {
				output = append(output, '_')
			}

			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}
