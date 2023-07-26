package data

import (
	"errors"
	"kratos-demo/internal/conf"
	"testing"
)

func TestNewDB(t *testing.T) {
	t.Run("test NewDB", func(t *testing.T) {
		db := NewDB(&conf.Data{
			Database: &conf.Data_Database{
				Source: "root:passwd@tcp(192.168.44.20:3306)/realworld?charset=utf8mb4&parseTime=True&loc=Local",
			},
		})
		if db == nil {
			panic(errors.New("db is nil"))
		}
	})
}
