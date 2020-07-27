import { gql } from "apollo-boost";

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
