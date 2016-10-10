package tumblrapiclient

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"errors"
	"strings"
	"github.com/dghubble/oauth1"
	"github.com/tumblr/tumblr.go"
	"golang.org/x/net/context"
)

const apiBase = "https://api.tumblr.com/v2/"

// The Tumblr API Client object
type Client struct {
	tumblrapi.ClientInterface
	consumer *oauth1.Config
	user *oauth1.Token
	client *http.Client
}

// Constructor with only the consumer key and secret
func NewClient(consumerKey string, consumerSecret string) *Client {
	c := Client{}
	c.SetConsumer(consumerKey, consumerSecret)
	return &c
}

// Constructor with consumer key/secret and user token/secret
func NewClientWithToken(consumerKey string, consumerSecret string, token string, tokenSecret string) *Client {
	c := NewClient(consumerKey, consumerSecret)
	c.SetToken(token, tokenSecret)
	return c
}

// Set consumer credentials, invalidates any previously cached client
func (c *Client) SetConsumer(consumerKey string, consumerSecret string) {
	c.consumer = oauth1.NewConfig(consumerKey, consumerSecret)
	c.client = nil
}

// Set user credentials, invalidates any previously cached client
func (c *Client) SetToken(token string, tokenSecret string) {
	c.user = oauth1.NewToken(token, tokenSecret)
	c.client = nil
}

// Issue GET request to Tumblr API
func (c *Client) Get(endpoint string) (tumblrapi.Response, error) {
	return c.GetWithParams(endpoint, url.Values{})
}

// Issue GET request to Tumblr API with param values
func (c *Client) GetWithParams(endpoint string, params url.Values) (tumblrapi.Response, error) {
	return getResponse(c.GetHttpClient().Get(createRequestURI(appendPath(apiBase,endpoint),params)))
}

// Issue POST request to Tumblr API
func (c *Client) Post(endpoint string) (tumblrapi.Response, error) {
	return c.PostWithParams(endpoint, url.Values{});
}

// Issue POST request to Tumblr API with param values
func (c *Client) PostWithParams(endpoint string, params url.Values) (tumblrapi.Response, error) {
	return getResponse(c.GetHttpClient().PostForm(appendPath(apiBase, endpoint), params))
}

// Issue PUT request to Tumblr API
func (c *Client) Put(endpoint string) (tumblrapi.Response, error) {
	return c.PutWithParams(endpoint, url.Values{});
}

// Issue PUT request to Tumblr API with param values
func (c *Client) PutWithParams(endpoint string, params url.Values) (tumblrapi.Response, error) {
	req, err := http.NewRequest("PUT", createRequestURI(appendPath(apiBase, endpoint), params), strings.NewReader(""))
	if err == nil {
		return getResponse(c.GetHttpClient().Do(req))
	}
	return tumblrapi.Response{}, err
}

// Issue DELETE request to Tumblr API
func (c *Client) Delete(endpoint string) (tumblrapi.Response, error) {
	return c.DeleteWithParams(endpoint, url.Values{});
}

// Issue DELETE request to Tumblr API with param values
func (c *Client) DeleteWithParams(endpoint string, params url.Values) (tumblrapi.Response, error) {
	req, err := http.NewRequest("DELETE", createRequestURI(appendPath(apiBase, endpoint), params), strings.NewReader(""))
	if err == nil {
		return getResponse(c.GetHttpClient().Do(req))
	}
	return tumblrapi.Response{}, err
}

// Retrieve the underlying HTTP client
func (c *Client) GetHttpClient() *http.Client {
	if c.consumer == nil {
		panic("Consumer credentials are not set")
	}
	if c.user == nil {
		c.SetToken("", "")
	}
	if c.client == nil {
		c.client = c.consumer.Client(context.TODO(), c.user)
		c.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	return c.client
}

// Helper function to ease appending path to a base URI
func appendPath(base string, path string) string {
	// if path starts with `/` shave it off
	if path[0] == '/' {
		path = path[1:]
	}
	return base + path
}

// Helper function to create a URI with query params
func createRequestURI(base string, params url.Values) string {
	if len(params) != 0 {
		base += "?" + params.Encode()
	}
	return base
}

// Standard way of receiving data from the API response
func getResponse(resp *http.Response, e error) (tumblrapi.Response, error) {
	response := tumblrapi.Response{}
	if e != nil {
		return response, e
	}
	defer resp.Body.Close()
	response.Headers = resp.Header
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return response, e
	}
	response = *tumblrapi.NewResponse(body, resp.Header)
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return response, errors.New(resp.Status)
	}
	return response, nil
}

func (c *Client) GetPost(id uint64, blogName string) (*tumblrapi.PostRef) {
	return tumblrapi.NewPostRef(c, &tumblrapi.MiniPost{
		Id: id,
		BlogName: blogName,
	})
}

func (c *Client) GetBlog(name string) (*tumblrapi.BlogRef) {
	return tumblrapi.NewBlogRef(c, name)
}

func (c *Client) GetUser() (*tumblrapi.User, error) {
	return tumblrapi.GetUserInfo(c)
}

func (c *Client) GetDashboard() ([]tumblrapi.PostInterface, error) {
	return c.GetDashboardWithParams(url.Values{})
}

func (c *Client) GetDashboardWithParams(params url.Values) ([]tumblrapi.PostInterface, error) {
	return tumblrapi.GetDashboard(c, params)
}

func (c *Client) GetLikes() (*tumblrapi.Likes, error) {
	return c.GetLikesWithParams(url.Values{})
}

func (c *Client) GetLikesWithParams(params url.Values) (*tumblrapi.Likes, error) {
	return tumblrapi.GetLikes(c, params)
}

func (c *Client) TaggedSearch(tag string) (*tumblrapi.SearchResults, error) {
	return tumblrapi.TaggedSearch(c, tag, url.Values{})
}

func (c *Client) TaggedSearchWithParams(tag string, params url.Values) (*tumblrapi.SearchResults, error) {
	return tumblrapi.TaggedSearch(c, tag, params)
}