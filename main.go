package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type City struct {
	Name       string  `json:"name"`
	Country    string  `json:"country"`
	Population int     `json:"population"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Language   string  `json:"language"`
}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header missing", http.StatusForbidden)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetCities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityName := vars["city"]
	cities := []City{
		{"Tunis", "Tunisia", 1200000, 36.8065, 10.1815, "Arabic"},
		{"New York", "USA", 8500000, 40.7128, -74.0060, "English"},
		{"Paris", "France", 2100000, 48.8566, 2.3522, "French"},
		{"Tokyo", "Japan", 13929286, 35.6895, 139.6917, "Japanese"},
		{"London", "UK", 8982000, 51.5074, -0.1278, "English"},
		{"Berlin", "Germany", 3748148, 52.5200, 13.4050, "German"},
		{"Madrid", "Spain", 3223334, 40.4168, -3.7038, "Spanish"},
		{"Rome", "Italy", 2873000, 41.9028, 12.4964, "Italian"},
		{"Beijing", "China", 21540000, 39.9042, 116.4074, "Mandarin"},
		{"Moscow", "Russia", 11920000, 55.7558, 37.6173, "Russian"},
		{"São Paulo", "Brazil", 12325232, -23.5505, -46.6333, "Portuguese"},
		{"Mumbai", "India", 12442373, 19.0760, 72.8777, "Hindi"},
		{"Mexico City", "Mexico", 9209944, 19.4326, -99.1332, "Spanish"},
		{"Cairo", "Egypt", 9200000, 30.0444, 31.2357, "Arabic"},
		{"Istanbul", "Turkey", 15462452, 41.0082, 28.9784, "Turkish"},
		{"Jakarta", "Indonesia", 10187595, -6.2088, 106.8456, "Indonesian"},
		{"Lagos", "Nigeria", 14500000, 6.5244, 3.3792, "English"},
		{"Buenos Aires", "Argentina", 2890151, -34.6037, -58.3816, "Spanish"},
		{"Rio de Janeiro", "Brazil", 6748000, -22.9068, -43.1729, "Portuguese"},
		{"Bangkok", "Thailand", 10539000, 13.7563, 100.5018, "Thai"},
		{"Hong Kong", "China", 7451000, 22.3193, 114.1694, "Chinese"},
		{"Seoul", "South Korea", 9776000, 37.5665, 126.978, "Korean"},
		{"Singapore", "Singapore", 5617000, 1.3521, 103.8198, "English"},
		{"Kuala Lumpur", "Malaysia", 1840000, 3.139, 101.6869, "Malay"},
		{"Lisbon", "Portugal", 505526, 38.7223, -9.1399, "Portuguese"},
		{"Vienna", "Austria", 1897480, 48.2082, 16.3738, "German"},
		{"Zurich", "Switzerland", 440000, 47.3769, 8.5417, "German"},
		{"Stockholm", "Sweden", 975904, 59.3293, 18.0686, "Swedish"},
		{"Oslo", "Norway", 159000, 59.9139, 10.7522, "Norwegian"},
		{"Helsinki", "Finland", 655000, 60.1695, 24.9354, "Finnish"},
		{"Brussels", "Belgium", 185000, 50.8503, 4.3517, "Dutch"},
		{"Amsterdam", "Netherlands", 872680, 52.3676, 4.9041, "Dutch"},
		{"Warsaw", "Poland", 1790658, 52.2297, 21.0122, "Polish"},
		{"Prague", "Czech Republic", 1300000, 50.0755, 14.4378, "Czech"},
		{"Budapest", "Hungary", 1750000, 47.4979, 19.0402, "Hungarian"},
		{"Athens", "Greece", 664046, 37.9838, 23.7275, "Greek"},
		{"Dublin", "Ireland", 1173179, 53.3331, -6.2489, "English"},
		{"Edinburgh", "UK", 540000, 55.9533, -3.1883, "English"},
		{"Glasgow", "UK", 630000, 55.8642, -4.2518, "English"},
		{"Bucharest", "Romania", 1836000, 44.4268, 26.1025, "Romanian"},
		{"Sofia", "Bulgaria", 1241675, 42.6977, 23.3219, "Bulgarian"},
		{"Santiago", "Chile", 5816282, -33.4489, -70.6693, "Spanish"},
		{"Montevideo", "Uruguay", 1336000, -34.9037, -56.1882, "Spanish"},
		{"Lima", "Peru", 9430000, -12.0464, -77.0428, "Spanish"},
		{"Bogotá", "Colombia", 7735000, 4.6110, -74.0827, "Spanish"},
		{"Quito", "Ecuador", 276000, -0.2317, -78.5243, "Spanish"},
		{"Caracas", "Venezuela", 2920000, 10.4918, -66.8983, "Spanish"},
		{"Managua", "Nicaragua", 1100000, 12.6344, -86.2087, "Spanish"},
		{"San Salvador", "El Salvador", 570000, 13.6929, -89.2182, "Spanish"},
		{"Guatemala City", "Guatemala", 3000000, 14.6349, -90.5069, "Spanish"},
		{"Port-au-Prince", "Haiti", 987000, 18.5944, -72.3074, "Haitian Creole"},
	}

	if cityName == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cities)
		return
	}

	for _, city := range cities {
		if city.Name == cityName {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(city)
			return
		}
	}

	http.Error(w, "City not found", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/login", Login).Methods("POST")

	// Protected routes
	api := router.PathPrefix("/api").Subrouter()
	api.Use(JwtVerify)
	api.HandleFunc("/cities", GetCities).Methods("GET")
	api.HandleFunc("/cities/{city:[a-zA-Z ]+}", GetCities).Methods("GET") // Added route for city-specific queries

	log.Fatal(http.ListenAndServe(":8000", router))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	// Parse the JSON body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate credentials
	if req.Username != "admin" || req.Password != "password" {
		http.Error(w, "Invalid credentials", http.StatusForbidden)
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Respond with token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
