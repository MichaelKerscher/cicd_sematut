FROM golang:1.20-alpine

# Set maintainer label: maintainer=[YOUR-EMAIL]
LABEL maintainer="michael.kerscher@outlook.com"

# Set working directory: `/src`
WORKDIR /src

# Copy go.mod first and download dependencies
COPY go.mod ./
RUN go mod download

# Copy main.go (and test file just in case)
COPY main.go .
COPY main_test.go .

# List items in the working directory (ls)
RUN ls -la

# Build the GO app as myapp binary and move it to /usr/
RUN go build -o myapp main.go && mv myapp /usr/

#Expose port 8888
EXPOSE 8888

# Run the service myapp when a container of this image is launched
CMD ["/usr/myapp"]
