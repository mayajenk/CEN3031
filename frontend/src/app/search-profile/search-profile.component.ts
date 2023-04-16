import { Component, OnInit } from '@angular/core';
import { User } from '../user';
import { AuthService } from '../auth/auth.service';
import { ActivatedRoute } from '@angular/router';
import { ProfileService } from '../profile.service';

@Component({
  selector: 'app-search-profile',
  templateUrl: './search-profile.component.html',
  styleUrls: ['./search-profile.component.sass']
})
export class SearchProfileComponent implements OnInit {
  user: User = {
    id: 0,
    username: '',
    first_name: '',
    last_name: '',
    is_tutor: true,
    rating: 0,
    subjects: [],
    email: '',
    phone: '',
    contact: '',
    about: '',
    grade: 0,
    profile_picture: '',
    title: '',
    price: 0,
    connections: [],
    reviews: []
  }

  profilePictureURL = "/assets/img/avatar.webp"

  constructor(private route: ActivatedRoute, private profileService: ProfileService) {}

  ngOnInit() {
    const id: number = this.route.snapshot.params['id'];
    this.profileService.getUser(id).subscribe(user => {
      this.user = user;
      this.profilePictureURL = `/api/users/${id}/profile-picture`
    })
  }

  setRatingBackground(rating: number) {
    if (rating < 3) {
      return `#e84c3f`
    }
    else if (rating< 6) {
      return `#b5b500`
    }
    else if (rating < 8) {
      return `#95b500`
    }
    else {
      return `#2ac325`
    }
  }

  rate() {

  }
}
