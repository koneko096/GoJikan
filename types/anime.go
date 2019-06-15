package types

type Anime struct {
	ID            int32        `json:"mal_id"`
	URL           string       `json:"url"`
	ImageURL      string       `json:"image_url"`
	TrailerURL    string       `json:"trailer_url"`
	Title         string       `json:"title"`
	TitleEnglish  string       `json:"title_english"`
	TitleJapanese string       `json:"title_japanese"`
	Type          string       `json:"type"`
	Source        string       `json:"source"`
	Status        string       `json:"status"`
	Episodes      int8         `json:"episodes"`
	Airing        bool         `json:"airing"`
	Duration      string       `json:"duration"`
	Rating        string       `json:"rating"`
	Score         float32      `json:"score"`
	ScoredBy      int32        `json:"scored_by"`
	Rank          int32        `json:"rank"`
	Popularity    int32        `json:"popularity"`
	Members       int32        `json:"members"`
	Favorites     int32        `json:"favorites"`
	Synopsis      string       `json:"synopsis"`
	Background    string       `json:"background"`
	Premiered     string       `json:"premiered"`
	Broadcast     string       `json:"broadcast"`
	Related       AnimeRelated `json:"related"`
	Producers     []Reference  `json:"producers"`
	Licencors     []Reference  `json:"licencors"`
	Studios       []Reference  `json:"studios"`
	Genres        []Reference  `json:"genres"`
	OpeningThemes []string     `json:"opening_themes"`
	EndingThemes  []string     `json:"ending_themes"`
	R18           bool         `json:"r18"`
	Kids          bool         `json:"kids"`
	Continuing    bool         `json:"continuing"`

	// TitleSyn string `json:"title"`
}

type AnimeRelated struct {
	Adaptation []Reference `json:"Adaptation"`
	SideStory  []Reference `json:"Side story"`
	Summary    []Reference `json:"Summary"`
}

type Reference struct {
	ID       int32  `json:"mal_id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Role     string `json:"role"`
}

type Picture struct {
	Large string `json:"large"`
	Small string `json:"small"`
}

type PictureResponse struct {
	Pictures []Picture `json:"pictures"`
}

type Video struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	VideoURL string `json:"video_url"`
}

type Episode struct {
	ID            int8   `json:"episode_id"`
	Episode       string `json:"episode"`
	URL           string `json:"url"`
	ImageURL      string `json:"image_url"`
	VideoURL      string `json:"video_url"`
	ForumURL      string `json:"forum_url"`
	Title         string `json:"title"`
	TitleRomanji  string `json:"title_romanji"`
	TitleJapanese string `json:"title_japanese"`
	Filler        bool   `json:"filler"`
	Recap         bool   `json:"recap"`
}

type VideoResponse struct {
	Promo    []Video   `json:"promo"`
	Episodes []Episode `json:"episodes"`
}
