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

export const CREATE_PROCEDURE = gql`
  mutation createProcedure($input: CreateProcedureInput!) {
    createProcedure(input: $input) {
      id
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
