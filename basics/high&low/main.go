package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card はトランプのカードを表す構造体
type Card struct {
	Suit  string // スート（ハート、ダイヤなど）
	Value int    // 値（1-13: A, 2-10, J, Q, K）
}

// String はカードを文字列で表現する
func (c Card) String() string {
	valueName := ""
	switch c.Value {
	case 1:
		valueName = "A"
	case 11:
		valueName = "J"
	case 12:
		valueName = "Q"
	case 13:
		valueName = "K"
	default:
		valueName = fmt.Sprintf("%d", c.Value)
	}
	return fmt.Sprintf("%s%s", c.Suit, valueName)
}

// createDeck はトランプのデッキを作成する
func createDeck() []Card {
	suits := []string{"♠", "♥", "♦", "♣"}
	deck := []Card{}
	
	for _, suit := range suits {
		for value := 1; value <= 13; value++ {
			deck = append(deck, Card{Suit: suit, Value: value})
		}
	}
	
	return deck
}

// shuffle はデッキをシャッフルする
func shuffle(deck []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	shuffled := make([]Card, len(deck))
	copy(shuffled, deck)
	
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	
	return shuffled
}

func main() {
	fmt.Println("=== ハイ&ロー ゲーム ===")
	fmt.Println("次のカードが現在のカードより高いか低いかを当てましょう！")
	fmt.Println()
	
	// デッキを作成してシャッフル
	deck := createDeck()
	deck = shuffle(deck)
	
	score := 0
	
	// ゲームループ（10回勝負）
	for round := 1; round <= 10; round++ {
		if len(deck) < 2 {
			fmt.Println("カードが足りなくなりました！")
			break
		}
		
		// 現在のカードを引く
		currentCard := deck[0]
		deck = deck[1:]
		
		fmt.Printf("ラウンド %d/10\n", round)
		fmt.Printf("現在のカード: %s\n", currentCard)
		fmt.Print("次のカードは？ (h: 高い, l: 低い, s: 同じ): ")
		
		var input string
		fmt.Scan(&input)
		
		// 次のカードを引く
		nextCard := deck[0]
		deck = deck[1:]
		
		fmt.Printf("次のカード: %s\n", nextCard)
		
		// 判定
		correct := false
		if input == "h" && nextCard.Value > currentCard.Value {
			correct = true
		} else if input == "l" && nextCard.Value < currentCard.Value {
			correct = true
		} else if input == "s" && nextCard.Value == currentCard.Value {
			correct = true
		}
		
		if correct {
			fmt.Println("正解！ +1ポイント")
			score++
		} else {
			fmt.Println("不正解...")
		}
		
		fmt.Printf("現在のスコア: %d\n", score)
		fmt.Println("---")
	}
	
	// 最終結果
	fmt.Println()
	fmt.Println("=== ゲーム終了 ===")
	fmt.Printf("最終スコア: %d/10\n", score)
	
	if score >= 8 {
		fmt.Println("素晴らしい！")
	} else if score >= 5 {
		fmt.Println("まずまずです！")
	} else {
		fmt.Println("もう一度チャレンジ！")
	}
}