# cyboze_pre_issue
## ⚡️ Concept of this project
Cyboze 事前課題：[https://cybozu.co.jp/company/job/recruitment/intern/infrastructure.html](https://cybozu.co.jp/company/job/recruitment/intern/infrastructure.html)

## ✏️ Pre-Issues
読み込んだファイルの各行を並列に処理して, 処理結果を元の行の並び通りに出力するプログラムをGoで作ってください.

処理の内容は行データのSHA256チェックサムのHEXダンプとします.

提出はsecret gist のリンクを推奨する.

## 📝 Requirement

| Language/FrameWork | Version |
| :------------------ | ---------: |
| go                  |       1.18 |

## 🚀 Usage
```
### このリポジトリをクローン
$ git clone https://github.com/GotoRen/cyboze_pre_issue.git

### 実行
$ make

### デバッグモードの切り替え
.env >> DEBUG_MODE: true/false

### 読み込みファイルの切り替え
tests >> 01_input.txt ~ 05_input.txt
```
