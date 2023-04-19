import { HomeComponent } from "./home.component"

describe('Home page', () =>
{
  it('mounts', () =>
  {
    cy.mount(HomeComponent,
      {
      imports: [],
      providers: []
    })
  })
  it('displays the welcome message', () => {
    cy.mount(HomeComponent);

    cy.get('.title #h1title')
      .should('be.visible')
      .and('have.text', 'Welcome to Find a Tutor!');
  });

  it('displays the mission statement', () => {
    cy.mount(HomeComponent);

    cy.get('.statement h2')
      .should('be.visible')
      .and('have.text', 'Mission Statement');

    cy.get('.statement p').should('have.length', 3);
  });

  it('displays the description', () => {
    cy.mount(HomeComponent);

    cy.get('.description h2')
      .should('be.visible')
      .and('have.text', 'Get started in three easy steps:');

    cy.get('.description ul').should('have.length', 1);
    cy.get('.description li').should('have.length', 3);
  });
});
