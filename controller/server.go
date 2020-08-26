package controller

import (
	"log"
	"time"

	"github.com/ashishb26/rzpbool/models"
	"github.com/gin-gonic/gin"

	"github.com/go-redsync/redsync/v3"
	"github.com/go-redsync/redsync/v3/redis"
	"github.com/go-redsync/redsync/v3/redis/redigo"
	redigolib "github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql driver
)

// Server struct is used to represent the database connection and router
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
	Mutex  *redsync.Mutex
}

// NewServer creates a new database connection and router and returns a Server struct
func NewServer() *Server {

	dbConfig := models.GetDBConfig()

	db, err := gorm.Open("mysql", dbConfig.DBUser+":"+dbConfig.DBPassword+"@tcp("+dbConfig.DBHost+":"+dbConfig.DBPort+")/"+dbConfig.DBName+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}
	router := gin.Default()

	db.AutoMigrate(&models.BoolRecord{})

	newLock, err := GetRedisLock()

	if err != nil {
		log.Fatalln(err.Error())
	}

	return &Server{
		DB:     db,
		Router: router,
		Mutex:  newLock,
	}
}

// GetRedisLock creates a new redis client and reddis based mutex
func GetRedisLock() (*redsync.Mutex, error) {

	redisAddr, err := models.GetRedisAddr()
	if err != nil {
		return nil, err
	}
	pool := redigo.NewRedigoPool(&redigolib.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redigolib.Conn, error) {
			return redigolib.Dial("tcp", redisAddr)
		},
		TestOnBorrow: func(c redigolib.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	})

	rs := redsync.New([]redis.Pool{pool})

	mutex := rs.NewMutex("db_mutex")
	return mutex, nil
}
