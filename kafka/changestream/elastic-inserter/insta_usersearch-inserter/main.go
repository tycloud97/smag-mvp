package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"

	"github.com/codeuniversity/smag-mvp/kafka/changestream"
	esInserter "github.com/codeuniversity/smag-mvp/kafka/changestream/elastic-inserter"
	"github.com/codeuniversity/smag-mvp/service"
	"github.com/codeuniversity/smag-mvp/utils"
)

const esIndex = "insta_user"

func main() {
	kafkaAddress := utils.GetStringFromEnvWithDefault("KAFKA_ADDRESS", "my-kafka:9092")
	groupID := utils.GetStringFromEnvWithDefault("KAFKA_GROUPID", "insta_usersearch-inserter")
	changesTopic := utils.GetStringFromEnvWithDefault("KAFKA_CHANGE_TOPIC", "postgres.public.users")

	esHosts := utils.GetMultipliesStringsFromEnvDefault("ELASTIC_SEARCH_ADDRESS", []string{"http://localhost:9201"})

	// create and run esInserter
	i := esInserter.New(esHosts, esIndex, instaUserMapping, kafkaAddress, changesTopic, groupID, handleChangemessage)

	service.CloseOnSignal(i)
	waitUntilClosed := i.Start()

	waitUntilClosed()
}

// handleChangemessage filters relevant events and upserts them
func handleChangemessage(m *changestream.ChangeMessage, esClient *elasticsearch.Client) error {
	user := &user{}
	if err := json.Unmarshal(m.Payload.After, user); err != nil {
		return err
	}

	switch m.Payload.Op {
	case "c", "r", "u":
		return upsertComment(user, esClient)
	}

	return nil
}

func upsertComment(u *user, esClient *elasticsearch.Client) error {
	// not using esapi.NewJSONReader, because we will wrap the user into the upsert request
	jsonUser, err := json.Marshal(*u)
	if err != nil {
		return err
	}

	searchUser := fmt.Sprintf(instaUserUpsert, jsonUser, jsonUser)
	response, err := esClient.Update(
		esIndex,
		strconv.Itoa(u.ID),
		strings.NewReader(searchUser))
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("Failed to upsert user. StatusCode: %d", response.StatusCode)
	}

	return nil
}
