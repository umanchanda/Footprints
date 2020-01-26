# Footprints

## How to Run

1. Make sure Golang is installed. If not, follow instructions for your machine [here](https://golang.org/doc/install)
2. Clone the repo to your clone machine
3. Navigate to the repo and run `go run main.go` inside your terminal
4. Head over to `localhost:8000` in your browswer and enjoy!

## Things I have done and want to-do

### Done

1. Set up PostgreSQL database on Heroku
2. Sucessfully connected to it via Golang
3. Created `Footprints` table via Golang
4. Inserted data from `building.csv` into `Footprints` table via Golang
5. Set up Mux Routing

### To-Do

1. Add more routes via Mux
2. Configure landing page to display a few records
3. Decide what data should be displayed
4. Use React to enhance the Front End
5. Reorganize code into packages instead of everything inside `main.go`
6. Dockerize the whole app