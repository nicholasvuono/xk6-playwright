FROM golang:buster
RUN apt-get update
RUN apt install git
RUN go env -w GOPROXY=https://proxy.golang.org
RUN git clone https://github.com/grafana/k6.git
RUN go env -w GO111MODULE=off
RUN go get github.com/playwright-community/playwright-go
RUN go install github.com/playwright-community/playwright-go/cmd/playwright
RUN go env -w GO111MODULE=auto
RUN apt-get install -y \
    fonts-liberation \
    gconf-service \
    libappindicator1 \
    libasound2 \
    libatk1.0-0 \
    libcairo2 \
    libcups2 \
    libfontconfig1 \
    libgbm-dev \
    libgdk-pixbuf2.0-0 \
    libgtk-3-0 \
    libicu-dev \
    libjpeg-dev \
    libnspr4 \
    libnss3 \
    libpango-1.0-0 \
    libpangocairo-1.0-0 \
    libpng-dev \
    libx11-6 \
    libx11-xcb1 \
    libxcb1 \
    libxcomposite1 \
    libxcursor1 \
    libxdamage1 \
    libxext6 \
    libxfixes3 \
    libxi6 \
    libxrandr2 \
    libxrender1 \
    libxss1 \
    libxtst6 \
    xdg-utils
RUN playwright install
WORKDIR /go/k6
RUN CGO_ENABLED=0 go install -a -trimpath -ldflags "-s -w -X ./lib/consts.VersionDetails=$(date -u +"%FT%T%z")/$(git describe --always --long --dirty)"
RUN go install go.k6.io/xk6/cmd/xk6@latest
RUN xk6 build v0.39.0 --with github.com/wosp-io/xk6-playwright@latest
RUN cp k6 $GOPATH/bin/k6
WORKDIR /home/k6
ENTRYPOINT ["k6"]