import React from 'react'
import { shallow } from 'enzyme'
import Homepage from './index'

jest.mock('react-relay', () => ({ createFragmentContainer: component => component }))

describe('<Homepage />', () => {
  const viewer = { name: 'Gnusto' }
  const wrapper = shallow(<Homepage viewer={viewer} />)

  it('renders', () => {
    expect(wrapper).toMatchSnapshot()
  })
})
