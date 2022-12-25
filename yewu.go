package main

import (
	"fmt"
	"os"
)

type huowu struct {
	title   string
	author  string
	price   float32
	publish bool
}

func showmenu() {
	fmt.Println("欢迎登陆业务!")
	fmt.Println("1.添加货物")
	fmt.Println("2.修改货物")
	fmt.Println("3.展示所有货物")
	fmt.Println("4.退出")
}

func userInput() *huowu {
	var (
		title   string
		author  string
		price   float32
		publish bool
	)
	fmt.Println("请根据提示输入相关内容")
	fmt.Print("请输入货物名:")
	fmt.Scanln(&title)
	fmt.Print("请输入送货人:")
	fmt.Scanln(&author)
	fmt.Print("请输入价格:")
	fmt.Scanln(&price)
	fmt.Print("请输入是否上架(true|false):")
	fmt.Scanln(&publish)
	fmt.Println(title, author, price, publish)
	book := newbook(title, author, price, publish)
	return book
}

var allbooks = make([]*huowu, 0, 200)

func newbook(title, author string, price float32, publish bool) *huowu {
	return &huowu{
		title:   title,
		author:  author,
		price:   price,
		publish: publish,
	}
}

func addbook() {
	var (
		title   string
		author  string
		price   float32
		publish bool
	)
	fmt.Println("请根据提示输入相关内容")
	fmt.Print("请输入货物名:")
	fmt.Scanln(&title)
	fmt.Print("请输入送货人:")
	fmt.Scanln(&author)
	fmt.Print("请输入价格:")
	fmt.Scanln(&price)
	fmt.Print("请输入是否上架(true|false):")
	fmt.Scanln(&publish)
	fmt.Println(title, author, price, publish)
	book := newbook(title, author, price, publish)
	for _, b := range allbooks {
		if b.title == book.title {
			fmt.Printf("《%s》此货物已经存在", book.title)
			return
		}
	}
	allbooks = append(allbooks, book)
	fmt.Println("添加货物成功！")
}

func updatebook() {
	book := userInput()
	for index, b := range allbooks {
		if b.title == book.title {
			allbooks[index] = book
			fmt.Printf("货物:《%s》更新成功！", book.title)
			return
		}
	}
	fmt.Printf("货物:《%s》不存在！", book.title)
}

func showbook() {
	if len(allbooks) == 0 {
		fmt.Println("啥也么有")
	}
	for _, b := range allbooks {
		fmt.Printf("《%s》送货人：%s 价格:%.2f 是否上架销售: %t\n", b.title, b.author, b.price, b.publish)
	}
}

func main() {
	for {
		showmenu()
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			addbook()
		case 2:
			updatebook()
		case 3:
			showbook()
		case 4:
			os.Exit(0)
		}
	}
}
