package main

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

const outOfRange = 99999
const daysInLastSixMonths = 183
const weeksInLastSixMonths = 26

type column []int

func stats(email string) {
	commits := processRepositories(email)
	print(commits)
}

func processRepositories(email string) map[int]int {
	dotFile := getDotFilePath()
	repositories := parseFileLinesToSlice(dotFile)
	daysInMap := daysInLastSixMonths

	commits := make(map[int]int, daysInMap)

	for i := daysInMap; i > 0; i-- {
		commits[i] = 0
	}

	for _, path := range repositories {
		commits = fillCommits(email, path, commits)
	}

	return commits
}

func fillCommits(email string, path string, commits map[int]int) map[int]int {
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}

	head, err := repo.Head()
	if err != nil {
		panic(err)
	}

	iterator, err := repo.Log(&git.LogOptions{From: head.Hash()})
	if err != nil {
		panic(err)
	}

	err = iterator.ForEach(func(c *object.Commit) error {
		daysAgo := c.Author.When.Day()

		if c.Author.Email != email {
			return nil
		}

		commits[daysAgo]++

		return nil
	})
	if err != nil {
		panic(err)
	}

	return commits
}
