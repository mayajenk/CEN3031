import { Component, Input } from '@angular/core';

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
    // ignore capitilization and spaces
  }
}