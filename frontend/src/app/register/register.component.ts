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
    first_name: string,
    last_name: string,
    username: string,
    password: string,
    role: string
  } = {
    first_name: '',
    last_name: '',
    username: '',
    password: '',
    role: ''
  };
  
  constructor(private authService : AuthService, private router : Router) {}

  register(form: NgForm) {
    let is_tutor : boolean = this.formData.role == "tutor" ? true : false;
    this.authService.registerAndLogin(this.formData.first_name, this.formData.last_name, this.formData.username, this.formData.password, is_tutor)
      .subscribe(response => {
        console.log(response);
        this.router.navigate(["/"]);
      }, error => {
        console.error(error);
      });
  }
}
