package inserter

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	// necessary for sql :pointup:
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/codeuniversity/smag-mvp/models"
	"github.com/codeuniversity/smag-mvp/service"
	"github.com/codeuniversity/smag-mvp/utils"

	"github.com/segmentio/kafka-go"
)

// Inserter represents the scraper containing all clients it uses
type Inserter struct {
	qReader *kafka.Reader
	qWriter *kafka.Writer

	db *gorm.DB
	*service.Executor
}

// New returns an initilized scraper
func New(postgresHost, postgresPassword string, qReader *kafka.Reader, qWriter *kafka.Writer) *Inserter {
	i := &Inserter{}
	i.qReader = qReader
	i.qWriter = qWriter

	connectionString := fmt.Sprintf("host=%s user=postgres dbname=instascraper sslmode=disable", postgresHost)
	if postgresPassword != "" {
		connectionString += " " + "password=" + postgresPassword
	}

	db, err := gorm.Open("postgres", connectionString)
	utils.PanicIfErr(err)
	i.db = db

	db.AutoMigrate(&models.TwitterUser{})

	i.Executor = service.New()
	return i
}

// Run the inserter
func (i *Inserter) Run() {
	defer func() {
		i.MarkAsStopped()
	}()

	fmt.Println("starting inserter")
	for i.IsRunning() {
		m, err := i.qReader.FetchMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}
		info := &models.TwitterUserRaw{}
		err = json.Unmarshal(m.Value, info)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("inserting: ", info.Username)
		i.insertUser(models.ConvertTwitterUser(info))
		i.qReader.CommitMessages(context.Background(), m)
		fmt.Println("commited: ", info.Username)
	}
}

// Close the inserter
func (i *Inserter) Close() {
	i.Stop()
	i.WaitUntilStopped(time.Second * 3)

	i.db.Close()
	i.qReader.Close()
	if i.qWriter != nil {
		i.qWriter.Close()
	}

	i.MarkAsClosed()
}

func (i *Inserter) insertUser(user *models.TwitterUser) {
	var err error

	fromUser := models.TwitterUser{}
	filter := &models.TwitterUser{ID: user.ID}

	err = createOrUpdate(i.db, &fromUser, filter, user)
	utils.PanicIfErr(err)

	for _, follow := range p.Follows {
		var toUser models.User
		var d *gorm.DB
		d = i.db.Where("user_name = ?", follow.UserName).Select("ID").Find(&toUser)
		if err := d.Error; err != nil {
			if d.RecordNotFound() == true {
				d = i.db.Create(&models.User{
					UserName: follow.UserName,
				}).Scan(&toUser)
				utils.PanicIfErr(d.Error)

				i.handleCreatedUser(follow.UserName)
			} else {
				utils.PanicIfErr(err)
			}
		}
		err = i.db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING").Create(&models.Follow{
			From: fromUser.ID,
			To:   toUser.ID,
		}).Error
		utils.PanicIfErr(err)
	}

}

func (i *Inserter) handleCreatedUser(userName string) {
	// if qWriter is nil, user discovery is disabled
	if i.qWriter != nil {
		i.qWriter.WriteMessages(context.Background(), kafka.Message{
			Value: []byte(userName),
		})
	}
}

func createOrUpdate(db *gorm.DB, out interface{}, where interface{}, update interface{}) error {
	var err error

	tx := db.Begin()
	if tx.Where(where).First(out).RecordNotFound() {
		err = tx.Create(update).Scan(out).Error
	} else {
		err = tx.Model(out).Update(update).Scan(out).Error
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
