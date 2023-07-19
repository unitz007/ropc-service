package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"ropc-service/model/entities"
)

type Suite struct {
	suite.Suite
	DB     *gorm.DB
	mock   sqlmock.Sqlmock
	client *entities.Client
}

//
//func Test_clientRepository(t *testing.T) {
//
//	s := &Suite{}
//
//	var (
//		db  *sql.DB
//		err error
//	)
//
//	db, s.mock, err = sqlmock.New()
//	if err != nil {
//		t.Errorf("Failed to open mock sql db, error: %v", err)
//	}
//
//	if db == nil {
//		t.Error("mock db is nil")
//	}
//
//	if s.mock == nil {
//		t.Error("sqlmock is nil")
//	}
//
//	dialector := mysql.New(mysql.Config{
//		DriverName: "mysql",
//		DSN:        "sqlmock_db_0",
//		Conn:       db,
//	})
//
//	s.DB, err = gorm.Open(dialector, &gorm.Config{})
//
//	if err != nil {
//		t.Errorf("Failed to open gorm db, error: %v", err)
//	}
//
//	if s.DB == nil {
//		t.Error("gorm db is nil")
//	}
//
//	s.client = &model.Client{
//		ClientId:     "clientId",
//		ClientSecret: "clientSecret",
//	}
//
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//			t.Error("Could not close mock db:", err)
//		}
//	}(db)
//
//	s.mock.MatchExpectationsInOrder(false)
//	s.mock.ExpectBegin()
//
//	s.mock.ExpectQuery()
//
//}
