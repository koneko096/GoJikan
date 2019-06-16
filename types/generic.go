package types

import "encoding/json"

const (
	UPCOMING   = "upcoming"
	COMPLETED  = "completed"
	PUBLISHING = "publishing"
	AIRING     = "airing"
)

type GenreId int

const (
	ACTION GenreId = iota + 1
	ADVENTURE
	CARS
	COMEDY
	DEMENTIA
	DEMONS
	MYSTERY
	DRAMA
	ECCHI
	FANTASY
	GAME
	HENTAI
	HISTORICAL
	HORROR
	KIDS
	MAGIC
	MARTIAL_ARTS
	MECHA
	MUSIC
	PARODY
	SAMURAI
	ROMANCE
	SCHOOL
	SCI_FI
	SHOUJO
	SHOUJO_AI
	SHOUNEN
	SHOUNEN_AI
	SPACE
	SPORTS
	SUPER_POWER
	VAMPIRE
	YAOI
	YURI
	HAREM
	SLICE_OF
	SUPERNATURAL
	MILITARY
	POLICE
	PSYCHOLOGICAL
)

type DayOfWeek int

const (
	MONDAY DayOfWeek = iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
	OTHER
	UNKNOWN
)

func (t DayOfWeek) String() string {
	return [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday", "other", "unknown"}[t]
}

type ErrorResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
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

type ScheduleResponse struct {
	Monday    []Anime `json:"monday"`
	Tuesday   []Anime `json:"tuesday"`
	Wednesday []Anime `json:"wednesday"`
	Thursday  []Anime `json:"thursday"`
	Friday    []Anime `json:"friday"`
	Saturday  []Anime `json:"saturday"`
	Sunday    []Anime `json:"sunday"`
	Other     []Anime `json:"other"`
	Unknown   []Anime `json:"unknown"`
}
