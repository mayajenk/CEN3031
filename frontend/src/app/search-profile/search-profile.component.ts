import { Component, OnInit } from '@angular/core';
import { User } from '../user';
import { AuthService } from '../auth/auth.service';
import { ActivatedRoute } from '@angular/router';
import { ProfileService } from '../profile.service';
import { NgForm } from '@angular/forms';
import { ReviewService } from '../review.service';

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

  review: {
    reviewer_id: number,
    reviewee_id: number,
    review_text: string,
    rating: number
  } = {
    reviewer_id: 0,
    reviewee_id: 0,
    review_text: '',
    rating: 1,
  }

  loading: boolean = true;

  isConnection: boolean = false;

  profilePictureURL = "/assets/img/avatar.webp"

  constructor(private route: ActivatedRoute, private profileService: ProfileService, private authService: AuthService, private reviewService: ReviewService) {}

  ngOnInit() {
    const id: number = this.route.snapshot.params['id'];
    this.profileService.getUser(id).subscribe(user => {
      this.user = user;
      let currUserID = this.authService.getUser().id;
      for (let i = 0; i < user.connections.length; i++) {
        if (user.connections[i].ID === currUserID) {
          this.isConnection = true;
        }
      }
      this.profilePictureURL = `/api/users/${id}/profile-picture`
      this.loading = false;
    })
  }

  setRatingBackground(rating: number) {
    if (rating == 0) {
      return `#595959`
    }
    else if (rating < 3) {
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

  submitReview() {
    this.review.reviewer_id = Number(this.authService.getUser().id);
    this.review.reviewee_id = Number(this.route.snapshot.params['id']);
    this.reviewService.addReview(this.review).subscribe();
  }

  addTutor() {
    let user_2: number = Number(this.route.snapshot.params['id'])
    this.authService.updateUserConnections(user_2).subscribe(
      (response: any) => {
        console.log('User connection saved successfully!');
        this.authService.updateBrowserStorage().subscribe();
      },
      (error: any) => {
        console.error('Error saving user connection:', error);
      }
    );
  }
}
