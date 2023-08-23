package handler

import (
	"auction_backend/event"
	"auction_backend/model"
	"database/sql"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
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
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

	connStr := "postgresql://user:yoptupotop@localhost/auction?sslmode=disable"
	dbase, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	newContract := model.ContractCreated{Auction: common.Address{0x0000000000000000000000000000000000000010}, Owner: common.Address{0x0000000000000000000000000000000000000001}}
	event.ContractsCreated = append(event.ContractsCreated, newContract)

	type args struct {
		c  *fiber.Ctx
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{
			name: "first",
			args: args{
				c:  ctx,
				db: dbase,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PostHandler(tt.args.c, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("PostHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
