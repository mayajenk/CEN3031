import { Component } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { AuthService } from '../auth/auth.service';

@Component({
  selector: 'app-logout',
  templateUrl: './logout.component.html',
  styleUrls: ['./logout.component.sass']
})
export class LogoutComponent {

  constructor(private authService: AuthService, public dialogRef: MatDialogRef<LogoutComponent>, private router: Router) {}

  logout() {
    // if loggedin, logout
    this.dialogRef.close();
    this.authService.logout().subscribe(
      () => this.router.navigate(["/"])
    );
  }
}
