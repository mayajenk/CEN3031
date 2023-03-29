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
  it('should display a search bar and a "Search! button', () => {
    
    cy.mount(SearchComponent, {
      imports: [MatSelectModule, MatFormFieldModule, BrowserAnimationsModule]
    })

    cy.get('mat-form-field').should('exist');
    cy.get('input[name=search]').should('exist');
    cy.get('button[type=submit]').should('exist').contains('Search!');

  });
});
