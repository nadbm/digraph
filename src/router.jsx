import queryMiddleware from 'farce/lib/queryMiddleware'
import createRender from 'found/lib/createRender'
import makeRouteConfig from 'found/lib/makeRouteConfig'
import Route from 'found/lib/Route'
import { Resolver } from 'found-relay'
import React from 'react'
import { graphql } from 'react-relay'
import { Environment, Network, RecordSource, Store } from 'relay-runtime'

import { defaultOrganizationId } from './components/constants'
import Homepage, { query as homepageQuery } from './components/Homepage'
import TopicsPage, { query as topicsPageQuery } from './components/TopicsPage'
import TopicPage, { query as topicPageQuery } from './components/TopicPage'
import LinksPage, { query as linksPageQuery } from './components/LinksPage'
import Layout from './components/Layout'
import withErrorBoundary from './components/withErrorBoundary'

export const historyMiddlewares = [queryMiddleware]

export function createResolver(fetcher) {
  const environment = new Environment({
    network: Network.create((...args) => fetcher.fetch(...args)),
    store: new Store(new RecordSource()),
  })
  return new Resolver(environment)
}

const renderTopicPage = ({ props, error }: any) => {
  if (error)
    return <div>There was a problem.</div>
  if (!props)
    return <div>Loading ...</div>
  if (!props.view)
    return <div>You must log in and select an organization first.</div>
  return (
    <TopicPage
      topic={props.view.topic}
      {...props}
    />
  )
}

/* eslint function-paren-newline: 0 */
export const routeConfig = makeRouteConfig(
  <Route
    Component={Layout}
    path="/"
    query={
      graphql`
      query router_Query {
        viewer {
          name
          avatarUrl
        }
      }`
    }
    prepareVariables={(params, { location }) => {
      const { q } = location.query
      return {
        ...params,
        orgIds: [defaultOrganizationId],
        searchString: q,
      }
    }}
  >
    <Route
      Component={withErrorBoundary(Homepage)}
      query={homepageQuery}
    />
    <Route path="topics">
      <Route
        Component={withErrorBoundary(TopicsPage)}
        query={topicsPageQuery}
      />
      <Route
        path=":topicId"
        render={renderTopicPage}
        query={topicPageQuery}
      />
    </Route>
    <Route path="links">
      <Route
        Component={withErrorBoundary(LinksPage)}
        query={linksPageQuery}
      />
    </Route>
  </Route>,
)

export const render = createRender({})
