import { Component } from '@angular/core';
import { User } from '../user';
import { AuthService } from '../auth/auth.service';

@Component({
  selector: 'app-search-profile',
  templateUrl: './search-profile.component.html',
  styleUrls: ['./search-profile.component.sass']
})
export class SearchProfileComponent {
  user: User = this.authService.getUser();

  constructor(private authService: AuthService) {}

  rate() {

  }
}
