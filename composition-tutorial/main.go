package main

import (
	"fmt"
	"time"
)

// ========================================
// 例1: 基本的なコンポジション
// ========================================

// 共通の情報を持つ構造体
type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() string {
	return fmt.Sprintf("私は%sです。%d歳です。", p.Name, p.Age)
}

// Employee は Person を埋め込む(コンポジション)
type Employee struct {
	Person      // 埋め込み - Person のフィールドとメソッドが使える
	EmployeeID  string
	Department  string
}

func (e Employee) Work() string {
	return fmt.Sprintf("%sは%s部門で働いています。", e.Name, e.Department)
}

// Student も Person を埋め込む
type Student struct {
	Person
	StudentID string
	Grade     int
}

func (s Student) Study() string {
	return fmt.Sprintf("%sは%d年生で勉強中です。", s.Name, s.Grade)
}

// ========================================
// 例2: 複数の構造体を埋め込む
// ========================================

type Engine struct {
	Horsepower int
	FuelType   string
}

func (e Engine) Start() string {
	return fmt.Sprintf("%s エンジン始動! 馬力: %d", e.FuelType, e.Horsepower)
}

type GPS struct {
	CurrentLocation string
}

func (g GPS) Navigate(destination string) string {
	return fmt.Sprintf("%s から %s へナビゲーション中", g.CurrentLocation, destination)
}

type MusicPlayer struct {
	CurrentSong string
}

func (m MusicPlayer) Play() string {
	return fmt.Sprintf("♪ %s を再生中", m.CurrentSong)
}

// Car は複数の機能を持つ(多重コンポジション)
type Car struct {
	Engine
	GPS
	MusicPlayer
	Brand string
	Model string
}

func (c Car) Info() string {
	return fmt.Sprintf("%s %s", c.Brand, c.Model)
}

// ========================================
// 例3: インターフェースとコンポジション
// ========================================

// Speaker インターフェース - 「話せる」能力
type Speaker interface {
	Speak() string
}

// Walker インターフェース - 「歩ける」能力
type Walker interface {
	Walk() string
}

// 犬の構造体
type Dog struct {
	Name  string
	Breed string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s: ワンワン!", d.Name)
}

func (d Dog) Walk() string {
	return fmt.Sprintf("%sが散歩しています", d.Name)
}

// ロボットの構造体
type Robot struct {
	Model   string
	Battery int
}

func (r Robot) Speak() string {
	return fmt.Sprintf("%s: ピピピ！バッテリー残量%d%%", r.Model, r.Battery)
}

func (r Robot) Walk() string {
	return fmt.Sprintf("%sが移動しています", r.Model)
}

// インターフェースを使った関数 - 犬もロボットも同じように扱える
func MakeSpeak(s Speaker) {
	fmt.Println("  →", s.Speak())
}

func MakeWalk(w Walker) {
	fmt.Println("  →", w.Walk())
}

// ========================================
// 例4: 実践的な例 - ロギング機能の追加
// ========================================

type Logger struct {
	Prefix string
}

func (l Logger) Log(message string) {
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("[%s] %s: %s\n", timestamp, l.Prefix, message)
}

type Database struct {
	Logger // ロガーを埋め込む
	Name   string
}

func (db Database) Connect() {
	db.Log(fmt.Sprintf("%s データベースに接続しました", db.Name))
}

func (db Database) Query(sql string) {
	db.Log(fmt.Sprintf("クエリ実行: %s", sql))
}

type UserService struct {
	Logger   // ロガーを埋め込む
	Database // データベースを埋め込む
}

func (us UserService) CreateUser(username string) {
	us.Logger.Log(fmt.Sprintf("ユーザー作成開始: %s", username))
	us.Database.Connect()
	us.Database.Query(fmt.Sprintf("INSERT INTO users (name) VALUES ('%s')", username))
	us.Logger.Log(fmt.Sprintf("ユーザー作成完了: %s", username))
}

// ========================================
// 例5: フィールド名を明示的に指定する方法
// ========================================

type Address struct {
	Street string
	City   string
}

type Company struct {
	Name    string
	Address Address // 埋め込みではなく、フィールドとして持つ
}

func (c Company) FullAddress() string {
	// Address.Street のように明示的にアクセス
	return fmt.Sprintf("%s, %s", c.Address.Street, c.Address.City)
}

// ========================================
// メイン関数
// ========================================

func main() {
	fmt.Println("============================================")
	fmt.Println("例1: 基本的なコンポジション")
	fmt.Println("============================================")
	
	// Employee の作成
	emp := Employee{
		Person: Person{
			Name: "田中太郎",
			Age:  30,
		},
		EmployeeID: "E12345",
		Department: "開発",
	}
	
	// Person のメソッドも Employee のメソッドも使える
	fmt.Println(emp.Introduce())  // Person のメソッド
	fmt.Println(emp.Work())        // Employee のメソッド
	fmt.Printf("社員ID: %s\n", emp.EmployeeID)
	
	fmt.Println()
	
	// Student の作成
	student := Student{
		Person: Person{
			Name: "佐藤花子",
			Age:  20,
		},
		StudentID: "S98765",
		Grade:     2,
	}
	
	fmt.Println(student.Introduce())  // Person のメソッド
	fmt.Println(student.Study())      // Student のメソッド
	
	fmt.Println("\n============================================")
	fmt.Println("例2: 複数の構造体を埋め込む")
	fmt.Println("============================================")
	
	myCar := Car{
		Engine: Engine{
			Horsepower: 200,
			FuelType:   "ガソリン",
		},
		GPS: GPS{
			CurrentLocation: "東京",
		},
		MusicPlayer: MusicPlayer{
			CurrentSong: "夏の思い出",
		},
		Brand: "トヨタ",
		Model: "プリウス",
	}
	
	fmt.Println(myCar.Info())
	fmt.Println(myCar.Start())      // Engine のメソッド
	fmt.Println(myCar.Navigate("大阪"))  // GPS のメソッド
	fmt.Println(myCar.Play())       // MusicPlayer のメソッド
	
	fmt.Println("\n============================================")
	fmt.Println("例3: インターフェースとコンポジション")
	fmt.Println("============================================")
	
	dog := Dog{Name: "ポチ", Breed: "柴犬"}
	robot := Robot{Model: "RX-78", Battery: 85}
	
	fmt.Println("犬とロボットを同じインターフェースで扱う:")
	MakeSpeak(dog)    // 犬を Speaker として扱う
	MakeSpeak(robot)  // ロボットを Speaker として扱う
	
	fmt.Println("\n歩かせる:")
	MakeWalk(dog)     // 犬を Walker として扱う
	MakeWalk(robot)   // ロボットを Walker として扱う
	
	fmt.Println("\n============================================")
	fmt.Println("例4: 実践的な例 - ロギング機能")
	fmt.Println("============================================")
	
	userService := UserService{
		Logger: Logger{
			Prefix: "UserService",
		},
		Database: Database{
			Logger: Logger{
				Prefix: "Database",
			},
			Name: "UsersDB",
		},
	}
	
	userService.CreateUser("山田太郎")
	
	fmt.Println("\n============================================")
	fmt.Println("例5: フィールドとして持つ vs 埋め込み")
	fmt.Println("============================================")
	
	company := Company{
		Name: "ABC株式会社",
		Address: Address{
			Street: "渋谷1-1-1",
			City:   "東京都",
		},
	}
	
	fmt.Printf("会社名: %s\n", company.Name)
	fmt.Printf("住所: %s\n", company.FullAddress())
	// company.Street ではアクセスできない（埋め込みではないため）
	// company.Address.Street でアクセスする
	
	fmt.Println("\n============================================")
	fmt.Println("まとめ")
	fmt.Println("============================================")
	fmt.Println("✓ コンポジションは「持つ(has-a)」の関係")
	fmt.Println("✓ 埋め込むと、埋め込んだ型のメソッドが使える")
	fmt.Println("✓ 複数の構造体を埋め込める")
	fmt.Println("✓ インターフェースで共通の振る舞いを定義できる")
	fmt.Println("✓ 柔軟で保守しやすいコードが書ける")
}
