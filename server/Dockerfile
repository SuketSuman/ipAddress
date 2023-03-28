FROM golang:1.16.3-alpine3.13

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .
ENV CGO_ENABLED=0
# Build the binary executable
RUN go mod tidy
RUN go build -o main .

# Expose the port that the application will listen on
EXPOSE 8080

# Run the binary executable
CMD ["/app/main"]