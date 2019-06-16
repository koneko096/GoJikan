package gojikan

const (
	BASE_URL = "http://api.jikan.moe/v3"
)

const (
	GET_ANIME_ENDPOINT                   = "/anime/%d"
	GET_ANIME_PICTURES_ENDPOINT          = "/anime/%d/pictures"
	GET_ANIME_VIDEOS_ENDPOINT            = "/anime/%d/videos"
	GET_ANIME_CHARACTERS_STAFFS_ENDPOINT = "/anime/%d/characters_staff"
	GET_ANIME_FROM_GENRE_ENDPOINT        = "/genre/anime/%d/%d"
	GET_ANIME_FROM_PRODUCER_ENDPOINT     = "/producer/%d/%d"

	GET_MANGA_ENDPOINT               = "/manga/%d"
	GET_MANGA_PICTURES_ENDPOINT      = "/manga/%d/pictures"
	GET_MANGA_CHARACTERS_ENDPOINT    = "/manga/%d/characters"
	GET_MANGA_FROM_GENRE_ENDPOINT    = "/genre/manga/%d/%d"
	GET_MANGA_FROM_MAGAZINE_ENDPOINT = "/magazine/%d/%d"

	GET_CHARACTER_ENDPOINT   = "/character/%d"
	GET_SEASON_LIST_ENDPOINT = "/season/%d/%s"

	GET_TOP_ANIME_ENDPOINT = "/top/anime/%d"
	GET_TOP_MANGA_ENDPOINT = "/top/manga/%d"

	GET_SCHEDULE_ENDPOINT = "/schedule/"

	SEARCH_ANIME_ENDPOINT = "/search/anime"
	SEARCH_MANGA_ENDPOINT = "/search/manga"
)
