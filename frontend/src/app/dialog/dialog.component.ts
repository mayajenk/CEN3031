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
    phone: string,
    email: string,
    other: string
  } = {
    phone: '',
    email: '',
    other: ''
  };
  
  constructor(private authService: AuthService) {}

  saveContact(form: NgForm) {
    let user: User = this.authService.getUser();
    console.log("button closed");

    if (this.formData.phone !== '') {
      user.phone = this.formData.phone
    }
    if (this.formData.email !== '') {
      user.email = this.formData.email
    }
        if (this.formData.other !== '') {
      user.contact = this.formData.other
    }
    this.authService.updateUser(user).subscribe();
  }
}
