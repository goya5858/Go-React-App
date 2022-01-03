# フロントエンドをReact,　バックエンドをGolangでウェブアプリを作成する際のひな形です  
  
* フロントとバックそれぞれにコンテナを建てて通信するようにしています
  
* CORSエラー対応のために、Nginxのサーバーをプロキシとして建てて、プロキシサーバー経由でGolangのAPIを叩くようにしています

* データベースとしてMySQLのコンテナを建てて、Golangのコンテナとリンクをしています  
    
**実行方法**
```
docker compose up --build
```
で、必要なすべてのコンテナとネットワークが展開されます

適当なブラウザを開いて、
```
http://localhost
```
へアクセスするとReactで構築されたでもページへ遷移できます  

  
また、
```
http://localhost:1234
```
へのアクセスで、phpmyadminを用いてデータベースサーバーの詳細を確認できます

### 細かい仕様
```
./frontend/front-app/src/component/SamplePage.jsx
./backend/controllers/webserver.go
```
がフロントとバックそれぞれのメインファイル  

また、
```
./frontend/front-app/src/component/xxx.jsx
```
と
```
./backend/controllers/yyy.go
```
がCRUD操作のそれぞれに一対一で対応している