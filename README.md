## Chat Application using websockets 

## How to run (Needs golang installed on local machine)
1. Clone the project into your local machine
2. Install the packages for go using "go mod tidy"
3. start the go application using "go run ."
4. Hit url "localhost:8080/" from your browser in multiple tabs to get the html file and establish websocket connection with go backend server
5. Send messages from different tabs and you will see them appear in real time across all your other tabs


## Docker setup
1. Clone the project
2. Build docker image using commad "docker build go-chatroom ."
3. Run the docker using "docker run <container-id>"
4. You are all set to use the application