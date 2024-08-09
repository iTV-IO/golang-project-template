# iTV Services Template

APIs Template for iTV services

## Requirements

- Docker
- Docker Compose
- systemd (for service management)

You do not need to have Go installed locally.

### Environment Configuration

Create a `.env` file in the project root with the following content:
```
PORT=8000
POSTGRES_DB=db_name
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
```

Adjust the values according to your specific setup.

### Docker Setup

1. Build and start the Docker containers:
```
docker compose up --build -d
```
2. To stop and remove the containers:
```
docker compose down
```

### Using Makefile

To rebuild and restart the Docker containers:
```
make docker-rebuild
```

## Systemd Service Setup

1. Copy the provided systemd service example file to the system directory:
```
sudo cp systemctl.example.service /etc/systemd/system/todo.service
```
2. Edit the service file if necessary:
```
sudo nano /etc/systemd/system/todo.service
```
3. Reload the systemd daemon:
```
sudo systemctl daemon-reload
```
4. Start the service:
```
sudo systemctl start todo
```
5. Stop the service to start on boot:
```
sudo systemctl stop todo
```
6. Check the status of the service:
```
sudo systemctl status todo
```

