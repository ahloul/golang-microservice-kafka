FROM golang:1.16.8

# Setup the Go app

RUN mkdir /app
ADD ./code /app
WORKDIR /app


# Get Dependency
RUN go mod download

# The ENTRYPOINT defines the command that will be ran when the
RUN chmod +x /app/start.sh

#for development only
# Install Air for hot reload
RUN go get -u github.com/cosmtrek/air
# In this case air command for hot reload go apps on file changes
ENTRYPOINT air

# Expose the port the app will run on
EXPOSE 8080

# Add VOLUMEs to allow backup of config, logs and databases
HEALTHCHECK CMD ["wget", "-q", "0.0.0.0:8080"]




