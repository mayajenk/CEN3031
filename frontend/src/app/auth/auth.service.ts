import { HttpClient, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, tap } from 'rxjs';
import { CookieService } from 'ngx-cookie-service';
import { User } from '../user';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private isLoggedInSubject = new BehaviorSubject<boolean>(false);
  public isLoggedIn$ = this.isLoggedInSubject.asObservable();

  private isTutorSubject = new BehaviorSubject<boolean>(false);
  public isTutor$ = this.isTutorSubject.asObservable();

  private userSubject = new BehaviorSubject<User>({
    id: 0,
    username: "",
    first_name: "",
    last_name: "",
    is_tutor: false,
    rating: 0,
    subjects: [],
    email: "",
    phone: "",
    contact: "",
    about: "",
    grade: 0,
    });

  public user$ = this.userSubject.asObservable();


  constructor(private http: HttpClient, private cookieService: CookieService) {
    const isLoggedIn = this.cookieService.get('isLoggedIn');
    const isTutor = this.cookieService.get('isTutor');
    if (isLoggedIn === 'true') {
      this.isLoggedInSubject.next(true);
    }
    if (isTutor === 'true') {
      this.isTutorSubject.next(true);
    }
    else {
      this.isTutorSubject.next(false);
    }

    let options = {
      withCredentials: true
    };
    this.http.get<User>("/api/user", options).pipe(
      tap(response => {
        this.userSubject.next(response);
      })
    );    
  }

  register(username: string, password: string, is_tutor: boolean): Observable<HttpResponse<any>> {
    return this.http.post<any>("/api/users", {username, password, is_tutor});
  }


  login(username: string, password: string): Observable<any> {
    return this.http.post<any>("/api/login", {username, password}).pipe(
      tap(response => {
        if (response.status == 200) {
          this.setLoggedIn(true);
          this.cookieService.set('isLoggedIn', 'true');
          this.cookieService.set('id', response.user.id);
          // check if user is tutor and set value accordingly
          this.userSubject.next(response.user);
          if (response.user.is_tutor) {
            this.setTutor(true);
            this.cookieService.set('isTutor', 'true');
          }
          else {
            this.setTutor(false);
            this.cookieService.set('isTutor', 'false');
          }
        }
        return response;
      })
    );
  }

  registerAndLogin(username: string, password: string, is_tutor: boolean): Observable<HttpResponse<any>> {
    return this.register(username, password, is_tutor).pipe(
      tap(() => {
        return this.login(username, password).subscribe();
      })
    )
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
    this.isTutorSubject.next(value);
  }

  getIsTutor(): boolean {
    return this.isTutorSubject.value;
  }

  getUser(): User {
    return this.userSubject.value;
  }
}
