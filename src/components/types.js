// @flow

export type OrganizationType = {
  uid: string,
}

export type TopicType = {
  name: string,
  uid: string,
  externalId: string,
}

export type LinkType = {
  title: string,
  url: string,
  uid: string,
}

export type RelayProps = {
  organization: {
    id: string,
    uid: string,
  },
  relay: {
    environment: Object,
  },
  topic: {
    id: string,
    uid: string,
  },
}
