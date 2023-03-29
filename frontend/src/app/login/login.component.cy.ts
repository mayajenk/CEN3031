import { LoginComponent } from "./login.component"
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { FormsModule } from '@angular/forms';
import { AuthService } from "../auth/auth.service";
import { MatCardModule } from "@angular/material/card";
import { MatFormFieldModule } from "@angular/material/form-field";
import { MatInputModule } from "@angular/material/input";
import { MatButtonModule } from "@angular/material/button";

describe('LoginComponent', () => {
    it('mounts', () => {
        cy.mount(LoginComponent, {
            imports: [HttpClientTestingModule, FormsModule, MatCardModule, MatFormFieldModule, MatInputModule, MatButtonModule],
            providers: [AuthService]
        })
    })
    it('pass input to username and password', () => {
        cy.mount(LoginComponent, {
            imports: [HttpClientTestingModule, FormsModule, MatCardModule, MatFormFieldModule, MatInputModule, MatButtonModule],
            providers: [AuthService],
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
