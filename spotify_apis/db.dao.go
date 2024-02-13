package main

type DaoTrack struct {
	Album            DaoAlbum
	AlbumID          string
	Artists          []DaoArtist
	AvailableMarkets []string
	DiscNumber       int
	DurationMs       int
	Explicit         bool
	ExternalIds      ExternalId
	ExternalUrls     ExternalUrl
	Href             string
	ID               string
	IsLocal          bool
	Name             string
	Popularity       int
	PreviewURL       any
	TrackNumber      int
	Type             string
	URI              string
}

type DaoAlbum struct {
	AlbumType            string
	Artists              []DaoArtist
	AvailableMarkets     []string
	ExternalUrls         ExternalUrl
	Href                 string
	ID                   string
	Images               []Image
	Name                 string
	ReleaseDate          string
	ReleaseDatePrecision string
	TotalTracks          int
	Type                 string
	URI                  string
}

type DaoArtist struct {
	Sec          int64
	TrackID      string
	AlbumID      string
	ExternalUrls ExternalUrl
	Href         string
	ID           string
	Name         string
	Type         string
	URI          string
}

func getDaoTrack(item Item) DaoTrack {
	daoTrack := DaoTrack{}
	daoTrack.Album = getDaoAlbum(item.Album)
	daoTrack.Artists = getDaoArtistList(item.Artists)
	daoTrack.DiscNumber = item.DiscNumber
	daoTrack.DurationMs = item.DurationMs
	daoTrack.Explicit = item.Explicit
	daoTrack.ExternalIds = item.ExternalIds
	daoTrack.ExternalUrls = item.ExternalUrls
	daoTrack.Href = item.Href
	daoTrack.ID = item.ID
	daoTrack.IsLocal = item.IsLocal
	daoTrack.Name = item.Name
	daoTrack.Popularity = item.Popularity
	daoTrack.PreviewURL = item.PreviewURL
	daoTrack.TrackNumber = item.TrackNumber
	daoTrack.Type = item.Type
	daoTrack.URI = item.URI
	return daoTrack
}
func getDaoAlbum(item Album) DaoAlbum {
	daoAlbum := DaoAlbum{}
	daoAlbum.AlbumType = item.AlbumType
	daoAlbum.Artists = getDaoArtistList(item.Artists)
	daoAlbum.AvailableMarkets = item.AvailableMarkets
	daoAlbum.ExternalUrls = item.ExternalUrls
	daoAlbum.Href = item.Href
	daoAlbum.ID = item.ID
	daoAlbum.Images = item.Images
	daoAlbum.Name = item.Name
	daoAlbum.ReleaseDate = item.ReleaseDate
	daoAlbum.ReleaseDatePrecision = item.ReleaseDatePrecision
	daoAlbum.TotalTracks = item.TotalTracks
	daoAlbum.Type = item.Type
	daoAlbum.URI = item.URI

	return daoAlbum
}
func getDaoArtistList(items []Artist) []DaoArtist {
	daoArtist := []DaoArtist{}
	for _, item := range items {
		daoArtist = append(daoArtist, getDaoArtist(item))
	}
	return daoArtist
}

func getDaoArtist(item Artist) DaoArtist {
	daoArtist := DaoArtist{}
	daoArtist.ExternalUrls = item.ExternalUrls
	daoArtist.ID = item.ID
	daoArtist.Name = item.Name
	daoArtist.Type = item.Type
	daoArtist.URI = item.URI
	return daoArtist
}
