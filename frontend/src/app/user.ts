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
}