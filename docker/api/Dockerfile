FROM golang:1.16

# 実行時のカレントディレクトリを設定
WORKDIR /wwgt-diary-server

# go.mod, go.sumをコピーして依存パッケージを取得
COPY go.mod go.sum ./
RUN go mod download

# カレントディレクトリのソースコードをコピー
COPY . .

# ビルドして実行ファイルをmainという名前で書き出す
RUN go build -o main .

# コンテナ実行時に8080ポートでListen
EXPOSE 8080

COPY ./dev-runscript.sh .
CMD ["sh", "dev-runscript.sh"]
