import { Component } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { RegisterService } from '../register.service';
import { NgForm } from '@angular/forms';

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
  
  constructor(private registerService : RegisterService) {}

  register(form: NgForm) {
    // check if username and password entered are valid
    // if username or password is invalid -> display message
    // check if username is already taken
    // if user name is already taken -> display message
    // if successful -> display message
    let is_tutor : boolean = this.formData.role == "tutor" ? true : false;
    this.registerService.register(this.formData.username, this.formData.password, is_tutor)
      .subscribe(response => {
        console.log(response);
      }, error => {
        console.error(error);
      });
  }
}
