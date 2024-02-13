package main

type UnderlyingAsset struct {
}

type SearchRequestByIDOrName struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SearchRequestByArtistName struct {
	Name string `json:"artist_name"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type SearchTrackResponse struct {
	Tracks Tracks `json:"tracks"`
}

type Tracks struct {
	Href     string `json:"href"`
	Items    []Item `json:"items"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous any    `json:"previous"`
	Total    int    `json:"total"`
}

type Item struct {
	Album            Album       `json:"album"`
	Artists          []Artist    `json:"artists"`
	AvailableMarkets []string    `json:"available_markets"`
	DiscNumber       int         `json:"disc_number"`
	DurationMs       int         `json:"duration_ms"`
	Explicit         bool        `json:"explicit"`
	ExternalIds      ExternalId  `json:"external_ids"`
	ExternalUrls     ExternalUrl `json:"external_urls"`
	Href             string      `json:"href"`
	ID               string      `json:"id"`
	IsLocal          bool        `json:"is_local"`
	Name             string      `json:"name"`
	Popularity       int         `json:"popularity"`
	PreviewURL       any         `json:"preview_url"`
	TrackNumber      int         `json:"track_number"`
	Type             string      `json:"type"`
	URI              string      `json:"uri"`
}

type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type Album struct {
	AlbumType            string      `json:"album_type"`
	Artists              []Artist    `json:"artists"`
	AvailableMarkets     []string    `json:"available_markets"`
	ExternalUrls         ExternalUrl `json:"external_urls"`
	Href                 string      `json:"href"`
	ID                   string      `json:"id"`
	Images               []Image     `json:"images"`
	Name                 string      `json:"name"`
	ReleaseDate          string      `json:"release_date"`
	ReleaseDatePrecision string      `json:"release_date_precision"`
	TotalTracks          int         `json:"total_tracks"`
	Type                 string      `json:"type"`
	URI                  string      `json:"uri"`
}

type Artist struct {
	ExternalUrls ExternalUrl `json:"external_urls"`
	Href         string      `json:"href"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
}

type ExternalId struct {
	Isrc string `json:"isrc"`
}

type ExternalUrl struct {
	Spotify string `json:"spotify"`
}

func GetItemByDaoTrack(track DaoTrack) Item {
	item := Item{}
	item.Album = getAlbumByDaoAlbum(track.Album)
	item.Artists = getArtistListByDaoArtits(track.Artists)
	item.DiscNumber = track.DiscNumber
	item.DurationMs = track.DurationMs
	item.Explicit = track.Explicit
	item.ExternalIds = track.ExternalIds
	item.ExternalUrls = track.ExternalUrls
	item.Href = track.Href
	item.ID = track.ID
	item.IsLocal = track.IsLocal
	item.Name = track.Name
	item.Popularity = track.Popularity
	item.PreviewURL = track.PreviewURL
	item.TrackNumber = track.TrackNumber
	item.Type = track.Type
	item.URI = track.URI
	return item
}

func getAlbumByDaoAlbum(item DaoAlbum) Album {
	album := Album{}
	album.AlbumType = item.AlbumType
	album.Artists = getArtistListByDaoArtits(item.Artists)
	album.AvailableMarkets = item.AvailableMarkets
	album.ExternalUrls = item.ExternalUrls
	album.Href = item.Href
	album.ID = item.ID
	album.Images = item.Images
	album.Name = item.Name
	album.ReleaseDate = item.ReleaseDate
	album.ReleaseDatePrecision = item.ReleaseDatePrecision
	album.TotalTracks = item.TotalTracks
	album.Type = item.Type
	album.URI = item.URI

	return album
}
func getArtistListByDaoArtits(items []DaoArtist) []Artist {
	artistList := []Artist{}
	for _, item := range items {
		artistList = append(artistList, getArtistByDaoArtist(item))
	}
	return artistList
}

func getArtistByDaoArtist(item DaoArtist) Artist {
	artist := Artist{}
	artist.ExternalUrls = item.ExternalUrls
	artist.ID = item.ID
	artist.Name = item.Name
	artist.Type = item.Type
	artist.URI = item.URI
	return artist
}
