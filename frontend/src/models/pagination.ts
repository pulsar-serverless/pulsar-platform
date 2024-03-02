export interface Paginated<T> {
  limit: number;
  page: number;
  totalPages: number;
  rows: T[];
}
