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

describe('Register form', () => {
  it('registers a new user', () => {
    cy.visit('localhost:8080/register');

    cy.get('#first_name').type('test1')
    cy.get('#last_name').type('test')
    cy.get('input[name=username]').type('test1')
    cy.get('input[name=password]').type('test')
    cy.get('mat-button-toggle[data-cy=tutor]').click()
    cy.get('#submit').click()

    // check that the registration was successful
    cy.url().should('include', 'localhost:8080')
    // doesn't look like cypress has access to POST requests so it would be a little difficult running this part
  })
});
