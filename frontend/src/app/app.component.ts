import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent {
  title = 'frontend';
  swtichIntroScreen() {
    // if b2 == Create acount
    document.getElementById("intro").innerHTML = "Welcome";
    // switch button labels
    // add teacher or student check

    // else switch to login screen  
  }
}
