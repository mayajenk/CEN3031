import { LoginComponent } from "./login.component"
import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { LoginService } from "../login.service";
import { FormsModule } from '@angular/forms';

describe('LoginComponent', () => {
    beforeEach(() => TestBed.configureTestingModule({
        imports: [HttpClientTestingModule, FormsModule], 
        providers: [LoginService]
      }));

    it('mounts', () => {
        cy.mount(LoginComponent)
    })
})
