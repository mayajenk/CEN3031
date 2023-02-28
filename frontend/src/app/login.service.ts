import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { HttpResponse } from '@angular/common/http';
import { Observable } from 'rxjs/internal/Observable';

@Injectable({
  providedIn: 'root'
})
export class LoginService {
  private loginURL = "/api/login";

  constructor(private http: HttpClient) { }

  login(username: string, password: string) {
    const body = {username, password};
    const headers = new HttpHeaders().set('Content-Type', 'application/json');
    return this.http.post(this.loginURL, body, {headers});
  }

}
