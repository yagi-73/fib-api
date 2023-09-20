# Fibonacci API
このプロジェクトは、Fibonacci数列を計算するためのシンプルなRESTful APIです。

## エンドポイントURL
https://fibonacci-api-jqpp.onrender.com/fib?n=<計算したい数>

## ソースコードの構成
このアプリケーションのソースコードは主に以下の部分で構成されています。

| ファイル名 | 役割 |
| ---- | ---- |
| main.go | アプリケーションのエントリーポイントで、Echoフレームワークを使用してAPIサーバーを起動します。 |
| controllers/fibonacci_controller.go | フィボナッチ数列を計算するコントローラーで、APIのエンドポイントでリクエストを処理します。 |
| router/router.go | Echoルーターの設定とコントローラーのマッピングを行います。 |
| main_test.go | フィボナッチAPIのテストが記載されています。 |
| fib_benchmark_test.go | フィボナッチ関数の実行速度を比較するためのファイルです。 |
| Dockerfile | Dockerコンテナ内でアプリケーションを実行するためのファイルです。 |
| docker-compose.yml | Dockerコンテナのビルドおよび実行を簡素化するための設定ファイルです。 |

<br>

## ローカル環境で動作させる場合
1. 任意のディレクトリでこのプロジェクトをクローンします。
```
git clone https://github.com/yagi-73/fibonacci-api.git
```
2. アプリケーションをビルドおよび実行します。
```
cd fibonacci-api
docker-compose up -d --build
```
3. ブラウザやAPIクライアントを使用して、以下のエンドポイントにアクセスしてください。n番目のフィボナッチ数を返します。
http://localhost:8080/fib?n=<計算したい数>

<br>

## テストの実行
ユニットテスト
```
docker compose exec api go test
```
ベンチマークテスト
```
docker compose exec api go test -bench . -benchmem
```