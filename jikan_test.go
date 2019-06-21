package gojikan

import (
	"errors"
	"net/url"
	"testing"
)

const MOCK_BASE_URL = "http://base.url"

type MockClientHelper struct {
	Result []byte
}

type FailedClientHelper struct {
}

func NewMockClientHelper(result []byte) ClientHelper {
	return &MockClientHelper{
		Result: result,
	}
}

func NewFailedClientHelper() ClientHelper {
	return &FailedClientHelper{}
}

func NewJikanWithClientHelper(clientHelper ClientHelper) *Jikan {
	return &Jikan{
		baseUrl:      MOCK_BASE_URL,
		clientHelper: clientHelper,
	}
}

func (m *MockClientHelper) Get(endpoint string, params url.Values) ([]byte, error) {
	return m.Result, nil
}

func (m *FailedClientHelper) Get(endpoint string, params url.Values) ([]byte, error) {
	return nil, errors.New("Some error")
}

func TestGetAnime(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"mal_id\": 123}")))
	anime, _ := jikan.GetAnime(1)
	if anime.ID != 123 {
		t.Error("meong")
	}
}

func TestGetAnime_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetAnime(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetAnime_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetAnime(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetManga(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"mal_id\": 123}")))
	manga, _ := jikan.GetManga(1)
	if manga.ID != 123 {
		t.Error("meong")
	}
}

func TestGetManga_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetManga(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetManga_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetManga(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetCharacter(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"mal_id\": 123}")))
	char, _ := jikan.GetCharacter(1)
	if char.ID != 123 {
		t.Error("meong")
	}
}

func TestGetCharacter_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetCharacter(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetCharacter_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetCharacter(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetAnimeFromGenre(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"mal_url\": {\"mal_id\": 123}}")))
	resp, _ := jikan.GetAnimeFromGenre(1)
	if resp.MalURL.ID != 123 {
		t.Error("meong")
	}
}

func TestGetAnimeFromGenre_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetAnimeFromGenre(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetAnimeFromGenre_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetAnimeFromGenre(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetAnimeFromProducer(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"meta\": {\"mal_id\": 123}}")))
	resp, _ := jikan.GetAnimeFromProducer(1)
	if resp.Meta.ID != 123 {
		t.Error("meong")
	}
}

func TestGetAnimeFromProducer_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetAnimeFromProducer(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetAnimeFromProducer_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetAnimeFromProducer(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetMangaFromGenre(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"mal_url\": {\"mal_id\": 123}}")))
	resp, _ := jikan.GetMangaFromGenre(1)
	if resp.MalURL.ID != 123 {
		t.Error("meong")
	}
}

func TestGetMangaFromGenre_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetMangaFromGenre(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetMangaFromGenre_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetMangaFromGenre(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetMangaFromMagazine(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"meta\": {\"mal_id\": 123}}")))
	resp, _ := jikan.GetMangaFromMagazine(1)
	if resp.Meta.ID != 123 {
		t.Error("meong")
	}
}

func TestGetMangaFromMagazine_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetMangaFromMagazine(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetMangaFromMagazine_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetMangaFromMagazine(1)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetTopAnime(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"top\": [{\"mal_id\": 123}]}")))
	resp, _ := jikan.GetTopAnime()
	if resp.Top[0].ID != 123 {
		t.Error("meong")
	}
}

func TestGetTopAnime_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetTopAnime()
	if err == nil {
		t.Error("meong")
	}
}

func TestGetTopAnime_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetTopAnime()
	if err == nil {
		t.Error("meong")
	}
}

func TestGetTopManga(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"top\": [{\"mal_id\": 123}]}")))
	resp, _ := jikan.GetTopManga()
	if resp.Top[0].ID != 123 {
		t.Error("meong")
	}
}

func TestGetTopManga_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetTopManga()
	if err == nil {
		t.Error("meong")
	}
}

func TestGetTopManga_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetTopManga()
	if err == nil {
		t.Error("meong")
	}
}

func TestSearchAnime(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"results\": [{\"mal_id\": 123}]}")))
	resp, _ := jikan.SearchAnime("abc")
	if resp.Results[0].ID != 123 {
		t.Error("meong")
	}
}

func TestSearchAnime_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.SearchAnime("abc")
	if err == nil {
		t.Error("meong")
	}
}

func TestSearchAnime_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.SearchAnime("abc")
	if err == nil {
		t.Error("meong")
	}
}

func TestSearchManga(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"results\": [{\"mal_id\": 123}]}")))
	resp, _ := jikan.SearchManga("abc")
	if resp.Results[0].ID != 123 {
		t.Error("meong")
	}
}

func TestSearchManga_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.SearchManga("abc")
	if err == nil {
		t.Error("meong")
	}
}

func TestSearchManga_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.SearchManga("abc")
	if err == nil {
		t.Error("meong")
	}
}

func TestGetAnimePictures(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"pictures\": [{\"large\": \"Yoda\"}]}")))
	resp, _ := jikan.GetAnimePictures(123)
	if resp.Pictures[0].Large != "Yoda" {
		t.Error("meong")
	}
}

func TestGetAnimePictures_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetAnimePictures(123)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetAnimePictures_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetAnimePictures(123)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetMangaPictures(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"pictures\": [{\"large\": \"Yodi\"}]}")))
	resp, _ := jikan.GetMangaPictures(123)
	if resp.Pictures[0].Large != "Yodi" {
		t.Error("meong")
	}
}

func TestGetMangaPictures_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetMangaPictures(123)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetMangaPictures_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetMangaPictures(123)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetAnimeVideos(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"promo\": [{\"title\": \"Skywalker\"}]}")))
	resp, _ := jikan.GetAnimeVideos(123)
	if resp.Promo[0].Title != "Skywalker" {
		t.Error("meong")
	}
}

func TestGetAnimeVideos_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetAnimeVideos(123)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetAnimeVideos_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetAnimeVideos(123)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetAnimeCharactersAndStaffs(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"characters\": [{\"name\": \"Souma\"}]}")))
	resp, _ := jikan.GetAnimeCharactersAndStaffs(123)
	if resp.Characters[0].Name != "Souma" {
		t.Error("meong")
	}
}

func TestGetAnimeCharactersAndStaffs_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetAnimeCharactersAndStaffs(123)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetAnimeCharactersAndStaffs_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetAnimeCharactersAndStaffs(123)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetMangaCharactersAndStaffs(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"characters\": [{\"name\": \"Kuroko\"}]}")))
	resp, _ := jikan.GetMangaCharactersAndStaffs(123)
	if resp.Characters[0].Name != "Kuroko" {
		t.Error("meong")
	}
}

func TestGetMangaCharactersAndStaffs_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetMangaCharactersAndStaffs(123)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetMangaCharactersAndStaffs_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetMangaCharactersAndStaffs(123)
	if err == nil {
		t.Error("meong")
	}
}

func TestGetSchedule(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"monday\": [{\"mal_id\": 123}]}")))
	resp, _ := jikan.GetSchedule()
	if resp.Monday[0].ID != 123 {
		t.Error("meong")
	}
}

func TestGetSchedule_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetSchedule()
	if err == nil {
		t.Error("meong")
	}
}

func TestGetSchedule_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetSchedule()
	if err == nil {
		t.Error("meong")
	}
}

func TestGetSeasonList(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"season_year\": 2019}")))
	resp, _ := jikan.GetSeasonList(2019, "")
	if resp.Year != 2019 {
		t.Error("meong")
	}
}

func TestGetSeasonList_clientError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewFailedClientHelper())
	_, err := jikan.GetSeasonList(2019, "")
	if err == nil {
		t.Error("meong")
	}
}

func TestGetSeasonList_parseError(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("[]")))
	_, err := jikan.GetSeasonList(2019, "")
	if err == nil {
		t.Error("meong")
	}
}
