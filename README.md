# cyboze_pre_issue
[![Go](https://github.com/GotoRen/cyboze_pre_issue/actions/workflows/go.yml/badge.svg)](https://github.com/GotoRen/cyboze_pre_issue/actions/workflows/go.yml)
[![Language](https://img.shields.io/badge/Go-1.18.0-blue.svg)](https://github.com/GotoRen/cyboze_pre_issue)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## ⚡️ Concept of this project
Cyboze 事前課題：[https://cybozu.co.jp/company/job/recruitment/intern/infrastructure.html](https://cybozu.co.jp/company/job/recruitment/intern/infrastructure.html)

## ✏️ Pre-Issues
読み込んだファイルの各行を並列に処理して, 処理結果を元の行の並び通りに出力するプログラムをGoで作ってください.

処理の内容は行データのSHA256チェックサムのHEXダンプとします.

## 📝 Requirement

| Language/FrameWork | Version |
| :------------------ | ---------: |
| go                  |       1.18 |

## 🚀 Usage
```
### このリポジトリをクローン
$ git clone https://github.com/GotoRen/cyboze_pre_issue.git

### 実行
$ cd cyboze_pre_issue/
$ make

### デバッグモードの切り替え
.env >> DEBUG_MODE: true/false

### 読み込みファイルの切り替え
tests >> FILE: 01_input.txt ~ 05_input.txt
```

## Project managers

- [Issues](https://github.com/GotoRen/cyboze_pre_issue/issues)
- [Pull Request](https://github.com/GotoRen/cyboze_pre_issue/pulls)
- [Projects](https://github.com/GotoRen/cyboze_pre_issue/projects/1)
