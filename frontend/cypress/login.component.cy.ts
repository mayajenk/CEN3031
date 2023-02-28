import { LoginComponent } from "src/app/login/login.component";

describe('login.component.cy.ts', () => {
  it('login', () => {
    cy.visit('http://localhost:8080/login');
    cy.get(':nth-child(1) > .form-control').type('test@test.com');
    cy.get(':nth-child(2) > .form-control').type('test');
    cy.get('.btn').click();
  })
})
