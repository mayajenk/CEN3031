import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { User } from '../user';
import { MatChipEditedEvent, MatChipInputEvent } from '@angular/material/chips';
import {COMMA, ENTER} from '@angular/cdk/keycodes';
import { MatLabel } from '@angular/material/form-field';
import { MatFormField } from '@angular/material/form-field';
import { Router } from '@angular/router';
import { ProfileService } from '../profile.service';
import { AuthService } from '../auth/auth.service';
import { NgForm } from '@angular/forms';

@Component({
  selector: 'app-dialog',
  templateUrl: './dialog.component.html',
  styleUrls: ['./dialog.component.sass']
})
export class DialogComponent {
  formData: {
    first_name: string,
    last_name: string,
    price: number,
    title: string,
    phone: string,
    email: string,
    other: string,
    about: string
  } = {
    first_name: '',
    last_name: '',
    price: 0,
    title: '',
    phone: '',
    email: '',
    other: '',
    about: ''
  };
  user: User
  
  constructor(private authService: AuthService) {
    this.user = this.authService.getUser();
    this.formData.first_name = this.user.first_name;
    this.formData.last_name = this.user.last_name;
    this.formData.price = this.user.price;
    this.formData.title = this.user.title;
    this.formData.phone = this.user.phone;
    this.formData.email = this.user.email;
    this.formData.other = this.user.contact;
    this.formData.about = this.user.about;
  }

  saveInfo(form: NgForm) {
    this.user = this.authService.getUser();

    this.user.first_name = this.formData.first_name;
    this.user.last_name = this.formData.last_name;
    this.user.price = this.formData.price;
    this.user.title = this.formData.title;
    this.user.phone = this.formData.phone;
    this.user.email = this.formData.email;
    this.user.contact = this.formData.other;
    this.user.about = this.formData.about;

    this.authService.updateUser(this.user).subscribe();
  }
}
