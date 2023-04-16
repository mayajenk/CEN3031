import { Review } from "./review";

export interface User {
    id: number;
    username: string;
    first_name: string;
    last_name: string;
    is_tutor: boolean;
    rating: number;
    subjects: {name: string}[];
    email: string;
    phone: string;
    contact: string;
    about: string;
    grade: number;
    profile_picture: string;
    title: string;
    price: number;
    connections: User[];
    reviews: Review[];
}