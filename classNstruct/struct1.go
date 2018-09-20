package classNstruct

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "BrianZheng", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "wBrianZhengm", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", subject:"golang in action"})

	var Book1 Books /* 声明 Book1 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "Author"
	Book1.subject = "Subject Language"
	Book1.book_id = 6495407

	printBook(Book1)
	printBook2(&Book1)
	printBook(Book1)

}

func printBook(book Books) {  // value copy
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	book.subject = "changed"
	fmt.Printf("Book book_id : %d\n", book.book_id)
	fmt.Printf("---------\n")

}

func printBook2(book *Books) { // pass reference
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	book.author = "changed"
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
	fmt.Printf("---------\n")
}
