package domain

import "time"

// tipe data di struct ini jika masuk ke db, akan sesuai dengan default nya, int = 0, string  = " "
// jika mau null default value nya, maka menggunakan pointer, seperti *int, *string
type Exercise struct {
	ID          int
	Title       string
	Description string
	Questions []Question
}

type Question struct {
	ID            int
	ExerciseID    int
	Body          string
	OptionA       string
	OptionB       string
	OptionC       string
	OptionD       string
	CorrectAnswer string
	Score         int
	CreatorID     int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Answer struct {
	ID         int
	ExerciseID int
	QuestionID int
	UserID     int
	Answer     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
