describe('Check links', () => {
  it('Visit website and clicks login', () => {
    cy.visit('localhost:8080')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
  })
  it('Visit website and clicks Create Account', () => {
    cy.visit('localhost:8080')
    cy.contains('Register').click()
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
    cy.contains('Register').click()
    cy.url().should('include', '/register')
  })
})

describe('Login form', () => {
  it('logs in with correct credentials', () => {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('foo');
    cy.get('#password').type('bar');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');
  });
});
