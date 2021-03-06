// @flow
import React from 'react'
import { QueryRenderer, graphql } from 'react-relay'

import type { TopicType } from 'components/types'
import makeEditTopic from './EditTopic'

type Props = {
  isOpen: boolean,
  orgLogin: string,
  relay: {
    environment: Object,
  },
  toggleForm: Function,
  topic: TopicType,
}

const EditTopicContainer = ({ isOpen, orgLogin, topic, relay, toggleForm }: Props) => (
  <QueryRenderer
    environment={relay.environment}
    query={graphql`
      query EditTopicContainerQuery(
        $orgLogin: String!,
        $repoName: String,
        $repoIds: [ID!],
        $topicId: ID!,
      ) {
        viewer {
          ...EditTopicForm_viewer
        }

        view(
          currentOrganizationLogin: $orgLogin,
          currentRepositoryName: $repoName,
          repositoryIds: $repoIds,
        ) {
          topic(id: $topicId) {
            ...EditTopic_topic
          }
        }
      }
    `}
    variables={{
      orgLogin,
      repoName: null,
      repoIds: [],
      topicId: topic.id,
    }}
    render={makeEditTopic({ isOpen, orgLogin, relay, toggleForm })}
  />
)

export default EditTopicContainer
