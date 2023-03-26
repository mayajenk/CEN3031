import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { User } from '../user';
import { ProfileService } from '../profile.service';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.sass'],
  providers: [CookieService, ProfileService]
})
export class ProfileComponent {
  user: User = {
    id: 1,
    username: "foo",
    first_name: "Foo",
    last_name: "Bar",
    is_tutor: true,
    rating: 10.0,
    subjects: [{name: "Reading"}, {name: "Math"}],
    email: "foo@bar.com",
    phone: "000-000-0000",
    about: "Foo bar, foo bar. Foo foo foo, foo bar bar.",
    grade: 0
  }

  constructor() { 
  }
}
