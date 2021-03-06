// @flow
import React, { Component } from 'react'
import { createFragmentContainer, graphql } from 'react-relay'

import type { TopicType } from 'components/types'
import Input from 'components/ui/Input'
import SaveOrCancel from 'components/ui/SaveOrCancel'
import deleteTopicMutation from 'mutations/deleteTopicMutation'
import updateTopicMutation from 'mutations/updateTopicMutation'
import updateTopicTopicsMutation from 'mutations/updateTopicParentTopicsMutation'
import EditTopicList from 'components/ui/EditTopicList'
import DeleteButton from 'components/ui/DeleteButton'

type Props = {
  availableTopics: TopicType[],
  isOpen: boolean,
  relay: {
    environment: Object,
  },
  selectedTopics: TopicType[],
  toggleForm: Function,
  topic: TopicType,
}

type State = {
  description: ?string,
  name: string,
}

class EditTopicForm extends Component<Props, State> {
  constructor(props: Props) {
    super(props)
    this.state = {
      name: props.topic.name,
      description: props.topic.description,
    }
  }

  onSave = () => {
    updateTopicMutation(
      this.props.relay.environment,
      [],
      {
        topicIds: this.addTopicIds,
        description: this.state.description || '',
        id: this.props.topic.id,
        name: this.state.name,
      },
    )
    this.props.toggleForm()
  }

  onDelete = () => {
    deleteTopicMutation(
      this.props.relay.environment,
      [{
        type: 'NODE_DELETE',
        deletedIDFieldName: 'deletedTopicId',
      }],
      {
        topicId: this.props.topic.id,
      },
    )
  }

  // eslint-disable-next-line class-methods-use-this
  get addTopicIds(): string[] {
    return []
  }

  get topicId(): string {
    return this.props.topic.id
  }

  updateParentTopics = (parentTopicIds: string[]) => {
    updateTopicTopicsMutation(
      this.props.relay.environment,
      [],
      {
        topicId: this.props.topic.id,
        parentTopicIds,
      },
    )
  }

  updateDescription = (event: Object) => {
    this.setState({ description: event.currentTarget.value })
  }

  updateName = (event: Object) => {
    this.setState({ name: event.currentTarget.value })
  }

  render() {
    if (!this.props.isOpen)
      return null

    return (
      <div>
        <div className="d-flex col-12">
          <Input
            className="col-5 mr-3"
            id={`edit-link-title-${this.topicId}`}
            label="Name"
            onChange={this.updateName}
            value={this.state.name}
          />
          <Input
            className="col-6"
            id={`edit-topic-description-${this.topicId}`}
            label="Description"
            onChange={this.updateDescription}
            value={this.state.description}
          />
        </div>
        <div>
          <SaveOrCancel
            onSave={this.onSave}
            onCancel={this.props.toggleForm}
          />
          <DeleteButton
            className="float-right"
            onDelete={this.onDelete}
          />
        </div>
        <EditTopicList
          availableTopics={this.props.availableTopics}
          selectedTopics={this.props.selectedTopics}
          updateTopics={this.updateParentTopics}
        />
      </div>
    )
  }
}

export default createFragmentContainer(EditTopicForm, graphql`
  fragment EditTopicForm_viewer on User {
    defaultRepository {
      id
    }
  }

  fragment EditTopicForm_topic on Topic {
    description
    id
    name
  }
`)
