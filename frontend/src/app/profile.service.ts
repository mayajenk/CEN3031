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

  updateProfile(user: User) {
    let id: number = user.id
    let requestBody = Object.keys(user).forEach((k) => user[k as keyof User] == null && delete user[k as keyof User]);
    return this.http.put<any>("/api/users/" + id, user);
  }
}
