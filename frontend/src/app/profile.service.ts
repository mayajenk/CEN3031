import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CookieService } from 'ngx-cookie-service';
import { Observable} from 'rxjs';
import { User } from './user';


@Injectable({
  providedIn: 'root'
})
export class ProfileService {
  private profileURL = "/api/user";

  constructor(private cookieService: CookieService, private http: HttpClient) { 
  }

  getProfile(): Observable<User> {
    console.log(this.cookieService.get("session"));
    let session : string = this.cookieService.get("session");
    let options = {
      withCredentials: true
    };
    
    return this.http.get<User>("/api/user", options);
  }
}
