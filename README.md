# CryptoDetailsAPI-GoLang
 A project submission for Krypto


Libraries used:
---------------
* GORM - The fantastic ORM library for Golang
* Fiber - Express inspired web framework written in Go
* RabbitMQ - RabbitMQ is a message broker for cohesiveness between main and worker(to send emails).

Installation
---------------
Run the following go gets to install the necessary libraries.
```
go get gorm.io/gorm
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/jwt/v3
go get gorm.io/driver/mysql
go get github.com/streadway/amqp
go get github.com/golang-jwt/jwt/v4
```
Along with these libraries, Install [RabbitMQ](https://www.rabbitmq.com/download.html) as well as [MySQL](https://dev.mysql.com/downloads/installer/)

Once everything is setup, 
1.Create a MySQL database named "kryptodb"
2.Startup RabbitMQ
Everything is set and finished.


Run the project
---------------
To run the project, Simply cd into the folder:
```cd .\CurrencyAlertAPI\ ```
Run main.go:
```go run .\main.go```
And then, run the worker.go file inside the worker folder as such:
```cd .\worker\```
Run worker.go:
```go run .\worker.go```

Documentation
---------------
### The project has three subfolders namely backend,model and worker
* The backend consists of the email sender, object model relationship (along with driver to local MySQL database) and a route handler for ease of scalability.
* The model contains the currency details and user alerts models used to represent the structure of the information. Also kept seperately to aid in ease of readablitly and scalability.

### API Endpoint functions(only handlers):
* Login - Function to get JWT Token from given parameters
* AlertCreate - Function to create alert from user given parameters (email,target,currency aka symbol of the cryptocurrency from given API.example: btc,eth,etc)
* AlertDelete - Function to delete alert by given ID 
* FetchAlerts - Function to fetch all alerts from user
* FetchTriggeredAlerts - Function to fetch only triggered alerts from user


### Constants used:
* DNS is a constant.
```const DNS = "root:pass@tcp(127.0.0.1:3306)/kryptodb?charset=utf8mb4&parseTime=True&loc=Local"```
* JWT Token is a constant since nothing about user registration was specified.
```eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZW1haWwiOiJiLmJhbGF0YW1vZ2huYUBnbWFpbC5jb20iLCJleHAiOjE2MzE1NDM4ODZ9.Ahh7X-Z2y_gg80EGinmp7hbzsirHc7N_RNTkY-0ZGFw```
* Attain the JWT Token through login by using these parameters: 
```user=b.balatamoghna@gmail.com&pass=Krypto```
* Although it is best practice to not have credentials inside a git repo, the smtp email and password for a throwaway email has been included.(for obvious reasons, I won't be writing the creds here)


### Currency details Model
```
	ID           int     `gorm:"auto_increment" json:"id"`
	Symbol       string  `gorm:"primaryKey" json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
	Updated      int64   `gorm:"autoUpdateTime:milli"
 ```
 
 ### User alerts Model
 ```
 	ID        int     `gorm:"primaryKey;auto_increment" json:"id"`
	Email     string  `json:"email"`
	Currency  string  `json:"currency"`
	Target    float64 `json:"target"`
	Triggered string  `json:"triggered"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
 ```
 ## Screenshots
 go run main.go
 ![go run main.go](https://user-images.githubusercontent.com/480968/132880440-608fb0f3-b0fa-4edb-81be-8019d33b24fe.png)
 go run worker.go
 ![go run worker.go](https://user-images.githubusercontent.com/480968/132880515-2f7c35ea-d66b-4d27-8b9b-8bdd2f9c3fca.png)
 get JWT Token
 ![get JWT Token](https://user-images.githubusercontent.com/480968/132880604-6530340b-8af1-43ed-b36a-e7e6dae285fd.png)
 create alert
 ![create alert](https://user-images.githubusercontent.com/480968/132880709-82760a2f-7453-4b74-9ba6-4ecfeb4f74b5.png)
 fetch ALL alerts
 ![fetch ALL alerts](https://user-images.githubusercontent.com/480968/132880814-a3f9385b-3963-45d5-8646-5f9d750f1fce.png)
 fetch triggered alerts
 ![fetch triggered alerts](https://user-images.githubusercontent.com/480968/132880881-9e86c073-5a26-493c-8943-8bef367c264e.png)
 delete alert
 ![delete alert](https://user-images.githubusercontent.com/480968/132881585-b5ac9792-ca6f-4426-8c30-414014cfde1f.png)
 alert triggered
 ![alert triggered](https://user-images.githubusercontent.com/480968/132882911-5886e18c-8d7f-4067-8a82-1ac9819aeed1.png)
 sending mail from worker
 ![sending mail from worker](https://user-images.githubusercontent.com/480968/132882292-1c36956d-6837-4943-baea-cc977787b619.png)
 mail recieved
 ![mail recieved](https://user-images.githubusercontent.com/480968/132882855-2ca68f27-e236-4f03-ad07-5657f9b86187.png)
 fetchall shows alert has been triggered
 ![fetchall shows alert has been triggered](https://user-images.githubusercontent.com/480968/132883264-415a36f1-74a3-4986-9b18-b0132278b60b.png)
 RabbitMQ queue
 ![RabbitMQ queue](https://user-images.githubusercontent.com/480968/132883502-451251df-17fa-42cb-beb1-80952c77734f.png)

 ## Note
 The worker.go utilizes a function in the backend folder in email.go file. Although it is in the same folder as the backend for the main.go file, there are not shared functions between the main.go and worker.go where the only connection between them is the Rabbit MQ message broker thus ensuring independency.
 The backend can be a little cleaner, and worker can be in a seperate folder outside the main project in order to follow nomenclature and best practices.
 
 Email: b.balatamoghna@gmail.com
 VIT Email: b.balatamoghna2018@vitstudent.ac.in
 

# Problem statement
Incase the problem statement is required. [Here](https://drive.google.com/file/d/1ZXNHy1lH10ves3MnwOeoGxD8IhDn0rBT/view?usp=sharing) is the link.
