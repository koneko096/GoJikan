package gojikan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/koneko096/gojikan/types"
)

type Jikan struct {
	BaseUrl string       `json:"baseUrl"`
	Client  *http.Client `json:"-"`
	Debug   bool         `json:"debug"`
}

func NewJikan() (*Jikan, error) {
	return NewJikanWithBaseUrl(BASE_URL)
}

func NewJikanWithBaseUrl(baseUrl string) (*Jikan, error) {
	return NewJikanWithBaseUrlAndClient(baseUrl, &http.Client{})
}

func NewJikanWithBaseUrlAndClient(baseUrl string, client *http.Client) (*Jikan, error) {
	api := &Jikan{
		BaseUrl: baseUrl,
		Client:  client,
		Debug:   false,
	}

	return api, nil
}

// MakeRequest makes a request to a specific endpoint.
func (api *Jikan) makeRequest(endpoint string, params url.Values) ([]byte, error) {
	method := api.BaseUrl + endpoint

	resp, err := api.Client.PostForm(method, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if api.Debug {
		log.Printf("%s resp: %s", endpoint, bytes)
	}

	return bytes, nil
}

func (api *Jikan) GetAnime(id int) (types.Anime, error) {
	bytes, err := api.makeRequest(fmt.Sprintf(GET_ANIME_ENDPOINT, id), url.Values{})
	if err != nil {
		return types.Anime{}, err
	}

	var resp = types.Anime{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.Anime{}, err
	}

	return resp, nil
}

func (api *Jikan) GetAnimePictures(id int) (types.PictureResponse, error) {
	bytes, err := api.makeRequest(fmt.Sprintf(GET_ANIME_PICTURES_ENDPOINT, id), url.Values{})
	if err != nil {
		return types.PictureResponse{}, err
	}

	var resp = types.PictureResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.PictureResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) GetAnimeVideos(id int) (types.VideoResponse, error) {
	bytes, err := api.makeRequest(fmt.Sprintf(GET_ANIME_VIDEOS_ENDPOINT, id), url.Values{})
	if err != nil {
		return types.VideoResponse{}, err
	}

	var resp = types.VideoResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.VideoResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) GetAnimeCharactersAndStaffs(id int) (types.CharactersAndStaffsResponse, error) {
	bytes, err := api.makeRequest(fmt.Sprintf(GET_ANIME_CHARACTERS_STAFFS_ENDPOINT, id), url.Values{})
	if err != nil {
		return types.CharactersAndStaffsResponse{}, err
	}

	var resp = types.CharactersAndStaffsResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.CharactersAndStaffsResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) GetCharacter(id int) (types.Character, error) {
	bytes, err := api.makeRequest(fmt.Sprintf(GET_CHARACTER_ENDPOINT, id), url.Values{})
	if err != nil {
		return types.Character{}, err
	}

	var resp = types.Character{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.Character{}, err
	}

	return resp, nil
}

func (api *Jikan) GetSeasonList(year int, season string) (types.SeasonListResponse, error) {
	bytes, err := api.makeRequest(fmt.Sprintf(GET_SEASON_LIST_ENDPOINT, year, season), url.Values{})
	if err != nil {
		return types.SeasonListResponse{}, err
	}

	var resp = types.SeasonListResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.SeasonListResponse{}, err
	}

	return resp, nil
}
