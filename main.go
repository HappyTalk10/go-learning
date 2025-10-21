package main

import (
    "fmt"
    "strconv"
)

func main() {
    num := 123
    
    fmt.Printf("10進数: %d\n", num)
    
    // strconv.FormatIntを使用した変換
    binary := strconv.FormatInt(int64(num), 2)
    octal := strconv.FormatInt(int64(num), 8)
    hex := strconv.FormatInt(int64(num), 16)
    
    fmt.Printf("2進数:  %s\n", binary)
    fmt.Printf("8進数:  %s\n", octal)
    fmt.Printf("16進数: %s\n", hex)
    
    // fmt.Sprintfを使用した方法
    fmt.Printf("\n別の方法:\n")
    fmt.Printf("2進数:  %b\n", num)
    fmt.Printf("8進数:  %o\n", num)
    fmt.Printf("16進数: %x\n", num)
    fmt.Printf("16進数: %X\n", num) // 大文字
}

