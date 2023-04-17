import { Component, Input } from '@angular/core';
import { NgForm } from '@angular/forms';

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
  // @Input() formData: {
    // search : string
    // } = {
    // search : ''
  // };
  
  // search(form: NgForm) {
    // list available tutors based on search
    // ignore capitilization and spaces
  // }
  
  searchSubject: string = '';
  tutors: Tutor[] = [
    { name: 'John Doe', subject: 'Mathematics' },
    { name: 'Jane Smith', subject: 'English' },
    { name: 'Bob Johnson', subject: 'History' },
    { name: 'Sara Lee', subject: 'Chemistry' },
    { name: 'Mark Davis', subject: 'Physics' },
    { name: 'Emily Jones', subject: 'Biology' },
    { name: 'David Lee', subject: 'Mathematics' },
    { name: 'Lisa Chen', subject: 'Mathematics' },
    { name: 'Mike Brown', subject: 'Mathematics' },
  ];

  filteredTutors: Tutor[] = [];

  filterTutors() {
    this.filteredTutors = this.tutors.filter(tutor => tutor.subject === this.searchSubject);
  }
}
