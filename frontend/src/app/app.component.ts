import { Component } from '@angular/core';
import { AuthService } from './auth/auth.service';
import { LogoutComponent } from './logout/logout.component';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent {
  title = 'Find a Tutor';
  authService: AuthService;

  constructor(authService: AuthService, public dialog: MatDialog) {
    this.authService = authService
  }

  openDialog(): void {
    const dialogRef = this.dialog.open(LogoutComponent);
  }
}
