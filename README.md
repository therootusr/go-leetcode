# go-leetcode

`go-leetcode` is a Go library that implements the practical LeetCode
endpoints used by tooling:

- GraphQL `questionData` fetch
- REST submit
- REST polling of submission status

## Package layout

This repository is a conventional single-package Go library (`package leetcode`)
at the module root.

## Live snapshot (golden) tests

Golden snapshots live under:

- `testdata/leetcode/*.json`

They are excluded from the default test run (because they hit the real LeetCode
servers). To run them explicitly:

- `go test -tags=live ./... -run TestLiveSnapshot_QuestionData -count=1`

To update snapshots from the current live output:

- `GO_LEETCODE_UPDATE_SNAPSHOTS=1 go test -tags=live ./... -run TestLiveSnapshot_QuestionData -count=1`

Optional:

- Override the live base URL: `GO_LEETCODE_LIVE_BASE_URL=https://leetcode.com`

## Security note

LeetCode authentication relies on session cookies (and an optional CSRF token).
Treat these values as secrets: do not log them or print them.
