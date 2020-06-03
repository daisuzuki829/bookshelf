# 図書館システム

## 仕様
- 本は，タイトル(title)，著者(authors)，出版社(publisher)，出版年(publishYear)を持つ．
- 貸し出し履歴は，利用者(user)，本(book)，貸し出し日(lendDate)，返却日(returnDate)を持つ．


## CMS

### ＊ http://localhost:8080/index
- welcomeページ
- http://localhost:8080/books
と
http://localhost:8080/users
へのリンク

### ＊ http://localhost:8080/books
- 本の一覧、詳細、編集、削除、キーワード検索
- 本の追加
- 貸し出し履歴一覧

### ＊ http://localhost:8080/users
- ユーザー一覧、詳細、編集、削除
- ユーザー追加

## アプリ

### 本の予約
