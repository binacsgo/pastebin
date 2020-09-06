# pastebin
PasteBin implement, written in Go.


## 0 Quick Start

提供多语法解析为 `HTML` 的接口实现。


## 1 MarkDown

引用包 [blackfriday.v2](https://godoc.org/gopkg.in/russross/blackfriday.v2) ，其 `GitHub` 地址在 [这里](https://github.com/russross/blackfriday) 。

>   `bytes.Replace([]byte(input), []byte("\r"), nil, -1)` 消除 `\r` 换行带来的格式影响；
>
>   `blackfriday.HardLineBreak` 参数旨在消除两行连成一行的情况；

