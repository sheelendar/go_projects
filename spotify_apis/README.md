## 1. Install Go version if don't have and set GOPATH and GOROOT
## 2.Go to directory
cd $GOPATH/Go/src/go_projects/spotify_apis

## 3. Run the commnad on terminal for part 1 
docker-compose up

 ## 4. Run the commnad on terminal for part 2
 go run .

# 5  Request:
Get Request URL: localhost:8080/search_track
Body: 
{
    "id":"",
    "name":"Dreams - 2004 Remaster"
}

# 6 Request:
Get Request URL: localhost:8080/search_track_artist
{
    "artist_name":"Eagles"
}

# 7 Request:
GetRequest URL it fetch track from spotify and insert into db: 
localhost:8080/search_spotify?query=remaster