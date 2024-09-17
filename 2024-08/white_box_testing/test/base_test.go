package test

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
	"white_box_testing/logic"
	"white_box_testing/mocks"
	"white_box_testing/model"
)

/* 生命周期
SetupSuite
SetupTest
Test
TearDownTest
TearDownSuite
*/

const (
	DBUser = "root"
	DBPass = "123456"
	DBHost = "127.0.0.1"
	DBPort = "3306"
	DBName = "test"
)

// BaseSuite 是测试套件的基础结构。
type BaseSuite struct {
	suite.Suite
	db         *sql.DB
	calculator logic.PriceIncreaseCalculator
}

// TestBaseSuite 运行 BaseSuite 测试套件。
func TestBaseSuite(t *testing.T) {
	suite.Run(t, &BaseSuite{})
}

// SetupSuite 设置测试套件。
func (bs *BaseSuite) SetupSuite() {
	bs.T().Log("Setting up the suite...")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser, DBPass, DBHost, DBPort, DBName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		bs.FailNowf("unable to connect to database", err.Error())
		return
	}

	pp := model.NewPriceProvider(db)
	calculator := logic.NewPriceIncreaseCalculator(pp)

	setupDatabase(bs, db)
	bs.db = db
	bs.calculator = calculator
}

// 在测试之前进行设置
func (bs *BaseSuite) BeforeTest(suiteName, testName string) {
	if testName == "TestCalculate_Error" {
		return
	}
	seedTestTable(bs, bs.db) // ts -> price=1, ts+1min -> price=2
}

// 每个测试结束后清理表
func (bs *BaseSuite) TearDownTest() {
	cleanTable(bs)
}

// 所有测试结束后清理数据库
func (bs *BaseSuite) TearDownSuite() {
	tearDownDatabase(bs)
}

func (bs *BaseSuite) TestCalculate_Error() {
	actual, err := bs.calculator.PriceIncrease()

	bs.EqualError(err, "not enough data")
	bs.Equal(0.0, actual)
}

func (bs *BaseSuite) TestCalculate() {
	actual, err := bs.calculator.PriceIncrease()

	bs.Nil(err)
	bs.Equal(100.0, actual)
}

// 初始化数据库（创建表）
func setupDatabase(bs *BaseSuite, db *sql.DB) {
	bs.T().Log("setting up database")

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS stockprices (
    id INT NOT NULL AUTO_INCREMENT,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    price DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`)

	if err != nil {
		bs.FailNowf("unable to create table", err.Error())
	}
}

// 表插入数据
func seedTestTable(bs *BaseSuite, db *sql.DB) {
	bs.T().Log("seeding test table")

	for i := 1; i <= 2; i++ {
		_, err := db.Exec("INSERT INTO stockprices (timestamp, price) VALUES (?,?)",
			time.Now().Add(time.Duration(i)*time.Minute), float64(i))
		if err != nil {
			bs.FailNowf("unable to seed table", err.Error())
		}
	}
}

// 清空表
func cleanTable(bs *BaseSuite) {
	bs.T().Log("cleaning database")

	_, err := bs.db.Exec(`DELETE FROM stockprices`)
	if err != nil {
		bs.FailNowf("unable to clean table", err.Error())
	}
}

// 删除表
func tearDownDatabase(bs *BaseSuite) {
	bs.T().Log("tearing down database")

	_, err := bs.db.Exec(`DROP TABLE stockprices`)
	if err != nil {
		bs.FailNowf("unable to drop table", err.Error())
	}

	err = bs.db.Close()
	if err != nil {
		bs.FailNowf("unable to close database", err.Error())
	}
}

// Unit tests

type UnitTestSuite struct {
	suite.Suite
	calculator        logic.PriceIncreaseCalculator
	priceProviderMock *mocks.PriceProvider
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (uts *UnitTestSuite) SetupTest() {
	priceProviderMock := mocks.PriceProvider{}
	calculator := logic.NewPriceIncreaseCalculator(&priceProviderMock)

	uts.calculator = calculator
	uts.priceProviderMock = &priceProviderMock
}

func (uts *UnitTestSuite) TestCalculate() {
	//uts.priceProviderMock.On("List", mock.Anything).Return([]*model.PriceData{}, nil)
	//
	//actual, err := uts.calculator.PriceIncrease()
	//
	//uts.Equal(0.0, actual)
	//uts.EqualError(err, "not enough data")

	priceData := []*model.PriceData{
		{Timestamp: time.Now(), Price: 2.0},
		{Timestamp: time.Now().Add(time.Minute), Price: 1.0},
	}

	uts.priceProviderMock.On("List", mock.Anything).Return(priceData, nil)
	actual, err := uts.calculator.PriceIncrease()
	uts.Nil(err)
	uts.Equal(100.0, actual)
}

func (uts *UnitTestSuite) TestCalculate_ErrorFromPriceProvider() {
	expectedError := errors.New("oh my god")

	uts.priceProviderMock.On("List", mock.Anything).Return([]*model.PriceData{}, expectedError)

	actual, err := uts.calculator.PriceIncrease()

	uts.Equal(0.0, actual)
	uts.Equal(expectedError, err)

}
