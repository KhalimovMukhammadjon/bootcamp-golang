package repo

type BookData interface {
	GetBooks() []Library
	AddBook(Book)
	RemoveBookBook(Book)
	UpdateBook(Book)
	GetBookByld(Book)
	GetBookByCategory(Book)
}

type Library struct {
	Books []Book
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Status   string `json:"status"`
	Price    int    `json:"price"`
	Period   int    `json:"period"`
	Category string `json:"category"`
	Page     int    `json:"page"`
}
