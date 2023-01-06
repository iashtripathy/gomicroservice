
This project demonstrates a simple go microservice to fetch multiple stocks price.
The prices are fetched concurrently using go routines which decreases the overall time.
I tried to follow the Gokit tool kit while creating this microservice.

Command to run :
go run main.go

API Call :
http://localhost:8080/fetchstocks?stocks=msft,aapl

The above api call fetches the current stock prices of microsoft and apple.

