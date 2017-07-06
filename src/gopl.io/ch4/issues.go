package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"gopl.io/ch4/github"
)

const ONE_MONTH = 30 * 24 * time.Hour
const ONE_YEAR  = 12 * ONE_MONTH

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()
	var monthIssues []*github.Issue
	var yearIssues  []*github.Issue
	var olderIssues []*github.Issue

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if item.CreatedAt.After(t.Add(-ONE_MONTH)) {
			append(monthIssues, item)
		} else if item.CreatedAt.After(t.Add(-ONE_YEAR)) && item.CreatedAt.Before(t.Add(-ONE_MONTH)) {
			append(yearIssues, item)
		} else {
			append(olderIssues, item)
		}
	}

	fmt.Println("Create less then one month:")
	for _, monthItem := range monthIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			monthItem.Number, monthItem.User.Login, monthItem.Title)
	}

	fmt.Println("Create less then one year:")
	for _, yearItem := range yearIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			yearItem.Number, yearItem.User.Login, yearItem.Title)
	}

	fmt.Println("Create more then one year:")
	for _, olderItem := range olderIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			olderItem.Number, olderItem.User.Login, olderItem.Title)
	}
}
