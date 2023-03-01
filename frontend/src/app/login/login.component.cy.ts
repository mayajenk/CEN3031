import { LoginComponent } from "./login.component"
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { LoginService } from "../login.service";
import { FormsModule } from '@angular/forms';

describe('LoginComponent', () => {
    it('mounts', () => {
        cy.mount(LoginComponent, {
            imports: [HttpClientTestingModule, FormsModule],
            providers: [LoginService]
        })
    })
    it('pass input to username and password', () => {
        cy.mount(LoginComponent, {
            imports: [HttpClientTestingModule, FormsModule],
            providers: [LoginService],
            componentProperties: {
                formData: {
                    username: "foo",
                    password: "bar"
                }
            }
        })
        cy.get('input[name="username"]').should('have.value', "foo")
        cy.get('input[name="password"]').should('have.value', "bar")
    })
})
