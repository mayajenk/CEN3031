import { Component, OnInit } from '@angular/core';
import { CookieService } from 'ngx-cookie-service/public-api';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent implements OnInit {
  title = 'frontend';
  
  private cookieValue: string | undefined;

  constructor(private cookieService: CookieService) {}

  public ngOnInit(): void {
    this.cookieValue = this.cookieService.get('session-name');
  }
}
