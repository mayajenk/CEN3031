import { Component } from '@angular/core';
import { User } from '../user';
import { AuthService } from '../auth/auth.service';

@Component({
  selector: 'app-tprofile',
  templateUrl: './tprofile.component.html',
  styleUrls: ['./tprofile.component.sass']
})
export class TprofileComponent {
  user: User = this.authService.getUser();

  constructor(private authService: AuthService) { 
  }
}
