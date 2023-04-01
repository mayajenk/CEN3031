import { Component } from '@angular/core';
import { User } from '../user';
import { TprofileComponent } from '../Tprofile/tprofile.component';
import { SprofileComponent } from '../Sprofile/sprofile.component';
import { ProfileService } from '../profile.service';
import { AuthService } from '../auth/auth.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.sass'],
  providers: [ProfileService]
})
export class ProfileComponent {
  constructor(public authService: AuthService){}
}
