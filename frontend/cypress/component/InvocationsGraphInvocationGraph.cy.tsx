import InvocationGraph from '@/components/analytics/InvocationsGraph'
import React from 'react'

describe('<InvocationGraph />', () => {
  it('renders', () => {
    cy.mount(<InvocationGraph projectId='' />)
  })
})