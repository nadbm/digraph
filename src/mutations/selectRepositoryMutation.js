import { graphql } from 'react-relay'

import defaultMutation from './defaultMutation'

export default defaultMutation(graphql`
  mutation selectRepositoryMutation(
    $input: SelectRepositoryInput!
  ) {
    selectRepository(input: $input) {
      viewer {
        ...AddLink_viewer
        ...AddTopic_viewer
        ...SelectRepository_viewer
      }

      repository {
        id
        isPrivate
      }
    }
  }
`)