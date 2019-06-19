package gojikan

import (
	"net/url"
	"testing"
)

const MOCK_BASE_URL = "http://base.url"

type MockClientHelper struct {
	Result []byte
}

func NewMockClientHelper(result []byte) ClientHelper {
	return &MockClientHelper{
		Result: result,
	}
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

func TestGetAnime(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"mal_id\": 123}")))
	anime, _ := jikan.GetAnime(1)
	if anime.ID != 123 {
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

func TestGetCharacter(t *testing.T) {
	jikan := NewJikanWithClientHelper(NewMockClientHelper([]byte("{\"mal_id\": 123}")))
	char, _ := jikan.GetCharacter(1)
	if char.ID != 123 {
		t.Error("meong")
	}
}
