package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/daskol/2chai/api"
	"github.com/google/subcommands"
)

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

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	subcommands.Register(&listBoards{}, "")
	subcommands.Register(&listPosts{}, "")
	subcommands.Register(&listThreads{}, "")

	flag.Parse()

	os.Exit(int(subcommands.Execute(context.Background())))
}
