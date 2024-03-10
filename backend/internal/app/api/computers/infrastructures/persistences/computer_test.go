package persistences

import (
	"backend/internal/app/api/computers/domains/entities"
	"backend/internal/app/api/pkg/database"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	computer := &entities.Computer{
		HostName:   "Testing",
		IPAddress:  "0.0.0.0",
		MACAddress: "00:00:00:00:00:00",
	}

	db, mock, err := database.GetMock()
	if err != nil {
		t.Errorf(err.Error())
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "computers"`)).
		WithArgs(database.AnyTime{}, database.AnyTime{}, nil, computer.HostName, computer.IPAddress, computer.MACAddress).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	p := NewComputerPersistence()
	if err := p.Create(db, computer); err != nil {
		t.Errorf(err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf(err.Error())
	}
}

func TestUpdate(t *testing.T) {
	computer := &entities.Computer{
		Model: gorm.Model{
			ID: 1,
		},
		HostName:   "Testing",
		IPAddress:  "0.0.0.0",
		MACAddress: "00:00:00:00:00:00",
	}

	db, mock, err := database.GetMock()
	if err != nil {
		t.Errorf(err.Error())
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "computers" SET`)).
		WithArgs(database.AnyTime{}, database.AnyTime{}, nil, computer.HostName, computer.IPAddress, computer.MACAddress, computer.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	p := NewComputerPersistence()

	if err := p.Update(db, computer); err != nil {
		t.Errorf(err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf(err.Error())
	}
}

func TestDelete(t *testing.T) {
	computer := &entities.Computer{
		Model: gorm.Model{
			ID: 1,
		},
		HostName:   "Testing",
		IPAddress:  "0.0.0.0",
		MACAddress: "00:00:00:00:00:00",
	}

	db, mock, err := database.GetMock()
	if err != nil {
		t.Errorf(err.Error())
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "computers"`)).
		WithArgs(computer.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	p := NewComputerPersistence()

	if err := p.Delete(db, computer); err != nil {
		t.Errorf(err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf(err.Error())
	}
}

func TestFindOneById(t *testing.T) {
	var computer entities.Computer

	db, mock, err := database.GetMock()
	if err != nil {
		t.Errorf(err.Error())
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "computers" WHERE "computers"."id" = $1 AND "computers"."deleted_at" IS NULL ORDER BY "computers"."id" LIMIT 1`)).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "host_name", "ip_address", "mac_address"}).
				AddRow(1, time.Now(), time.Now(), nil, "Testing", "0.0.0.0", "00:00:00:00:00:00"),
		)

	p := NewComputerPersistence()
	if err := p.FindOneById(db, &computer, 1); err != nil {
		t.Errorf(err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf(err.Error())
	}
}

func TestFind(t *testing.T) {
	var computers []entities.Computer
	options := database.Options{
		Sort:  "updated_at",
		Order: "desc",
	}

	db, mock, err := database.GetMock()
	if err != nil {
		t.Errorf(err.Error())
	}

	mock.ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(`SELECT * FROM "computers" WHERE "computers"."deleted_at" IS NULL ORDER BY %s %s`, options.Sort, options.Order))).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "host_name", "ip_address", "mac_address"}).
				AddRow(1, time.Now(), time.Now(), nil, "Testing", "0.0.0.0", "00:00:00:00:00:00"),
		)

	p := NewComputerPersistence()
	if err := p.Find(db, &computers, &options); err != nil {
		t.Errorf(err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf(err.Error())
	}
}
