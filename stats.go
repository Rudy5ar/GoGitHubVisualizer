package main

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
	return nil
}
