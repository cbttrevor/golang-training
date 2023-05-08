package main

import (
	"log"
	"os"

	"golang.org/x/exp/slog"
)

func BasicLog() {
	logFile, _ := os.OpenFile("golog.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0660)
	defer logFile.Close()

	mylogger := log.New(logFile, "CBTNUGGETS: ", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)
	mylogger.Println("Starting main function")
	DoWork(mylogger)
	// mylogger.Fatalln("A critical error occurred, so I'm dying")
	mylogger.Println("Program finished, exiting main")
}

func main() {
	// BasicLog()
	TestSlog()
}

func DoWork(l *log.Logger) {
	l.Println("Doing some work ....")
}

func TestSlog() {
	slog.Info("Welcome to structured logging with CBT Nuggets!")
	slog.Debug("This is a debug message from CBT Nuggets!")
	slog.Error("Something bad happened! I can't go on.")

	hopts := slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}
	// thandler := hopts.NewTextHandler(os.Stdout)
	jhandler := hopts.NewJSONHandler(os.Stdout)
	myslogger := slog.New(jhandler)

	attr01 := slog.String("FirstName", "Trevor")
	attr02 := slog.String("Severity", "Urgent")
	attrGroup := slog.Group("CustomData", attr01, attr02)

	myslogger.Info("Welcome to my slogger!", attr01)
	myslogger.Warn("Hey, just so you know, that's a bad idea!", attrGroup)
	myslogger.Debug("This is a debug helpful message")
}
