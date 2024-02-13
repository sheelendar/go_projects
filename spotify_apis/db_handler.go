package main

import (
	"context"
	"fmt"
	"log"
)

const (
	trackInsertQuery  = "INSERT INTO TRACKS (id, album_id,name,disc_number,duration_ms,href,explicit,is_local,popularity,preview_url,track_number,type,uri) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) ON CONFLICT (id) DO NOTHING"
	albumInsertQuery  = "INSERT INTO ALBUMS (id,name,href,album_type,total_tracks,release_date,release_date_precision,type,uri) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) ON CONFLICT (id) DO NOTHING"
	artistInsertQuery = "INSERT INTO ARTISTS (id,name,track_id,album_id,href,type,uri) VALUES ($1,$2,$3,$4,$5,$6,$7) ON CONFLICT DO NOTHING"

	searchTrackByIDQuery         = `select id, album_id,name,disc_number,duration_ms,href,explicit,is_local,popularity,preview_url,track_number,type,uri from tracks where id=$1 OR name=$2`
	searchTrackByArtistNameQuery = `select TRACKS.id,TRACKS.album_id,TRACKS.name,TRACKS.disc_number,TRACKS.duration_ms,TRACKS.href,TRACKS.explicit,TRACKS.is_local,TRACKS.popularity,TRACKS.preview_url,TRACKS.track_number,TRACKS.type,TRACKS.uri from TRACKS INNER JOIN ARTISTS ON TRACKS.id=ARTISTS.track_id where ARTISTS.name=$1`

	searchAlbumByIDQuery = `select id,name,href,album_type,total_tracks,release_date,release_date_precision,type,uri from ALBUMS where id=$1`

	searchArtistByAlbumIDQuery = `select id,name,COALESCE(track_id,''),COALESCE(album_id,''),href,type,uri from ARTISTS where album_id=$1`
	searchArtistByTrackIDQuery = `select id,name,COALESCE(track_id,''),COALESCE(album_id,''),href,type,uri from ARTISTS where track_id=$1`
)

// insertTrackIntoDB insert all track into db.
// We can use bulk insert inplace of single insert here for performance.
func insertTrackIntoDB(ctx context.Context, daoTracklist []DaoTrack) map[string]error {
	errArtistIDMap := make(map[string]error)
	for _, daoTrack := range daoTracklist {
		args := getTrackArgs(daoTrack)
		res, err := dbPool.Exec(ctx, trackInsertQuery, args...)

		if err != nil {
			fmt.Println("unable to insert row: %w", err)
			errArtistIDMap[daoTrack.AlbumID] = err
		}
		fmt.Println(res)
	}
	return errArtistIDMap
}

// insertAlbumIntoDB insert album detail into db.
func insertAlbumIntoDB(ctx context.Context, daoAlbum DaoAlbum) error {
	args := getAlbumArgs(daoAlbum)
	res, err := dbPool.Exec(ctx, albumInsertQuery, args...)

	if err != nil {
		fmt.Println("unable to insert row: %w", err)
		return err
	}
	fmt.Println(res)
	return nil
}

// insertArtistIntoDB insert artist into db.
// here as well we can use bulk insert into db for performance.
func insertArtistIntoDB(ctx context.Context, daoArtistList []DaoArtist, trackID, albumID string) {
	for _, daoArtist := range daoArtistList {
		args := getArtistArgs(daoArtist, trackID, albumID)
		res, err := dbPool.Exec(ctx, artistInsertQuery, args...)

		if err != nil {
			fmt.Println("unable to insert row: %w", err)
		}
		fmt.Println(res)
	}
}

func getArtistArgs(daoArtist DaoArtist, trackID, albumID string) []interface{} {
	var arr []interface{}
	arr = append(arr, daoArtist.ID)
	arr = append(arr, daoArtist.Name)
	if trackID == "" {
		arr = append(arr, nil)
	} else {
		arr = append(arr, trackID)
	}
	if albumID == "" {
		arr = append(arr, nil)
	} else {
		arr = append(arr, albumID)
	}
	arr = append(arr, daoArtist.Href)
	arr = append(arr, daoArtist.Type)
	arr = append(arr, daoArtist.URI)
	return arr
}

func getAlbumArgs(daoAlbum DaoAlbum) []interface{} {
	var arr []interface{}
	arr = append(arr, daoAlbum.ID)
	arr = append(arr, daoAlbum.Name)
	arr = append(arr, daoAlbum.Href)
	arr = append(arr, daoAlbum.AlbumType)
	arr = append(arr, daoAlbum.TotalTracks)
	arr = append(arr, daoAlbum.ReleaseDate)
	arr = append(arr, daoAlbum.ReleaseDatePrecision)
	arr = append(arr, daoAlbum.Type)
	arr = append(arr, daoAlbum.URI)
	return arr
}

func getTrackArgs(daoTrack DaoTrack) []interface{} {
	var arr []interface{}
	arr = append(arr, daoTrack.ID)
	arr = append(arr, daoTrack.Album.ID)
	arr = append(arr, daoTrack.Name)
	arr = append(arr, daoTrack.DiscNumber)
	arr = append(arr, daoTrack.DurationMs)
	arr = append(arr, daoTrack.Href)
	arr = append(arr, daoTrack.Explicit)
	arr = append(arr, daoTrack.IsLocal)
	arr = append(arr, daoTrack.Popularity)
	arr = append(arr, daoTrack.PreviewURL)
	arr = append(arr, daoTrack.TrackNumber)
	arr = append(arr, daoTrack.Type)
	arr = append(arr, daoTrack.URI)
	return arr
}

// SearchTracksByArtistName search Tracks into db using artist name and return list.
func SearchTracksByArtistName(ctx context.Context, name string) []DaoTrack {
	rows, err := dbPool.Query(context.Background(), searchTrackByArtistNameQuery, name)
	if err != nil {
		log.Fatal("Error executing query in SearchTracksByArtistName :", err)
	}
	defer rows.Close()
	var trackList []DaoTrack
	for rows.Next() {
		track := DaoTrack{}
		err := rows.Scan(&track.ID, &track.AlbumID, &track.Name, &track.DiscNumber, &track.DurationMs, &track.Href, &track.Explicit, &track.IsLocal, &track.Popularity, &track.PreviewURL, &track.TrackNumber, &track.Type, &track.URI)
		if err != nil {
			log.Fatal("Error scanning row in  SearchTracksByArtistName:", err)
		} else {
			trackList = append(trackList, track)
		}
	}
	return trackList
}

// SearchTracksByIDOrName search track by Id or name into db and return Track.
func SearchTracksByIDOrName(ctx context.Context, id, name string) DaoTrack {
	rows, err := dbPool.Query(context.Background(), searchTrackByIDQuery, id, name)
	if err != nil {
		log.Fatal("Error executing query in SearchTracksByIDOrName :", err)
	}
	defer rows.Close()
	track := DaoTrack{}
	if rows.Next() {
		err := rows.Scan(&track.ID, &track.AlbumID, &track.Name, &track.DiscNumber, &track.DurationMs, &track.Href, &track.Explicit,
			&track.IsLocal, &track.Popularity, &track.PreviewURL, &track.TrackNumber, &track.Type, &track.URI)
		if err != nil {
			log.Fatal("Error scanning row in SearchTracksByIDOrName :", err)
		}
	}
	return track
}

// SearchAlbumByID search album into db and return album.
func SearchAlbumByID(ctx context.Context, id string) DaoAlbum {
	rows, err := dbPool.Query(context.Background(), searchAlbumByIDQuery, id)
	if err != nil {
		log.Fatal("Error executing query in SearchAlbumByID :", err)
	}
	defer rows.Close()
	album := DaoAlbum{}
	if rows.Next() {
		//select id,name,href,album_type,total_tracks,release_date,release_date_precision,type,uri
		err := rows.Scan(&album.ID, &album.Name, &album.Href, &album.AlbumType, &album.TotalTracks, &album.ReleaseDate, &album.ReleaseDatePrecision, &album.Type, &album.URI)
		if err != nil {
			log.Fatal("Error scanning row int SearchAlbumByID :", err)
		}
	}
	return album
}

// SearchArtistByAlbumID search Artist by albumID and return artistList.
func SearchArtistByAlbumID(ctx context.Context, albumID string) []DaoArtist {
	rows, err := dbPool.Query(context.Background(), searchArtistByAlbumIDQuery, albumID)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()
	var artistLsit []DaoArtist
	for rows.Next() {
		artist := DaoArtist{}
		err := rows.Scan(&artist.ID, &artist.Name, &artist.TrackID, &artist.AlbumID, &artist.Href, &artist.Type, &artist.URI)
		if err != nil {
			log.Fatal("Error scanning row in SearchArtistByAlbumID :", err)
		} else {
			artistLsit = append(artistLsit, artist)
		}
	}
	return artistLsit
}

// SearchArtistByTrackID search Artist by trackId into db and return list.
func SearchArtistByTrackID(ctx context.Context, trackID string) []DaoArtist {
	rows, err := dbPool.Query(context.Background(), searchArtistByTrackIDQuery, trackID)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()
	var artistLsit []DaoArtist
	for rows.Next() {
		artist := DaoArtist{}
		err := rows.Scan(&artist.ID, &artist.Name, &artist.TrackID, &artist.AlbumID, &artist.Href, &artist.Type, &artist.URI)
		if err != nil {
			log.Fatal("Error scanning row in SearchArtistByTrackID :", err)
		} else {
			artistLsit = append(artistLsit, artist)
		}
	}
	return artistLsit
}
