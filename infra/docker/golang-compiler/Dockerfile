FROM golang:1.8.1


# WORKDIR /go/src/app
# COPY . .

# Install dependancy managers
RUN go get github.com/Masterminds/glide \
 && go get github.com/kardianos/govendor

# Install test helpers and tools
RUN go get github.com/onsi/ginkgo/ginkgo \
 && go get github.com/onsi/gomega
