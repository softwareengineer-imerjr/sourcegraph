import { gql } from '@sourcegraph/http-client'

const REPO_FIELDS = gql`
    fragment ContextSelectorRepoFields on Repository {
        id
        name
        embeddingExists
        externalRepository {
            id
            serviceType
        }
    }
`

const EMBEDDING_JOB_FIELDS = gql`
    fragment ContextSelectorEmbeddingJobFields on RepoEmbeddingJob {
        id
        state
        failureMessage
    }
`

export const ReposSelectorSearchQuery = gql`
    query ReposSelectorSearch($query: String!, $includeJobs: Boolean!) {
        repositories(query: $query, first: 15) {
            nodes {
                ...ContextSelectorRepoFields
                embeddingJobs(first: 1) @include(if: $includeJobs) {
                    nodes {
                        ...ContextSelectorEmbeddingJobFields
                    }
                }
            }
        }
    }

    ${REPO_FIELDS}
    ${EMBEDDING_JOB_FIELDS}
`

export const SuggestedReposQuery = gql`
    query SuggestedRepos($names: [String!]!, $includeJobs: Boolean!) {
        byName: repositories(names: $names, first: 10) {
            nodes {
                ...ContextSelectorRepoFields
                embeddingJobs(first: 1) @include(if: $includeJobs) {
                    nodes {
                        ...ContextSelectorEmbeddingJobFields
                    }
                }
            }
        }
        # We also grab the first 10 embedded repos available on the site to
        # show in suggestions as a backup.
        firstTen: repositories(first: 10, embedded: true) {
            nodes {
                ...ContextSelectorRepoFields
                embeddingJobs(first: 1) @include(if: $includeJobs) {
                    nodes {
                        ...ContextSelectorEmbeddingJobFields
                    }
                }
            }
        }
    }

    ${REPO_FIELDS}
    ${EMBEDDING_JOB_FIELDS}
`

export const ReposStatusQuery = gql`
    query ReposStatus($repoNames: [String!]!, $first: Int!, $includeJobs: Boolean!) {
        repositories(names: $repoNames, first: $first) {
            nodes {
                ...ContextSelectorRepoFields
                embeddingJobs(first: 1) @include(if: $includeJobs) {
                    nodes {
                        ...ContextSelectorEmbeddingJobFields
                    }
                }
            }
        }
    }

    ${REPO_FIELDS}
    ${EMBEDDING_JOB_FIELDS}
`
