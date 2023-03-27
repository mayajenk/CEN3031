import { Component } from '@angular/core';
import { NgForm } from '@angular/forms';
import { AuthService } from '../auth/auth.service';

@Component({
  selector: 'app-logout',
  templateUrl: './logout.component.html',
  styleUrls: ['./logout.component.sass']
})
export class LogoutComponent {

  constructor(private authService: AuthService) {}

  logout(form: NgForm) {
    // if loggedin, logout
    this.authService.logout();
  }
}
