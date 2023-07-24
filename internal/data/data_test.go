package data

import "testing"

func TestNewDB(t *testing.T) {
	t.Run("test NewDB", func(t *testing.T) {
		err := NewDB(nil)
		if err != nil {
			panic(err)
			return
		}
	})
}
