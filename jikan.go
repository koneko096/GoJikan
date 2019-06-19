package gojikan

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/koneko096/gojikan/types"
)

type Jikan struct {
	baseUrl      string
	clientHelper ClientHelper
}

func NewJikan() *Jikan {
	return NewJikanWithBaseUrl(BASE_URL)
}

func NewJikanWithBaseUrl(baseUrl string) *Jikan {
	return NewJikanWithBaseUrlAndClient(baseUrl, &http.Client{})
}

func NewJikanWithBaseUrlAndClient(baseUrl string, client HttpClient) *Jikan {
	return &Jikan{
		baseUrl:      baseUrl,
		clientHelper: NewClientHelper(client),
	}
}

func (api *Jikan) get(path string) ([]byte, error) {
	return api.getWithParam(path, url.Values{})
}

func (api *Jikan) getWithParam(path string, params url.Values) ([]byte, error) {
	return api.clientHelper.Get(BASE_URL+path, params)
}

func (api *Jikan) GetAnime(id int) (types.Anime, error) {
	bytes, err := api.get(fmt.Sprintf(GET_ANIME_ENDPOINT, id))
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
	bytes, err := api.get(fmt.Sprintf(GET_ANIME_PICTURES_ENDPOINT, id))
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
	bytes, err := api.get(fmt.Sprintf(GET_ANIME_VIDEOS_ENDPOINT, id))
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
	bytes, err := api.get(fmt.Sprintf(GET_ANIME_CHARACTERS_STAFFS_ENDPOINT, id))
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
	bytes, err := api.get(fmt.Sprintf(GET_CHARACTER_ENDPOINT, id))
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
	bytes, err := api.get(fmt.Sprintf(GET_SEASON_LIST_ENDPOINT, year, season))
	if err != nil {
		return types.SeasonListResponse{}, err
	}

	var resp = types.SeasonListResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.SeasonListResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) GetAnimeFromGenre(genreId types.GenreId) (types.AnimeGenreResponse, error) {
	return api.GetAnimeFromGenreWithPage(genreId, 1)
}

func (api *Jikan) GetAnimeFromGenreWithPage(genreId types.GenreId, page int) (types.AnimeGenreResponse, error) {
	bytes, err := api.get(fmt.Sprintf(GET_ANIME_FROM_GENRE_ENDPOINT, genreId, page))
	if err != nil {
		return types.AnimeGenreResponse{}, err
	}

	var resp = types.AnimeGenreResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.AnimeGenreResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) GetAnimeFromProducer(producerId int) (types.AnimeProducerResponse, error) {
	return api.GetAnimeFromProducerWithPage(producerId, 1)
}

func (api *Jikan) GetAnimeFromProducerWithPage(producerId int, page int) (types.AnimeProducerResponse, error) {
	bytes, err := api.get(fmt.Sprintf(GET_ANIME_FROM_PRODUCER_ENDPOINT, producerId, page))
	if err != nil {
		return types.AnimeProducerResponse{}, err
	}

	var resp = types.AnimeProducerResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.AnimeProducerResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) GetManga(id int) (types.Manga, error) {
	bytes, err := api.get(fmt.Sprintf(GET_MANGA_ENDPOINT, id))
	if err != nil {
		return types.Manga{}, err
	}

	var resp = types.Manga{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.Manga{}, err
	}

	return resp, nil
}

func (api *Jikan) GetMangaPictures(id int) (types.PictureResponse, error) {
	bytes, err := api.get(fmt.Sprintf(GET_MANGA_PICTURES_ENDPOINT, id))
	if err != nil {
		return types.PictureResponse{}, err
	}

	var resp = types.PictureResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.PictureResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) GetMangaCharactersAndStaffs(id int) (types.CharactersAndStaffsResponse, error) {
	bytes, err := api.get(fmt.Sprintf(GET_MANGA_CHARACTERS_ENDPOINT, id))
	if err != nil {
		return types.CharactersAndStaffsResponse{}, err
	}

	var resp = types.CharactersAndStaffsResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.CharactersAndStaffsResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) GetMangaFromGenre(genreId types.GenreId) (types.MangaGenreResponse, error) {
	return api.GetMangaFromGenreWithPage(genreId, 1)
}

func (api *Jikan) GetMangaFromGenreWithPage(genreId types.GenreId, page int) (types.MangaGenreResponse, error) {
	bytes, err := api.get(fmt.Sprintf(GET_MANGA_FROM_GENRE_ENDPOINT, genreId, page))
	if err != nil {
		return types.MangaGenreResponse{}, err
	}

	var resp = types.MangaGenreResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.MangaGenreResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) GetMangaFromMagazine(magazineId int) (types.MangaMagazineResponse, error) {
	return api.GetMangaFromMagazineWithPage(magazineId, 1)
}

func (api *Jikan) GetMangaFromMagazineWithPage(magazineId int, page int) (types.MangaMagazineResponse, error) {
	bytes, err := api.get(fmt.Sprintf(GET_MANGA_FROM_MAGAZINE_ENDPOINT, magazineId, page))
	if err != nil {
		return types.MangaMagazineResponse{}, err
	}

	var resp = types.MangaMagazineResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.MangaMagazineResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) GetTopAnime() (types.TopAnimeResponse, error) {
	return api.GetTopAnimeWithPage(1)
}

func (api *Jikan) GetTopAnimeWithPage(page int) (types.TopAnimeResponse, error) {
	bytes, err := api.get(fmt.Sprintf(GET_TOP_ANIME_ENDPOINT, page))
	if err != nil {
		return types.TopAnimeResponse{}, err
	}

	var resp = types.TopAnimeResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.TopAnimeResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) GetTopManga() (types.TopMangaResponse, error) {
	return api.GetTopMangaWithPage(1)
}

func (api *Jikan) GetTopMangaWithPage(page int) (types.TopMangaResponse, error) {
	bytes, err := api.get(fmt.Sprintf(GET_TOP_MANGA_ENDPOINT, page))
	if err != nil {
		return types.TopMangaResponse{}, err
	}

	var resp = types.TopMangaResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.TopMangaResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) GetSchedule() (types.ScheduleResponse, error) {
	return api.GetScheduleOfDay("")
}

func (api *Jikan) GetScheduleOfDay(day string) (types.ScheduleResponse, error) {
	path := GET_SCHEDULE_ENDPOINT + day
	bytes, err := api.get(path)
	if err != nil {
		return types.ScheduleResponse{}, err
	}

	var resp = types.ScheduleResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.ScheduleResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) SearchAnime(key string) (types.AnimeSearchResultResponse, error) {
	q := url.Values{}
	q.Add("q", key)
	q.Add("page", strconv.Itoa(1))
	q.Add("limit", strconv.Itoa(10))

	bytes, err := api.getWithParam(SEARCH_ANIME_ENDPOINT, q)
	if err != nil {
		return types.AnimeSearchResultResponse{}, err
	}

	var resp = types.AnimeSearchResultResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.AnimeSearchResultResponse{}, err
	}

	return resp, nil
}

func (api *Jikan) SearchManga(key string) (types.MangaSearchResultResponse, error) {
	q := url.Values{}
	q.Add("q", key)
	q.Add("page", strconv.Itoa(1))
	q.Add("limit", strconv.Itoa(10))

	bytes, err := api.getWithParam(SEARCH_MANGA_ENDPOINT, q)
	if err != nil {
		return types.MangaSearchResultResponse{}, err
	}

	var resp = types.MangaSearchResultResponse{}
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return types.MangaSearchResultResponse{}, err
	}

	return resp, nil
}
