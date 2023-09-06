# cybozu-pre-issue

[![License](https://img.shields.io/badge/license-MIT-orange.svg)](https://opensource.org/licenses/MIT)
[![Go](https://img.shields.io/badge/Go-1.21.0-blue.svg)](https://tip.golang.org/doc/go1.21)

[![go](https://github.com/GotoRen/cybozu-pre-issue/actions/workflows/go.yaml/badge.svg)](https://github.com/GotoRen/cybozu-pre-issue/actions/workflows/go.yaml)
[![reviewdog](https://github.com/GotoRen/cybozu-pre-issue/actions/workflows/reviewdog.yaml/badge.svg)](https://github.com/GotoRen/cybozu-pre-issue/actions/workflows/reviewdog.yaml)

## Concept of this project

cybozu 事前課題内容

- https://cybozu.co.jp/company/job/recruitment/intern/infrastructure.html

## Pre-Issues

読み込んだファイルの各行を並列に処理して, 処理結果を元の行の並び通りに出力するプログラムを Go で作ってください.

処理の内容は行データの SHA256 チェックサムの HEX ダンプとします.

## Requirement

| Language/FrameWork | Version |
| :----------------- | ------: |
| Go                 |  1.20.0 |
| direnv             |  2.32.3 |

## Usage

```shell
### このリポジトリをクローン
$ git clone https://github.com/GotoRen/cybozu-pre-issue.git

### 実行
$ cd cybozu-pre-issue/
$ make run

### .envrc: デバッグモードの切り替え
DEBUG_MODE: true/false

### .envrc: 読み込みファイルの切り替え
FILE_NAME: 01_input.txt ~ 05_input.txt
```

## Example

```
raw text: At a crossroads brimming with people,
00000000  80 b4 84 a3 47 48 62 a6  72 46 02 c7 88 a5 e4 70  |....GHb.rF.....p|
00000010  8d 83 7c 8f 83 8a 9a b4  f9 2a 39 7e f2 60 54 74  |..|......*9~.`Tt|

raw text: where will you go? (Being washed away)
00000000  51 b1 a9 08 e5 1f ad 7f  66 ba ba b0 19 88 2f 5b  |Q.......f...../[|
00000010  90 a2 9d c6 1d b9 e0 d0  3e 24 40 90 39 5a 21 8f  |........>$@.9Z!.|

raw text: Wearing the same clothes
00000000  d5 51 05 51 47 7a e3 c7  37 78 4e 62 21 b1 84 c2  |.Q.QGz..7xNb!...|
00000010  87 e8 44 8e b8 e4 a8 ae  b5 69 f6 5a be 88 1c 66  |..D......i.Z...f|

raw text: wearing the same expression
00000000  10 5b b6 e9 48 04 1a 28  3c fb 60 84 3a 0f b8 25  |.[..H..(<.`.:..%|
00000010  23 13 d3 60 2b 94 ac 26  c5 e0 bb c3 d8 44 07 2d  |#..`+..&.....D.-|

raw text:
00000000  e3 b0 c4 42 98 fc 1c 14  9a fb f4 c8 99 6f b9 24  |...B.........o.$|
00000010  27 ae 41 e4 64 9b 93 4c  a4 95 99 1b 78 52 b8 55  |'.A.d..L....xR.U|

raw text: Walking in a way that will get you inside the flock (without suspecting anything)
00000000  38 4a 36 86 e1 ca c8 8a  88 1c f7 5e 77 69 ff ca  |8J6........^wi..|
00000010  4b b8 9b a5 4e 37 98 95  bc 13 07 6f 18 42 f5 e8  |K...N7.....o.B..|

raw text: Why are you worrying about being different from someone else?
00000000  d5 8b 7e 28 cf 1a fc 6d  dd b9 e0 18 b9 9b e7 2a  |..~(...m.......*|
00000010  f5 1e 19 f2 8c e5 28 00  62 36 a7 34 15 bd 81 97  |......(.b6.4....|

raw text:
00000000  e3 b0 c4 42 98 fc 1c 14  9a fb f4 c8 99 6f b9 24  |...B.........o.$|
00000010  27 ae 41 e4 64 9b 93 4c  a4 95 99 1b 78 52 b8 55  |'.A.d..L....xR.U|

raw text: The people ahead turn to face you
00000000  ac 75 56 92 07 a0 4a 76  e7 9a 9b 8d 17 f7 21 c1  |.uV...Jv......!.|
00000010  12 ce 93 4f c4 78 bc c7  6d 2f 31 54 60 f0 94 ec  |...O.x..m/1T`...|

raw text: and tell you to keep in line
00000000  3f f9 bd d7 38 16 2b 9b  2f 72 51 00 67 12 c1 69  |?...8.+./rQ.g..i|
00000010  65 ba 89 80 7e 61 21 b8  d8 34 54 c2 c8 f6 9e 4e  |e...~a!..4T....N|

raw text: They preach these rules,
00000000  1d 3d 3c db 4f 29 28 f9  47 69 46 f3 92 4b 1e a7  |.=<.O)(.GiF..K..|
00000010  c5 45 e7 df a3 22 0a 17  e9 58 cd 0d 53 c1 78 c0  |.E..."...X..S.x.|

raw text: but their eyes are dead
00000000  a5 1a 6e a9 07 2b c2 ec  19 e9 71 80 e1 1d fd e8  |..n..+....q.....|
00000010  2c 87 f5 b8 19 9e 9e 0e  c4 47 83 53 d9 3b 5b eb  |,........G.S.;[.|

raw text:
00000000  e3 b0 c4 42 98 fc 1c 14  9a fb f4 c8 99 6f b9 24  |...B.........o.$|
00000010  27 ae 41 e4 64 9b 93 4c  a4 95 99 1b 78 52 b8 55  |'.A.d..L....xR.U|

raw text: You have the freedom to be the way you are
00000000  2e 82 53 bf ca 05 11 01  cc 11 f9 c8 3d 95 f4 05  |..S.........=...|
00000010  2d 06 a6 c1 62 20 b1 cd  30 fd 80 14 00 ee 3e 52  |-...b ..0.....>R|

raw text: Don't be controlled by adults
00000000  23 29 8f c6 9e 0d 90 a4  ea 4f 89 36 aa 5e d2 72  |#).......O.6.^.r|
00000010  f2 c1 60 2b 17 6a 57 9b  03 9b 52 36 d0 0d 62 67  |..`+.jW...R6..bg|

raw text: If people give up like that from the start
00000000  ee e7 7d 41 67 54 65 f4  05 79 ea 7f d6 19 f8 b7  |..}AgTe..y......|
00000010  9d 47 37 9d 32 d8 b8 80  05 39 15 7e 59 40 ea 2a  |.G7.2....9.~Y@.*|

raw text: Then why were we even born in the first place
00000000  34 5d f2 a9 ce f0 4a 2c  04 0e d7 18 5f f6 e8 16  |4]....J,...._...|
00000010  c5 eb 03 4f 71 ab c7 3c  fb 2b f8 91 b9 10 c0 33  |...Oq..<.+.....3|

raw text: Having dreams means at times you'll be faced with loneliness
00000000  61 6a bb 20 ef 2b 7b 85  2f 53 0a 41 dc 5d 7e e6  |aj. .+{./S.A.]~.|
00000010  9d 68 fd c6 f2 7d bc be  28 e7 8a 91 16 00 a1 7f  |.h...}..(.......|

raw text: You have to walk an empty path
00000000  e9 91 18 cc 96 26 2e 60  47 2c 35 59 60 cb 19 1a  |.....&.`G,5Y`...|
00000010  8b 63 60 5b 8d 86 b1 ac  e6 d5 85 e2 5b 5a ee 05  |.c`[........[Z..|

raw text: You won't get there even if the world is all the same
00000000  5f 7c 88 25 c5 d8 92 44  e9 44 6d 05 e6 1e d0 d7  |_|.%...D.Dm.....|
00000010  9a 22 e4 32 95 8b da 52  88 6b 9d e0 b3 8a 4b 55  |.".2...R.k....KU|

raw text: Are you really okay with Yes
00000000  28 85 fb 92 66 83 9e 62  8c 5d 5c 0e 4b 2a 39 19  |(...f..b.]\.K*9.|
00000010  dd 08 a7 fe 04 d5 93 35  5a 1f dd 3e 77 64 b4 a8  |.......5Z..>wd..|

raw text: Silent Majority
00000000  4f 17 fc 7e 3e 7a 8b 2b  48 c1 5c d5 1f 40 7e d2  |O..~>z.+H.\..@~.|
00000010  18 03 00 ea 9b 88 f3 df  58 07 ee db af ba 6d fc  |........X.....m.|

raw text:
00000000  e3 b0 c4 42 98 fc 1c 14  9a fb f4 c8 99 6f b9 24  |...B.........o.$|
00000010  27 ae 41 e4 64 9b 93 4c  a4 95 99 1b 78 52 b8 55  |'.A.d..L....xR.U|

raw text: The president of some country once said (he lied)
00000000  ff 98 74 ab 73 67 a7 4d  a4 41 a4 a4 e6 22 ab 96  |..t.sg.M.A..."..|
00000010  64 5d db 8e bb 2b ff fa  ff be f9 e4 98 94 90 d3  |d]...+..........|

raw text: Those who don't raise their voices are agreeing
00000000  7e 2b ee fc 2f f2 61 3e  4a d0 b4 15 7d e1 67 c6  |~+../.a>J...}.g.|
00000010  dd 22 7f a3 e5 73 27 0f  90 67 f0 c0 9a 2d 98 94  |."...s'..g...-..|

raw text:
00000000  e3 b0 c4 42 98 fc 1c 14  9a fb f4 c8 99 6f b9 24  |...B.........o.$|
00000010  27 ae 41 e4 64 9b 93 4c  a4 95 99 1b 78 52 b8 55  |'.A.d..L....xR.U|

raw text: Choosing is important
00000000  ab 1e 08 a9 f6 cf 13 0f  e2 8e fa a0 fa ef 56 f1  |..............V.|
00000010  31 62 85 61 86 2f 95 ec  89 45 df 32 26 b9 c5 e1  |1b.a./...E.2&...|

raw text: Don't leave it to others
00000000  01 4e 40 08 c7 ca 10 31  bc 1a 9c ec 62 c6 82 b5  |.N@....1....b...|
00000010  e9 79 e3 7b 0c 9e d5 ec  15 15 77 15 24 b4 3b 0a  |.y.{......w.$.;.|

raw text: If you don't take action,
00000000  97 84 2e 7b 7a 82 01 18  77 81 bc d1 4e b0 a9 2f  |...{z...w...N../|
00000010  55 a6 d0 51 b0 35 90 17  5a 01 84 e1 68 85 1a 19  |U..Q.5..Z...h...|

raw text: they won't hear your No
00000000  ea 92 30 75 cc 02 7e 92  76 6b 04 d4 f4 94 90 35  |..0u..~.vk.....5|
00000010  75 86 a6 ca db b4 ba 5e  b7 55 66 4d f9 cf 06 07  |u......^.UfM....|

raw text:
00000000  e3 b0 c4 42 98 fc 1c 14  9a fb f4 c8 99 6f b9 24  |...B.........o.$|
00000010  27 ae 41 e4 64 9b 93 4c  a4 95 99 1b 78 52 b8 55  |'.A.d..L....xR.U|

raw text: You can be yourself, all you need is to do what you want
00000000  a3 04 fa 86 41 79 4e e3  0c 53 40 48 f0 0c 72 35  |....AyN..S@H..r5|
00000010  e2 92 77 25 ee 83 dc 6a  ae 53 83 44 23 e8 7f 7a  |..w%...j.S.D#..z|

raw text: Don't be swayed by one of them
00000000  fc fd 28 4a fe 95 02 b6  bf a4 e9 8d 87 c1 20 e1  |..(J.......... .|
00000010  71 6e bc 57 ac 5d 15 f5  19 4e ae 49 75 00 1d f6  |qn.W.]...N.Iu...|

raw text: There are as many paths as there are people
00000000  59 53 31 a9 83 d8 e2 ce  5d 08 d2 13 fd 21 9f d8  |YS1.....]....!..|
00000010  1e 17 cc c1 89 6f 9a 4f  ea b6 1e 62 32 ad a8 8d  |.....o.O...b2...|

raw text: You just have to walk your own way
00000000  7f f4 86 43 c7 a0 a5 d7  19 35 f3 29 1e 67 85 34  |...C.....5.).g.4|
00000010  10 53 e7 78 e5 65 c5 43  38 ff 27 7d 31 31 01 40  |.S.x.e.C8.'}11.@|

raw text: Leave behind the boring adults who are kept in chains
00000000  26 c5 25 2b ab ef 03 5f  18 ba 57 9d 6a e2 59 ef  |&.%+..._..W.j.Y.|
00000010  4e e1 fb 37 2f 84 27 1b  25 98 de 4f 7d 8c 4c 5c  |N..7/.'.%..O}.L\|

raw text: By their appearances and pride
00000000  6c c0 ec 4d 4a e1 b1 32  4f 20 26 60 83 98 f7 da  |l..MJ..2O &`....|
00000010  e1 85 22 ad fd af 7e 09  e2 26 ae 67 e4 e3 63 15  |.."...~..&.g..c.|

raw text: The future is for you all
00000000  84 c8 d8 5e 40 e9 34 62  f0 0a ec 09 fb 6b 6c 2b  |...^@.4b.....kl+|
00000010  85 8c 53 08 f3 5c 6c a2  45 58 45 83 47 7d 39 53  |..S..\l.EXE.G}9S|

raw text: Say No to them
00000000  be bd 0a 01 c3 ef 92 c6  cd cc 96 7a 66 4d 18 23  |...........zfM.#|
00000010  f2 09 a3 08 ff fb 97 14  fb ab 65 e4 56 65 a1 94  |..........e.Ve..|

raw text: Silent Majority
00000000  4f 17 fc 7e 3e 7a 8b 2b  48 c1 5c d5 1f 40 7e d2  |O..~>z.+H.\..@~.|
00000010  18 03 00 ea 9b 88 f3 df  58 07 ee db af ba 6d fc  |........X.....m.|

raw text:
00000000  e3 b0 c4 42 98 fc 1c 14  9a fb f4 c8 99 6f b9 24  |...B.........o.$|
00000010  27 ae 41 e4 64 9b 93 4c  a4 95 99 1b 78 52 b8 55  |'.A.d..L....xR.U|

raw text: If you're just following someone
00000000  ae 80 9c 4a 5e a5 a7 f0  f3 bc 1a 5a 4b 4a 07 08  |...J^......ZKJ..|
00000010  d8 0f c3 88 73 07 cb a6  91 d2 12 06 16 64 b2 54  |....s........d.T|

raw text: You won't get hurt, but
00000000  ab 0c 70 6f e8 60 a7 66  bf c8 53 7c 06 27 49 1f  |..po.`.f..S|.'I.|
00000010  f8 27 5c 7a a2 23 e6 d4  e7 97 94 6b 32 4f 0b b9  |.'\z.#.....k2O..|

raw text: That crowd is just one mind
00000000  80 f6 9c b2 2b ff 3a f8  50 e4 2c f8 fb 92 f1 44  |....+.:.P.,....D|
00000010  56 83 85 4f 42 23 dc 93  d6 e3 3b 4f 6e 01 bb 2a  |V..OB#....;On..*|

raw text: You'll be made one of them
00000000  23 58 b6 62 ac 14 a8 bf  34 92 ce 0a 59 d8 6f 76  |#X.b....4...Y.ov|
00000010  89 e2 9f 94 3b db 05 e8  ba 7f b4 b3 3a 71 0f 89  |....;.......:q..|

raw text:
00000000  e3 b0 c4 42 98 fc 1c 14  9a fb f4 c8 99 6f b9 24  |...B.........o.$|
00000010  27 ae 41 e4 64 9b 93 4c  a4 95 99 1b 78 52 b8 55  |'.A.d..L....xR.U|

raw text: You have the freedom to be the way you are
00000000  2e 82 53 bf ca 05 11 01  cc 11 f9 c8 3d 95 f4 05  |..S.........=...|
00000010  2d 06 a6 c1 62 20 b1 cd  30 fd 80 14 00 ee 3e 52  |-...b ..0.....>R|

raw text: Don't be controlled by adults
00000000  23 29 8f c6 9e 0d 90 a4  ea 4f 89 36 aa 5e d2 72  |#).......O.6.^.r|
00000010  f2 c1 60 2b 17 6a 57 9b  03 9b 52 36 d0 0d 62 67  |..`+.jW...R6..bg|

raw text: If people give up like that from the start
00000000  ee e7 7d 41 67 54 65 f4  05 79 ea 7f d6 19 f8 b7  |..}AgTe..y......|
00000010  9d 47 37 9d 32 d8 b8 80  05 39 15 7e 59 40 ea 2a  |.G7.2....9.~Y@.*|

raw text: Then why were we even born in the first place
00000000  34 5d f2 a9 ce f0 4a 2c  04 0e d7 18 5f f6 e8 16  |4]....J,...._...|
00000010  c5 eb 03 4f 71 ab c7 3c  fb 2b f8 91 b9 10 c0 33  |...Oq..<.+.....3|

raw text: Having dreams means at times you'll be faced with loneliness
00000000  61 6a bb 20 ef 2b 7b 85  2f 53 0a 41 dc 5d 7e e6  |aj. .+{./S.A.]~.|
00000010  9d 68 fd c6 f2 7d bc be  28 e7 8a 91 16 00 a1 7f  |.h...}..(.......|

raw text: You have to walk an empty path
00000000  e9 91 18 cc 96 26 2e 60  47 2c 35 59 60 cb 19 1a  |.....&.`G,5Y`...|
00000010  8b 63 60 5b 8d 86 b1 ac  e6 d5 85 e2 5b 5a ee 05  |.c`[........[Z..|

raw text: You won't get there even if the world is all the same
00000000  5f 7c 88 25 c5 d8 92 44  e9 44 6d 05 e6 1e d0 d7  |_|.%...D.Dm.....|
00000010  9a 22 e4 32 95 8b da 52  88 6b 9d e0 b3 8a 4b 55  |.".2...R.k....KU|

raw text: Are you really okay with Yes
00000000  28 85 fb 92 66 83 9e 62  8c 5d 5c 0e 4b 2a 39 19  |(...f..b.]\.K*9.|
00000010  dd 08 a7 fe 04 d5 93 35  5a 1f dd 3e 77 64 b4 a8  |.......5Z..>wd..|

raw text: Silent Majority
00000000  4f 17 fc 7e 3e 7a 8b 2b  48 c1 5c d5 1f 40 7e d2  |O..~>z.+H.\..@~.|
00000010  18 03 00 ea 9b 88 f3 df  58 07 ee db af ba 6d fc  |........X.....m.|
```

## Check: SHA256 converter tool

- https://dencode.com/en/hash/sha256
