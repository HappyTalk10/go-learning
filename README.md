# Go Learning 🚀

Go言語の学習用リポジトリへようこそ！
このプロジェクトはGitHub Codespacesを使用して、ブラウザ上で簡単にGo言語を学習できる環境を提供します。

## 📋 目次

- [はじめに](#はじめに)
- [環境構築](#環境構築)
- [プロジェクト構成](#プロジェクト構成)
- [学習内容](#学習内容)
- [使い方](#使い方)
- [リソース](#リソース)

## はじめに

このリポジトリは、Go言語の基礎から実践的な内容まで段階的に学習できるように構成されています。GitHub Codespacesを使用することで、ローカル環境の構築なしにすぐに学習を開始できます。

## 環境構築

### GitHub Codespacesで開く

1. このリポジトリのページで「Code」ボタンをクリック
2. 「Codespaces」タブを選択
3. 「Create codespace on main」をクリック

数分で開発環境が起動します！

### ローカル環境

```bash
# リポジトリをクローン
git clone https://github.com/yourusername/go-learning.git
cd go-learning

# Goのバージョン確認
go version
```

## プロジェクト構成

```
go-learning/
├── .devcontainer/          # Codespaces設定
│   └── devcontainer.json
├── basics/                 # 基礎編
│   ├── 01-hello/
│   ├── 02-variables/
│   └── 03-functions/
├── intermediate/           # 中級編
│   ├── 01-structs/
│   ├── 02-interfaces/
│   └── 03-goroutines/
├── advanced/              # 応用編
│   ├── 01-web-server/
│   ├── 02-database/
│   └── 03-testing/
├── exercises/             # 演習問題
└── README.md
```

## 学習内容

### 基礎編 (basics/)

- **Hello World**: 最初のGoプログラム
- **変数と型**: データ型、変数宣言、定数
- **制御構文**: if, for, switch
- **関数**: 関数定義、引数、戻り値
- **配列とスライス**: データ構造の基礎
- **マップ**: キー・バリュー型のデータ構造

### 中級編 (intermediate/)

- **構造体**: カスタム型の定義
- **メソッド**: レシーバー付き関数
- **インターフェース**: 抽象化と多態性
- **エラーハンドリング**: エラー処理のベストプラクティス
- **ゴルーチン**: 並行処理の基礎
- **チャネル**: ゴルーチン間の通信

### 応用編 (advanced/)

- **Webサーバー**: net/httpパッケージ
- **データベース接続**: database/sqlの使用
- **テスティング**: testing パッケージ
- **RESTful API**: 実践的なAPI開発
- **パッケージ管理**: Go Modules

## 使い方

### プログラムの実行

```bash
# ディレクトリに移動
cd basics/01-hello

# プログラムを実行
go run main.go

# ビルドして実行
go build
./01-hello
```

### テストの実行

```bash
# テストを実行
go test

# カバレッジ付きで実行
go test -cover

# 詳細表示
go test -v
```

### コードのフォーマット

```bash
# ファイルをフォーマット
go fmt ./...

# コードの静的解析
go vet ./...
```

## 学習の進め方

1. **basics/** から順番に進めてください
2. 各ディレクトリの `README.md` で概念を理解
3. `main.go` のコードを読んで実行
4. `exercises/` で演習問題に挑戦
5. 理解できたら次のトピックへ

### おすすめの学習順序

```
基礎編 → 中級編 → 応用編 → プロジェクト作成
```

## リソース

### 公式ドキュメント

- [Go公式サイト](https://go.dev/)
- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go標準パッケージ](https://pkg.go.dev/std)

### おすすめの学習教材

- [Go by Example](https://gobyexample.com/)
- [Go言語による並行処理](https://www.oreilly.co.jp/)
- [The Go Programming Language](https://www.gopl.io/)

### コミュニティ

- [Go Forum](https://forum.golangbridge.org/)
- [Gopher Slack](https://gophers.slack.com/)
- [Reddit r/golang](https://www.reddit.com/r/golang/)

## トラブルシューティング

### よくある問題

**Q: Codespacesが起動しない**
A: ブラウザのキャッシュをクリアして再試行してください

**Q: `go: command not found`**
A: Codespacesの再起動を試してください。`.devcontainer/devcontainer.json`の設定を確認してください

**Q: モジュールのエラーが出る**
A: `go mod tidy`を実行してください

## コントリビューション

改善提案やバグ報告は、Issueまたはプルリクエストでお願いします！

1. このリポジトリをフォーク
2. 新しいブランチを作成 (`git checkout -b feature/amazing-feature`)
3. 変更をコミット (`git commit -m 'Add some amazing feature'`)
4. ブランチにプッシュ (`git push origin feature/amazing-feature`)
5. プルリクエストを作成

## ライセンス

MIT License - 自由に学習・改変してください

## 著者

HappyTalk10 ([@HappyTalk10](https://github.com/HappyTalk10))

---

Happy Coding! 🎉 質問があれば、Issueでお気軽にどうぞ！