package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Response modules the JSON structure
// that we get back from the YouTube API
type Response struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

// Items stores the ID + Statistics for
// a given channel
type Items struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

// Stats stores the information we care about
// so how many views the channel has, how many subscribers
// how many video etc.
type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

func GetSubscribers() (Items, error) {
	proxyUrl, _ := url.Parse("http://localhost:12333")
	var response Response
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println(err)
		return Items{}, err
	}

	// here we define the query parameters and their respective values
	q := req.URL.Query()
	q.Add("key", os.Getenv("YOUTUBE_KEY"))
	q.Add("id", os.Getenv("CHANNEL_ID"))
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

	// finally we make the request to the URL that we have just
	// constructed
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	resp, err := client.Do(req)
	if err != nil {
		return Items{}, err
	}
	defer resp.Body.Close()
	fmt.Println("Response Status: ", resp.Status)

	// we then read in all of the body of the
	// JSON response
	body, _ := ioutil.ReadAll(resp.Body)
	// and finally unmarshal it into an Response struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Items{}, nil
	}

	// we only care about the first Item in our
	// Items array, so we just send that back
	return response.Items[0], nil
}
