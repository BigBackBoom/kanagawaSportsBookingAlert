# bookingSystemSakazuki

## 開発環境
- ubuntu 18.04.2 (Bionic Beaver)

※golangでの開発なのでmacなどでもできますが、
OS似合わせてSeleniumの環境を整える必要があります。小一時間ほど試してみたら何故かmacでは動かず。
なにかと融通がきくLinuxでの開発にシフトしました

## 依存環境

### Golang
- [Goの公式サイト](https://golang.org/)でgoの環境を整える
- GOMODULEを使えるようにする
```bash
export GO111MODULE=on 
```

### ChromeDriverのインストール
```bash
wget https://chromedriver.storage.googleapis.com/76.0.3809.25/chromedriver_linux64.zip
unzip chromedriver_linux64.zip -d ~/bin/
```

※chromdriverのバージョンとchromeのバージョンを合わせること

### ビルド方法
- 本プロジェクトのディレクトリに移動して、下記のコマンドを入力する。
```bash
go build
./AutoReservationSys
```
