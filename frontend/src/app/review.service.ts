import { Injectable } from '@angular/core';
import { Review } from './review';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ReviewService {

  constructor(private http: HttpClient) { }

  addReview(  review: {
    reviewer_id: number,
    reviewee_id: number,
    review_text: string,
    rating: number
  }) {
    return this.http.post("/api/review", review);
  }

  deleteReview(id: number) {
    return this.http.delete(`/api/review/${id}`);
  }
}
