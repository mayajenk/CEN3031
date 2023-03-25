import { Component, Input } from '@angular/core';
import { NgForm } from '@angular/forms';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-logout',
  templateUrl: './logout.component.html',
  styleUrls: ['./logout.component.sass']
})
export class LogoutComponent {
  @Input() formData: {
    username: string,
    password: string
  } = {
    username: '',
    password: ''
  };

  private cookieValue: string | undefined;

  constructor(private cookieService: CookieService) {}

  logout(form: NgForm) {
    // if loggedin, logout
  }
}
