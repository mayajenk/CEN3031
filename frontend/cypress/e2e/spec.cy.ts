describe('Check links', () =>
{
  it('Visit website and clicks login', () =>
  {
    cy.visit('localhost:8080')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
  })

  it('Visit website and clicks Create Account', () =>
  {
    cy.visit('localhost:8080')
    cy.contains('Register').click()
    cy.url().should('include', '/register')
  })
})

describe('Check links', () =>
{

  it('Visit website and clicks login', () =>
  {
    cy.visit('localhost:8080')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
  })
  it('Visit website and clicks Create Account', () =>
  {
    cy.visit('localhost:8080')
    cy.contains('Register').click()
    cy.url().should('include', '/register')
  })
})

describe('Login form', () =>
{
  it('logs in with valid tutor credentials', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('foo');
    cy.get('#password').type('bar');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');
  });

  it('logs in with valid student credentials', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');
  });

  it('does not log in with invalid credentials', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('wee');
    cy.get('#password').type('woo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080/login');
  });
});

describe('Register form', () =>
{
  it('registers a new tutor user', () =>
  {
    cy.visit('localhost:8080/register');

    cy.get('#first_name').type('test1')
    cy.get('#last_name').type('test')
    cy.get('input[name=username]').type('test1')
    cy.get('input[name=password]').type('test')
    cy.get('mat-button-toggle[data-cy=tutor]').click()
    cy.get('#submit').click()

    cy.url().should('include', 'localhost:8080')
  })

  it('registers a new student user', () =>
  {
    cy.visit('localhost:8080/register');

    cy.get('#first_name').type('test2')
    cy.get('#last_name').type('test')
    cy.get('input[name=username]').type('test2')
    cy.get('input[name=password]').type('test')
    cy.get('mat-button-toggle[data-cy=student]').click()
    cy.get('#submit').click()

    cy.url().should('include', 'localhost:8080')
  })

  it('does not create a new account if the username is not unique', () => {
    cy.visit('localhost:8080/register');

    cy.get('#first_name').type('foo')
    cy.get('#last_name').type('bar')
    cy.get('input[name=username]').type('foo')
    cy.get('input[name=password]').type('bar')
    cy.get('mat-button-toggle[data-cy=tutor]').click()
    cy.get('#submit').click()

    cy.url().should('include', 'localhost:8080/register')
  })
});

describe('Home page', () =>
{
  it('shows all of the correct parts', () =>
  {
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

describe('Search page', () =>
{
  beforeEach(() =>
  {
    cy.visit('localhost:8080');
    cy.contains('Login').click();

  });
  it('makes sure the search page exists', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/search');
  });

  it('should allow searching for mathematics tutors after logging in', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/search');

    cy.get('#subject').type('Math');

    cy.get('button[color="primary"]').click();

    cy.get('.card-container mat-card').should('have.length.gt', 0);
  });

  it('should display tutor information correctly', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

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

  it('should allow searching for physics tutors after logging in', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/search');

    cy.get('#subject').type('Physics');

    cy.get('button[color="primary"]').click();

    cy.get('.card-container mat-card').should('have.length.gt', 0);
  });

  it('goes to the right page when clicking the link in the name', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/search');

    cy.get('#subject').type('Math');

    cy.get('button[color="primary"]').click();

    cy.get('.card-container mat-card').should('have.length.gt', 0);

    cy.get('.card-container mat-card:first-child a').click();

    cy.url().should('include', '/users/1');
  });

  it('displays nothing if no tutors have the desired subject', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/search');

    cy.get('#subject').type('English');
    cy.get('.card-container mat-card').should('not.exist');
  });
});

describe('Search profile pages', () =>
{
  beforeEach(() => {
    cy.visit('localhost:8080');
    cy.contains('Login').click();

  });
  it('displays correct tutor information when student views tutor profile', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/search');

    cy.get('#subject').type('Math');

    cy.get('button[color="primary"]').click();

    cy.get('.card-container mat-card').should('have.length.gt', 0);

    cy.get('.card-container mat-card:first-child a').click();

    cy.get('#title1').should('contain', 'Math Teacher');
    cy.get('#full-name').should('contain', 'foo bar');
    cy.get('#rating-number').should('contain', '6');
    cy.get('#addTutor').should('not.exist');
    cy.get('#profile-picture').should('have.attr', 'src', '/assets/img/avatar.webp');
    cy.get('.about').should('exist');
    cy.get('.about').should('contain', 'I like math');
    cy.get('.price').should('exist');
    cy.get('.price').should('contain', '$23/hr');
    cy.get('.contact').should('exist');
    cy.get('.contact').should('contain', '1234567890');
    cy.get('.courses').should('exist');
  });

  it('Adding a tutor as a student', () => {
    cy.url().should('include', '/login')
    cy.get('#username').type('bar');
    cy.get('#password').type('foo');
    cy.get('form').submit();
    cy.url().should('include', 'localhost:8080');
    cy.visit('localhost:8080/search');
    cy.get('#subject').type('Calculus');
    cy.get('button[color="primary"]').click();
    cy.get('.card-container mat-card').should('have.length.gt', 0);
    cy.get('.card-container mat-card:first-child a').click();
    cy.get('#full-name').should('contain', 'test1 test');
    cy.get('button[color="primary"]').click();
  });
});

describe('Tutor profile page', () =>
{
  beforeEach(() =>
  {
    cy.visit('localhost:8080');
    cy.contains('Login').click();
  });
  it('should display user information', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('foo');
    cy.get('#password').type('bar');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/profile');

    cy.get('.name').should('contain', 'foo bar');
    cy.get('.rating b').should('contain', '6');
    cy.get('.about p').should('contain', 'I like math');
    cy.get('.price h2').should('contain', '$23/hr');
    cy.get('.contact p').should('contain', 'Phone Number: 1234567890');
    cy.get('.contact p').should('contain', 'Email: foo@bar.com');
    cy.get('.contact p').should('contain', 'Other: N/A');
  });

  it('displays the review section for the tutor on their own page', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('foo');
    cy.get('#password').type('bar');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/profile');

    cy.get('.reviews').should('exist');
  });

  it('displays individual reviews for the tutor on their own page', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('foo');
    cy.get('#password').type('bar');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/profile');

    cy.get('.reviews').should('exist');

    cy.get('.reviews').should('contain', 'Meh...');
  });

  it('allows the tutor edit the info of themselves', () => {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('foo');
    cy.get('#password').type('bar');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/profile');

    cy.contains('Edit your info').click();
  });

  it('edits and saves the info of the tutor', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('foo');
    cy.get('#password').type('bar');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/profile');

    cy.contains('Edit your info').click();

    cy.get('input[name="first_name"]').clear();
    cy.get('input[name="first_name"]').type('foo');
    cy.get('input[name="last_name"]').clear();
    cy.get('input[name="last_name"]').type('bar');
    cy.get('input[name="price"]').clear();
    cy.get('input[name="price"]').type('$23/hr');
    cy.get('input[name="title"]').clear();
    cy.get('input[name="title"]').type('Math Teacher');
    cy.get('input[name="phone"]').clear();
    cy.get('input[name="phone"]').type('1234567890');
    cy.get('input[name="email"]').clear();
    cy.get('input[name="email"]').type('foo@bar.com');
    cy.get('input[name="other"]').clear();
    cy.get('input[name="other"]').type('N/A');

    cy.get('button#done').click();

    cy.get('.name').should('contain', 'foo bar');
    cy.get('.rating b').should('contain', '6');
    cy.get('.about p').should('contain', 'I like math');
    cy.get('.price h2').should('contain', '$23/hr');
    cy.get('.contact p').should('contain', 'Phone Number: 1234567890');
    cy.get('.contact p').should('contain', 'Email: foo@bar.com');
    cy.get('.contact p').should('contain', 'Other: N/A');
  });
});

describe('Student profile page', () =>
{
  beforeEach(() =>
  {
    cy.visit('localhost:8080');
    cy.contains('Login').click();
  });
  it('should display user information', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/profile');

    cy.get('.name').should('contain', 'bar foo');
    cy.get('.contact h1').should('contain', 'Contact Me');
    cy.get('.contact p').should('contain', 'Phone Number: 0987654321');
    cy.get('.contact p').should('contain', 'Email: bar@foo.com');
    cy.get('.contact p').should('contain', 'Other: N/A');
  });

  it('displays the reviews section for the student on their own page', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/profile');

    cy.get('.reviews').should('exist');
  });

  it('displays individual reviews for the tutor on their own page', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/profile');

    cy.get('.reviews').should('exist');

    cy.get('.reviews').should('contain', 'kinda bad...');
  });

  it('edits the info of the student', () =>
  {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/profile');

    cy.contains('Edit your info').click();
  });

  it('edits and saves the info of the student', () => {
    cy.visit('localhost:8080/login');

    cy.get('#username').type('bar');
    cy.get('#password').type('foo');

    cy.get('form').submit();

    cy.url().should('include', 'localhost:8080');

    cy.visit('localhost:8080/profile');

    cy.contains('Edit your info').click();

    cy.get('input[name="first_name"]').clear();
    cy.get('input[name="first_name"]').type('bar');
    cy.get('input[name="last_name"]').clear();
    cy.get('input[name="last_name"]').type('foo');
    cy.get('input[name="phone"]').clear();
    cy.get('input[name="phone"]').type('0987654321');
    cy.get('input[name="email"]').clear();
    cy.get('input[name="email"]').type('bar@foo.com');
    cy.get('input[name="other"]').clear();
    cy.get('input[name="other"]').type('N/A');

    cy.get('button#done').click();

    cy.get('.name').should('contain', 'bar foo');
    cy.get('.contact p').should('contain', 'Phone Number: 0987654321');
    cy.get('.contact p').should('contain', 'Email: bar@foo.com');
    cy.get('.contact p').should('contain', 'Other: N/A');
  });
});

describe('Logout feature', () =>
{
  beforeEach(() =>
  {
    cy.visit('localhost:8080');
    cy.contains('Login').click();
  });
  it('displays the logout confirmation message', () =>
  {
    cy.visit('localhost:8080')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
    cy.get('#username').type('foo');
    cy.get('#password').type('bar');
    cy.get('form').submit();
    cy.url().should('include', 'localhost:8080');

    cy.get('#logout-button').click();
    cy.contains('Are you sure you want to logout?').should('be.visible');
  });

  it('logs out after logging in', () =>
  {
    cy.visit('localhost:8080')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
    cy.get('#username').type('foo');
    cy.get('#password').type('bar');
    cy.get('form').submit();
    cy.url().should('include', 'localhost:8080');

    cy.get('#logout-button').click();
    cy.get('#logout').click();
  })

  it('does not log out after pressing No after clicking the logout button', () =>
  {
    cy.visit('localhost:8080')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
    cy.get('#username').type('foo');
    cy.get('#password').type('bar');
    cy.get('form').submit();
    cy.url().should('include', 'localhost:8080');

    cy.get('#logout-button').click();
    cy.get('#no').click();
  })
})
