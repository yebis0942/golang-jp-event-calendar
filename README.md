# golang-jp-event-calendar

connpassのGoコミュニティのイベントを集めて、iCalendar形式で出力するアプリケーションです。

## 仕組み

* 収集対象のconnpassグループの直近のイベントを`https://{group_id}.connpass.com/ja.atom`から取得して、iCalendar形式で出力する
    * 収集対象のconnpassのグループは`config.go`で設定する

## 使い方

```
# 依存パッケージをインストールする
go mod tidy

# 2024-07-01以降のイベントをcalendar.icsに出力する
go run ./cmd/build_calendar -since 2024-07-01 > calendar.ics
```

## ライセンス

`LICENSE`ファイルを参照してください。
