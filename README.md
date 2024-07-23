# 事前準備
テキストエディタや IDE は特に指定しません。ご自身が使い慣れているものを使用しください。

おすすめは [VSCode](https://azure.microsoft.com/ja-jp/products/visual-studio-code) か [Goland](https://www.jetbrains.com/ja-jp/go/) です。


### Goland 向け
Goland を使用している人はこのリポジトリをダウンロードして開くだけで準備が整っていると思うので事前準備をする必要はありません。


### VSCode 向け
Dev Container という仕組みを使うのであらかじめ以下をダウンロードしてください
- [Docker](https://docs.docker.com/get-docker/)
- [Dev Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

上記のツールをダウンロードできたら、自身の PC にこのリポジトリを clone や fork などをしてください。

Dev Container の設定はあらかじめ用意しているので起動させてください。

無事 Dev Container が立ち上がったら以下のコマンドを実行して Go が入っている確認しください。

```shell
$ go version
# ↓な感じの出力があれば OK
go version go1.22.1 linux/arm64
```