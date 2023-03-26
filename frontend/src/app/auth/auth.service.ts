import { HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, tap } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private isLoggedInSubject = new BehaviorSubject<boolean>(false);
  public isLoggedIn$ = this.isLoggedInSubject.asObservable();

  constructor(private http: HttpClient) { }

  login(username: string, password: string): Observable<HttpResponse<any>> {
    return this.http.post<any>("/api/login", {username, password}).pipe(
      tap(response => {
        console.log(response.status);
        if (response.status == 200) {
          this.setLoggedIn(true);
        }
      })
    );
  }

  logout(): Observable<any> {
    return this.http.post<any>("/api/logout", {}).pipe(
      tap(response => {
        console.log(response);
        if (response.status == 200) {
          this.setLoggedIn(false);
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
