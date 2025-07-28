package main

import (
    "fmt"
    "flag"
    "os"
    //"net/http"

    //"github.com/jean0t/AuraEvents/internal"
)


func migrate() {
    //var db = database.ConnectDB()
    //database.Migrate(db)
    fmt.Println("migration selected")
}

func startServer(arguments []string) {
    var (
        startFlag *flag.FlagSet = flag.NewFlagSet("server", flag.ExitOnError)
        port *string = startFlag.String("p", "8000", "Port to serve the requests")
    )
    startFlag.Parse(arguments)

    //router = routing.NewRouter()
    //_ = http.ListenAndServe(":"+*port, router)
    fmt.Println("Port is: ", *port)
}

func help() {
    fmt.Println("Help selected")
}

func main() {
    fmt.Println("os.Args: ", os.Args)
    if len(os.Args) < 2 {
        fmt.Println("Please select at least one command.")
        return
    }

    var command []string = os.Args
    fmt.Println("command: ", command)
    switch command[1] {
        case "migrate":
            migrate()

        case "start", "serve":
            startServer(command[2:])

        default:
            help()
    }
}
