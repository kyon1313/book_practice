package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kyon1313/books/model"
)

type HttpMethod string

const (
	POST   HttpMethod = "POST"
	GET    HttpMethod = "GET"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

func consumeEndpoint(method HttpMethod, url string, body []byte) (*http.Response, []byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(string(method), url, bytes.NewReader(body))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, responseBody, nil
}

func CallGetAuthorBookEndpoint() ([]model.AuthorsBookResponse, error) {
	resp, responseBody, err := consumeEndpoint(GET, "http://127.0.0.1:3000/endpoint/v1/getBooks", nil)

	if err != nil {
		fmt.Println(err)
		return []model.AuthorsBookResponse{}, err
	}
	defer resp.Body.Close()
	var authorBook []model.AuthorsBookResponse
	json.Unmarshal(responseBody, &authorBook)

	if len(authorBook) == 0 {
		return []model.AuthorsBookResponse{}, errors.New("no data available")
	}

	return authorBook, nil

}

func CallGetAuthorBookByIsbnEndpoint(isbn string) (model.AuthorsBookResponse, error) {
	resp, responseBody, err := consumeEndpoint(GET, "http://127.0.0.1:3000/endpoint/v1/getBookByIbsn/"+isbn, nil)

	if err != nil {
		fmt.Println(err)
		return model.AuthorsBookResponse{}, err
	}
	defer resp.Body.Close()
	var authorBook model.AuthorsBookResponse
	json.Unmarshal(responseBody, &authorBook)

	return authorBook, nil

}

func CallAddAuthorBookEndpoint(c *fiber.Ctx, marshallBookRequest []byte) (*http.Response, map[string]interface{}, error) {
	var bookResponse map[string]interface{}
	resp, responseBody, err := consumeEndpoint(POST, "http://127.0.0.1:3000/endpoint/v1/addBook", marshallBookRequest)
	if err != nil {
		return nil, map[string]interface{}{}, err
	}
	defer resp.Body.Close()

	err = json.Unmarshal(responseBody, &bookResponse)
	if err != nil {
		return nil, map[string]interface{}{}, err
	}

	return resp, bookResponse, nil
}

func CallUpdateAuthorBookEndpoint(c *fiber.Ctx, marshallBookRequest []byte, isbn string) (*http.Response, map[string]interface{}, error) {
	var bookResponse map[string]interface{}
	resp, responseBody, err := consumeEndpoint(POST, "http://127.0.0.1:3000/endpoint/v1/updateAuthorBook/"+isbn, marshallBookRequest)
	if err != nil {
		return nil, map[string]interface{}{}, err
	}
	defer resp.Body.Close()

	err = json.Unmarshal(responseBody, &bookResponse)
	if err != nil {
		return nil, map[string]interface{}{}, err
	}

	return resp, bookResponse, nil
}

func CallDeleteAuthorBook(isbn string) (*http.Response, error) {

	resp, _, err := consumeEndpoint(GET, "http://127.0.0.1:3000/endpoint/v1/deleteBookByIbsn/"+isbn, nil)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}
	defer resp.Body.Close()
	return resp, nil
}
