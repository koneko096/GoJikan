package gojikan

const (
	BASE_URL = "http://api.jikan.moe/v3"
)

const (
	GET_ANIME_ENDPOINT                   = "/anime/%d"
	GET_ANIME_PICTURES_ENDPOINT          = "/anime/%d/pictures"
	GET_ANIME_VIDEOS_ENDPOINT            = "/anime/%d/videos"
	GET_ANIME_CHARACTERS_STAFFS_ENDPOINT = "/anime/%d/characters_staff"
	GET_CHARACTER_ENDPOINT               = "/character/%d"
	GET_SEASON_LIST_ENDPOINT             = "/season/%d/%s"
	GET_ANIME_FROM_GENRE_ENDPOINT        = "/genre/anime/%d/%d"
)
