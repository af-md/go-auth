run:
	@echo "Running the application"
	@export MYSQL_USER=user 
	@export MYSQL_PASSWORD=password

	go run main.go 