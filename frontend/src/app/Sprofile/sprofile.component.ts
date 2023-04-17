import { Component } from '@angular/core';
import { User } from '../user';
import { AuthService } from '../auth/auth.service';
import { MatDialog } from '@angular/material/dialog';
  import { StudentDialogComponent } from '../student-dialog/student-dialog.component';

@Component({
  selector: 'app-sprofile',
  templateUrl: './sprofile.component.html',
  styleUrls: ['./sprofile.component.sass']
})
export class SprofileComponent {
  user: User = this.authService.getUser();
  selectedFile: any = null;
  cacheBuster: string = '?cache=' + Math.random();
  profilePictureURL: string = `/api/users/${this.user.id}/profile-picture${this.cacheBuster}`

  constructor(private authService: AuthService, public dialog: MatDialog) { 
  }
  addOnBlur = true;

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
    if (rating < 3) {
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

  openDialog() {
    this.dialog.open(StudentDialogComponent);
  }
}
