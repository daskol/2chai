package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/daskol/2chai/api"
	"github.com/google/subcommands"
	_ "github.com/lib/pq"
)

var dsn string

func init() {
	dsn = os.Getenv("DSN")
	if dsn == "" {
		dsn = "postgresql://2chai@127.0.0.1/2chai?sslmode=disable"
	}
}

// listBoards реализует интерфейс subcommands.Commander для тго, чтобы вывести
// список возможныйх досок.
type listBoards struct{}

func (l *listBoards) Name() string     { return "list-boards" }
func (l *listBoards) Synopsis() string { return "List avaliable boards." }
func (l *listBoards) Usage() string {
	return "list-boards\n"
}

func (l *listBoards) SetFlags(_ *flag.FlagSet) {}

func (l *listBoards) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if len(f.Args()) > 0 {
		log.Println("Too many arguments.")
		return subcommands.ExitUsageError
	}

	if lst, err := api.ListBoards(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(lst)

		for i, board := range lst.Boards {
			log.Printf("[%03d] %s\n", i+1, board)
		}
	}

	return subcommands.ExitSuccess
}

// listThreads реализует интерфейс subcommands.Commander для того, чтобы можно
// было пролистать список нитей заданной доски.
type listThreads struct{}

func (l *listThreads) Name() string     { return "list-threads" }
func (l *listThreads) Synopsis() string { return "List threads of the given board." }
func (l *listThreads) Usage() string {
	return "list-threads BOARD\n"
}

func (l *listThreads) SetFlags(_ *flag.FlagSet) {}

func (l *listThreads) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	switch {
	case len(f.Args()) < 1:
		log.Println("Too few arguments.")
		return subcommands.ExitUsageError
	case len(f.Args()) > 1:
		log.Println("Too many arguments.")
		return subcommands.ExitUsageError
	}

	board := f.Args()[0]

	if lst, err := api.ListThreadCatalog(board); err != nil {
		log.Fatal(err)
	} else {
		log.Println(lst)

		for i, thread := range lst.Threads {
			log.Printf("[%03d] %s\n", i+1, thread)
		}
	}

	return subcommands.ExitSuccess
}

// listPosts реализует интерфейс subcommands.Commander для того, чтобы вывести
// список постов, принадлежащих заданной ните.
type listPosts struct{}

func (l *listPosts) Name() string     { return "list-posts" }
func (l *listPosts) Synopsis() string { return "list posts of the given thread" }
func (l *listPosts) Usage() string {
	return "list-posts BOARD THREAD\n"
}

func (l *listPosts) SetFlags(_ *flag.FlagSet) {}

func (l *listPosts) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	switch {
	case len(f.Args()) < 2:
		log.Println("Too few arguments.")
		return subcommands.ExitUsageError
	case len(f.Args()) > 2:
		log.Println("Too many arguments.")
		return subcommands.ExitUsageError
	}

	board := f.Args()[0]
	thread := f.Args()[1]

	if posts, err := api.ListPosts(board, thread); err != nil {
		log.Fatal(err)
	} else {
		for i, post := range posts {
			log.Printf("[%03d] %s\n", i+1, post)
		}
	}

	return subcommands.ExitSuccess
}

// syncBoards реализует интерфейс subcommands.Commander для того, чтобы
// наполнить базу данных списком доступных досок.
type syncBoards struct{}

func (s *syncBoards) Name() string     { return "sync-boards" }
func (s *syncBoards) Synopsis() string { return "Synchronize avaliable boards." }
func (s *syncBoards) Usage() string {
	return "sync-boards\n"
}

func (s *syncBoards) SetFlags(_ *flag.FlagSet) {}

func (s *syncBoards) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if len(f.Args()) > 0 {
		log.Println("Too many arguments.")
		return subcommands.ExitUsageError
	}

	if lst, err := api.ListBoards(); err != nil {
		log.Fatal(err)
	} else if err := s.upsertBoards(lst); err != nil {
		log.Fatal(err)
	}

	return subcommands.ExitSuccess
}

func (s *syncBoards) upsertBoards(boards *api.Boards) error {
	log.Println("connect to database")
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	log.Println("prepare insert or update statement")
	stmt := sq.
		Insert("boards").
		Columns("abbr", "name", "description").
		Suffix("" +
			"ON CONFLICT (abbr) " +
			"DO UPDATE " +
			"SET name = $2," +
			"    description = $3," +
			"    updated_at = CLOCK_TIMESTAMP()").
		PlaceholderFormat(sq.Dollar).
		RunWith(db)

	for _, board := range boards.Boards {
		stmt = stmt.Values(board.ID, board.Name, board.Info)
	}

	log.Println("execute statement")

	if _, err := stmt.Exec(); err != nil {
		return err
	}

	log.Println("done.")

	return nil
}

// syncPosts реализует интерфейс subcommands.Commander для того, чтобы
// добавить новые нити или обновить существующие.
type syncPosts struct{}

func (s *syncPosts) Name() string     { return "sync-posts" }
func (s *syncPosts) Synopsis() string { return "Synchronize posts of specified thread." }
func (s *syncPosts) Usage() string {
	return "sync-posts BOARD THREAD\n"
}

func (s *syncPosts) SetFlags(_ *flag.FlagSet) {}

func (s *syncPosts) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	switch {
	case len(f.Args()) < 2:
		log.Println("Too few arguments.")
		return subcommands.ExitUsageError
	case len(f.Args()) > 2:
		log.Println("Too many arguments.")
		return subcommands.ExitUsageError
	}

	board := f.Args()[0]
	thread := f.Args()[1]

	if lst, err := api.ListPosts(board, thread); err != nil {
		log.Fatal(err)
	} else if err := s.upsertPosts(board, thread, lst); err != nil {
		log.Fatal(err)
	}

	return subcommands.ExitSuccess
}

func (s *syncPosts) upsertPosts(board, thread string, posts []*api.Post) error {
	log.Println("connect to database")
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	log.Println("find identifier of board `" + board + "`")
	boardID := 0
	threadID, _ := strconv.Atoi(thread)
	rowScanner := sq.
		Select("board_id").
		From("boards").
		Where(sq.Eq{"abbr": board}).
		PlaceholderFormat(sq.Dollar).
		RunWith(db).
		QueryRow()

	if err := rowScanner.Scan(&boardID); err != nil {
		return err
	}

	log.Println("prepare insert or update statement")
	stmt := sq.
		Insert("posts").
		Columns("post_id", "thread_id", "board_id",
			"ordinal", "op",
			"author", "email", "subject", "comment",
			"created_at").
		Suffix("ON CONFLICT (board_id, post_id) DO NOTHING").
		PlaceholderFormat(sq.Dollar).
		RunWith(db)

	for _, post := range posts {
		createdAt := time.Unix(post.Timestamp, 0)

		stmt = stmt.Values(post.Num, threadID, boardID,
			post.Number, post.Op,
			post.Name, post.Email, post.Subject, post.Comment,
			createdAt)
	}

	log.Println("execute statement")

	if _, err := stmt.Exec(); err != nil {
		return err
	}

	log.Println("update thread update timestamp")
	createdAt := time.Unix(posts[0].LastHit, 0)
	_, err = sq.Update("threads").
		Set("updated_at", createdAt).
		Where(sq.Eq{"board_id": boardID, "thread_id": threadID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(db).
		Exec()

	if err != nil {
		return err
	}

	log.Println("done.")

	return nil
}

// syncThreads реализует интерфейс subcommands.Commander для того, чтобы
// добавить новые нити или обновить существующие.
type syncThreads struct{}

func (s *syncThreads) Name() string     { return "sync-threads" }
func (s *syncThreads) Synopsis() string { return "Synchronize threads of specified board." }
func (s *syncThreads) Usage() string {
	return "sync-threads BOARD\n"
}

func (s *syncThreads) SetFlags(_ *flag.FlagSet) {}

func (s *syncThreads) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	switch {
	case len(f.Args()) < 1:
		log.Println("Too few arguments.")
		return subcommands.ExitUsageError
	case len(f.Args()) > 1:
		log.Println("Too many arguments.")
		return subcommands.ExitUsageError
	}

	board := f.Args()[0]

	if lst, err := api.ListThreadCatalog(board); err != nil {
		log.Fatal(err)
	} else if err := s.upsertThreads(board, lst); err != nil {
		log.Fatal(err)
	}

	return subcommands.ExitSuccess
}

func (s *syncThreads) upsertThreads(board string, threads *api.Threads) error {
	log.Println("connect to database")
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	log.Println("find identifier of board `" + board + "`")
	boardID := 0
	rowScanner := sq.
		Select("board_id").
		From("boards").
		Where(sq.Eq{"abbr": board}).
		PlaceholderFormat(sq.Dollar).
		RunWith(db).
		QueryRow()

	if err := rowScanner.Scan(&boardID); err != nil {
		return err
	}

	log.Println("prepare insert or update statement")
	stmt := sq.
		Insert("threads").
		Columns("thread_id", "board_id", "subject", "created_at").
		Suffix("ON CONFLICT (board_id, thread_id) DO NOTHING").
		PlaceholderFormat(sq.Dollar).
		RunWith(db)

	for _, thread := range threads.Threads {
		createdAt := time.Unix(thread.Timestamp, 0)
		stmt = stmt.Values(thread.Num, boardID, thread.Subject, createdAt)
	}

	log.Println("execute statement")

	if _, err := stmt.Exec(); err != nil {
		return err
	}

	log.Println("done.")

	return nil
}

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	subcommands.Register(&listBoards{}, "")
	subcommands.Register(&listPosts{}, "")
	subcommands.Register(&listThreads{}, "")

	subcommands.Register(&syncBoards{}, "")
	subcommands.Register(&syncPosts{}, "")
	subcommands.Register(&syncThreads{}, "")

	flag.Parse()

	os.Exit(int(subcommands.Execute(context.Background())))
}
