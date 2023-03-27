import { HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, tap } from 'rxjs';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private isLoggedInSubject = new BehaviorSubject<boolean>(false);
  public isLoggedIn$ = this.isLoggedInSubject.asObservable();

  constructor(private http: HttpClient, private cookieService: CookieService) {
    const isLoggedIn = this.cookieService.get('isLoggedIn');
    if (isLoggedIn === 'true') {
      this.isLoggedInSubject.next(true);
    }
   }

  login(username: string, password: string): Observable<HttpResponse<any>> {
    return this.http.post<any>("/api/login", {username, password}).pipe(
      tap(response => {
        if (response.status == 200) {
          this.setLoggedIn(true);
          this.cookieService.set('isLoggedIn', 'true');
        }
      })
    );
  }

  logout(): Observable<HttpResponse<any>> {
    return this.http.post<any>("/api/logout", {}).pipe(
      tap(response => {
        if (response.status == 200) {
          this.setLoggedIn(false);
          this.cookieService.delete('isLoggedIn');
        }
      })
    );
  }

  setLoggedIn(value: boolean) {
    this.isLoggedInSubject.next(value);
  }

  getIsLoggedIn(): boolean {
    return this.isLoggedInSubject.value;
  }
}
