import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { User } from '../user';
import { AuthService } from '../auth/auth.service';
import { MatChipEditedEvent, MatChipInputEvent } from '@angular/material/chips';
import {COMMA, ENTER} from '@angular/cdk/keycodes';
import { MatLabel } from '@angular/material/form-field';
import { MatFormField } from '@angular/material/form-field';
import { DialogComponent } from '../dialog/dialog.component';

export interface Subject {
  name: string;
}

@Component({
  selector: 'app-tprofile',
  templateUrl: './tprofile.component.html',
  styleUrls: ['./tprofile.component.sass']
})
export class TprofileComponent {
  user: User = this.authService.getUser();

  constructor(private authService: AuthService, public dialog: MatDialog) { 
  }
  addOnBlur = true;
  readonly separatorKeysCodes = [ENTER, COMMA] as const;
  subjects: Subject[] = [];

  add(event: MatChipInputEvent): void {
    const value = (event.value || '').trim();

    // Add subject
    if (value) {
      this.subjects.push({name: value});
    }

    // Clear the input value
    event.chipInput!.clear();
  }

  remove(subject: Subject): void {
    const index = this.subjects.indexOf(subject);

    if (index >= 0) {
      this.subjects.splice(index, 1);
    }
  }

  edit(subject: Subject, event: MatChipEditedEvent) {
    const value = event.value.trim();

    // Remove subject if it no longer has a name
    if (!value) {
      this.remove(subject);
      return;
    }

    // Edit existing subjects
    const index = this.subjects.indexOf(subject);
    if (index >= 0) {
      this.subjects[index].name = value;
    }
  }

  openDialog() {
    this.dialog.open(DialogComponent);
  }
}