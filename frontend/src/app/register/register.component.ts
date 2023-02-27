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
    tutor: boolean
  } = {
    username: '',
    password: '',
    tutor: false
  };
  
  constructor(private registerService : RegisterService) {}

  register(form: NgForm) {
    // check if username and password entered are valid
    // if username or password is invalid -> display message
    // check if username is already taken
    // if user name is already taken -> display message
    // if successful -> display message
    this.registerService.register(this.formData.username, this.formData.password, this.formData.tutor)
      .subscribe(response => {
        console.log(response);
      }, error => {
        console.error(error);
      });
  }
}
