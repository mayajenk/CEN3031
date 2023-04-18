import { SearchComponent } from "./search.component"
import { AppModule } from '../app.module';

describe('search.component.cy.ts', () => {
  beforeEach(() => {
  });

  it('mounts', () => {
    cy.mount(SearchComponent, {
      imports: [AppModule]
    })
  })
  it('should display search form and find tutors on search for math', () => {
    //cy.contains('button', 'Find Tutors').click(); will change later with static code since it can't work get access to other parts
  });
});
