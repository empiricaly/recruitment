import { gql } from "@apollo/client/core";

export const MTURK_QUALIFICATION_TYPES = gql`
  {
    mturkQualificationTypes {
      id
      name
      description
      type
    }
  }
`;

export const MTURK_LOCALES = gql`
  {
    mturkLocales {
      country
      subdivision
    }
  }
`;

export const CREATE_TEMPLATE = gql`
  mutation createTemplate($input: CreateTemplateInput!) {
    createTemplate(input: $input) {
      id
    }
  }
`;

export const UPDATE_TEMPLATE = gql`
  mutation updateTemplate($input: UpdateTemplateInput!) {
    updateTemplate(input: $input) {
      id
    }
  }
`;

export const CREATE_RUN = gql`
  mutation createRun($input: CreateRunInput!) {
    createRun(input: $input) {
      id
    }
  }
`;

export const UPDATE_RUN = gql`
  mutation updateRun($input: UpdateRunInput!) {
    updateRun(input: $input) {
      id
    }
  }
`;

export const GET_RUNS = gql`
  query getRuns($projectID: ID!, $limit: Int) {
    project(projectID: $projectID) {
      runs(limit: $limit) {
        id
        name
        status
        startAt
        startedAt
        endedAt

        template {
          steps {
            id
          }
        }
      }
    }
  }
`;

export const GET_RUN = gql`
  query getRun($projectID: ID!, $runID: ID!) {
    project(projectID: $projectID) {
      id
      runs(runID: $runID) {
        id
        name
        status
        startAt
        startedAt
        endedAt

        template {
          id
          name
          selectionType
          participantCount
          adult
          steps {
            id
            index
            type
            duration
            msgArgs {
              url
              message
              messageType
              lobby
              lobbyType
              lobbyExpiration
            }
            hitArgs {
              title
              description
              keywords
              microbatch
              reward
              timeout
              duration
              workersCount
            }
            filterArgs {
              type
              filter
              js
              condition {
                and {
                  and {
                    key
                    comparator
                    values {
                      int
                      float
                      string
                      boolean
                    }
                  }
                  or {
                    key
                    comparator
                    values {
                      int
                      float
                      string
                      boolean
                    }
                  }
                  key
                  comparator
                  values {
                    int
                    float
                    string
                    boolean
                  }
                }
                or {
                  and {
                    key
                    comparator
                    values {
                      int
                      float
                      string
                      boolean
                    }
                  }
                  or {
                    key
                    comparator
                    values {
                      int
                      float
                      string
                      boolean
                    }
                  }
                  key
                  comparator
                  values {
                    int
                    float
                    string
                    boolean
                  }
                }
                key
                comparator
                values {
                  int
                  float
                  string
                  boolean
                }
              }
            }
          }
          mturkCriteria {
            qualifications {
              id
              comparator
              values
              locales {
                country
                subdivision
              }
            }
          }
          internalCriteria {
            all
            condition {
              and {
                and {
                  key
                  comparator
                  values {
                    int
                    float
                    string
                    boolean
                  }
                }
                or {
                  key
                  comparator
                  values {
                    int
                    float
                    string
                    boolean
                  }
                }
                key
                comparator
                values {
                  int
                  float
                  string
                  boolean
                }
              }
              or {
                and {
                  key
                  comparator
                  values {
                    int
                    float
                    string
                    boolean
                  }
                }
                or {
                  key
                  comparator
                  values {
                    int
                    float
                    string
                    boolean
                  }
                }
                key
                comparator
                values {
                  int
                  float
                  string
                  boolean
                }
              }
              key
              comparator
              values {
                int
                float
                string
                boolean
              }
            }
          }
          steps {
            id
          }
        }
      }
    }
  }
`;

export const CREATE_PROJECT = gql`
  mutation createProject($input: CreateProjectInput!) {
    createProject(input: $input) {
      id
      projectID
    }
  }
`;

export const GET_PROJECTS = gql`
  {
    projects {
      id
      projectID
      name
    }
  }
`;

export const GET_PROJECT = gql`
  query getProject($projectID: ID!) {
    project(projectID: $projectID) {
      id
      projectID
      name
    }
  }
`;

export const AUTH = gql`
  mutation Auth($input: AuthInput) {
    auth(input: $input) {
      token
    }
  }
`;

export const ME = gql`
  {
    me {
      ... on Admin {
        id
        name
      }
    }
  }
`;
