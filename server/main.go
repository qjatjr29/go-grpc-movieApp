package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	moviepb "github.com/qjatjr29/go-grpc-movieapp/proto/movie"
	"google.golang.org/grpc"
)

const portNumber = "9001"

type movieServer struct {
	moviepb.MovieServer
}

// 인기,상영중,개봉할 영화 목록
type MovieLists struct {
	Results []struct {
		ID           int     `json:"id"`
		Title        string  `json:"title"`
		Release_date string  `json:"release_date"`
		Poster_path  string  `json:"poster_path"`
		Overview     string  `json:"overview"`
		Vote_average float32 `json:"vote_average`
	} `json:"results`
}
// 검색한 영화 목록
type SearchMovies struct {
	Results []struct {
		ID           int     `json:"id"`
		Title        string  `json:"title"`
		Release_date string  `json:"release_date"`
		Poster_path  string  `json:"poster_path"`
		Vote_average float32 `json:"vote_average`
	} `json:"results`
}
// 영화 detail 
type DetailMovies struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Release_date string  `json:"release_date"`
	Poster_path  string  `json:"poster_path"`
	Overview string `json:"overview"`
	Homepage string `json:"homepage"`
	Genres []struct {
		ID int `json:"id"`
		Name string `json="name"`
	}
	Backdrop_path string `json:"backdrop_path"`
	Runtime int32 `json:"runtime"`
	Imdb_id string `json:"imdb_id"`
	Production []struct{
		ID int `json:"id"`
		Logo_path string `json:"logo_path"`
		Name string `json:"name"`
	} `json:"production_companies"`
}
// 배우, 스태프 

type Cast struct {
	CastList []struct {
		ID int32 `json: "id"`
		Name string `json:"name"`
		Profile_path string `json:"profile_path"`
		Character string `json:"character"`
	} `json:"cast"`
}
type Crew struct {
	CrewList []struct {
		ID int32 `json:"id"`
		Name string `json:"name"`
		Profile_path string `json:"profile_path"`
		Job string `json:"job"`
	}`json:"crew"`
}

type Video struct{
	Results []struct {
		Key string `json:"key"`
		Name string `json:"name"`
	}
}


// Get Movie Detail
func (s *movieServer) GetMovie(ctx context.Context, req *moviepb.GetMovieRequest) (*moviepb.GetMovieResponse, error) {
	
	movieID := req.MovieId
	fmt.Println(movieID)
	url := "https://api.themoviedb.org/3/movie/"
	url += movieID
	url += "?api_key=b71e7ddc0337840f4f46be79b18e4c41&language=en-US"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result DetailMovies
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	// MovieMessages := make([]*moviepb.MovieMessage, len(result.Results))

	
	var MovieDetail = &moviepb.MovieMessage{}
	MovieDetail.Id = int32(result.ID)
	MovieDetail.Title = result.Title
	MovieDetail.PosterPath = result.Poster_path
	MovieDetail.ReleaseDate = result.Release_date
	MovieDetail.BackdropPath = result.Backdrop_path
	MovieDetail.ImdbId = result.Imdb_id
	MovieDetail.Runtime = result.Runtime

	genresList := make([]*moviepb.Genres,len(result.Genres))
	for i,rec := range result.Genres{
		var genre = &moviepb.Genres{}
		genre.Id = int32(rec.ID)
		genre.Name = rec.Name
		genresList[i]=genre
	}
	productionList := make([]*moviepb.Production,len(result.Production))
	for i,rec := range result.Production{
		var production = &moviepb.Production{}
		production.Id = int32(rec.ID)
		production.Name = rec.Name
		production.LogoPath=rec.Logo_path
		productionList[i]=production
	}
	MovieDetail.Genres=genresList
	MovieDetail.Production=productionList
	MovieDetail.Homepage=result.Homepage
	MovieDetail.Overview=result.Overview
	
	return &moviepb.GetMovieResponse{
		MovieMessage: MovieDetail,
	}, nil
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// List Popular Movies
func (s *movieServer) ListPopularMovies(ctx context.Context, req *moviepb.ListPopularMovieRequest) (*moviepb.ListPopularMovieResponse, error) {

	resp, err := http.Get("https://api.themoviedb.org/3/movie/popular?api_key=b71e7ddc0337840f4f46be79b18e4c41&language=en-US&page=1")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result MovieLists
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	popularMovieMessages := make([]*moviepb.MovieListsMessage, len(result.Results))

	for i, rec := range result.Results {
		var popularMovie = &moviepb.MovieListsMessage{}
		popularMovie.Id = int32(rec.ID)
		popularMovie.Title = rec.Title
		popularMovie.PosterPath = rec.Poster_path
		popularMovie.ReleaseDate = rec.Release_date
		popularMovie.Overview = rec.Overview
		popularMovie.VoteAverage = rec.Vote_average
		popularMovieMessages[i] = popularMovie
	}
	return &moviepb.ListPopularMovieResponse{
		PopularmovieMessage: popularMovieMessages,
	}, nil
}
// List playing Movies
func (s *movieServer) ListPlayingMovies(ctx context.Context, req *moviepb.ListPlayingMovieRequest) (*moviepb.ListPlayingMovieResponse, error) {

	resp, err := http.Get("https://api.themoviedb.org/3/movie/now_playing?api_key=b71e7ddc0337840f4f46be79b18e4c41&language=en-US&page=1")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result MovieLists
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	playingMovieMessages := make([]*moviepb.MovieListsMessage, len(result.Results))

	for i, rec := range result.Results {
		var playingMovie = &moviepb.MovieListsMessage{}
		playingMovie.Id = int32(rec.ID)
		playingMovie.Title = rec.Title
		playingMovie.PosterPath = rec.Poster_path
		playingMovie.ReleaseDate = rec.Release_date
		playingMovie.Overview = rec.Overview
		playingMovie.VoteAverage = rec.Vote_average
		playingMovieMessages[i] = playingMovie
	}
	return &moviepb.ListPlayingMovieResponse{
		PlayingmovieMessage: playingMovieMessages,
	}, nil
}
// List upcoming Movies
func (s *movieServer) ListUpcomingMovies(ctx context.Context, req *moviepb.ListUpcomingMovieRequest) (*moviepb.ListUpcomingMovieResponse, error) {

	resp, err := http.Get("https://api.themoviedb.org/3/movie/upcoming?api_key=b71e7ddc0337840f4f46be79b18e4c41&language=en-US&page=1")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result MovieLists
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	upcomingMovieMessages := make([]*moviepb.MovieListsMessage, len(result.Results))

	for i, rec := range result.Results {
		var upcomingMovie = &moviepb.MovieListsMessage{}
		upcomingMovie.Id = int32(rec.ID)
		upcomingMovie.Title = rec.Title
		upcomingMovie.PosterPath = rec.Poster_path
		upcomingMovie.ReleaseDate = rec.Release_date
		upcomingMovie.Overview = rec.Overview
		upcomingMovie.VoteAverage = rec.Vote_average
		upcomingMovieMessages[i] = upcomingMovie
	}
	return &moviepb.ListUpcomingMovieResponse{
		UpcomingmovieMessage: upcomingMovieMessages,
	}, nil
}



// List Search Result Movies
func (s *movieServer) SearchMovies(ctx context.Context, req *moviepb.ListSearchMovieRequest) (*moviepb.ListSearchMovieResponse, error) {
	fmt.Println("check")
	movieID := req.SearchMovieId
	fmt.Println(movieID)
	url := "https://api.themoviedb.org/3/search/movie?api_key=b71e7ddc0337840f4f46be79b18e4c41&language=en-US&query="
	url += movieID
	url += "&page=1&include_adult=false"
	fmt.Println(url)
	resp, err := http.Get("https://api.themoviedb.org/3/search/movie?api_key=b71e7ddc0337840f4f46be79b18e4c41&language=en-US&query=" + movieID + "&page=1&include_adult=false")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result SearchMovies
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	SearchMovieMessages := make([]*moviepb.SearchMoviesMessage, len(result.Results))

	for i, rec := range result.Results {
		var searchMovie = &moviepb.SearchMoviesMessage{}
		searchMovie.Id = int32(rec.ID)
		searchMovie.Title = rec.Title
		searchMovie.PosterPath = rec.Poster_path
		searchMovie.ReleaseDate = rec.Release_date
		searchMovie.VoteAverage = rec.Vote_average
		SearchMovieMessages[i] = searchMovie
	}
	return &moviepb.ListSearchMovieResponse{
		SearchmovieMessage: SearchMovieMessages,
	}, nil
}

// List Actors in Movies
func (s *movieServer) CastMovies(ctx context.Context, req *moviepb.CastMovieRequest) (*moviepb.CastMovieResponse, error) {
	fmt.Println("check")
	movieID := req.MovieId
	fmt.Println(movieID)
	url := "https://api.themoviedb.org/3/movie/"
	url += movieID
	url += "/credits?api_key=b71e7ddc0337840f4f46be79b18e4c41&language=en-US"
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result Cast
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	CastMovieMessages := make([]*moviepb.CastMovieMessage, len(result.CastList))
	for i, rec := range result.CastList {
		var castActor = &moviepb.CastMovieMessage{}
		castActor.Id = int32(rec.ID)
		castActor.Name = rec.Name
		castActor.ProfilePath = rec.Profile_path
		castActor.Character = rec.Character
		CastMovieMessages[i] = castActor
	}
	return &moviepb.CastMovieResponse{
		CastmovieMessage: CastMovieMessages,
	}, nil
}

// List staff in Movies
func (s *movieServer) CrewMovies(ctx context.Context, req *moviepb.CrewMovieRequest) (*moviepb.CrewMovieResponse, error) {
	fmt.Println("check")
	movieID := req.MovieId
	fmt.Println(movieID)
	url := "https://api.themoviedb.org/3/movie/"
	url += movieID
	url += "/credits?api_key=b71e7ddc0337840f4f46be79b18e4c41&language=en-US"
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result Crew
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	CrewMovieMessages := make([]*moviepb.CrewMovieMessage, len(result.CrewList))

	for i, rec := range result.CrewList {
		fmt.Println(rec)
		var crew = &moviepb.CrewMovieMessage{}
		crew.Id = int32(rec.ID)
		crew.Name = rec.Name
		crew.ProfilePath=rec.Profile_path
		crew.Job = rec.Job
		CrewMovieMessages[i] = crew
	}
	return &moviepb.CrewMovieResponse{
		CrewmovieMessage: CrewMovieMessages,
	}, nil
}

// List video 
func (s *movieServer) VideoMovies(ctx context.Context, req *moviepb.VideoMovieRequest) (*moviepb.VideoMovieResponse, error) {
	movieID := req.MovieId
	fmt.Println(movieID)
	url := "https://api.themoviedb.org/3/movie/"
	url += movieID
	url += "/videos?api_key=b71e7ddc0337840f4f46be79b18e4c41&language=en-US"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result Video
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	VideoMessages := make([]*moviepb.VideosMessage, len(result.Results))

	for i, rec := range result.Results {
		// fmt.Println(rec)
		var video = &moviepb.VideosMessage{}
		video.Key = rec.Key
		video.Name = rec.Name
		
		VideoMessages[i] = video
	}
	return &moviepb.VideoMovieResponse{
		VideoMessage: VideoMessages,
	}, nil
}



func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	moviepb.RegisterMovieServer(grpcServer, &movieServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// https://api.themoviedb.org/3/movie/popular?api_key=b71e7ddc0337840f4f46be79b18e4c41&language=en-US&page=1
