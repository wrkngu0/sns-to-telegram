package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handler(ctx context.Context, snsEvent events.SNSEvent) {

	var ApiKey = os.Getenv("API_KEY")
	var ChatId = os.Getenv("CHAT_ID")

	consUrl := []string{"https://api.telegram.org/bot", ApiKey, "/sendMessage"}
	requestUrl := strings.Join(consUrl, "")

	client := &http.Client{}

	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		message := fmt.Sprintf("%s \n", snsRecord.Message)

		req, err := http.NewRequest("GET", requestUrl, nil)
		check(err)
		q := req.URL.Query()
		q.Add("chat_id", ChatId)
		q.Add("text", message)
		req.URL.RawQuery = q.Encode()
		resp, err := client.Do(req)
		check(err)

		if resp.StatusCode != 200 {
			msg := fmt.Sprintf("Error response from telegramm api [ Status code: %d] \n", resp.StatusCode)
			log.Print(msg)
		}
	}

}

func main() {
	lambda.Start(handler)
}
