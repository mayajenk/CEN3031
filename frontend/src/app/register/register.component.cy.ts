import { HttpClientModule } from "@angular/common/http"
import { FormsModule } from "@angular/forms"
import { RegisterService } from "../register.service"
import { RegisterComponent } from "./register.component"

describe('register.component.cy.ts', () => {
  it('mounts', () => {
    cy.mount(RegisterComponent, {
      imports: [HttpClientModule, FormsModule],
      providers: [RegisterService],
    })
  })
  it('pass in username, password, and is_tutor input', () => {
    cy.mount(RegisterComponent, {
      imports: [HttpClientModule, FormsModule],
      providers: [RegisterService],
      componentProperties: {
        formData: {
          username: "foo",
          password: "bar",
          is_tutor: true
        }
      }
    })
    cy.get('input[name="username"]').should('have.value', "foo")
    cy.get('input[name="password"]').should('have.value', "bar")
    cy.get('input[name="is_tutor"]').should('have.value', "on")
  })
})