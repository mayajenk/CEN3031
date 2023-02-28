import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

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
