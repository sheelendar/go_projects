package main

import (
	"context"
)

// insertDaoHandler handle external insert request and call DB handler for insert.
func insertDaoHandler(ctx context.Context, items []Item) {
	daoTrackList := make([]DaoTrack, len(items))
	for index, item := range items {
		daoTrackList[index] = getDaoTrack(item)
	}

	errArtistIDMap := insertTrackIntoDB(ctx, daoTrackList)
	for _, daoTack := range daoTrackList {
		if _, ok := errArtistIDMap[daoTack.AlbumID]; !ok {
			err := insertAlbumIntoDB(ctx, daoTack.Album)
			if err == nil {
				insertArtistIntoDB(ctx, daoTack.Album.Artists, "", daoTack.Album.ID)
			}
			insertArtistIntoDB(ctx, daoTack.Album.Artists, daoTack.ID, "")
		}
	}
}

// GetTrackByIDOrName handle external Get request and call DB handler for Get.
func GetTrackByIDOrName(ctx context.Context, id, name string) Item {
	track := SearchTracksByIDOrName(ctx, id, name)
	FetAlbumAndArtists(ctx, &track)
	return GetItemByDaoTrack(track)
}

// GetTracksByArtistName handle external Get tracks request and call DB handler for Get.
func GetTracksByArtistName(ctx context.Context, name string) []Item {
	trackList := SearchTracksByArtistName(ctx, name)
	itemList := make([]Item, len(trackList))
	for i := 0; i < len(trackList); i++ {
		FetAlbumAndArtists(ctx, &trackList[i])
		itemList[i] = GetItemByDaoTrack(trackList[i])
	}
	return itemList
}

// FetAlbumAndArtists fetch all required detail for a Tracks from DB.
func FetAlbumAndArtists(ctx context.Context, track *DaoTrack) {
	album := SearchAlbumByID(ctx, track.AlbumID)
	album.Artists = SearchArtistByAlbumID(ctx, track.AlbumID)
	track.Album = album
	track.Artists = SearchArtistByTrackID(ctx, track.ID)
}
