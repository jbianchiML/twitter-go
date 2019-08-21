package service

import (
	"errors"

	"github.com/jbianchiML/twitter-go/src/domain"
)

var mytweet map[string][]*domain.Tweet

func InitializeService() {
	mytweet = make(map[string][]*domain.Tweet)
}

//PublishTweet set tweet
func PublishTweet(tweet *domain.Tweet) (int, error) {

	error := verifyTweet(tweet)
	if error == nil {
		//_, existe := mytweet[tweet.User] //mytweet = append(mytweet, tweet)
		//if existe {
		userTweets := mytweet[tweet.User]
		if userTweets == nil {
			userTweets = make([]*domain.Tweet, 0)
		}
		mytweet[tweet.User] = append(userTweets, tweet)
		/*else{
			mytweet[tweet.User]=
		}*/
		return len(mytweet[tweet.User]) - 1, error
	}
	return -1, error
}

func verifyTweet(tweet *domain.Tweet) error {
	if tweet.User == "" {
		return errors.New("user is required")
	}
	if tweet.Text == "" {
		return errors.New("text is required")
	}
	//println(len(tweet.Text))
	if len(tweet.Text) > 140 {
		return errors.New("tweet exceeding 140 characters")
	}
	return nil
}

// GetTweet Devuelve el tweet
/*func GetTweet() *domain.Tweet {
	return mytweet[0]
}
*/
// GetTweets
func GetTweets() map[string][]*domain.Tweet {
	return mytweet
}

func GetTweetById(id int, user string) *domain.Tweet {

	return mytweet[user][id]
}

func CountTweetsByUser(user string) int {
	/*count := 0
	for _, valor := range mytweet {
		if valor.User == user {
			count++
		}
	}
	return count*/

	return len(mytweet[user])

}

// GetTweetsByUser
func GetTweetsByUser(user string) []*domain.Tweet {
	return mytweet[user]
}
