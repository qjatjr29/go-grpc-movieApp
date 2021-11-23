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

// "github.com/qjatjr29/go-grpc-movieapp/data"

const portNumber = "9001"

type movieServer struct {
	moviepb.MovieServer
}

// 인기 영화 목록

type PopularMovies struct {
	Results []struct {
		ID           int    `json:"id"`
		Title        string `json:"title"`
		Release_date string `json:"release_date"`
		Poster_path  string `json:"poster_path"`
		Overview     string `json:"overview"`
	} `json:"results`
}

// GetUser returns user message by user_id
func (s *movieServer) GetMovie(ctx context.Context, req *moviepb.GetMovieRequest) (*moviepb.GetMovieResponse, error) {
	// userID := req.UserId

	var movieMessage *moviepb.MovieMessage
	movieMessage.MovieId = 1
	movieMessage.PosterPath = "1"
	movieMessage.ReleaseDate = "1"
	movieMessage.Title = "1"

	// for _, u := range data.UsersV2 {
	// 	if u.UserId != userID {
	// 		continue
	// 	}
	// 	movieMessage = u
	// 	break
	// }

	return &moviepb.GetMovieResponse{
		// MovieMessage: movieMessage,
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

// ListUsers returns all user messages
func (s *movieServer) ListPopularMovies(ctx context.Context, req *moviepb.ListPopularMovieRequest) (*moviepb.ListPopularMovieResponse, error) {

	resp, err := http.Get("https://api.themoviedb.org/3/movie/popular?api_key=b71e7ddc0337840f4f46be79b18e4c41&language=en-US&page=1")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result PopularMovies
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	popularMovieMessage := make([]*moviepb.PopularMovieMessage, 0)

	// var popularMovie *moviepb.PopularMovieMessage
	var popularMovie = &moviepb.PopularMovieMessage{
		Id:          1,
		Title:       "",
		PosterPath:  "",
		ReleaseDate: "",
		Overview:    "",
	}
	for _, rec := range result.Results {
		popularMovie.Id = int32(rec.ID)
		popularMovie.Title = rec.Title
		popularMovie.PosterPath = rec.Poster_path
		popularMovie.ReleaseDate = rec.Release_date
		popularMovie.Overview = rec.Overview
		// fmt.Println(rec.ID)
		// _popularmoive := *moviepb.PopularMovieMessage{
		// 	MovieId:     int32(rec.ID),
		// 	Title:       rec.Title,
		// 	PosterPath:  rec.Poster_path,
		// 	Overview:    rec.Overview,
		// 	ReleaseDate: rec.Release_date,
		// }
		fmt.Println(rec.Title)
		popularMovieMessage = append(popularMovieMessage, popularMovie)
	}

	return &moviepb.ListPopularMovieResponse{
		PopularmovieMessage: popularMovieMessage,
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
