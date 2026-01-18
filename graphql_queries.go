package leetcode

const kQuestionDataQuery = `
query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionId
    questionFrontendId
    title
    titleSlug
    difficulty
    content
    exampleTestcases
    sampleTestCase
    hints
    topicTags { name slug }
    codeSnippets { lang langSlug code }
  }
}
`
