# pastebin
PasteBin implement, written in Go.


## 0 Quick Start

提供多语法解析为 `HTML` 的接口实现。


## 1 MarkDown

引用包 [blackfriday.v2](https://godoc.org/gopkg.in/russross/blackfriday.v2) ，其 `GitHub` 地址在 [这里](https://github.com/russross/blackfriday) 。

>   `bytes.Replace([]byte(input), []byte("\r"), nil, -1)` 消除 `\r` 换行带来的格式影响；
>
>   `blackfriday.HardLineBreak` 参数旨在消除两行连成一行的情况；

## 2 GoldMark

Mermaid:

```mermaid
flowchart LR
	nodeA((nodeA))
	nodeB((nodeB))
	nodeA-- dp1/pod1 -->nodeB
```

Code Block:

```c++
int main(int argc, char* argv[]) {
    int a = 666;
    vector<int> b({1, 2, 3});
    string c = "hello world";

    cout << "a=" << a << ", b=" << b << ", c=" << c
         << endl;
    debug(a, b, c);  // a=666, b=[ 1, 2, 3, ], c=hello world

    return 0;
}
```