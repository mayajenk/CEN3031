import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { User } from '../user';
import { AuthService } from '../auth/auth.service';
import { MatChipEditedEvent, MatChipInputEvent } from '@angular/material/chips';
import {COMMA, ENTER} from '@angular/cdk/keycodes';
import { MatLabel } from '@angular/material/form-field';
import { MatFormField } from '@angular/material/form-field';
import { Router } from '@angular/router';
import { ProfileService } from '../profile.service';
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
  
  constructor(private pService : ProfileService) {}

  saveContact(form: NgForm) {
    // this.pService.updateProfile(this.formData.phone, this.formData.email, this.formData.other)
    //   .subscribe(response => {
    //     console.log(response);
    //   }, error => {
    //     console.error(error);
    //   });
  }
}
