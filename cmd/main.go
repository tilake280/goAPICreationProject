package main

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}
	api := application{
		config: cfg,
	}
}
