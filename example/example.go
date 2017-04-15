package main

import (
	"encoding/xml"
	"fmt"

	"net/url"

	"net/http"

	"time"

	post "github.com/ariefrahmansyah/postman"
)

func printOutput(response *http.Response, body []byte, err error) {
	fmt.Println("HTTP Response:", response, "\nBody:", string(body), "\nError:", err)
}

func main() {
	fmt.Println("Custom Client Request Example")
	customClientExample()

	fmt.Println("\nGET Request Example")
	getExample()

	fmt.Println("\nPOST Request Example")
	postExample()

	fmt.Println("\nPOST Raw Byte Request Example")
	postByteExample()

	fmt.Println("\nPOST Raw String Request Example")
	postStringExample()

	fmt.Println("\nPOST JSON Request Example")
	postJSONExample()

	fmt.Println("\nPOST XML Request Example")
	postXMLExample()
}

func customClientExample() {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	header := http.Header{
		"Cache-Control": {"no-cache", "no-store", "must-revalidate"},
	}
	params := url.Values{
		"show_env": {"0"},
	}

	postman := post.NewPostman()
	postman.Client = client
	postman.Header = header
	postman.Method = post.MethodGet
	postman.URL = "http://httpbin.org/get?show_env=1"
	postman.Params = params

	printOutput(postman.Send())
}

func getExample() {
	header := http.Header{
		"Cache-Control": {"no-cache", "no-store", "must-revalidate"},
	}
	params := url.Values{
		"show_env": {"0"},
	}

	postman := post.NewPostman()
	postman.Header = header
	postman.Method = post.MethodGet
	postman.URL = "http://httpbin.org/get?show_env=1"
	postman.Params = params

	printOutput(postman.Send())
}

func postExample() {
	header := http.Header{
		"Cache-Control": {"no-cache", "no-store", "must-revalidate"},
	}
	params := url.Values{
		"show_env": {"0"},
	}

	postman := post.NewPostman()
	postman.Header = header
	postman.Method = post.MethodPost
	postman.URL = "http://httpbin.org/post"
	postman.Params = params

	printOutput(postman.Send())
}

func postByteExample() {
	header := http.Header{
		"Cache-Control": {"no-cache", "no-store", "must-revalidate"},
	}

	postman := post.NewPostman()
	postman.Header = header
	postman.Method = post.MethodPostRaw
	postman.URL = "http://httpbin.org/post"
	postman.Body = []byte("this is byte")

	printOutput(postman.Send())
}

func postStringExample() {
	header := http.Header{
		"Cache-Control": {"no-cache", "no-store", "must-revalidate"},
	}

	postman := post.NewPostman()
	postman.Header = header
	postman.Method = post.MethodPostRaw
	postman.URL = "http://httpbin.org/post"
	postman.Body = "this is string"

	printOutput(postman.Send())
}

func postJSONExample() {
	header := http.Header{
		"Cache-Control": {"no-cache", "no-store", "must-revalidate"},
	}
	params := url.Values{
		"show_env": {"0"},
	}

	data := struct {
		UserID   int64  `json:"user_id"`
		Username string `json:"username"`
	}{int64(230394), "ariefrahmansyah"}

	postman := post.NewPostman()
	postman.Header = header
	postman.Method = post.MethodPostJSON
	postman.URL = "http://httpbin.org/post"
	postman.Params = params
	postman.Body = data

	printOutput(postman.Send())
}

func postXMLExample() {
	header := http.Header{
		"Cache-Control": {"no-cache", "no-store", "must-revalidate"},
		"Content-Type":  {"application/xml"},
	}
	params := url.Values{
		"show_env": {"0"},
	}

	data := struct {
		XMLName  xml.Name `xml:"user"`
		UserID   int64    `xml:"user_id"`
		Username string   `xml:"username"`
	}{UserID: int64(230394), Username: "ariefrahmansyah"}

	postman := post.NewPostman()
	postman.Header = header
	postman.Method = post.MethodPostXML
	postman.URL = "http://httpbin.org/post"
	postman.Params = params
	postman.Body = data

	printOutput(postman.Send())
}
