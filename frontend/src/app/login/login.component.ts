import { Component } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { LoginService } from '../login.service';
import { NgForm } from '@angular/forms';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.sass']
})
export class LoginComponent {
  formData: {
    username: string,
    password: string
  } = {
    username: '',
    password: ''
  };

  private cookieValue: string | undefined;

  constructor(private loginService : LoginService, private cookieService: CookieService) {}

  login(form: NgForm) {
    // check if username and password entered are valid
    // if username or password is invalid -> display message
    // if successful -> display message
    this.loginService.login(this.formData.username, this.formData.password)
      .subscribe(response => {
        console.log(response);
      }, error => {
        console.error(error);
      });
    this.cookieValue = this.cookieService.get('session-name')
  }
}
