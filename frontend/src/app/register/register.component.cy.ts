import { HttpClientModule } from "@angular/common/http"
import { FormsModule } from "@angular/forms"
import { RegisterService } from "../register.service"
import { RegisterComponent } from "./register.component"
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatCardModule } from '@angular/material/card';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('register.component.cy.ts', () => {
  it('mounts', () => {
    cy.mount(RegisterComponent, {
      imports: [HttpClientModule, FormsModule, MatButtonToggleModule, MatFormFieldModule, MatButtonModule, MatInputModule, MatCardModule, BrowserAnimationsModule],
      providers: [RegisterService],
    })
  })
  it('pass in username, password, and is_tutor input', () => {
    cy.mount(RegisterComponent, {
      imports: [HttpClientModule, FormsModule, MatButtonToggleModule, MatFormFieldModule, MatButtonModule, MatInputModule, MatCardModule, BrowserAnimationsModule],
      providers: [RegisterService],
      componentProperties: {
        formData: {
          username: "foo",
          password: "bar",
          role: "tutor"
        }
      }
    })
    cy.get('input[name="username"]').should('have.value', "foo");
    cy.get('input[name="password"]').should('have.value', "bar");
    cy.get('[data-cy="tutor"]').find('button').should('be.enabled');
  })
})