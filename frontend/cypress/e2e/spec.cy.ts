describe('Check links', () => {
  it('Visit website and clicks login', () => {
    cy.visit('localhost:8080')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
  })
  it('Visit website and clicks Create Account', () => {
    cy.visit('localhost:8080')
    cy.contains('Create Account').click()
    cy.url().should('include', '/register')
  })
})

describe('Check links', () => {
  it('Visit website and clicks login', () => {
    cy.visit('localhost:8080')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
  })
  it('Visit website and clicks Create Account', () => {
    cy.visit('localhost:8080')
    cy.contains('Create Account').click()
    cy.url().should('include', '/register')
  })
})