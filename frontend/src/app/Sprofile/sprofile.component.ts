import { Component } from '@angular/core';
import { User } from '../user';
import { ProfileService } from '../profile.service';

@Component({
  selector: 'app-sprofile',
  templateUrl: './sprofile.component.html',
  styleUrls: ['./sprofile.component.sass']
})
export class SprofileComponent {
  user: User = {
    id: 1,
    username: "foo",
    first_name: "Foo",
    last_name: "Bar",
    is_tutor: false,
    rating: 10.0,
    subjects: [{name: "Reading"}, {name: "Math"}],
    email: "foo@bar.com",
    phone: "000-000-0000",
    about: "Foo bar, foo bar. Foo foo foo, foo bar bar.",
    grade: 0
  }

  constructor(private profileService: ProfileService) { 
    this.profileService.getProfile()
      .subscribe(user => this.user = user);
  }
}
