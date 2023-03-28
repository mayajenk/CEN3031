import { Component, Input } from '@angular/core';
import { SearchService } from '../search.service';
import { NgForm } from '@angular/forms';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.sass']
})
export class SearchComponent {
  @Input() formData: {
    search : string
    } = {
    search : ''
  };
  
  search() {
    // list available tutors based on search
  }
}