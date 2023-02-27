import { Component } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.sass']
})
export class LoginComponent {
  login() {
    // check if username and password entered are valid
    // if username or password is invalid -> display message
    // if successful -> display message
  }
}
