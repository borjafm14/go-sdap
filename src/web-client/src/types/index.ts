export interface User {
    username: string;
    email: string;
    createdAt: string;
}

export interface ApiResponse<T> {
    status: Status;
    data?: T;
}

export enum Status {
    STATUS_OK = 'STATUS_OK',
    STATUS_ERROR = 'STATUS_ERROR',
}