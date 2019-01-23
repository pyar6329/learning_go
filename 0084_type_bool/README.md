# 【TRY】組み込み型(真偽値)

次のプログラムでtrueとなるケースを考えてください

- trueと表示される場合のa, b, cの値にはどのようなパターンがあるか
- 真理値表を作成し埋めて下さい

```go
package main
func main() {
    var a, b, c bool
    if a && b || !c {
        println("true")
    } else {
        println("false")
    }
}
```

answer: [main.go](./main.go)

## trueと表示される場合のa, b, cの値にはどのようなパターンがあるか

以下のときはtrueと表示される
- a = true, b = true, c = true
- a = true, b = true, c = false
- a = true, b = false, c = false
- a = false, b = true, c = false
- a = false, b = false, c = false

## 真理値表を作成し埋めて下さい

| a    | b    | c    | a && b | !c   | output |
| ---- | ---- | ---- | ------ | ---- | ------ |
| F    | F    | F    | F      | T    | T      |
| F    | F    | T    | F      | F    | F      |
| F    | T    | F    | F      | T    | T      |
| F    | T    | T    | F      | F    | F      |
| T    | F    | F    | F      | T    | T      |
| T    | F    | T    | F      | F    | F      |
| T    | T    | F    | T      | T    | T      |
| T    | T    | T    | T      | F    | T      |

