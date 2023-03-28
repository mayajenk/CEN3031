describe('profile.component.cy.ts', () => {
  
    it('displays user name', () => {
      cy.get('.name').should('contain', 'Welcome foo!');
    });
  
    it('displays user rating', () => {
      cy.get('.rating').should('contain', 'Rating: 4.5');
    });
  
    it('displays user contact info', () => {
      cy.get('.info').should('contain', 'Contact Me:');
      cy.get('p').should('contain', 'Phone: 000-000-0000');
      cy.get('p').should('contain', 'Email: foo@bar.com');
    });
  
    it('displays user subjects', () => {
      cy.get('.courses').should('contain', 'My Subjects:');
      cy.get('.mat-chip').should('have.length', 2); 
    });
  
    it('displays current connections', () => {
      cy.get('.connections').should('contain', 'Current Connections:');
      cy.get('.mat-list-item').should('have.length', 3); 
    });
  
    it('allows user to edit contact info', () => {
      cy.get('[type="edit"]').click();
    });
  
    it('allows user to add a subject', () => {
      cy.get('[type="addSubject"]').click();
    });
  
    it('allows user to add a person', () => {
      cy.get('[type="addPerson"]').click();
    });
  });
  