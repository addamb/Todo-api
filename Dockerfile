FROM golang:bookworm

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

WORKDIR /app/cmd/todo
# Build the Go app
RUN go build -o todo

# Expose port 8080 to the outside world
EXPOSE 8000

# Run the executable
CMD ["todo"]