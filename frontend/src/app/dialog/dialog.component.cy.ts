import { DialogComponent } from "./dialog.component"
import { FormsModule } from '@angular/forms';
import { AuthService } from "../auth/auth.service";
import { MatCardModule } from "@angular/material/card";
import { MatFormFieldModule } from "@angular/material/form-field";
import { MatInputModule } from "@angular/material/input";
import { MatButtonModule } from "@angular/material/button";

describe('DialogComponent', () =>
{
  it('mounts', () =>
  {
    cy.mount(DialogComponent, {
      imports: [FormsModule, MatCardModule, MatFormFieldModule, MatInputModule, MatButtonModule],
      providers: [AuthService]
    })
  })

  it('should fill out the form fields', () => {
    const formData = {
      first_name: 'John',
      last_name: 'Doe',
      price: '50',
      title: 'Math Teacher',
      phone: '123-456-7890',
      email: 'john.doe@example.com',
      other: 'Some other contact info',
      about: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.'
    };

    cy.get('input[name="first_name"]').type(formData.first_name);
    cy.get('input[name="last_name"]').type(formData.last_name);
    cy.get('input[name="price"]').type(formData.price);
    cy.get('input[name="title"]').type(formData.title);
    cy.get('input[name="phone"]').type(formData.phone);
    cy.get('input[name="email"]').type(formData.email);
    cy.get('input[name="other"]').type(formData.other);
    cy.get('textarea[name="about"]').type(formData.about);

    cy.get('input[name="first_name"]').should('have.value', formData.first_name);
    cy.get('input[name="last_name"]').should('have.value', formData.last_name);
    cy.get('input[name="price"]').should('have.value', formData.price.toString());
    cy.get('input[name="title"]').should('have.value', formData.title);
    cy.get('input[name="phone"]').should('have.value', formData.phone);
    cy.get('input[name="email"]').should('have.value', formData.email);
    cy.get('input[name="other"]').should('have.value', formData.other);
    cy.get('textarea[name="about"]').should('have.value', formData.about);
  });

  it('should submit the form with the correct data', () => {
    const formData = {
      first_name: 'John',
      last_name: 'Doe',
      price: '50',
      title: 'Math Teacher',
      phone: '123-456-7890',
      email: 'john.doe@example.com',
      other: 'Some other contact info',
      about: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.'
    };

    cy.get('input[name="first_name"]').type(formData.first_name);
    cy.get('input[name="last_name"]').type(formData.last_name);
    cy.get('input[name="price"]').type(formData.price);
    cy.get('input[name="title"]').type(formData.title);
    cy.get('input[name="phone"]').type(formData.phone);
    cy.get('input[name="email"]').type(formData.email);
    cy.get('input[name="other"]').type(formData.other);
    cy.get('textarea[name="about"]').type(formData.about);

    cy.get('button#done').click();

    cy.get('@saveInfo').should('be.calledWithExactly', {
      first_name: formData.first_name,
      last_name: formData.last_name,
      price: formData.price,
      title: formData.title,
      phone: formData.phone,
      email: formData.email,
      other: formData.other,
      about: formData.about
    });
  });
});
