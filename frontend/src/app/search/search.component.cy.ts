import { SearchComponent } from "./search.component"
import { MatSelectModule } from '@angular/material/select'
import { MatFormFieldModule } from '@angular/material/form-field'
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('search.component.cy.ts', () => {
  it('mounts', () => {
    cy.mount(SearchComponent, {
      imports: [MatSelectModule, MatFormFieldModule, BrowserAnimationsModule]
    })
  })
  it('should display three options: reading, math, and science', () => {
    
    cy.mount(SearchComponent, {
      imports: [MatSelectModule, MatFormFieldModule, BrowserAnimationsModule]
    })

    // Get the drop-down menu element
    const select = cy.get('mat-select');

    // Click on the drop-down menu to open it
    select.click();

    // Get the options within the drop-down menu
    const options = select.get('mat-option');

    // Check that there are three options
    options.should('have.length', 3);

    // Check that the options have the correct text
    options.eq(0).should('have.text', 'Reading');
    options.eq(1).should('have.text', '');
    options.eq(2).should('have.text', 'Science');
  });

  it('should allow the user to select an option', () => {    

    cy.mount(SearchComponent, {
      imports: [MatSelectModule, MatFormFieldModule, BrowserAnimationsModule]
    })

    // Get the drop-down menu element
    const select = cy.get('mat-select');

    // Select the "Math" option
    select.click();
    cy.get('mat-option').contains('Math').click();

    // Check that the selected value is "math"
    select.should('have.value', 'math');
  });
});
