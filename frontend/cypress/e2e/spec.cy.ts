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

describe('Home page', () => {
  it('shows all of the correct parts', () => {
    cy.visit('localhost:8080');

    cy.get('.title').should('exist');
    cy.contains('h1', 'Welcome to Find a Tutor!').should('exist');

    cy.get('.statement').should('exist');
    cy.contains('h2', 'Mission Statement').should('exist');
    cy.contains('Our mission at Find a Tutor is to provide accessible').should('exist');

    cy.get('.description').should('exist');
    cy.contains('h2', 'Get started in three easy steps:').should('exist');
    cy.contains('1. Create an account').should('exist');
    cy.contains('2. Search by subjects').should('exist');
    cy.contains('3. Connect with a tutor you like').should('exist');
  });
});

describe('Search page', () => {
  it('makes sure the search page exists', () => {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('foo');
    cy.get('#password').type('bar');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/search');
  });

  it('should allow searching for math tutors after logging in', () => {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('foo');
    cy.get('#password').type('bar');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/search');

    cy.get('#subject').type('Math');

    cy.get('button[color="primary"]').click();

    cy.get('.card-container mat-card').should('have.length.gt', 0);
  });

  it('should display tutor information correctly', () => {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('foo');
    cy.get('#password').type('bar');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/search');

    cy.get('#subject').type('Math');

    cy.get('button[color="primary"]').click();

    cy.get('.card-container mat-card').should('have.length.gt', 0);

    cy.get('.card-container mat-card:first-child h1.name').should('contain', 'foo bar');

    cy.get('.card-container mat-card:first-child button#price').should('contain', '$23/hr');

    cy.get('.card-container mat-card:first-child button.rating:last-child').should('contain', '6/10');

    cy.get('.card-container mat-card:first-child mat-card-content h2:first-child').should('contain', 'Math Teacher');

    cy.get('.card-container mat-card:first-child mat-card-content mat-chip-set').should('contain', 'Physics');
    cy.get('.card-container mat-card:first-child mat-card-content mat-chip-set').should('contain', 'Mathematics');
  });

  it('should allow searching for physics tutors after logging in', () => {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('foo');
    cy.get('#password').type('bar');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/search');

    cy.get('#subject').type('Physics');

    cy.get('button[color="primary"]').click();

    cy.get('.card-container mat-card').should('have.length.gt', 0);
  });
});
