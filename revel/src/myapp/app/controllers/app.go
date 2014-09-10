package controllers

import (
	// console "fmt"
	"github.com/revel/revel"
	_l "log"
	"unicode/utf8"
)

//http://qiita.com/tenntenn/items/0e33a4959250d1a55045
//Go言語の初心者が見ると幸せになれる場所

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// 関数内では以下のように型つけしなくてもいい
	//var greeting string = "dig" と同じ
	greeting := "dig"
	return c.Render(greeting)
}

//getやpostで入ってくるものは引数？
func (c App) Hello(myName string) revel.Result {
	// http://qiita.com/suin/items/d952fb963956ac31b243
	_l.Printf("%+v", myName)

	c.Validation.Required(myName).Message("Your name is required!")
	//日本語だとエラーならない
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")
	//のでこっち
	// http://golang.org/pkg/unicode/utf8/#example_RuneCountInString
	_l.Println("bytes =", len(myName))
	_l.Println("runes =", utf8.RuneCountInString(myName))
	// 上をc.Validationにいれたい

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}
