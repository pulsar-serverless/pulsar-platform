export interface Log {
    id: string;
    createdAt: string;
    message: string;
    type: 'Info' | 'Warning' | 'Error'
}