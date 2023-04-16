export interface Review {
    id: number;
    reviewer_id: number;
    reviewee_id: number;
    review_text: string;
    rating: number;
}