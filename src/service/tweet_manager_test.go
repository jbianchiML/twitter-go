package service_test

import (
	"testing"

	"github.com/jbianchiML/twitter-go/src/domain"
	"github.com/jbianchiML/twitter-go/src/service"
)

/*func TestPublishedTweetIsSaved(t *testing.T) {

	var tweet string = "This is my first tweet"

	service.PublishTweet(tweet)

	//if service.Tweet != tweet {
	//	t.Error("Expected tweet is", tweet)
	//}

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}

}*/

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ := service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetById(id, user)
	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestPublishedTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	user := "Usuario"
	var text string
	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	user := "Usuario"
	text := "Tweet de mas de 140 caracteres, Test de Go para probar si el tweet tiene mas de 140 caracteres.	bool string int int8 int16 int32 int64. (si escribo solo int, automaticamente toma el tamaño máximo que soporte la arquitectura de la pc)  uint uint8 uint16 uint32 uint64 uintptr 	byte //alias de uint8 		rune // alias de int32. Representa un code point de Unicode  loat32 float64  omplex64 complex128"
	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "tweet exceeding 140 characters" {
		t.Error("Expected error is Tweet exceeding 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet *domain.Tweet // Fill the tweets with data
	user01 := "Usuario"
	text01 := "Tweet01"
	tweet = domain.NewTweet(user01, text01)

	user02 := "Usuario"
	text02 := "Tweet02"
	secondTweet = domain.NewTweet(user02, text02)

	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[user01][0]
	secondPublishedTweet := publishedTweets[user02][0]
	if !isValidTweet(t, firstPublishedTweet, 0, user01, text01) {
		return
	}
	// Same for secondPublishedTweet
	if !isValidTweet(t, secondPublishedTweet, 1, user02, text02) {
		return
	}
}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user string, text string) bool {
	if tweet.Text != text || tweet.User != user {
		t.Errorf("Invalid tweet")
		return false
	}
	return true
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	println(id)
	// Validation
	publishedTweet := service.GetTweetById(id, user)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)
	// Operation
	count := service.CountTweetsByUser(user)
	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	// publish the 3 tweets
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)
	// Operation
	tweets := service.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 { /* handle error */
	}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet
	if !isValidTweet(t, firstPublishedTweet, 0, user, text) {
		return
	}
	// Same for secondPublishedTweet
	if !isValidTweet(t, secondPublishedTweet, 1, user, secondText) {
		return
	}
}
