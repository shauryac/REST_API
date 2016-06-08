package main

import (
    "github.com/drone/routes"
    "log"
    "net/http"
    "encoding/json"
)
type Box struct {
    Email string `json:"email"`
    Zip  string `json:"zip"`
    Country string `json:"country"`
    Profession  string `json:"profession"`
    Favourite_color   string `json:"favourite_color"`
    Is_smoking string `json:"is_smoking"`   
    Favorite_sport string `json:"favorite_sport"`
    Food `json:"food"`
    Music `json:"music"`
    Movie`json:"movie"`
    Travel `json:"travel"`
}

type Food struct { 
        Type string `json:"type"` 
        Drink_alcohol string `json:"drink_alcohol"`}

type Music struct {
        Spotify_user_id string `json:"spotify_user_id"`}

type Movie struct {
        Tv_shows []string `json:"tv_shows"`
        Movies []string  `json:"movies"` } 
        
type Travel struct {
        Flight `json:"flight"` } 
        
type Flight struct {
        Seat string `json:"seat"` } 

func main() {
    mux := routes.New()
    mux.Get("/profile/:email", GetProfile)
    mux.Post("/profile", PostProfile)
    mux.Del("/profile/:email", DeleteProfile)
    mux.Put("/profile/:email",PutProfile)

    http.Handle("/", mux)
    log.Println("Listening...")
    http.ListenAndServe(":3000", nil)
    
}

var profileitems map[string][]string  //travel
var TVprofileitems map[string][]string
var Moviesprofileitems map[string][]string
var counter = 1


func GetProfile(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    email := params.Get(":email")
    
   var box Box
    
    /* test if entry is present in the map or not*/
   items, ok := profileitems[email]
   /* if ok is true, entry is present otherwise entry is absent*/
   if(ok){
      box.Email = email
      box.Zip = items[0]
      box.Country = items[1]
      box.Profession = items[2]
      box.Favourite_color = items[3]
      box.Is_smoking = items[4]
      box.Favorite_sport = items[5]
      box.Food.Type = items[6]
      box.Food.Drink_alcohol = items[7]
      box.Music.Spotify_user_id = items[8]
      box.Movie.Tv_shows  = []string{}
      box.Movie.Movies  =  []string{}
      box.Travel.Flight.Seat = items[9]
      var i int
      for i = 0; i < len(TVprofileitems[email]); i++ {
      box.Movie.Tv_shows = append(box.Movie.Tv_shows, TVprofileitems[email][i])
      }
      
      for i = 0; i < len(Moviesprofileitems[email]); i++ {
      box.Movie.Movies = append(box.Movie.Movies, Moviesprofileitems[email][i])
      }
         
   }else {
      http.Error(w, "Requested email is not present", 404)
      return
   }
    
    js, err := json.Marshal(box)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func  PostProfile(w http.ResponseWriter, r *http.Request) {

    var u Box
        if r.Body == nil {
            http.Error(w, "Please send a request body", 400)
            return
        }
        err := json.NewDecoder(r.Body).Decode(&u)
        if err != nil {
            http.Error(w, err.Error(), 400)
            return
        }
        if (counter == 1){
    
       /* create a map*/
       profileitems = make(map[string][]string)
       /* create a map*/
       TVprofileitems = make(map[string][]string)
       /* create a map*/
       Moviesprofileitems = make(map[string][]string) 
       counter++

       }

       TVShowitems := u.Movie.Tv_shows
       Moviesitems := u.Movie.Movies
       
       TVprofileitems[u.Email] = TVShowitems
       Moviesprofileitems[u.Email] = Moviesitems
   
       /* insert key-value pairs in the map*/
       profileitems[u.Email] = []string{u.Zip, u.Country, u.Profession, u.Favourite_color, u.Is_smoking, u.Favorite_sport, u.Food.Type, u.Food.Drink_alcohol, u.Music.Spotify_user_id, u.Travel.Flight.Seat }
       
 
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
}

func  DeleteProfile(w http.ResponseWriter, r *http.Request) {
    
    params := r.URL.Query()
    email := params.Get(":email")

    delete(profileitems,email)
    delete(TVprofileitems,email)
    delete(Moviesprofileitems, email)

    w.WriteHeader(204)
}
 func  PutProfile(w http.ResponseWriter, r *http.Request) {
    
    params := r.URL.Query()
    email := params.Get(":email")
    var box Box
    var changed Box
	var i int
    if r.Body == nil {
            http.Error(w, "Please send a request body", 400)
            return
        }
        err := json.NewDecoder(r.Body).Decode(&box)
    if err != nil {
            http.Error(w, err.Error(), 400)
            return
    }
    /* test if entry is present in the map or not*/
   items, ok := profileitems[email]
   /* if ok is true, entry is present otherwise entry is absent*/
   if(ok){
    
	if (len(box.Zip)!=0) {
      changed.Zip=box.Zip
    }
    if (len(box.Country)!=0) {
      changed.Country=box.Country
    }
	if (len(box.Profession)!=0) {
      changed.Profession=box.Profession
    }
	if (len(box.Favourite_color)!=0) {
      changed.Favourite_color=box.Favourite_color
    }
	if (len(box.Is_smoking)!=0) {
      changed.Is_smoking=box.Is_smoking
    }
	if (len(box.Favorite_sport)!=0) {
      changed.Favorite_sport=box.Favorite_sport
    }
	if (len(box.Food.Type)!=0) {
      changed.Food.Type=box.Food.Type
    }
	if (len(box.Food.Drink_alcohol)!=0) {
      changed.Food.Drink_alcohol=box.Food.Drink_alcohol
    }
	if (len(box.Music.Spotify_user_id)!=0) {
      changed.Music.Spotify_user_id=box.Music.Spotify_user_id
    }
	if (len(box.Movie.Tv_shows)!=0) {
      changed.Movie.Tv_shows=box.Movie.Tv_shows
    }
	if (len(box.Movie.Movies)!=0) {
      changed.Movie.Movies=box.Movie.Movies
    }
	if (len(box.Travel.Flight.Seat)!=0) {
      changed.Travel.Flight.Seat=box.Travel.Flight.Seat
    }
	
      changed.Email = email
	  if(len(changed.Zip)==0){
      changed.Zip = items[0]
	  }
	  if(len(changed.Country)==0){
      changed.Country = items[1]
	  }
	  if(len(changed.Profession)==0){
      changed.Profession = items[2]
	  }
	  if(len(changed.Favourite_color)==0){
      changed.Favourite_color = items[3]
	  }
	  if(len(changed.Is_smoking)==0){
      changed.Is_smoking = items[4]
	  }
	  if(len(changed.Favorite_sport)==0){
      changed.Favorite_sport = items[5]
	  }
	  if(len(changed.Food.Type)==0){
      changed.Food.Type = items[6]
	  }
	  if(len(changed.Food.Drink_alcohol)==0){
      changed.Food.Drink_alcohol = items[7]
	  }
	  if(len(changed.Music.Spotify_user_id)==0){
      changed.Music.Spotify_user_id = items[8]
	  }
	  if(len(changed.Movie.Tv_shows)==0){
      changed.Movie.Tv_shows = []string{}
	  for i = 0; i < len(TVprofileitems[email]); i++ {
      changed.Movie.Tv_shows = append(changed.Movie.Tv_shows, TVprofileitems[email][i])
      }
	  }
	  if(len(changed.Movie.Movies)==0){
      changed.Movie.Movies = []string{}
	  for i = 0; i < len(Moviesprofileitems[email]); i++ {
      changed.Movie.Movies = append(changed.Movie.Movies, Moviesprofileitems[email][i])
      }
	  }
	  if(len(changed.Travel.Flight.Seat)==0){
      changed.Travel.Flight.Seat= items[9]
	  }

       TVShowitems := changed.Movie.Tv_shows
       Moviesitems := changed.Movie.Movies
       
       TVprofileitems[changed.Email] = TVShowitems
       Moviesprofileitems[changed.Email] = Moviesitems
   
       /* insert key-value pairs in the map*/
       profileitems[changed.Email] = []string{changed.Zip, changed.Country, changed.Profession, changed.Favourite_color, changed.Is_smoking, changed.Favorite_sport, changed.Food.Type, changed.Food.Drink_alcohol, changed.Music.Spotify_user_id, changed.Travel.Flight.Seat }


    }

    w.WriteHeader(204)
}
