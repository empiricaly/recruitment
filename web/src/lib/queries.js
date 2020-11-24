import { gql } from "@apollo/client/core";

export const MTURK_QUALIFICATION_TYPES = gql`
  query mturkQualificationTypes($sandbox: Boolean) {
    mturkQualificationTypes(sandbox: $sandbox) {
      id
      name
      description
      type
    }
  }
`;

export const MTURK_LOCALES = gql`
  query mturkLocales($sandbox: Boolean) {
    mturkLocales(sandbox: $sandbox) {
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

export const DUPLICATE_RUN = gql`
  mutation duplicateRun($input: DuplicateRunInput!) {
    duplicateRun(input: $input) {
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

export const SCHEDULE_RUN = gql`
  mutation scheduleRun($input: ScheduleRunInput!) {
    scheduleRun(input: $input) {
      id
    }
  }
`;

export const UNSCHEDULE_RUN = gql`
  mutation unscheduleRun($input: UnscheduleRunInput!) {
    unscheduleRun(input: $input) {
      id
    }
  }
`;

export const START_RUN = gql`
  mutation startRun($input: StartRunInput!) {
    startRun(input: $input) {
      id
    }
  }
`;

export const CANCEL_RUN = gql`
  mutation cancelRun($input: CancelRunInput!) {
    cancelRun(input: $input) {
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

function recursiveComparator(level, indentation = 0) {
  let indent = "  ";
  if (level === 0) {
    indent = "";
    for (let index = 0; index < indentation; index++) {
      indent += "  ";
    }
  }

  return `{
  ${level < 1 ? "" : `and ${recursiveComparator(level - 1)}`}
  ${level < 1 ? "" : `or ${recursiveComparator(level - 1)}`}
  key
  comparator
  values {
    int
    float
    string
    boolean
  }
}`
    .split("\n")
    .map((l) => indent + l)
    .filter((l) => l.trim() !== "")
    .join("\n");
}

export const GET_RUN = gql`
  query getRun($projectID: ID!, $runID: ID!) {
    project(projectID: $projectID) {
      id
      projectID
      runs(runID: $runID) {
        id
        name
        status
        startAt
        startedAt
        endedAt
        error

        template {
          id
          name
          selectionType
          participantCount
          adult
          sandbox
          steps {
            id
            index
            type
            duration
            msgArgs {
              subject
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
              condition ${recursiveComparator(4, 14)}
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
            uninitialized
            all
            condition ${recursiveComparator(4, 7)}
          }
          steps {
            id
          }
        }
      }
    }
  }
`;

export const GET_RUNNING_RUN = gql`
  query getRuningRun($projectID: ID!, $runID: ID!) {
    project(projectID: $projectID) {
      id
      runs(runID: $runID) {
        id
        status
        error
        steps {
          id
          index
          status
          startedAt
          endedAt
          participantsCount
        }
      }
    }
  }
`;

export const GET_RUN_PARTICIPANTS = gql`
  query getRunParticipants($projectID: ID!, $runID: ID!) {
    project(projectID: $projectID) {
      id
      runs(runID: $runID) {
        id
        status
        steps {
          id
          participants {
            id
            mturkWorkerID
            data {
              key
              val
            }
          }
        }
      }
    }
  }
`;

export const GET_PROJECT_PARTICIPANTS = gql`
  query getProjectParticipants($projectID: ID!) {
    project(projectID: $projectID) {
      id
      participants {
        id
        mturkWorkerID
        createdAt
        data {
          key
          val
        }
      }
    }
  }
`;

export const GET_ALL_PARTICIPANTS = gql`
  query getAllParticipants {
    participants {
      id
      mturkWorkerID
      data {
        id
        key
        val
      }
    }
  }
`;

export const ADD_PARTICIPANTS = gql`
  mutation addParticipants($input: AddParticipantsInput!) {
    addParticipants(input: $input) {
      id
      mturkWorkerID
      data {
        key
        val
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
