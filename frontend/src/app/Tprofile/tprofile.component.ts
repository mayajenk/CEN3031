import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { User } from '../user';
import { AuthService } from '../auth/auth.service';
import { MatChipEditedEvent, MatChipInputEvent } from '@angular/material/chips';
import {COMMA, ENTER} from '@angular/cdk/keycodes';
import { DialogComponent } from '../dialog/dialog.component';
import { NgForm } from '@angular/forms';

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
  selectedFile: any = null;
  cacheBuster: string = '?cache=' + Math.random();
  profilePictureURL: string = `/api/users/${this.user.id}/profile-picture${this.cacheBuster}`

  constructor(private authService: AuthService, public dialog: MatDialog) { 
  }
  addOnBlur = true;
  readonly separatorKeysCodes = [ENTER, COMMA] as const;
  subjects: Subject[] = this.user.subjects;

  onFileSelected(event: any): void {
    this.selectedFile = event.target.files[0] ?? null;
    if (this.selectedFile != null) {
      const formData = new FormData();
      formData.append('file', this.selectedFile);
      this.authService.setProfilePicture(formData).subscribe(
        (response: any) => {
          this.cacheBuster = '?cache=' + Math.random();
          console.log('Cache buster updated to', this.cacheBuster);
          this.profilePictureURL = `/api/users/${this.user.id}/profile-picture${this.cacheBuster}`
        }
      )
    }
  }

  setRatingBackground(rating: number) {
    if (rating == 0) {
      return `#595959`
    }
    else if (rating < 3) {
      return `#e84c3f`
    }
    else if (rating < 6) {
      return `#b5b500`
    }
    else if (rating < 8) {
      return `#95b500`
    }
    else {
      return `#2ac325`
    }
  }

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

  saveSubjects() {
    // Create an array of subject names from the Subject objects
    const subjectNames: { name: string; }[] = this.subjects.map(subject => ({ name: subject.name }));
    // Call the API to update the user's subjects
    this.authService.updateUserSubjects(subjectNames).subscribe(
      (response: any) => {
        console.log('Subjects saved successfully!');
      },
      (error: any) => {
        console.error('Error saving subjects:', error);
      }
    );
  }
}
