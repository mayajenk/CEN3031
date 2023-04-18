import { Component, Input } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { User } from '../user';

interface Tutor {
  name: string;
  subject: string;
}


@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.sass']
})

export class SearchComponent {
  subject: string = '';
  tutors: User[] = [];

  constructor(private http: HttpClient) { }

  searchTutors() {
    const url = `api/search?subject=${this.subject}`;
    this.http.get(url).subscribe((data: any) => {
      this.tutors = data;
    });
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
}
