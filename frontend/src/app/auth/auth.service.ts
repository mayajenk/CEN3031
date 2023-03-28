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

  private isTutor = new BehaviorSubject<boolean>(false);
  public isTutor$ = this.isTutor.asObservable();


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
          // check if user is tutor and set value accordingly
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

  setTutor(value: boolean) {
    this.isTutor.next(value);
  }

  getIsTutor(): boolean {
    return this.isTutor.value;
  }
}
