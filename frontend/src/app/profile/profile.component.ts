import { Component } from '@angular/core';
import { User } from '../user';
import { ProfileService } from '../profile.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.sass'],
  providers: [ProfileService]
})
export class ProfileComponent {
  // Dummy data if request does not go through
  user: User = {
    id: 0,
    username: "",
    first_name: "",
    last_name: "",
    is_tutor: true,
    rating: 0,
    subjects: [{name: ""}],
    email: "",
    phone: "",
    about: "",
    grade: 0
  }

  constructor(private profileService: ProfileService) { 
  }
    
  ngOnInit(): void {
    this.profileService.getProfile().subscribe(user => this.user = user);
  }
}
