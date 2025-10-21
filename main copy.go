package main

import "fmt"

func main() {
    // 変数宣言
    var name string = "田中太郎"
    age := 25
    height := 175.5
    isStudent := true
    
    // 出力
    fmt.Println("=== プロフィール ===")
    fmt.Println("名前:", name)
    fmt.Println("年齢:", age)
    fmt.Println("身長:", height, "cm")
    fmt.Println("学生:", isStudent)
}
