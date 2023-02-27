import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpResponse } from '@angular/common/http';
import { LoginComponent } from './login/login.component';
import { Observable } from 'rxjs/internal/Observable';

@Injectable({
  providedIn: 'root'
})
export class LoginService {
  constructor(private http: HttpClient) { }

  // Create a get http request (get product information in json format)
  getProduct(id: username): Observable<LoginComponent> {
    return this.http.get(`$this.http_product_url/${id}`)
        .map((response: Response) => response.json());
  }
// Create a post http request (post/add product data to server)
addProduct(context: any) {
    return this.http.post(`$this.http_product_url`, JSON.stringify(context))
        .map((response: Response) => response.json());
}
// Create a put http request (put/update product data to server)
updateProduct(id:number, context: any) {
    return this.http.put(`$this.http_product_url/${id}`, JSON.stringify(context))
        .map((response: Response) => response.json());
}
// Create a delete http request (delete product to server)
deleteProduct(id: number) {
    return this.http.delete(`$this.http_product_url/${id}`)
        .map((response: Response) => response.json());
}
}
