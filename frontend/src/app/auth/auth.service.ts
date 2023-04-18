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
    profile_picture: "",
    title: "",
    price: 0,
    connections: [],
    reviews: []
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
    const userData: string | null = sessionStorage.getItem('userData');
    if (userData != null) {
      this.userSubject.next(JSON.parse(userData));
    }
  }

  register(first_name: string, last_name: string, username: string, password: string, is_tutor: boolean): Observable<HttpResponse<any>> {
    return this.http.post<any>("/api/users", {first_name, last_name, username, password, is_tutor});
  }


  login(username: string, password: string): Observable<any> {
    return this.http.post<any>("/api/login", {username, password}).pipe(
      tap(response => {
        if (response.status == 200) {
          this.setLoggedIn(true);
          this.cookieService.set('isLoggedIn', 'true');
          this.userSubject.next(response.user);
          sessionStorage.setItem('userData', JSON.stringify(response.user));
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

  registerAndLogin(first_name: string, last_name: string, username: string, password: string, is_tutor: boolean): Observable<HttpResponse<any>> {
    return this.register(first_name, last_name, username, password, is_tutor).pipe(
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
          this.cookieService.deleteAll();
          sessionStorage.clear();
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

  updateUser(user: User) {
    let id: number = this.userSubject.value.id
    return this.http.put<any>("/api/users/" + id, user).pipe(
      tap((response) => {
        this.userSubject.next(response);
        sessionStorage.setItem('userData', JSON.stringify(response));
      })
    )
  }

  getProfilePicture(): any {
    let id = this.userSubject.value.id;
    return this.http.get<any>("/api/users/" + id + "/profile-picture")
  }

  setProfilePicture(formData: any) {
    let id = this.userSubject.value.id;
    return this.http.post<any>("/api/users/" + id + "/profile-picture", formData).pipe(
      tap((response) => {
        let user: User = this.userSubject.value;
        user.profile_picture = response.filename;
        sessionStorage.setItem('userData', JSON.stringify(user));
        return response;
      })
    )
  }

  updateBrowserStorage() {
    let options = {
      withCredentials: true
    }
    let id: number = this.userSubject.value.id;
    return this.http.get(`/api/users/${id}`, options)
      .pipe(
        tap((response) => {
          sessionStorage.setItem('userData', JSON.stringify(response));
        })
      );
  }

  updateUserSubjects(subjects: {name: string}[]): Observable<User> {
    let id = this.userSubject.value.id;
    const user: User = this.getUser();
    user.subjects = subjects;
    return this.http.put<any>('/api/users/' + id + '/subjects', user.subjects);
  }

  updateUserConnections(user_2: number): Observable<User> {
    let user_1 = this.userSubject.value.id;

    let request_body = {
      user_1, user_2
    }
    const user: User = this.getUser();
    return this.http.post<any>('/api/connection', request_body);
  }
}
