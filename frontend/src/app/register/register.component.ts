import { Component } from '@angular/core';
import { NgForm } from '@angular/forms';
import { AuthService } from '../auth/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.sass']
})
export class RegisterComponent {
  formData: {
    username: string,
    password: string,
    role: string
  } = {
    username: '',
    password: '',
    role: ''
  };
  
  constructor(private authService : AuthService, private router : Router) {}

  register(form: NgForm) {
    // check if username and password entered are valid
    // if username or password is invalid -> display message
    // check if username is already taken
    // if user name is already taken -> display message
    // if successful -> display message
    let is_tutor : boolean = this.formData.role == "tutor" ? true : false;
    this.authService.registerAndLogin(this.formData.username, this.formData.password, is_tutor)
      .subscribe(response => {
        console.log(response);
        this.router.navigate(["/"]);
      }, error => {
        console.error(error);
      });
  }
}
