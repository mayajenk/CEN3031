import { Component } from '@angular/core';
import { User } from '../user';
import { ProfileService } from '../profile.service';
import { AuthService } from '../auth/auth.service';

@Component({
  selector: 'app-sprofile',
  templateUrl: './sprofile.component.html',
  styleUrls: ['./sprofile.component.sass']
})
export class SprofileComponent {
  user: User = this.authService.getUser();

  constructor(private authService: AuthService) { 
  }
}
