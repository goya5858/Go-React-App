# SQLを用いない簡易版です　APIの接続確認用です 
  
* フロントとバックそれぞれにコンテナを建てて通信するようにしています
  
* CORSエラー対応のために、Nginxのサーバーをプロキシとして建てて、プロキシサーバー経由でGolangのAPIを叩くようにしています
    
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