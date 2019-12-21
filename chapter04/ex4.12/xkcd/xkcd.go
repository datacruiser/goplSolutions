package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ComicIndex struct {
	Comics []*Comic
}

type Comic struct {
	Alt        string
	Day        string
	Img        string
	Link       string
	Month      string
	News       string
	Num        int
	SafeTitle  string
	Title      string
	Transcript string
	Year       string
}

func getComicURL(comicID int) string {
	return fmt.Sprintf("https://xkcd.com/%s/info.0.json", strconv.Itoa(comicID))
}

func NewComicIndex() ComicIndex {
	return ComicIndex{[]*Comic{}}
}

func GetComic(comicID int) (*Comic, error) {
	resp, err := http.Get(getComicURL(comicID))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get comic failed: %s", resp.Status)
	}

	var result Comic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
