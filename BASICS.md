# Go言語 基礎編ガイド 📚

このガイドでは、Go言語の基礎を段階的に学習します。各セクションには具体的なコード例と演習問題が含まれています。

## 目次

1. [Hello World](#1-hello-world)
2. [変数と型](#2-変数と型)
3. [制御構文](#3-制御構文)
4. [関数](#4-関数)
5. [配列とスライス](#5-配列とスライス)
6. [マップ](#6-マップ)

---

## 1. Hello World

### ファイル構成

```
basics/01-hello/
├── main.go
└── README.md
```

### `main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Welcome to Go Programming!")
}
```

### 実行方法

```bash
cd basics/01-hello
go run main.go
```

### 学習ポイント

- `package main`: 実行可能なプログラムはmainパッケージに属する
- `import "fmt"`: 標準ライブラリのインポート
- `func main()`: プログラムのエントリーポイント
- `fmt.Println()`: 標準出力に文字列を表示

### 演習問題

1. 自分の名前を表示するプログラムを作成
2. 複数行にわたるメッセージを表示
3. `fmt.Printf()` を使って「Hello, %s!」の形式で表示

---

## 2. 変数と型

### ファイル構成

```
basics/02-variables/
├── main.go
├── types.go
└── README.md
```

### `main.go`

```go
package main

import "fmt"

func main() {
    // 変数宣言の方法
    
    // 方法1: var キーワード
    var name string = "太郎"
    var age int = 25
    
    // 方法2: 型推論
    var city = "東京"
    
    // 方法3: 短縮宣言（関数内のみ）
    country := "日本"
    
    // 複数変数の同時宣言
    var x, y, z int = 1, 2, 3
    
    // 定数の宣言
    const Pi = 3.14159
    const AppName = "GoLearning"
    
    // 出力
    fmt.Println("名前:", name)
    fmt.Println("年齢:", age)
    fmt.Println("都市:", city)
    fmt.Println("国:", country)
    fmt.Printf("座標: (%d, %d, %d)\n", x, y, z)
    fmt.Println("円周率:", Pi)
}
```

### `types.go`

```go
package main

import "fmt"

func demonstrateTypes() {
    // 基本型
    
    // 整数型
    var intVar int = 42
    var int8Var int8 = 127
    var int16Var int16 = 32767
    var int32Var int32 = 2147483647
    var int64Var int64 = 9223372036854775807
    
    // 符号なし整数型
    var uintVar uint = 42
    var uint8Var uint8 = 255
    
    // 浮動小数点型
    var float32Var float32 = 3.14
    var float64Var float64 = 3.141592653589793
    
    // 真偽値型
    var boolVar bool = true
    
    // 文字列型
    var stringVar string = "こんにちは"
    
    // 文字型（rune = int32のエイリアス）
    var runeVar rune = 'あ'
    
    // バイト型（byte = uint8のエイリアス）
    var byteVar byte = 65 // 'A'
    
    // 型の表示
    fmt.Printf("int: %d, 型: %T\n", intVar, intVar)
    fmt.Printf("float64: %f, 型: %T\n", float64Var, float64Var)
    fmt.Printf("bool: %t, 型: %T\n", boolVar, boolVar)
    fmt.Printf("string: %s, 型: %T\n", stringVar, stringVar)
    fmt.Printf("rune: %c, 型: %T\n", runeVar, runeVar)
    fmt.Printf("byte: %c, 型: %T\n", byteVar, byteVar)
    
    // ゼロ値（初期化されない変数のデフォルト値）
    var zeroInt int       // 0
    var zeroFloat float64 // 0.0
    var zeroBool bool     // false
    var zeroString string // ""
    
    fmt.Println("\nゼロ値:")
    fmt.Printf("int: %d, float64: %f, bool: %t, string: '%s'\n", 
        zeroInt, zeroFloat, zeroBool, zeroString)
}
```

### 実行方法

```bash
cd basics/02-variables
go run main.go
go run types.go
```

### 演習問題

1. 自分の情報（名前、年齢、住所）を変数に格納して表示
2. 異なる型の変数を使った計算プログラムを作成
3. 定数を使って消費税率を定義し、価格計算を実装

---

## 3. 制御構文

### ファイル構成

```
basics/03-control-flow/
├── if.go
├── for.go
├── switch.go
└── README.md
```

### `if.go`

```go
package main

import "fmt"

func main() {
    // 基本的なif文
    age := 20
    if age >= 20 {
        fmt.Println("成人です")
    }
    
    // if-else文
    score := 75
    if score >= 80 {
        fmt.Println("優秀です")
    } else {
        fmt.Println("もう少し頑張りましょう")
    }
    
    // if-else if-else文
    temperature := 25
    if temperature >= 30 {
        fmt.Println("暑いです")
    } else if temperature >= 20 {
        fmt.Println("快適です")
    } else if temperature >= 10 {
        fmt.Println("涼しいです")
    } else {
        fmt.Println("寒いです")
    }
    
    // 初期化付きif文（スコープ内でのみ有効）
    if num := 10; num%2 == 0 {
        fmt.Println(num, "は偶数です")
    }
    // fmt.Println(num) // エラー: numはスコープ外
}
```

### `for.go`

```go
package main

import "fmt"

func main() {
    // 基本的なforループ
    fmt.Println("1. 基本的なfor:")
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    
    // 条件のみのfor（whileのような使い方）
    fmt.Println("\n2. 条件のみのfor:")
    count := 0
    for count < 3 {
        fmt.Println("カウント:", count)
        count++
    }
    
    // 無限ループ（breakで抜ける）
    fmt.Println("\n3. 無限ループとbreak:")
    n := 0
    for {
        if n >= 3 {
            break
        }
        fmt.Println("n =", n)
        n++
    }
    
    // continueの使用
    fmt.Println("\n4. continueの使用:")
    for i := 0; i < 5; i++ {
        if i == 2 {
            continue // i=2の時はスキップ
        }
        fmt.Println(i)
    }
    
    // 範囲ループ（range）
    fmt.Println("\n5. rangeを使ったループ:")
    numbers := []int{10, 20, 30, 40, 50}
    for index, value := range numbers {
        fmt.Printf("インデックス: %d, 値: %d\n", index, value)
    }
    
    // インデックスのみ
    fmt.Println("\n6. インデックスのみ:")
    for index := range numbers {
        fmt.Println("インデックス:", index)
    }
    
    // 値のみ（_でインデックスを無視）
    fmt.Println("\n7. 値のみ:")
    for _, value := range numbers {
        fmt.Println("値:", value)
    }
}
```

### `switch.go`

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 基本的なswitch
    day := 3
    switch day {
    case 1:
        fmt.Println("月曜日")
    case 2:
        fmt.Println("火曜日")
    case 3:
        fmt.Println("水曜日")
    case 4:
        fmt.Println("木曜日")
    case 5:
        fmt.Println("金曜日")
    case 6, 7:
        fmt.Println("週末")
    default:
        fmt.Println("不正な日")
    }
    
    // 条件式付きswitch
    score := 85
    switch {
    case score >= 90:
        fmt.Println("A評価")
    case score >= 80:
        fmt.Println("B評価")
    case score >= 70:
        fmt.Println("C評価")
    case score >= 60:
        fmt.Println("D評価")
    default:
        fmt.Println("F評価")
    }
    
    // 型スイッチ
    var x interface{} = "hello"
    switch v := x.(type) {
    case int:
        fmt.Println("整数:", v)
    case string:
        fmt.Println("文字列:", v)
    case bool:
        fmt.Println("真偽値:", v)
    default:
        fmt.Println("不明な型")
    }
    
    // 時刻によるswitch
    hour := time.Now().Hour()
    switch {
    case hour < 12:
        fmt.Println("おはようございます")
    case hour < 18:
        fmt.Println("こんにちは")
    default:
        fmt.Println("こんばんは")
    }
}
```

### 演習問題

1. 1から100までの数字のうち、3の倍数と5の倍数を表示するプログラム
2. じゃんけんゲーム（ユーザー入力と判定）を実装
3. 九九の表を表示するプログラム（ネストしたforループ）

---

## 4. 関数

### ファイル構成

```
basics/04-functions/
├── basic.go
├── advanced.go
└── README.md
```

### `basic.go`

```go
package main

import "fmt"

// 引数なし、戻り値なし
func sayHello() {
    fmt.Println("こんにちは！")
}

// 引数あり、戻り値なし
func greet(name string) {
    fmt.Println("こんにちは、", name, "さん！")
}

// 引数あり、戻り値あり
func add(a int, b int) int {
    return a + b
}

// 同じ型の引数は省略可能
func multiply(a, b int) int {
    return a * b
}

// 複数の戻り値
func divmod(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

// 名前付き戻り値
func rectangle(width, height int) (area int, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return // 名前付き戻り値は return だけでOK
}

// 可変長引数
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    // 関数の呼び出し
    sayHello()
    greet("太郎")
    
    // 戻り値の受け取り
    result := add(10, 20)
    fmt.Println("10 + 20 =", result)
    
    product := multiply(5, 6)
    fmt.Println("5 × 6 =", product)
    
    // 複数の戻り値
    q, r := divmod(17, 5)
    fmt.Printf("17 ÷ 5 = %d 余り %d\n", q, r)
    
    // 名前付き戻り値
    area, perimeter := rectangle(5, 10)
    fmt.Printf("面積: %d, 周囲: %d\n", area, perimeter)
    
    // 戻り値の一部を無視（_を使用）
    _, remainder := divmod(20, 3)
    fmt.Println("余りのみ:", remainder)
    
    // 可変長引数
    fmt.Println("合計:", sum(1, 2, 3, 4, 5))
    fmt.Println("合計:", sum(10, 20))
}
```

### `advanced.go`

```go
package main

import "fmt"

// 高階関数: 関数を引数に取る
func apply(f func(int, int) int, a, b int) int {
    return f(a, b)
}

// 無名関数とクロージャ
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// 再帰関数
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

// フィボナッチ数列（再帰）
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

// defer文: 関数の終了時に実行
func deferExample() {
    defer fmt.Println("3. 最後に実行")
    fmt.Println("1. 最初に実行")
    fmt.Println("2. 次に実行")
}

func main() {
    // 高階関数の使用
    addFunc := func(a, b int) int { return a + b }
    multiplyFunc := func(a, b int) int { return a * b }
    
    fmt.Println("Apply add:", apply(addFunc, 10, 5))
    fmt.Println("Apply multiply:", apply(multiplyFunc, 10, 5))
    
    // クロージャの使用
    fmt.Println("\nクロージャ:")
    increment := counter()
    fmt.Println(increment()) // 1
    fmt.Println(increment()) // 2
    fmt.Println(increment()) // 3
    
    // 新しいカウンター
    newIncrement := counter()
    fmt.Println(newIncrement()) // 1
    
    // 再帰関数
    fmt.Println("\n階乗:")
    fmt.Println("5! =", factorial(5)) // 120
    
    fmt.Println("\nフィボナッチ数列:")
    for i := 0; i < 10; i++ {
        fmt.Printf("fib(%d) = %d\n", i, fibonacci(i))
    }
    
    // defer文
    fmt.Println("\ndefer文の実行順序:")
    deferExample()
}
```

### 演習問題

1. 温度変換関数（摂氏→華氏、華氏→摂氏）を作成
2. 素数判定関数を実装
3. 文字列を反転させる関数を作成
4. 最大公約数を求める関数（ユークリッドの互除法）を実装

---

## 5. 配列とスライス

### ファイル構成

```
basics/05-arrays-slices/
├── arrays.go
├── slices.go
└── README.md
```

### `arrays.go`

```go
package main

import "fmt"

func main() {
    // 配列の宣言と初期化
    var arr1 [5]int // ゼロ値で初期化: [0 0 0 0 0]
    fmt.Println("arr1:", arr1)
    
    // 初期値を指定
    arr2 := [5]int{1, 2, 3, 4, 5}
    fmt.Println("arr2:", arr2)
    
    // 一部のみ初期化
    arr3 := [5]int{1, 2, 3} // [1 2 3 0 0]
    fmt.Println("arr3:", arr3)
    
    // サイズを自動推論
    arr4 := [...]int{10, 20, 30, 40}
    fmt.Println("arr4:", arr4)
    
    // インデックス指定で初期化
    arr5 := [5]int{0: 100, 4: 500} // [100 0 0 0 500]
    fmt.Println("arr5:", arr5)
    
    // 要素へのアクセス
    fmt.Println("\n要素へのアクセス:")
    fmt.Println("arr2[0] =", arr2[0])
    fmt.Println("arr2[4] =", arr2[4])
    
    // 要素の変更
    arr2[2] = 99
    fmt.Println("arr2[2]を変更後:", arr2)
    
    // 配列の長さ
    fmt.Println("arr2の長さ:", len(arr2))
    
    // 配列のループ
    fmt.Println("\n配列のループ:")
    for i := 0; i < len(arr2); i++ {
        fmt.Printf("arr2[%d] = %d\n", i, arr2[i])
    }
    
    // rangeを使ったループ
    fmt.Println("\nrangeを使ったループ:")
    for index, value := range arr2 {
        fmt.Printf("インデックス %d: 値 %d\n", index, value)
    }
    
    // 多次元配列
    var matrix [3][3]int = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println("\n行列:")
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("%d ", matrix[i][j])
        }
        fmt.Println()
    }
}
```

### `slices.go`

```go
package main

import "fmt"

func main() {
    // スライスの作成方法
    
    // 1. リテラルから作成
    slice1 := []int{1, 2, 3, 4, 5}
    fmt.Println("slice1:", slice1)
    
    // 2. makeで作成（長さとキャパシティを指定）
    slice2 := make([]int, 5)      // 長さ5、キャパシティ5
    slice3 := make([]int, 3, 10)  // 長さ3、キャパシティ10
    fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))
    fmt.Println("slice3:", slice3, "len:", len(slice3), "cap:", cap(slice3))
    
    // 3. 配列からスライス
    arr := [5]int{10, 20, 30, 40, 50}
    slice4 := arr[1:4] // インデックス1から3まで
    fmt.Println("slice4:", slice4)
    
    // スライスの操作
    fmt.Println("\n--- スライスの操作 ---")
    
    // 要素の追加（append）
    numbers := []int{1, 2, 3}
    fmt.Println("初期:", numbers)
    
    numbers = append(numbers, 4)
    fmt.Println("4を追加:", numbers)
    
    numbers = append(numbers, 5, 6, 7)
    fmt.Println("5,6,7を追加:", numbers)
    
    // スライスの結合
    moreNumbers := []int{8, 9, 10}
    numbers = append(numbers, moreNumbers...)
    fmt.Println("別のスライスを結合:", numbers)
    
    // スライスのコピー
    fmt.Println("\n--- スライスのコピー ---")
    source := []int{1, 2, 3, 4, 5}
    destination := make([]int, len(source))
    copy(destination, source)
    fmt.Println("コピー元:", source)
    fmt.Println("コピー先:", destination)
    
    // スライスの部分取得
    fmt.Println("\n--- スライシング ---")
    data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println("元のスライス:", data)
    fmt.Println("data[2:5]:", data[2:5])   // [2 3 4]
    fmt.Println("data[:5]:", data[:5])     // [0 1 2 3 4]
    fmt.Println("data[5:]:", data[5:])     // [5 6 7 8 9]
    fmt.Println("data[:]:", data[:])       // 全要素
    
    // スライスの容量と長さ
    fmt.Println("\n--- 容量と長さ ---")
    s := make([]int, 0, 5)
    fmt.Printf("長さ: %d, 容量: %d, 値: %v\n", len(s), cap(s), s)
    
    for i := 0; i < 10; i++ {
        s = append(s, i)
        fmt.Printf("長さ: %d, 容量: %d, 値: %v\n", len(s), cap(s), s)
    }
    
    // 多次元スライス
    fmt.Println("\n--- 多次元スライス ---")
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println("行列:", matrix)
    fmt.Println("matrix[1][2] =", matrix[1][2])
    
    // 実用例: フィルタリング
    fmt.Println("\n--- フィルタリング ---")
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    var evens []int
    for _, num := range nums {
        if num%2 == 0 {
            evens = append(evens, num)
        }
    }
    fmt.Println("偶数のみ:", evens)
}
```

### 演習問題

1. 配列の要素の合計と平均を計算する関数を作成
2. スライスから重複を除去する関数を実装
3. スライスを逆順にする関数を作成
4. 2つのスライスをマージして昇順にソートする関数を実装

---

## 6. マップ

### ファイル構成

```
basics/06-maps/
├── main.go
└── README.md
```

### `main.go`

```go
package main

import "fmt"

func main() {
    // マップの作成方法
    
    // 1. makeを使った作成
    ages := make(map[string]int)
    ages["太郎"] = 25
    ages["花子"] = 23
    ages["次郎"] = 30
    fmt.Println("ages:", ages)
    
    // 2. リテラルでの初期化
    scores := map[string]int{
        "数学": 85,
        "英語": 92,
        "国語": 78,
    }
    fmt.Println("scores:", scores)
    
    // マップの操作
    fmt.Println("\n--- 要素へのアクセス ---")
    fmt.Println("太郎の年齢:", ages["太郎"])
    fmt.Println("数学の点数:", scores["数学"])
    
    // 存在しないキー
    fmt.Println("存在しないキー:", ages["三郎"]) // 0（ゼロ値）
    
    // キーの存在確認
    fmt.Println("\n--- キーの存在確認 ---")
    age, exists := ages["太郎"]
    if exists {
        fmt.Println("太郎は存在します。年齢:", age)
    } else {
        fmt.Println("太郎は存在しません")
    }
    
    age, exists = ages["三郎"]
    if exists {
        fmt.Println("三郎は存在します。年齢:", age)
    } else {
        fmt.Println("三郎は存在しません")
    }
    
    // 要素の追加と更新
    fmt.Println("\n--- 要素の追加と更新 ---")
    ages["三郎"] = 28 // 追加
    fmt.Println("三郎を追加:", ages)
    
    ages["太郎"] = 26 // 更新
    fmt.Println("太郎の年齢を更新:", ages)
    
    // 要素の削除
    fmt.Println("\n--- 要素の削除 ---")
    delete(ages, "次郎")
    fmt.Println("次郎を削除:", ages)
    
    // マップの長さ
    fmt.Println("\n--- マップの長さ ---")
    fmt.Println("agesの要素数:", len(ages))
    
    // マップのループ
    fmt.Println("\n--- マップのループ ---")
    for name, age := range ages {
        fmt.Printf("%sさんは%d歳です\n", name, age)
    }
    
    // キーのみのループ
    fmt.Println("\n--- キーのみ ---")
    for name := range ages {
        fmt.Println("名前:", name)
    }
    
    // 値のみのループ（実用性は低い）
    fmt.Println("\n--- 値のみ ---")
    for _, age := range ages {
        fmt.Println("年齢:", age)
    }
    
    // ネストしたマップ
    fmt.Println("\n--- ネストしたマップ ---")
    students := map[string]map[string]int{
        "太郎": {
            "数学": 85,
            "英語": 90,
            "国語": 75,
        },
        "花子": {
            "数学": 95,
            "英語": 88,
            "国語": 92,
        },
    }
    
    fmt.Println("全学生の成績:", students)
    fmt.Println("太郎の数学:", students["太郎"]["数学"])
    
    // マップのコピー（手動でコピーが必要）
    fmt.Println("\n--- マップのコピー ---")
    original := map[string]int{"a": 1, "b": 2, "c": 3}
    copied := make(map[string]int)
    
    for key, value := range original {
        copied[key] = value
    }
    
    fmt.Println("元のマップ:", original)
    fmt.Println("コピー:", copied)
    
    // 実用例: 単語の出現回数
    fmt.Println("\n--- 単語の出現回数 ---")
    text := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
    wordCount := make(map[string]int)
    
    for _, word := range text {
        wordCount[word]++
    }
    
    fmt.Println("単語の出現回数:")
    for word, count := range wordCount {
        fmt.Printf("%s: %d回\n", word, count)
    }
    
    // 実用例: グループ化
    fmt.Println("\n--- 年齢でグループ化 ---")
    people := []struct {
        name string
        age  int
    }{
        {"太郎", 25},
        {"花子", 25},
        {"次郎", 30},
        {"三郎", 30},
        {"四郎", 25},
    }
    
    ageGroups := make(map[int][]string)
    for _, person := range people {
        ageGroups[person.age] = append(ageGroups[person.age], person.name)
    }
    
    fmt.Println("年齢別グループ:")
    for age, names := range ageGroups {
        fmt.Printf("%d歳: %v\n", age, names)
    }
}
```

### 演習問題

1. 文字列内の各文字の出現回数を数えるプログラムを作成
2. 学生の成績管理システム（追加、削除、検索、平均点計算）を実装
3. 電話帳アプリ（名前から電話番号を検索）を作成
4. 2つのマップをマージする関数を実装

---

## 学習の進め方

### ステップバイステップガイド

1. **各ディレクトリを順番に作成**
   ```bash
   mkdir -p basics/01-hello
   mkdir -p basics/02-variables
   mkdir -p basics/03-control-flow
   mkdir -p basics/04-functions
   mkdir -p basics/05-arrays-slices
   mkdir -p basics/06-maps
   ```

2. **各ファイルを作成して実行**
   - 各セクションのコードをコピー
   - ファイルを作成して保存
   - `go run` コマンドで実行
   - 動作を確認して理解を深める

3. **コードを改変して実験**
   - 値を変えて実行
   - 新しい機能を追加
   - エラーを意図的に起こして学習

4. **演習問題に挑戦**
   - 各セクションの演習問題を解く
   - 解答例は自分で考える
   - 行き詰まったら前のコードを参考に

### 実行コマンド一覧

```bash
# Hello World
cd basics/01-hello && go run main.go

# 変数と型
cd basics/02-variables && go run main.go
cd basics/02-variables && go run types.go

# 制御構文
cd basics/03-control-flow && go run if.go
cd basics/03-control-flow && go run for.go
cd basics/03-control-flow && go run switch.go

# 関数
cd basics/04-functions && go run basic.go
cd basics/04-functions && go run advanced.go

# 配列とスライス
cd basics/05-arrays-slices && go run arrays.go
cd basics/05-arrays-slices && go run slices.go

# マップ
cd basics/06-maps && go run main.go
```

---

## 総合演習問題

基礎編の全てを復習するための総合問題です。

### プロジェクト1: シンプルなタスク管理システム

```
basics/exercises/task-manager/
├── main.go
└── README.md
```

**要件:**
- タスクの追加、削除、一覧表示
- タスクの完了状態の管理
- タスクのフィルタリング（完了/未完了）

**使用する概念:**
- スライス（タスクのリスト）
- マップ（タスクIDとタスク情報の紐付け）
- 構造体（タスクの情報）
- 関数（各機能の実装）
- 制御構文（メニュー表示とループ）

### プロジェクト2: 成績計算プログラム

```
basics/exercises/grade-calculator/
├── main.go
└── README.md
```

**要件:**
- 複数の学生の成績を入力
- 各学生の平均点を計算
- クラス全体の平均点を計算
- 最高点と最低点を表示
- 成績順にソート

**使用する概念:**
- マップ（学生名と成績の紐付け）
- スライス（成績のリスト）
- 関数（計算ロジック）
- ループ（データの処理）

### プロジェクト3: 簡易電卓

```
basics/exercises/calculator/
├── main.go
└── README.md
```

**要件:**
- 四則演算（+、-、×、÷）
- 複数の演算の連続実行
- 計算履歴の表示
- エラー処理（ゼロ除算など）

**使用する概念:**
- 関数（各演算の実装）
- スイッチ文（演算子の判定）
- スライス（履歴の保存）
- エラーハンドリング

---

## よくあるエラーと解決方法

### 1. インデックスの範囲外エラー

```go
// ❌ 誤り
arr := [3]int{1, 2, 3}
fmt.Println(arr[3]) // panic: index out of range

// ✅ 正しい
if len(arr) > 3 {
    fmt.Println(arr[3])
} else {
    fmt.Println("インデックスが範囲外です")
}
```

### 2. マップのゼロ値エラー

```go
// ❌ 誤り
var m map[string]int
m["key"] = 1 // panic: assignment to entry in nil map

// ✅ 正しい
m := make(map[string]int)
m["key"] = 1
```

### 3. スライスの容量不足

```go
// ⚠️ 注意が必要
s := make([]int, 3)
s[3] = 4 // panic: index out of range

// ✅ 正しい
s := make([]int, 3)
s = append(s, 4)
```

### 4. 変数のスコープ

```go
// ❌ 誤り
if x := 10; x > 5 {
    fmt.Println(x)
}
fmt.Println(x) // エラー: x is undefined

// ✅ 正しい
x := 10
if x > 5 {
    fmt.Println(x)
}
fmt.Println(x)
```

---

## 次のステップ

基礎編を修了したら、以下のトピックに進みましょう：

1. **中級編へ進む**
   - 構造体とメソッド
   - インターフェース
   - エラーハンドリング
   - ゴルーチンとチャネル

2. **標準パッケージを学ぶ**
   - `fmt`: フォーマット入出力
   - `strings`: 文字列操作
   - `strconv`: 文字列変換
   - `os`: オペレーティングシステム機能
   - `io`: 入出力操作

3. **実践的なプロジェクトを作る**
   - CLIツールの開発
   - ファイル操作プログラム
   - 簡単なWebサーバー

---

## 参考リソース

### 公式ドキュメント
- [Go Tour](https://go.dev/tour/) - インタラクティブな学習
- [Effective Go](https://go.dev/doc/effective_go) - ベストプラクティス
- [Go by Example](https://gobyexample.com/) - サンプルコード集

### おすすめの練習サイト
- [Exercism - Go Track](https://exercism.org/tracks/go)
- [LeetCode](https://leetcode.com/) - アルゴリズム問題
- [HackerRank - Go](https://www.hackerrank.com/domains/go)

### コミュニティ
- [Go Forum](https://forum.golangbridge.org/)
- [r/golang](https://www.reddit.com/r/golang/)
- [Gophers Slack](https://gophers.slack.com/)

---

## チェックリスト

基礎編を修了する前に、以下の項目を確認しましょう：

- [ ] Hello Worldプログラムを実行できる
- [ ] 変数の宣言と初期化の違いを理解している
- [ ] if、for、switchの基本的な使い方を理解している
- [ ] 関数の定義と呼び出しができる
- [ ] 配列とスライスの違いを説明できる
- [ ] マップの基本操作（追加、取得、削除）ができる
- [ ] rangeを使ったループを書ける
- [ ] 基本的なエラーを自分で解決できる
- [ ] 演習問題を自力で解くことができる
- [ ] 簡単なプログラムを一から書ける

---

## まとめ

お疲れ様でした！この基礎編では以下のことを学びました：

1. **Hello World**: Goプログラムの基本構造
2. **変数と型**: データの扱い方と型システム
3. **制御構文**: プログラムのフロー制御
4. **関数**: コードの再利用と構造化
5. **配列とスライス**: データのコレクション管理
6. **マップ**: キー・バリュー型データの扱い方

これらの知識は、Go言語プログラミングの基盤となります。しっかりと理解してから次のステップに進みましょう！

質問や不明点があれば、[Issues](https://github.com/yourusername/go-learning/issues)で気軽に聞いてください。

Happy Coding! 🎉