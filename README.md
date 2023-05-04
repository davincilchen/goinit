# DB

* Use MariaDB 10.7

## Create DB and Migrate

* 依序執行

* 1.createDB.exe
* 2.migrate.exe
* 3.loadSeed.exe

# App 註冊
* 圖片存放路徑: web/static 資料夾下
* DB註冊路徑: static/執行檔名稱 [欄位:image_url]

# 清理不要的資料夾後
* 砍掉readme.me此檔
* 執行go mod tidy
* vscode設定
* cmd不要的

# 專案分佈
* cmd下可建置其他執行檔
* pkg下app外層 是其他共用的,含內層通用model (不含)deliverymodel
* pkg下app內ctxcache,deliverymodel,errordef是內層通用
* pkg下app內server內含進入點
* pkg下app內含delivery, usecase , repo(db orm, http, grpc ...)