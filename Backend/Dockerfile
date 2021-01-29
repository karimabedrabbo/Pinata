# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# FROM golang:1.10.0
# RUN go get github.com/codegangsta/negroni \
#            github.com/gorilla/mux \
#            github.com/xyproto/simpleredis
# WORKDIR /app
# ADD ./main.go .
# RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# FROM scratch
# WORKDIR /app
# COPY --from=0 /app/main .
# COPY ./public/index.html public/index.html
# COPY ./public/script.js public/script.js
# COPY ./public/style.css public/style.css
# CMD ["/app/main"]
# EXPOSE 3000


# Start by building the application.
# FROM golang:1.13 as build-env

# WORKDIR /go/src/app
# ADD . /go/src/app

# RUN go get -d -v ./...

# RUN go build -o /go/bin/app

# FROM gcr.io/distroless/base
# COPY --from=build-env /go/bin/app /
# CMD ["/app"]
# EXPOSE 3000

FROM golang:1.13 as build-env
LABEL eyobuild=true
LABEL builder=true

# All these steps will be cached
WORKDIR /go/src/app
# COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary (might need -a later)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app

FROM gcr.io/distroless/base
LABEL builder=false
COPY --from=build-env /go/bin/app /
CMD ["/app"]
EXPOSE 3000


