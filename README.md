# cybozu-pre-issue

[![License](https://img.shields.io/badge/license-MIT-orange.svg)](https://opensource.org/licenses/MIT)
[![Go](https://img.shields.io/badge/Go-1.21.0-blue.svg)](https://tip.golang.org/doc/go1.21)

[![Go](https://github.com/GotoRen/cybozu-pre-issue/actions/workflows/go.yml/badge.svg)](https://github.com/GotoRen/cybozu-pre-issue/actions/workflows/go.yml)

## Concept of this project

Cyboze 事前課題内容

- https://cybozu.co.jp/company/job/recruitment/intern/infrastructure.html

## Pre-Issues

読み込んだファイルの各行を並列に処理して, 処理結果を元の行の並び通りに出力するプログラムを Go で作ってください.

処理の内容は行データの SHA256 チェックサムの HEX ダンプとします.

## Requirement

| Language/FrameWork | Version |
| :----------------- | ------: |
| Go                 |  1.20.0 |

## Usage

```shell
### このリポジトリをクローン
$ git clone https://github.com/GotoRen/cybozu-pre-issue.git

### 実行
$ cd cybozu-pre-issue/
$ make

### デバッグモードの切り替え
.env >> DEBUG_MODE: true/false

### 読み込みファイルの切り替え
tests >> FILE: 01_input.txt ~ 05_input.txt
```
