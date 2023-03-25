import { Component, Input } from '@angular/core';
import { AuthService } from '../auth/auth.service';
import { NgForm } from '@angular/forms';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.sass']
})
export class LoginComponent {
  @Input() formData: {
    username: string,
    password: string
  } = {
    username: '',
    password: ''
  };

  private cookieValue: string | undefined;

  constructor(private cookieService: CookieService, private authService: AuthService) {}

  login(form: NgForm) {
    // check if username and password entered are valid
    // if username or password is invalid -> display message
    // if successful -> display message
    this.authService.login(this.formData.username, this.formData.password).subscribe(
      response => {
        console.log(response);
      }
    );
  }
}
