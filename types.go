package main

import "fmt"

//go:generate ffjson $GOFILE

type Thread struct {
	Banned     int           `json:"banned"`
	Closed     int           `json:"closed"`
	Comment    string        `json:"comment"`
	Date       string        `json:"date"`
	Email      string        `json:"email"`
	Endless    int           `json:"endless"`
	Files      []interface{} `json:"files"`
	FilesCount int           `json:"files_count"`
	PostsCount int           `json:"posts_count"`
	LastHit    int64         `json:"lasthit"`
	Name       string        `json:"name"`
	Num        string        `json:"num"`
	Op         int           `json:"op"`
	Parent     string        `json:"parent"`
	Score      float64       `json:"score"`
	Sticky     int           `json:"sticky"`
	Subject    string        `json:subject"`
	Tags       string        `json:"tags"`
	Timestamp  int64         `json:"timestamp"`
	Trip       string        `json:"trip"`
	Views      int           `json:"views"`
}

func (t *Thread) String() string {
	return fmt.Sprintf("<Thread[%s] subj=%s;>", t.Num, t.Subject)
}

// https://2ch.hk/доска/номерстраницы.json (первая страница: index).
type Threads struct {
	AdvertBottomImage string `json:"advert_bottom_image"`
	AdvertBottomLink  string `json:"advert_bottom_link"`
	AdvertMobileImage string `json:"advert_mobile_image"`
	AdvertMobileLink  string `json:"advert_mobile_link"`
	AdvertTopLink     string `json:"advert_top_link"`
	AdvertTopImage    string `json:"advert_top_image"`

	Board            string `json:"Board"`
	BoardBannerLink  string `json:"board_banner_link"`
	BoardBannerImage string `json:"board_banner_image"`
	BoardInfo        string `json:"BoardInfo"`
	BoardInfoOuter   string `json:"BoardInfoOuter"`
	BoardName        string `json:"BoardName"`
	BoardSpeed       int    `json:"board_speed"`

	BumpLimit int `json:"bump_limit"`

	CurrentPage   int `json:"current_page"`
	CurrentThread int `json:"current_thread"`

	DefaultName string `json:"default_name"`

	EnableDices       int `json:"enable_dices"`
	EnableFlags       int `json:"enable_flags"`
	EnableIcons       int `json:"enable_icons"`
	EnableImages      int `json:"enable_images"`
	EnableNames       int `json:"enable_names"`
	EnableLikes       int `json:"enable_likes"`
	EnableOekaki      int `json:"enable_oekaki"`
	EnablePosting     int `json:"enable_posting"`
	EnableSage        int `json:"enable_sage"`
	EnableShield      int `json:"enable_shield"`
	EnableSubject     int `json:"enable_subject"`
	EnableThread_tags int `json:"enable_thread_tags"`
	EnableTrips       int `json:"enable_trips"`
	EnableVideo       int `json:"enable_video"`

	Filter string `json:"filter"`

	IsIndex int `json:"is_index"`
	IsBoard int `json:"is_board"`

	MaxComment   int           `json:"max_comment"`
	MaxFilesSize int           `json:"max_files_size"`
	NewsAbu      []interface{} `json:"news_abu"`
	Pages        []int         `json:"pages"`
	Threads      []*Thread     `json:"threads"`
	Top          []interface{} `json:"top"`
}

func (t *Threads) String() string {
	return fmt.Sprintf("<Threads[%s] nothreads=%d; page=%d; pages=%v;>",
		t.BoardName, len(t.Threads), t.CurrentPage, t.Pages)
}
