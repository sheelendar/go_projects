
CREATE TABLE TRACKS(
	id VARCHAR(64) PRIMARY KEY  NOT NULL,
    album_id VARCHAR(64),
	name VARCHAR ( 255 ),
	disc_number bigint ,
    duration_ms bigint,
    href text,
    explicit boolean,
    is_local boolean,
    popularity smallint,
    preview_url text,
    track_number int,
    type VARCHAR(40),
	uri text
);


CREATE TABLE ALBUMS(
	id VARCHAR(64) PRIMARY KEY  NOT NULL,
    album_type VARCHAR(40),
	name VARCHAR ( 255 ),
    href text,
    total_tracks int,
    release_date VARCHAR(40),
    release_date_precision VARCHAR(40),
    type VARCHAR(40),
	uri text
);

CREATE TABLE ARTISTS(
	seq SERIAL PRIMARY KEY,
    id VARCHAR(64),
    track_id VARCHAR(64),
    album_id VARCHAR(64),
	name VARCHAR ( 255 ),
    href text,
    type VARCHAR(40),
	uri text,
    CONSTRAINT artist_id_track_id UNIQUE (id, track_id),
    CONSTRAINT album_id_id_track_id UNIQUE (id, album_id)
);

select id, album_id,name,disc_number,duration_ms,href,explicit,is_local,popularity,preview_url,track_number,type,uri from tracks where id=$1
select id, album_id,name,disc_number,duration_ms,href,explicit,is_local,popularity,preview_url,track_number,type,uri from tracks where name=$1
select TRACKS.id, TRACKS.album_id,TRACKS.name,TRACKS.disc_number,TRACKS.duration_ms,TRACKS.href,TRACKS.explicit,TRACKS.is_local,TRACKS.popularity,TRACKS.preview_url,TRACKS.track_number,TRACKS.type,TRACKS.uri from TRACKS INNER JOIN ARTISTS ON TRACKS.id=ARTISTS.track_id where ARTISTS.name='Eagles'; 