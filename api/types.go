package api

import "fmt"

//go:generate ffjson $GOFILE

type Tag struct {
	Board string `json:"board"` //	vg
	Tag   string `json:"tag"`   //	aion
}

// Board описывает глобальную доску.
type Board struct {
	BumpLimit     int    `json:"bump_limit"`     //	500
	Category      string `json:"category"`       //	"Разное"
	DefaultName   string `json:"default_name"`   //	"Аноним"
	EnableNames   int    `json:"enable_names"`   //	0
	EnableSage    int    `json:"enable_sage"`    //	1
	ID            string `json:"id"`             //	"b"
	Info          string `json:"info"`           //	"бред"
	LastNum       int    `json:"last_num"`       //	167025613
	Name          string `json:"name"`           //	"Бред"
	Speed         int    `json:"speed"`          //	2703
	Threads       int    `json:"threads"`        //	210
	UniquePosters int    `json:"unique_posters"` //	3727
}

func (b *Board) String() string {
	return fmt.Sprintf("<Board[%s] name=%s; nothreads=%d; speed=%d;>",
		b.ID, b.Name, b.Threads, b.Speed)
}

// Boards представляет список доступных глобальных досок.
type Boards struct {
	Boards       []*Board `json:"boards"`
	GlobalBoards int      `json:"global_boards"` //	156
	GlobalPosts  string   `json:"global_posts"`  //	"318,533,624\u0000"
	GlobalSpeed  string   `json:"global_speed"`  //	"6,107\u0000"
	IsIndex      int      `json:"is_index"`      //	1
	Tags         []*Tag   `json:"tags"`
	Type         int      `json:"type"` //	0
}

func (b *Boards) String() string {
	return fmt.Sprintf("<Boards[%d] noboards=%d; noglobal=%d; index=%d;>",
		b.Type, len(b.Boards), b.GlobalBoards, b.IsIndex)
}

// Top представляет запись, содержащую популярную нить на данный моммент.
type Top struct {
	board string `json:"board"`
	info  string `json:"info"`
	name  string `json:"name"`
}

// File описывает файловую запись в ответе.
type File struct {
	Height          int    `json:"height"`
	Width           int    `json:"width"`
	MD5             string `json:"md5"`
	Name            string `json:"name"`
	Path            string `json:"path"`
	Size            int    `json:"int"`
	Thumbnail       string `json:"thumbnail"`
	ThumbnailHeight int    `json:"tn_height"`
	ThumbnailWidth  int    `json:"tn_width"`
	Type            int    `json:"type"`
}

type Post struct {
	Banned     int     `json:"banned"`
	Closed     int     `json:"closed"`
	Comment    string  `json:"comment"`
	Date       string  `json:"date"`
	Email      string  `json:"email"`
	Endless    int     `json:"endless"`
	Files      []*File `json:"files"`
	FilesCount int     `json:"files_count"`
	PostsCount int     `json:"posts_count"`
	LastHit    int64   `json:"lasthit"`
	Name       string  `json:"name"`
	Num        int     `json:"num"`
	Number     int     `json:"number"`
	Op         int     `json:"op"`
	Parent     string  `json:"parent"`
	Sticky     int     `json:"sticky"`
	Subject    string  `json:subject"`
	Timestamp  int64   `json:"timestamp"`
	Trip       string  `json:"trip"`
}

func (p *Post) String() string {
	return fmt.Sprintf("<Post[%s:%d:%d] subj=%s; comment=%s;>",
		p.Parent, p.Number, p.Num, p.Subject, p.Comment)
}

type Thread struct {
	Banned     int     `json:"banned"`
	Closed     int     `json:"closed"`
	Comment    string  `json:"comment"`
	Date       string  `json:"date"`
	Email      string  `json:"email"`
	Endless    int     `json:"endless"`
	Files      []*File `json:"files"`
	FilesCount int     `json:"files_count"`
	PostsCount int     `json:"posts_count"`
	LastHit    int64   `json:"lasthit"`
	Name       string  `json:"name"`
	Num        string  `json:"num"`
	Op         int     `json:"op"`
	Parent     string  `json:"parent"`
	Score      float64 `json:"score"`
	Sticky     int     `json:"sticky"`
	Subject    string  `json:subject"`
	Tags       string  `json:"tags"`
	Timestamp  int64   `json:"timestamp"`
	Trip       string  `json:"trip"`
	Views      int     `json:"views"`
}

func (t *Thread) String() string {
	return fmt.Sprintf("<Thread[%s] subj=%s;>", t.Num, t.Subject)
}

// CommonAttributes содержит общие отрибуты ответы, такие как реклама или
// настройки доски.
//
// TODO: add missing
type CommonAttributes struct {
	Top []*Top `json:"top"`
}

type Posts struct {
	Posts []*Post `json:"posts"`
}

// ThreadResponse это результат ответа на запрос к API вернуть нить целиком.
// http(s)://2ch.hk/доска/res/номертреда.json
type ThreadResponse struct {
	*CommonAttributes

	Threads       []*Posts `json:"threads"`
	Title         string   `json:"title"`
	UniquePosters string   `json:"unique_posters"` // TODO: fix type
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
