package handler

import (
	"database/sql"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetHandler(t *testing.T) {
	type args struct {
		c  *fiber.Ctx
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetHandler(tt.args.c, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("GetHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostHandler(t *testing.T) {
	type args struct {
		c  *fiber.Ctx
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PostHandler(tt.args.c, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("PostHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
