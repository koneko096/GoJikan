package types

type Staff struct {
	ID        int32    `json:"mal_id"`
	Name      string   `json:"name"`
	URL       string   `json:"url"`
	ImageURL  string   `json:"image_url"`
	Positions []string `json:"positions"`
}

type Character struct {
	ID              int32        `json:"mal_id"`
	Name            string       `json:"name"`
	NameKanji       string       `json:"name_kanji"`
	Nicknames       []string     `json:"nicknames"`
	About           string       `json:"about"`
	MemberFavorites int32        `json:"member_favorites"`
	URL             string       `json:"url"`
	ImageURL        string       `json:"image_url"`
	Role            string       `json:"role"`
	VoiceActors     []VoiceActor `json:"voice_actors"`
	Animeography    []Reference  `json:"animeograpy"`
	Mangaography    []Reference  `json:"mangaograpy"`
}

type VoiceActor struct {
	ID       int32  `json:"mal_id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Language string `json:"language"`
}

type CharactersAndStaffsResponse struct {
	Characters []Character `json:"characters"`
	Staffs     []Staff     `json:"staff"`
}
