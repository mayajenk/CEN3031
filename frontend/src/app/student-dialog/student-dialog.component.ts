import { Component } from '@angular/core';
import { NgForm } from '@angular/forms';
import { AuthService } from '../auth/auth.service';
import { User } from '../user';

@Component({
  selector: 'app-student-dialog',
  templateUrl: './student-dialog.component.html',
  styleUrls: ['./student-dialog.component.sass']
})
export class StudentDialogComponent {
  formData: {
    first_name: string,
    last_name: string,
    phone: string,
    email: string,
    other: string,
  } = {
    first_name: '',
    last_name: '',
    phone: '',
    email: '',
    other: '',
  };
  user: User
  
  constructor(private authService: AuthService) {
    this.user = this.authService.getUser();
    this.formData.first_name = this.user.first_name;
    this.formData.last_name = this.user.last_name;
    this.formData.phone = this.user.phone;
    this.formData.email = this.user.email;
    this.formData.other = this.user.contact;
  }

  saveInfo(form: NgForm) {
    this.user = this.authService.getUser();

    this.user.first_name = this.formData.first_name;
    this.user.last_name = this.formData.last_name;
    this.user.phone = this.formData.phone;
    this.user.email = this.formData.email;
    this.user.contact = this.formData.other;

    this.authService.updateUser(this.user).subscribe();
  }
}
