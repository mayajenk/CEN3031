import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { User } from '../user';
import { AuthService } from '../auth/auth.service';
import { MatChipEditedEvent, MatChipInputEvent } from '@angular/material/chips';
import {COMMA, ENTER} from '@angular/cdk/keycodes';
import { MatLabel } from '@angular/material/form-field';
import { MatFormField } from '@angular/material/form-field';

@Component({
  selector: 'app-dialog',
  templateUrl: './dialog.component.html',
  styleUrls: ['./dialog.component.sass']
})
export class DialogComponent {
  saveContact() {
    
  }
}
